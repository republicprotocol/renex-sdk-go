package client

import (
	"context"
	"fmt"
	"math/big"
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
	case "testnet":
		return Network{
			URL:   "https://kovan.infura.io",
			Chain: "kovan",
			DarknodeRegistryAddress: "0x372b6204263c6867f81e2a9e11057ff43efea14b",
			OrderbookAddress:        "0xa7caa4780a39d8b8acd6a0bdfb5b906210bc76cd",
			RenExBalancesAddress:    "0xc5b98949AB0dfa0A7d4c07Bb29B002D6d6DA3e25",
			RenExSettlementAddress:  "0xfa0938e3c9a5e33b5084dfbffaca9241aef39be8",
			RenExAtomicInfoAddress:  "0xCf85e1Ee45df67Bc9a65B9707a24b64Ac3174259",
			RenExTokensAddress:      "0x160ECA47935be4139eC5B94D99B678d6f7e18f95",
		}, nil
	case "falcon":
		return Network{
			URL:   "https://kovan.infura.io",
			Chain: "kovan",
			DarknodeRegistryAddress: "0xfafd5c83d1e21763b79418c4ecb5d62b4970df8e",
			OrderbookAddress:        "0x044b08eec761c39ac32aee1d6ef0583812f21699",
			RenExBalancesAddress:    "0x3083e5ba36c6b42ca93c22c803013a4539eedc7f",
			RenExSettlementAddress:  "0x8617dcd709bb8660602ef70ade78626b7408a210",
			RenExAtomicInfoAddress:  "0xafe5539b40b17404e3cb6cf5013fc7dab3c54163",
			RenExTokensAddress:      "0x160ECA47935be4139eC5B94D99B678d6f7e18f95",
		}, nil
	case "nightly":
		return Network{
			URL:   "https://kovan.infura.io",
			Chain: "kovan",
			DarknodeRegistryAddress: "0x8a31d477267A5af1bc5142904ef0AfA31D326E03",
			OrderbookAddress:        "0x376127aDc18260fc238eBFB6626b2F4B59eC9b66",
			RenExBalancesAddress:    "0xa95dE870dDFB6188519D5CC63CEd5E0FBac1aa8E",
			RenExSettlementAddress:  "0x5f25233ca99104D31612D4fB937B090d5A2EbB75",
			RenExAtomicInfoAddress:  "0xe1A660657A32053fe83B19B1177F6B56C6F37b1f",
			RenExTokensAddress:      "0x160ECA47935be4139eC5B94D99B678d6f7e18f95",
		}, nil
	default:
		return Network{}, fmt.Errorf("Unknown Network")
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

	// Why is there no ethclient.Transfer?
	bound := bind.NewBoundContract(to, abi.ABI{}, nil, b.client, nil)
	tx, err := bound.Transfer(transactor)
	if err != nil {
		return err
	}
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
		reciept, err := bind.WaitMined(ctx, b.client, tx)
		if err != nil {
			return nil, err
		}
		if reciept.Status != 1 {
			return nil, fmt.Errorf("Transaction reverted")
		}
		return reciept, nil
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
