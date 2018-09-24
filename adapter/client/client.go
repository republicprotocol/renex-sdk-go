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
	case "mainnet":
		return Network{
			URL:   "https://mainnet.infura.io",
			Chain: "mainnet",
			DarknodeRegistryAddress: "0x3799006a87fde3ccfc7666b3e6553b03ed341c2f",
			OrderbookAddress:        "0x6b8bb175c092de7d81860b18db360b734a2598e0",
			RenExBalancesAddress:    "0x9636f9ac371ca0965b7c2b4ad13c4cc64d0ff2dc",
			RenExSettlementAddress:  "0x908262de0366e42d029b0518d5276762c92b21e1",
			RenExTokensAddress:      "0x7cade4fbc8761817bb62a080733d1b6cad744ec4",
		}, nil
	case "testnet":
		return Network{
			URL:   "https://kovan.infura.io",
			Chain: "kovan",
			DarknodeRegistryAddress: "0xf7daA0Baf257547A6Ad3CE7FFF71D55cb7426F76",
			OrderbookAddress:        "0xA53Da4093c682a4259DE38302341BFEf7e9f7a4f",
			RenExBalancesAddress:    "0x97073d0d654ebb71dd9efd1dfa777c73f56d4021",
			RenExSettlementAddress:  "0x68FE2088A321A42DE11Aba93D32C81C9f20b1Abe",
			RenExTokensAddress:      "0xedFF6E7C072fA0018720734F6d5a4f4DC30f9869",
		}, nil
	case "falcon":
		return Network{
			URL:   "https://kovan.infura.io",
			Chain: "kovan",
			DarknodeRegistryAddress: "0xDaA8C30AF85070506F641E456aFDB84d4bA972Bd",
			OrderbookAddress:        "0x592d16f8C5FA8f1E074ab3C2cd1ACD087ADcdc0B",
			RenExBalancesAddress:    "0xb3E632943fA995FC75692e46b62383BE49cDdbc4",
			RenExSettlementAddress:  "0xBE936cb23DD9a84E4D9358810f7F275e93CCD770",
			RenExTokensAddress:      "0x9a898c8148131eF189B1c8575692376403780325",
		}, nil
	case "nightly":
		return Network{
			URL:   "https://kovan.infura.io",
			Chain: "kovan",
			DarknodeRegistryAddress: "0x8a31d477267A5af1bc5142904ef0AfA31D326E03",
			OrderbookAddress:        "0x376127aDc18260fc238eBFB6626b2F4B59eC9b66",
			RenExBalancesAddress:    "0xa95dE870dDFB6188519D5CC63CEd5E0FBac1aa8E",
			RenExSettlementAddress:  "0x5f25233ca99104D31612D4fB937B090d5A2EbB75",
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
		if b.client == nil {
			return nil, fmt.Errorf("Nil Client")
		}
		if tx == nil {
			return nil, fmt.Errorf("Nil Tx")
		}
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
