package renex

import (
	"github.com/republicprotocol/renex-sdk-go/adapter/client"
	fundsAdapter "github.com/republicprotocol/renex-sdk-go/adapter/funds"
	obAdapter "github.com/republicprotocol/renex-sdk-go/adapter/orderbook"
	"github.com/republicprotocol/renex-sdk-go/adapter/trader"

	"github.com/republicprotocol/renex-sdk-go/core/funds"
	"github.com/republicprotocol/renex-sdk-go/core/orderbook"
)

type RenEx struct {
	orderbook.Orderbook
	funds.Funds
}

func NewRenEx(ingressAddress, configPath, keystorePath, passphrase string) (RenEx, error) {
	newTrader, err := trader.NewTrader(keystorePath, passphrase)
	if err != nil {
		return RenEx{}, err
	}

	newClient, err := client.NewClient(configPath)
	if err != nil {
		return RenEx{}, err
	}

	fAdapter, err := fundsAdapter.NewAdapter(ingressAddress, newClient, newTrader)
	if err != nil {
		return RenEx{}, err
	}

	oAdapter, err := obAdapter.NewAdapter(ingressAddress, newClient, newTrader)
	if err != nil {
		return RenEx{}, err
	}

	return RenEx{
		orderbook.NewService(oAdapter),
		funds.NewService(fAdapter),
	}, nil
}
