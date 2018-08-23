package trader

import (
	"crypto/rand"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

type trader struct {
	*keystore.Key
}

// Trader represents an individual entity that opens orders.
// Traders can be either honest participants, malicious
// or even participants with insufficient funds.
type Trader interface {
	Sign([]byte) ([]byte, error)
	TransactOpts() *bind.TransactOpts
	Address() common.Address
}

func NewTrader(path string, passphrase string) (Trader, error) {
	keyin, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	json, err := ioutil.ReadAll(keyin)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(json, passphrase)
	if err != nil {
		return nil, err
	}
	return &trader{
		key,
	}, nil
}

func (t *trader) TransactOpts() *bind.TransactOpts {
	return bind.NewKeyedTransactor(t.PrivateKey)
}

func (t *trader) Address() common.Address {
	return t.Key.Address
}

func (t *trader) Sign(data []byte) ([]byte, error) {
	return t.PrivateKey.Sign(rand.Reader, data, nil)
}
