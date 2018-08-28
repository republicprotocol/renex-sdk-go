package trader

import (
	"crypto/ecdsa"
	"crypto/rand"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/republicprotocol/republic-go/crypto"
)

type trader struct {
	*ecdsa.PrivateKey
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
	ks := crypto.Keystore{}
	keyin, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	json, err := ioutil.ReadAll(keyin)
	if err != nil {
		return nil, err
	}
	if err := ks.DecryptFromJSON(json, passphrase); err != nil {
		key, err := keystore.DecryptKey(json, passphrase)
		if err != nil {
			return nil, err
		}
		return &trader{
			key.PrivateKey,
		}, nil
	}
	return &trader{
		ks.EcdsaKey.PrivateKey,
	}, nil
}

func (t *trader) TransactOpts() *bind.TransactOpts {
	return bind.NewKeyedTransactor(t.PrivateKey)
}

func (t *trader) Address() common.Address {
	return bind.NewKeyedTransactor(t.PrivateKey).From
}

func (t *trader) Sign(data []byte) ([]byte, error) {
	return t.PrivateKey.Sign(rand.Reader, data, nil)
}
