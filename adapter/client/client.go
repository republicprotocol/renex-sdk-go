package client

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client interface {
	Network() string
	Client() *ethclient.Client
	OrderbookAddress() common.Address
	DarknodeRegistryAddress() common.Address
	RenExSettlementAddress() common.Address
	RenExBalancesAddress() common.Address
	RenExTokensAddress() common.Address
	WaitTillMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error)
	Transfer(to common.Address, from *bind.TransactOpts, value *big.Int) error
}

type Network struct {
	URL                     string `json:"url"`
	Chain                   string `json:"chain"`
	OrderbookAddress        string `json:"orderbook"`
	DarknodeRegistryAddress string `json:"darknodeRegistry"`
	RenExBalancesAddress    string `json:"renExBalances"`
	RenExTokensAddress      string `json:"renExTokens"`
	RenExSettlementAddress  string `json:"renExSettlement"`
	RenExAtomicInfoAddress  string `json:"renExAtomicInfo"`
}

type client struct {
	network          string
	client           *ethclient.Client
	orderbook        common.Address
	darknodeRegistry common.Address
	renExBalances    common.Address
	renExTokens      common.Address
	renExSettlement  common.Address
}

func GetNetwork(network string) (Network, error) {
	switch network {
	case "mainnet":
		return Network{
			URL:                     "https://mainnet.infura.io",
			Chain:                   "mainnet",
			DarknodeRegistryAddress: "0x3799006a87fde3ccfc7666b3e6553b03ed341c2f",
			OrderbookAddress:        "0x6b8bb175c092de7d81860b18db360b734a2598e0",
			RenExBalancesAddress:    "0x5eC18B477B20aF940807B5478dB5A64Cd4a77EFd",
			RenExSettlementAddress:  "0x908262dE0366E42d029B0518D5276762c92B21e1",
			RenExTokensAddress:      "0x7cade4fbc8761817bb62a080733d1b6cad744ec4",
		}, nil
	case "testnet":
		return Network{
			URL:                     "https://kovan.infura.io",
			Chain:                   "kovan",
			DarknodeRegistryAddress: "0x75Fa8349fc9C7C640A4e9F1A1496fBB95D2Dc3d5",
			OrderbookAddress:        "0xA9b453FC64b4766Aab8a867801d0a4eA7b1474E0",
			RenExBalancesAddress:    "0xb0E21B869D6f741a8A8F5075BA59E496593B881A",
			RenExSettlementAddress:  "0x65A699E555cf93e4e115FfC2DE2F41d5A21DF3Fb",
			RenExTokensAddress:      "0x481b39E2000a117CBA417473DC1E7cdAf4EAd98F",
		}, nil
	default:
		return Network{}, fmt.Errorf("unknown network %s", network)
	}
}

// NewClient creates a new ethereum client.
func NewClient(net string) (Client, error) {
	network, err := GetNetwork(net)
	if err != nil {
		return nil, err
	}

	ethclient, err := ethclient.Dial(network.URL)
	if err != nil {
		return nil, err
	}

	return &client{
		client:           ethclient,
		network:          network.Chain,
		orderbook:        common.HexToAddress(network.OrderbookAddress),
		darknodeRegistry: common.HexToAddress(network.DarknodeRegistryAddress),
		renExBalances:    common.HexToAddress(network.RenExBalancesAddress),
		renExTokens:      common.HexToAddress(network.RenExTokensAddress),
		renExSettlement:  common.HexToAddress(network.RenExSettlementAddress),
	}, nil
}

func NewClientFromNetwork(network Network) (Client, error) {
	ethclient, err := ethclient.Dial(network.URL)
	if err != nil {
		return nil, err
	}

	return &client{
		client:           ethclient,
		network:          network.Chain,
		orderbook:        common.HexToAddress(network.OrderbookAddress),
		darknodeRegistry: common.HexToAddress(network.DarknodeRegistryAddress),
		renExBalances:    common.HexToAddress(network.RenExBalancesAddress),
		renExTokens:      common.HexToAddress(network.RenExTokensAddress),
		renExSettlement:  common.HexToAddress(network.RenExSettlementAddress),
	}, nil
}

// NewAccount creates a new account and funds it with ether
func (b *client) NewAccount(value int64, from *bind.TransactOpts) (common.Address, *bind.TransactOpts, error) {
	account, err := crypto.GenerateKey()
	if err != nil {
		return common.Address{}, &bind.TransactOpts{}, err
	}

	accountAddress := crypto.PubkeyToAddress(account.PublicKey)
	accountAuth := bind.NewKeyedTransactor(account)

	return accountAddress, accountAuth, b.Transfer(accountAddress, from, big.NewInt(value))
}

// Transfer is a helper function for sending ETH to an address
func (b *client) Transfer(to common.Address, from *bind.TransactOpts, value *big.Int) error {
	transactor := &bind.TransactOpts{
		From:     from.From,
		Nonce:    from.Nonce,
		Signer:   from.Signer,
		Value:    value,
		GasPrice: from.GasPrice,
		GasLimit: 30000,
		Context:  from.Context,
	}

	nonce, nonceErr := b.Client().PendingNonceAt(context.Background(), transactor.From)
	if nonceErr != nil {
		return fmt.Errorf("cannot get nonce: %v", nonceErr)
	}
	transactor.Nonce = big.NewInt(int64(nonce))

	// Why is there no ethclient.Transfer?
	bound := bind.NewBoundContract(to, abi.ABI{}, nil, b.client, nil)
	tx, err := bound.Transfer(transactor)
	if err != nil {
		for try := 0; try < 60 && strings.Contains(err.Error(), "nonce"); try++ {
			time.Sleep(time.Second)
			nonce, nonceErr = b.Client().PendingNonceAt(context.Background(), transactor.From)
			if nonceErr != nil {
				continue
			}
			transactor.Nonce = big.NewInt(int64(nonce))
			tx, err = bound.Transfer(transactor)
			if err == nil {
				break
			}
		}
		if err != nil {
			return err
		}
	}
	from.Nonce = big.NewInt(0).Add(transactor.Nonce, big.NewInt(1))

	_, err = b.WaitTillMined(context.Background(), tx)
	return err
}

// WaitTillMined waits for tx to be mined on the blockchain.
// It stops waiting when the context is canceled.
//
// TODO: THIS DOES NOT WORK WITH PARITY, WHICH SENDS A TRANSACTION RECEIPT UPON
// RECEIVING A TX, NOT AFTER IT'S MINED
func (b *client) WaitTillMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	switch b.network {
	case "ganache":
		time.Sleep(100 * time.Millisecond)
		return nil, nil
	default:
		receipt, err := bind.WaitMined(ctx, b.client, tx)
		if err != nil {
			return nil, err
		}
		// if receipt.Status != types.ReceiptStatusSuccessful {
		// 	return receipt, errors.New("transaction reverted")
		// }
		return receipt, nil
	}
}

// PatchedWaitDeployed waits for a contract deployment transaction and returns the on-chain
// contract address when it is mined. It stops waiting when ctx is canceled.
//
// TODO: THIS DOES NOT WORK WITH PARITY, WHICH SENDS A TRANSACTION RECEIPT UPON
// RECEIVING A TX, NOT AFTER IT'S MINED
func (b *client) PatchedWaitDeployed(ctx context.Context, tx *types.Transaction) (common.Address, error) {
	switch b.network {
	case "ganache":
		time.Sleep(100 * time.Millisecond)
		return common.Address{}, nil
	default:
		return bind.WaitDeployed(ctx, b.client, tx)
	}
}

func (client *client) OrderbookAddress() common.Address {
	return client.orderbook
}

func (client *client) DarknodeRegistryAddress() common.Address {
	return client.darknodeRegistry
}

func (client *client) RenExBalancesAddress() common.Address {
	return client.renExBalances
}
func (client *client) RenExSettlementAddress() common.Address {
	return client.renExSettlement
}

func (client *client) RenExTokensAddress() common.Address {
	return client.renExTokens
}

func (client *client) Network() string {
	return client.network
}

func (client *client) Client() *ethclient.Client {
	return client.client
}
