package trader

import (
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type trader struct {
	path string
}

// Trader represents an individual entity that opens orders.
// Traders can be either honest participants, malicious
// or even participants with insufficient funds.
type Trader interface {
	keystore crypto.Keystore
}

func New(path string) Trader {
	return &trader{
		path: path,
	}
}

func (t *trader) GetTransactOpts(passphrase string) (*bind.TransactOpts, error) {
	keyin, err := os.Open(t.path)
	defer keyin.Close()
	if err != nil {
		return nil, err
	}
	return bind.NewTransactor(keyin, passphrase)
}
