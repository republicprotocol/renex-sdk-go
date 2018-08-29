package funds

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/republicprotocol/renex-ingress-go/httpadapter"
	"github.com/republicprotocol/renex-sdk-go/adapter/bindings"
	"github.com/republicprotocol/renex-sdk-go/adapter/client"
	"github.com/republicprotocol/renex-sdk-go/adapter/store"
	"github.com/republicprotocol/renex-sdk-go/adapter/trader"
	"github.com/republicprotocol/renex-sdk-go/core/funds"
	"github.com/republicprotocol/republic-go/order"
)

type adapter struct {
	renExBalancesContract *bindings.RenExBalances
	renExTokensContract   *bindings.RenExTokens
	client                client.Client
	trader                trader.Trader
	httpAddress           string
	store.Store
}

func NewAdapter(httpAddress string, client client.Client, trader trader.Trader, store store.Store) (funds.Adapter, error) {
	renExBalances, err := bindings.NewRenExBalances(client.RenExBalancesAddress(), bind.ContractBackend(client.Client()))
	if err != nil {
		return nil, err
	}
	renExTokens, err := bindings.NewRenExTokens(client.RenExTokensAddress(), bind.ContractBackend(client.Client()))
	if err != nil {
		return nil, err
	}
	return &adapter{
		renExBalancesContract: renExBalances,
		renExTokensContract:   renExTokens,
		httpAddress:           httpAddress,
		trader:                trader,
		Store:                 store,
	}, nil
}

func (adapter *adapter) RequestWithdrawalWithSignature(tokenCode order.Token, value *big.Int, signature []byte) error {
	token, err := adapter.renExTokensContract.Tokens(&bind.CallOpts{}, uint32(tokenCode))
	if err != nil {
		return err
	}

	if !token.Registered {
		return fmt.Errorf("Unregistered token")
	}

	tx, err := adapter.renExBalancesContract.Withdraw(adapter.trader.TransactOpts(), token.Addr, value, signature)
	if err != nil {
		return err
	}

	if _, err := adapter.client.WaitTillMined(context.Background(), tx); err != nil {
		return err
	}
	return nil
}

func (adapter *adapter) RequestWithdrawalFailSafeTrigger(tokenCode order.Token) (*funds.IdempotentKey, error) {
	token, err := adapter.renExTokensContract.Tokens(&bind.CallOpts{}, uint32(tokenCode))
	if err != nil {
		return nil, err
	}

	if !token.Registered {
		return nil, fmt.Errorf("Unregistered token")
	}

	tx, err := adapter.renExBalancesContract.SignalBackupWithdraw(adapter.trader.TransactOpts(), token.Addr)
	if err != nil {
		return nil, err
	}

	if _, err := adapter.client.WaitTillMined(context.Background(), tx); err != nil {
		return nil, err
	}

	hash32, err := toBytes32(tx.Hash().Bytes())
	if err != nil {
		return nil, err
	}
	key := funds.IdempotentKey(hash32)
	return &key, nil
}

func (adapter *adapter) RequestWithdrawalFailSafe(tokenCode order.Token, value *big.Int) error {
	token, err := adapter.renExTokensContract.Tokens(&bind.CallOpts{}, uint32(tokenCode))
	if err != nil {
		return err
	}

	if !token.Registered {
		return fmt.Errorf("Unregistered token")
	}

	tx, err := adapter.renExBalancesContract.Withdraw(adapter.trader.TransactOpts(), token.Addr, value, []byte{})
	if err != nil {
		return err
	}

	if _, err := adapter.client.WaitTillMined(context.Background(), tx); err != nil {
		return err
	}
	return nil
}

func (adapter *adapter) RequestWithdrawalSignature(tokenCode order.Token, value *big.Int) ([]byte, error) {
	req := httpadapter.ApproveWithdrawalRequest{
		Trader:  adapter.trader.Address().String(),
		TokenID: uint32(tokenCode),
	}

	data, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(data)

	resp, err := http.DefaultClient.Post(fmt.Sprintf("%s/withdrawals", adapter.httpAddress), "application/json", buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Signature string `json:"signature"`
	}

	respose := Response{}
	if err := json.Unmarshal(respBytes, respose); err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("Unexpected status code %d", resp.StatusCode)
	}

	return base64.StdEncoding.DecodeString(respose.Signature)
}

func (adapter *adapter) RequestDeposit(tokenCode order.Token, value *big.Int) error {
	token, err := adapter.renExTokensContract.Tokens(&bind.CallOpts{}, uint32(tokenCode))
	if err != nil {
		return err
	}

	if !token.Registered {
		return fmt.Errorf("Unregistered token")
	}

	addr, err := adapter.renExBalancesContract.ETHEREUM(&bind.CallOpts{})
	if err != nil {
		return err
	}

	if addr.String() == token.Addr.String() {
		auth := adapter.trader.TransactOpts()
		auth.Value = value
		tx, err := adapter.renExBalancesContract.Deposit(auth, token.Addr, value)
		if err != nil {
			return err
		}
		if _, err := adapter.client.WaitTillMined(context.Background(), tx); err != nil {
			return err
		}
		return nil
	}

	tokenContract, err := bindings.NewERC20(token.Addr, bind.ContractBackend(adapter.client.Client()))
	if err != nil {
		return err
	}

	tx, err := tokenContract.Approve(adapter.trader.TransactOpts(), adapter.client.RenExBalancesAddress(), value)
	if err != nil {
		return err
	}
	if _, err := adapter.client.WaitTillMined(context.Background(), tx); err != nil {
		return err
	}

	tx2, err := adapter.renExBalancesContract.Deposit(adapter.trader.TransactOpts(), token.Addr, value)
	if err != nil {
		return err
	}

	if _, err := adapter.client.WaitTillMined(context.Background(), tx2); err != nil {
		return err
	}
	return nil
}

func (adapter *adapter) CheckStatus(key *funds.IdempotentKey) uint8 {
	if key == nil {
		return uint8(0)
	}

	tx, status, err := adapter.client.Client().TransactionByHash(context.Background(), common.Hash(*key))
	if err != nil {
		return uint8(7)
	}

	if !status {
		return uint8(1)
	}

	txData := tx.Data()
	if len(tx.Data()) != 36 {
		return uint8(7)
	}

	traderWithdrawalFailSafeTime, err := adapter.renExBalancesContract.TraderWithdrawalSignals(&bind.CallOpts{}, adapter.trader.Address(), common.BytesToAddress(txData[4:36]))
	if err != nil {
		return uint8(7)
	}

	if time.Now().Unix() > traderWithdrawalFailSafeTime.Int64() {
		return uint8(3)
	}

	return uint8(2)
}

func (adapter *adapter) Balance(tokenCode order.Token) (*big.Int, error) {
	token, err := adapter.renExTokensContract.Tokens(&bind.CallOpts{}, uint32(tokenCode))
	if err != nil {
		return nil, err
	}

	if !token.Registered {
		return nil, fmt.Errorf("Unregistered token")
	}

	return adapter.renExBalancesContract.TraderBalances(&bind.CallOpts{}, adapter.trader.Address(), token.Addr)
}

func (adapter *adapter) Address() string {
	return adapter.trader.Address().String()
}

func (adapter *adapter) TransferEth(address string, value *big.Int) error {
	return adapter.client.Transfer(common.HexToAddress(address), adapter.trader.TransactOpts(), value)
}

func (adapter *adapter) TransferERC20(address string, tokenCode order.Token, value *big.Int) error {
	token, err := adapter.renExTokensContract.Tokens(&bind.CallOpts{}, uint32(tokenCode))
	if err != nil {
		return err
	}

	erc20, err := bindings.NewERC20(token.Addr, bind.ContractBackend(adapter.client.Client()))
	if err != nil {
		return err
	}

	tx, err := erc20.Transfer(adapter.trader.TransactOpts(), common.HexToAddress(address), value)
	if err != nil {
		return err
	}

	if _, err := adapter.client.WaitTillMined(context.Background(), tx); err != nil {
		return err
	}

	return nil
}

func toBytes32(b []byte) ([32]byte, error) {
	bytes32 := [32]byte{}
	if len(b) != 32 {
		return bytes32, errors.New("Length mismatch")
	}
	for i := range b {
		bytes32[i] = b[i]
	}
	return bytes32, nil
}
