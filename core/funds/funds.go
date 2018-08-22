package funds

import (
	"math/big"

	"github.com/republicprotocol/renex-sdk-go/adapter/funds"
)

type service struct {
	funds.Adapter
}

type Service interface {
	Deposit(token uint32, value *big.Int)
	Withdraw(token uint32, value *big.Int, forced bool, key *IdempotentKey)
}

func (service *service) Withdraw(token uint32, value *big.Int, forced bool, key *IdempotentKey) error {
	sig, err := service.RequestWithdrawalSignature(token, value);
	if err != nil {
		if !forced {
			return err
		}
		if err := service.RequestWithdrawalFailSafeTrigger(); err != nil {
			return err
		}
		if err := service.RequestWithdrawalFailSafe(tokenCode uint32, value *big.Int); err != nil {
			return err
		}
	}
}
