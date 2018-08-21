package wallet

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type wallet struct {
	passphrase string
}

type Output interface {
	GetTransactOpts(passphrase string) (*bind.TransactOpts, error)
}

type Input interface {
	Deposit(Token common.Address, value *big.Int)
	Withdraw(Token common.Address, value *big.Int)
}
