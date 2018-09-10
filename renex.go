package renex

import (
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
)

type RenEx struct {
	orderbook.Orderbook
	funds.Funds
}

func NewRenEx(network, keystorePath, passphrase string) (RenEx, error) {
	ingressAddress := fmt.Sprintf("https://renex-ingress-%s.herokuapp.com", network)
	newTrader, err := trader.NewTrader(keystorePath, passphrase)
	if err != nil {
		return RenEx{}, err
	}

	newClient, err := client.NewClient(network)
	if err != nil {
		return RenEx{}, err
	}

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

	fAdapter, err := fundsAdapter.NewAdapter(ingressAddress, newClient, newTrader, newStore)
	if err != nil {
		return RenEx{}, err
	}

	fService := funds.NewService(fAdapter)

	oAdapter, err := obAdapter.NewAdapter(ingressAddress, newClient, newTrader, fService, newStore, network)
	if err != nil {
		return RenEx{}, err
	}

	return RenEx{
		orderbook.NewService(oAdapter),
		fService,
	}, nil
}
