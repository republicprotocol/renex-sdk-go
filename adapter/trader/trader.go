package trader

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/republicprotocol/republic-go/crypto"
)

type trader struct {
	transactOpts *bind.TransactOpts
	address      common.Address
	*ecdsa.PrivateKey
	*sync.RWMutex
}

// Trader represents an individual entity that opens orders.
// Traders can be either honest participants, malicious
// or even participants with insufficient funds.
type Trader interface {
	Sign([]byte) ([]byte, error)
	SendTx(f func() (*types.Transaction, error)) (*types.Transaction, error)
	TransactOpts() *bind.TransactOpts
	Address() common.Address
}

func NewTrader(path string, passphrase string) (Trader, error) {
	ks := crypto.Keystore{}

	fmt.Println("Location ", path)
	keyin, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fmt.Println("Sucessfully opened the file")

	json, err := ioutil.ReadAll(keyin)
	if err != nil {
		return nil, err
	}

	fmt.Println("Decrypting the json file")
	if err := ks.DecryptFromJSON(json, passphrase); err != nil {
		key, err := keystore.DecryptKey(json, passphrase)
		if err != nil {
			return nil, err
		}
		transactOpts := bind.NewKeyedTransactor(key.PrivateKey)
		return &trader{
			transactOpts: transactOpts,
			address:      transactOpts.From,
			PrivateKey:   key.PrivateKey,
			RWMutex:      new(sync.RWMutex),
		}, nil
	}

	transactOpts := bind.NewKeyedTransactor(ks.EcdsaKey.PrivateKey)
	return &trader{
		transactOpts: transactOpts,
		address:      transactOpts.From,
		PrivateKey:   ks.EcdsaKey.PrivateKey,
		RWMutex:      new(sync.RWMutex),
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

func (t *trader) SendTx(f func() (*types.Transaction, error)) (*types.Transaction, error) {
	t.Lock()
	defer t.Unlock()
	return t.sendTx(f)
}

func (t *trader) sendTx(f func() (*types.Transaction, error)) (*types.Transaction, error) {
	tx, err := f()
	if err == nil {
		t.transactOpts.Nonce.Add(t.transactOpts.Nonce, big.NewInt(1))
		return tx, nil
	}

	if err == core.ErrNonceTooLow || err == core.ErrReplaceUnderpriced || strings.Contains(err.Error(), "nonce is too low") {
		t.transactOpts.Nonce.Add(t.transactOpts.Nonce, big.NewInt(1))
		return t.sendTx(f)
	}

	if err == core.ErrNonceTooHigh {
		t.transactOpts.Nonce.Sub(t.transactOpts.Nonce, big.NewInt(1))
		return t.sendTx(f)
	}
	return tx, err
}
