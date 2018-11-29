package funds

import (
	"context"
	"fmt"
	"math/big"

	"github.com/republicprotocol/republicprotocol-go/foundation/order"
)

type IdempotentKey [32]byte

type service struct {
	Adapter
}

type Adapter interface {
	Address() string
	BalanceEth() (*big.Int, error)
	BalanceErc20(tokenCode order.Token) (*big.Int, error)
	RenExBalance(tokenCode order.Token) (*big.Int, error)
	TransferEth(ctx context.Context, address string, value *big.Int) error
	TransferERC20(ctx context.Context, address string, token order.Token, value *big.Int) error
	RequestLockedBalance(tokenCode order.Token) (*big.Int, error)
	RequestDeposit(ctx context.Context, tokenCode order.Token, value *big.Int) error
	RequestWithdrawalSignature(tokenCode order.Token, value *big.Int) ([]byte, error)
	RequestWithdrawalWithSignature(ctx context.Context, tokenCode order.Token, value *big.Int, signature []byte) error
	RequestWithdrawalFailSafe(ctx context.Context, tokenCode order.Token, value *big.Int) error
	RequestWithdrawalFailSafeTrigger(ctx context.Context, tokenCode order.Token) (*IdempotentKey, error)
	OpenOrdersExist(tokenCode order.Token) (bool, error)
	CheckStatus(key *IdempotentKey) uint8
}

type Funds interface {
	Address() string
	Transfer(ctx context.Context, address string, token order.Token, value *big.Int) error
	Balance(token order.Token) (*big.Int, error)
	RenExBalance(token order.Token) (*big.Int, error)
	UsableRenExBalance(token order.Token) (*big.Int, error)
	Deposit(ctx context.Context, token order.Token, value *big.Int) error
	Withdraw(ctx context.Context, token order.Token, value *big.Int, forced bool, key *IdempotentKey) (*IdempotentKey, error)
}

func NewService(adapter Adapter) Funds {
	return &service{
		adapter,
	}
}

func (service *service) Withdraw(ctx context.Context, token order.Token, value *big.Int, forced bool, key *IdempotentKey) (*IdempotentKey, error) {
	switch service.CheckStatus(key) {
	case 0:
		sig, err := service.RequestWithdrawalSignature(token, value)
		if err != nil {
			if !forced {
				return nil, err
			}
			failSafeKey, err := service.RequestWithdrawalFailSafeTrigger(ctx, token)
			if err != nil {
				return nil, err
			}
			return failSafeKey, nil
		}
		return nil, service.RequestWithdrawalWithSignature(ctx, token, value, sig)
	case 1:
		return key, fmt.Errorf("Transaction Pending")
	case 2:
		return key, fmt.Errorf("Time left to trigger withdrawal without signature")
	case 3:
		return key, service.RequestWithdrawalFailSafe(ctx, token, value)
	default:
		return key, fmt.Errorf("Corrupted Idempotent Key")
	}
}

func (service *service) Deposit(ctx context.Context, tokenCode order.Token, value *big.Int) error {
	return service.RequestDeposit(ctx, tokenCode, value)
}

func (service *service) UsableRenExBalance(tokenCode order.Token) (*big.Int, error) {
	balance, err := service.RenExBalance(tokenCode)
	if err != nil {
		return nil, err
	}

	lockedBalance, err := service.RequestLockedBalance(tokenCode)
	if err != nil {
		return nil, err
	}

	return balance.Sub(balance, lockedBalance), nil
}

func (service *service) Transfer(ctx context.Context, address string, tokenCode order.Token, value *big.Int) error {
	switch tokenCode {
	case order.TokenREN:
		return service.TransferERC20(ctx, address, tokenCode, value)
	case order.TokenDGX:
		return service.TransferERC20(ctx, address, tokenCode, value)
	case order.TokenTUSD:
		return service.TransferERC20(ctx, address, tokenCode, value)
	case order.TokenZRX:
		return service.TransferERC20(ctx, address, tokenCode, value)
	case order.TokenOMG:
		return service.TransferERC20(ctx, address, tokenCode, value)
	case order.TokenETH:
		return service.TransferEth(ctx, address, value)
	default:
		return fmt.Errorf("unsupported currency")
	}
}

func (service *service) Balance(tokenCode order.Token) (*big.Int, error) {
	switch tokenCode {
	case order.TokenREN:
		return service.BalanceErc20(tokenCode)
	case order.TokenDGX:
		return service.BalanceErc20(tokenCode)
	case order.TokenTUSD:
		return service.BalanceErc20(tokenCode)
	case order.TokenZRX:
		return service.BalanceErc20(tokenCode)
	case order.TokenOMG:
		return service.BalanceErc20(tokenCode)
	case order.TokenETH:
		return service.BalanceEth()
	default:
		return nil, fmt.Errorf("unsupported currency")
	}
}
