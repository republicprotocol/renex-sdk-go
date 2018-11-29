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

	"github.com/republicprotocol/beth-go"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/republicprotocol/renex-sdk-go/adapter/bindings"
	"github.com/republicprotocol/renex-sdk-go/adapter/client"
	"github.com/republicprotocol/renex-sdk-go/adapter/store"
	"github.com/republicprotocol/renex-sdk-go/core/funds"
	"github.com/republicprotocol/republicprotocol-go/foundation/order"
)

type adapter struct {
	renExBalancesContract *bindings.RenExBalances
	renExTokensContract   *bindings.RenExTokens
	client                client.Client
	trader                beth.Account
	httpAddress           string
	store.Store
}

type approveWithdrawalRequest struct {
	Trader  string `json:"address"`
	TokenID uint32 `json:"tokenID"`
}

func NewAdapter(httpAddress string, client client.Client, trader beth.Account, store store.Store) (funds.Adapter, error) {
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
		client:                client,
	}, nil
}

func (adapter *adapter) RequestWithdrawalWithSignature(ctx context.Context, tokenCode order.Token, value *big.Int, signature []byte) error {
	token, err := adapter.renExTokensContract.Tokens(&bind.CallOpts{}, uint32(tokenCode))
	if err != nil {
		return err
	}
	if !token.Registered {
		return fmt.Errorf("Unregistered token")
	}

	txFunc := func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return adapter.renExBalancesContract.Withdraw(opts, token.Addr, value, signature)
	}
	return adapter.trader.Transact(ctx, nil, txFunc, nil, 1)
}

func (adapter *adapter) RequestWithdrawalFailSafeTrigger(ctx context.Context, tokenCode order.Token) (*funds.IdempotentKey, error) {
	token, err := adapter.renExTokensContract.Tokens(&bind.CallOpts{}, uint32(tokenCode))
	if err != nil {
		return nil, err
	}

	if !token.Registered {
		return nil, fmt.Errorf("Unregistered token")
	}

	var hash32 [32]byte
	txFunc := func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
		tx, err := adapter.renExBalancesContract.SignalBackupWithdraw(transactOpts, token.Addr)
		if err != nil {
			return nil, err
		}
		hash32, err = toBytes32(tx.Hash().Bytes())
		if err != nil {
			return nil, err
		}
		return tx, nil
	}
	if err := adapter.trader.Transact(ctx, nil, txFunc, nil, 1); err != nil {
		return nil, err
	}

	key := funds.IdempotentKey(hash32)
	return &key, nil
}

func (adapter *adapter) RequestWithdrawalFailSafe(ctx context.Context, tokenCode order.Token, value *big.Int) error {
	token, err := adapter.renExTokensContract.Tokens(&bind.CallOpts{}, uint32(tokenCode))
	if err != nil {
		return err
	}

	if !token.Registered {
		return fmt.Errorf("Unregistered token")
	}

	txFunx := func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
		return adapter.renExBalancesContract.Withdraw(transactOpts, token.Addr, value, []byte{})
	}
	return adapter.trader.Transact(ctx, nil, txFunx, nil, 1)
}

func (adapter *adapter) RequestWithdrawalSignature(tokenCode order.Token, value *big.Int) ([]byte, error) {
	req := approveWithdrawalRequest{
		Trader:  adapter.trader.Address().String()[2:],
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

	if !(resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK) {
		return nil, fmt.Errorf("Unexpected status code %d: %s %s", resp.StatusCode, resp.Status, string(respBytes))
	}

	type Response struct {
		Signature string `json:"signature"`
	}

	response := Response{}
	if err := json.Unmarshal(respBytes, &response); err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(response.Signature)
}

func (adapter *adapter) RequestDeposit(ctx context.Context, tokenCode order.Token, value *big.Int) error {
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
		txFunc := func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
			transactOpts.Value = value
			return adapter.renExBalancesContract.Deposit(transactOpts, token.Addr, value)
		}
		return adapter.trader.Transact(ctx, nil, txFunc, nil, 1)
	}

	tokenContract, err := bindings.NewERC20(token.Addr, bind.ContractBackend(adapter.client.Client()))
	if err != nil {
		return err
	}

	txFunc := func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
		return tokenContract.Approve(transactOpts, adapter.client.RenExBalancesAddress(), value)
	}
	if err := adapter.trader.Transact(ctx, nil, txFunc, nil, 1); err != nil {
		return err
	}

	txFunc = func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
		return adapter.renExBalancesContract.Deposit(transactOpts, token.Addr, value)
	}
	return adapter.trader.Transact(ctx, nil, txFunc, nil, 1)
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

func (adapter *adapter) RenExBalance(tokenCode order.Token) (*big.Int, error) {
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

func (adapter *adapter) TransferEth(ctx context.Context, address string, value *big.Int) error {
	return adapter.trader.Transfer(ctx, common.HexToAddress(address), value, 1)
}

func (adapter *adapter) TransferERC20(ctx context.Context, address string, tokenCode order.Token, value *big.Int) error {
	token, err := adapter.renExTokensContract.Tokens(&bind.CallOpts{}, uint32(tokenCode))
	if err != nil {
		return err
	}

	erc20, err := bindings.NewERC20(token.Addr, bind.ContractBackend(adapter.client.Client()))
	if err != nil {
		return err
	}

	txFunc := func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
		return erc20.Transfer(transactOpts, common.HexToAddress(address), value)
	}
	return adapter.trader.Transact(ctx, nil, txFunc, nil, 1)
}

func (adapter *adapter) BalanceEth() (*big.Int, error) {
	return adapter.client.Client().BalanceAt(context.Background(), adapter.trader.Address(), nil)
}

func (adapter *adapter) BalanceErc20(tokenCode order.Token) (*big.Int, error) {
	token, err := adapter.renExTokensContract.Tokens(&bind.CallOpts{}, uint32(tokenCode))
	if err != nil {
		return nil, err
	}

	erc20, err := bindings.NewERC20(token.Addr, bind.ContractBackend(adapter.client.Client()))
	if err != nil {
		return nil, err
	}

	return erc20.BalanceOf(&bind.CallOpts{}, adapter.trader.Address())
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
