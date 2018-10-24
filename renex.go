package renex

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/republicprotocol/renex-sdk-go/adapter/client"
	fundsAdapter "github.com/republicprotocol/renex-sdk-go/adapter/funds"
	"github.com/republicprotocol/renex-sdk-go/adapter/leveldb"
	obAdapter "github.com/republicprotocol/renex-sdk-go/adapter/orderbook"
	"github.com/republicprotocol/renex-sdk-go/adapter/store"
	"github.com/republicprotocol/renex-sdk-go/adapter/trader"
	"github.com/republicprotocol/renex-sdk-go/core/funds"
	"github.com/republicprotocol/renex-sdk-go/core/orderbook"
	"github.com/republicprotocol/republic-go/contract"
)

type RenEx struct {
	orderbook.Orderbook
	funds.Funds
	client client.Client
}

func NewRenEx(network, keystorePath, passphrase string) (RenEx, error) {
	newTrader, err := trader.NewTrader(keystorePath, passphrase)
	if err != nil {
		return RenEx{}, err
	}
	newClient, err := client.NewClient(network)
	if err != nil {
		return RenEx{}, err
	}
	conn, err := contract.Connect(contract.Config{Network: contract.Network(network)})
	if err != nil {
		return RenEx{}, err
	}
	return newRenEx(network, newTrader, newClient, conn)
}

func NewRenExWithPrivKey(network string, privKey *ecdsa.PrivateKey) (RenEx, error) {
	newClient, err := client.NewClient(network)
	if err != nil {
		return RenEx{}, err
	}
	conn, err := contract.Connect(contract.Config{Network: contract.Network(network)})
	if err != nil {
		return RenEx{}, err
	}
	return newRenEx(network, trader.NewTraderFromPrivateKey(privKey), newClient, conn)
}

func NewRenExWithNetwork(network string, clientNetwork client.Network, privKey *ecdsa.PrivateKey, conn contract.Conn) (RenEx, error) {
	newTrader := trader.NewTraderFromPrivateKey(privKey)
	newClient, err := client.NewClientFromNetwork(clientNetwork)
	if err != nil {
		return RenEx{}, err
	}
	return newRenEx(network, newTrader, newClient, conn)
}

func newRenEx(network string, t trader.Trader, c client.Client, conn contract.Conn) (RenEx, error) {
	ingressAddress := fmt.Sprintf("https://renex-ingress-%s.herokuapp.com", network)

	randomDBID := make([]byte, 32)
	if _, err := rand.Read(randomDBID); err != nil {
		return RenEx{}, err
	}

	newStoreAdapter, err := leveldb.NewLDBStore(os.Getenv("HOME") + fmt.Sprintf("/.renex/db%s", hex.EncodeToString(randomDBID)))
	if err != nil {
		return RenEx{}, err
	}

	newStore := store.NewStore(newStoreAdapter)
	if err != nil {
		return RenEx{}, err
	}

	fAdapter, err := fundsAdapter.NewAdapter(ingressAddress, c, t, newStore)
	if err != nil {
		return RenEx{}, err
	}

	fService := funds.NewService(fAdapter)

	oAdapter, err := obAdapter.NewAdapterFromConn(ingressAddress, c, t, fService, newStore, network, conn)
	if err != nil {
		return RenEx{}, err
	}

	return RenEx{
		orderbook.NewService(oAdapter),
		fService,
		c,
	}, nil
}
