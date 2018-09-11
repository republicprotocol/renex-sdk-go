package funds

import (
	"fmt"
	"math/big"

	"github.com/republicprotocol/republic-go/order"
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
	TransferEth(address string, value *big.Int) error
	TransferERC20(address string, token order.Token, value *big.Int) error
	RequestLockedBalance(tokenCode order.Token) (*big.Int, error)
	RequestDeposit(tokenCode order.Token, value *big.Int) error
	RequestWithdrawalSignature(tokenCode order.Token, value *big.Int) ([]byte, error)
	RequestWithdrawalWithSignature(tokenCode order.Token, value *big.Int, signature []byte) error
	RequestWithdrawalFailSafe(tokenCode order.Token, value *big.Int) error
	RequestWithdrawalFailSafeTrigger(tokenCode order.Token) (*IdempotentKey, error)
	OpenOrdersExist(tokenCode order.Token) (bool, error)
	CheckStatus(key *IdempotentKey) uint8
}

type Funds interface {
	Address() string
	Transfer(address string, token order.Token, value *big.Int) error
	Balance(token order.Token) (*big.Int, error)
	RenExBalance(token order.Token) (*big.Int, error)
	UsableRenExBalance(token order.Token) (*big.Int, error)
	Deposit(token order.Token, value *big.Int) error
	Withdraw(token order.Token, value *big.Int, forced bool, key *IdempotentKey) (*IdempotentKey, error)
}

func NewService(adapter Adapter) Funds {
	return &service{
		adapter,
	}
}

func (service *service) Withdraw(token order.Token, value *big.Int, forced bool, key *IdempotentKey) (*IdempotentKey, error) {
	switch service.CheckStatus(key) {
	case 0:
		sig, err := service.RequestWithdrawalSignature(token, value)
		if err != nil {
			if !forced {
				return nil, err
			}
			failSafeKey, err := service.RequestWithdrawalFailSafeTrigger(token)
			if err != nil {
				return nil, err
			}
			return failSafeKey, nil
		}
		return nil, service.RequestWithdrawalWithSignature(token, value, sig)
	case 1:
		return key, fmt.Errorf("Transaction Pending")
	case 2:
		return key, fmt.Errorf("Time left to trigger withdrawal without signature")
	case 3:
		return key, service.RequestWithdrawalFailSafe(token, value)
	default:
		return key, fmt.Errorf("Corrupted Idempotent Key")
	}
}

func (service *service) Deposit(tokenCode order.Token, value *big.Int) error {
	return service.RequestDeposit(tokenCode, value)
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

func (service *service) Transfer(address string, tokenCode order.Token, value *big.Int) error {
	switch tokenCode {
	case order.TokenREN:
		return service.TransferERC20(address, tokenCode, value)
	case order.TokenDGX:
		return service.TransferERC20(address, tokenCode, value)
	case order.TokenABC:
		return service.TransferERC20(address, tokenCode, value)
	case order.TokenPQR:
		return service.TransferERC20(address, tokenCode, value)
	case order.TokenXYZ:
		return service.TransferERC20(address, tokenCode, value)
	case order.TokenETH:
		return service.TransferEth(address, value)
	default:
		return fmt.Errorf("Unsupported Currency")
	}
}

func (service *service) Balance(tokenCode order.Token) (*big.Int, error) {
	switch tokenCode {
	case order.TokenREN:
		return service.BalanceErc20(tokenCode)
	case order.TokenDGX:
		return service.BalanceErc20(tokenCode)
	case order.TokenABC:
		return service.BalanceErc20(tokenCode)
	case order.TokenPQR:
		return service.BalanceErc20(tokenCode)
	case order.TokenXYZ:
		return service.BalanceErc20(tokenCode)
	case order.TokenETH:
		return service.BalanceEth()
	default:
		return nil, fmt.Errorf("Unsupported Currency")
	}
}
