package ethclient

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
	RenExBalancesAddress() common.Address
	RenExTokensAddress() common.Address
}

type Network struct {
	URL                     string `json:"url"`
	Chain                   string `json:"chain"`
	OrderbookAddress        string `json:"orderbook"`
	DarknodeRegistryAddress string `json:"darknodeRegistry"`
	RenExBalancesAddress    string `json:"renExBalances"`
	RenExTokensAddress      string `json:"renExTokens"`
}

type client struct {
	network          string
	client           *ethclient.Client
	orderbook        common.Address
	darknodeRegistry common.Address
	renExBalances    common.Address
	renExTokens      common.Address
}

// NewClient creates a new ethereum client.
func NewClient(network Network) (Client, error) {
	ethclient, err := ethclient.Dial(network.URL)
	if err != nil {
		return client{}, err
	}

	return &client{
		client:           ethclient,
		network:          network.Chain,
		orderbook:        common.HexToAddress(network.OrderbookAddress),
		darknodeRegistry: common.HexToAddress(network.DarknodeRegistryAddress),
		renExBalances:    common.HexToAddress(network.RenExBalancesAddress),
		renExTokens:      common.HexToAddress(network.RenExTokensAddress),
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

	return accountAddress, accountAuth, b.Transfer(accountAddress, from, value)
}

// Transfer is a helper function for sending ETH to an address
func (b *client) Transfer(to common.Address, from *bind.TransactOpts, value int64) error {
	transactor := &bind.TransactOpts{
		From:     from.From,
		Nonce:    from.Nonce,
		Signer:   from.Signer,
		Value:    big.NewInt(value),
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
	_, err = b.PatchedWaitMined(context.Background(), tx)
	return err
}

// PatchedWaitMined waits for tx to be mined on the blockchain.
// It stops waiting when the context is canceled.
//
// TODO: THIS DOES NOT WORK WITH PARITY, WHICH SENDS A TRANSACTION RECEIPT UPON
// RECEIVING A TX, NOT AFTER IT'S MINED
func (b *client) PatchedWaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
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

func (client *client) RenExTokensAddress() common.Address {
	return client.renExTokens
}

func (client *client) Network() string {
	return client.network
}

func (client *client) Client() *ethclient.Client {
	return client.client
}
