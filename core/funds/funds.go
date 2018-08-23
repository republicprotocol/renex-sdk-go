package funds

import (
	"fmt"
	"math/big"
)

type IdempotentKey [32]byte

type service struct {
	Adapter
}

type Adapter interface {
	RequestDeposit(tokenCode uint32, value *big.Int) error
	RequestWithdrawalSignature(tokenCode uint32, value *big.Int) ([]byte, error)
	RequestWithdrawalWithSignature(tokenCode uint32, value *big.Int, signature []byte) error
	RequestWithdrawalFailSafe(tokenCode uint32, value *big.Int) error
	RequestWithdrawalFailSafeTrigger(tokenCode uint32) (*IdempotentKey, error)
	CheckStatus(key *IdempotentKey) uint8
}

type Funds interface {
	Deposit(token uint32, value *big.Int) error
	Withdraw(token uint32, value *big.Int, forced bool, key *IdempotentKey) (*IdempotentKey, error)
}

func NewService(adapter Adapter) Funds {
	return &service{
		adapter,
	}
}

func (service *service) Withdraw(token uint32, value *big.Int, forced bool, key *IdempotentKey) (*IdempotentKey, error) {
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

func (service *service) Deposit(tokenCode uint32, value *big.Int) error {
	return service.RequestDeposit(tokenCode, value)
}
