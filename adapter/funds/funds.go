package funds

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/republicprotocol/renex-sdk-go/adapter/bindings"
	"github.com/republicprotocol/renex-sdk-go/adapter/trader"
	"github.com/republicprotocol/renex-sdk-go/adapter/client"
)

type adapter struct {
	renExBalancesContract *bindings.RenExBalances
	renExTokensContract   *bindings.RenExTokens

	trader      trader.Trader
	httpAddress string
}

type Adapter interface {
	Deposit()

	RequestWithdrawalSignature(tokenCode uint32, value *big.Int)
	RequestWithdrawalWithSignature(tokenCode uint32, value *big.Int, signature []byte)
	RequestWithdrawalFailSafe(tokenCode uint32, value *big.Int)
	RequestWithdrawalFailSafeTrigger()
}

func NewAdapter(httpAddress string, client client.Client, trader trader.Trader) Adapter {
	renExBalances, err := bindings.NewRenExBalances(client.(), client.Client())
	if err != nil {
		return nil, err
	}
	return &adapter{
		renExBalancesContract: renExBalances,
		httpAddress:           httpAddress,
		trader:                trader,
	}
}

func (adapter *adapter) RequestWithdrawalWithSignature(tokenCode uint32, value *big.Int, signature []byte) error {
	address, err := adapter.renExTokensContract.TokenAddresses(&bind.CallOpts{}, tokenCode)
	if err != nil {
		return err
	}

	auth, err := adapter.trader.TransactOpts()
	if err != nil {
		return err
	}

	tx, err := adapter.renExBalancesContract.Withdraw(auth, address, value, signature)
	if err != nil {
		return err
	}

	_, err := adapter.client.PatchWaitMined(tx, context.Background())
	return err
}

func (adapter *adapter) RequestWithdrawalFailSafeTrigger() error {
	address, err := adapter.renExTokensContract.TokenAddresses(&bind.CallOpts{}, tokenCode)
	if err != nil {
		return err
	}

	auth, err := adapter.trader.TransactOpts()
	if err != nil {
		return err
	}

	tx, err := adapter.renExBalancesContract.SignalBackupWithdraw(auth, address)
	if err != nil {
		return err
	}

	_, err := adapter.client.PatchWaitMined(tx, context.Background())
	return err
}

func (adapter *adapter) RequestWithdrawalFailSafe(tokenCode uint32, value *big.Int) error {
	adapter.renExBalancesContract.
		adapter.renExBalancesContract.TraderWithdrawalSignals()

	address, err := adapter.renExTokensContract.TokenAddresses(&bind.CallOpts{}, tokenCode)
	if err != nil {
		return err
	}

	auth, err := adapter.trader.TransactOpts()
	if err != nil {
		return err
	}

	tx, err := adapter.renExBalancesContract.Withdraw(auth, address, value, []byte{})
	if err != nil {
		return err
	}

	_, err := adapter.client.PatchWaitMined(tx, context.Background())
	return err
}
