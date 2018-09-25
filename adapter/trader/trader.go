package trader

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/republicprotocol/renex-sdk-go/adapter/client"
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
	SendTx(f func() (client.Client, *types.Transaction, error)) (*types.Transaction, error)
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

func (t *trader) SendTx(f func() (client.Client, *types.Transaction, error)) (*types.Transaction, error) {
	t.Lock()
	defer t.Unlock()
	return t.sendTx(f)
}

func (t *trader) sendTx(f func() (client.Client, *types.Transaction, error)) (*types.Transaction, error) {
	opts := t.transactOpts
	opts.GasLimit = 3000000
	client, tx, err := f()

	nonce, nonceErr := client.Client().PendingNonceAt(context.Background(), t.address)
	if nonceErr != nil {
		return tx, nonceErr
	}

	opts.Nonce = big.NewInt(int64(nonce))

	if err == nil {
		opts.Nonce.Add(opts.Nonce, big.NewInt(1))
		return tx, nil
	}

	if err == core.ErrNonceTooLow || err == core.ErrReplaceUnderpriced || strings.Contains(err.Error(), "nonce is too low") {
		opts.Nonce.Add(opts.Nonce, big.NewInt(1))
		return t.sendTx(f)
	}

	if err == core.ErrNonceTooHigh {
		opts.Nonce.Sub(opts.Nonce, big.NewInt(1))
		return t.sendTx(f)
	}

	// If any other type of nonce error occurs we will refresh the nonce and
	// try again for up to 1 minute
	for try := 0; try < 60 && strings.Contains(err.Error(), "nonce"); try++ {
		time.Sleep(time.Second)
		nonce, nonceErr = client.Client().PendingNonceAt(context.Background(), opts.From)
		if nonceErr != nil {
			continue
		}
		opts.Nonce = big.NewInt(int64(nonce))
		if _, tx, err = f(); err == nil {
			opts.Nonce.Add(opts.Nonce, big.NewInt(1))
			return tx, nil
		}
	}
	t.transactOpts.Nonce = opts.Nonce

	return tx, err
}
