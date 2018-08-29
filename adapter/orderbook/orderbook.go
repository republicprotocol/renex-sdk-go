package orderbook

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/republicprotocol/renex-ingress-go/httpadapter"
	"github.com/republicprotocol/renex-sdk-go/adapter/bindings"
	"github.com/republicprotocol/renex-sdk-go/adapter/client"
	"github.com/republicprotocol/renex-sdk-go/adapter/store"
	"github.com/republicprotocol/renex-sdk-go/adapter/trader"
	"github.com/republicprotocol/renex-sdk-go/core/funds"
	"github.com/republicprotocol/renex-sdk-go/core/orderbook"
	"github.com/republicprotocol/republic-go/crypto"
	"github.com/republicprotocol/republic-go/identity"
	"github.com/republicprotocol/republic-go/order"
	"github.com/republicprotocol/republic-go/registry"
)

type adapter struct {
	httpAddress              string
	orderbookContract        *bindings.Orderbook
	darknodeRegistryContract *bindings.DarknodeRegistry
	renexSettlementContract  *bindings.RenExSettlement
	trader                   trader.Trader
	client                   client.Client
	funds                    funds.Funds
	store                    store.Store
}

func NewAdapter(httpAddress string, client client.Client, trader trader.Trader, funds funds.Funds, store store.Store) (orderbook.Adapter, error) {
	orderbook, err := bindings.NewOrderbook(client.OrderbookAddress(), bind.ContractBackend(client.Client()))
	if err != nil {
		return nil, err
	}
	darknodeRegistry, err := bindings.NewDarknodeRegistry(client.DarknodeRegistryAddress(), bind.ContractBackend(client.Client()))
	if err != nil {
		return nil, err
	}
	renexSettlement, err := bindings.NewRenExSettlement(client.RenExSettlementAddress(), bind.ContractBackend(client.Client()))
	if err != nil {
		return nil, err
	}
	return &adapter{
		orderbookContract:        orderbook,
		darknodeRegistryContract: darknodeRegistry,
		renexSettlementContract:  renexSettlement,
		httpAddress:              httpAddress,
		trader:                   trader,
		client:                   client,
		funds:                    funds,
		store:                    store,
	}, nil
}

func (adapter *adapter) RequestOpenOrder(order order.Order) error {
	balance, err := adapter.funds.UsableBalance(getTokenCode(order))
	if balance.Uint64() < order.Volume {
		return fmt.Errorf("Order volume exceeded usable balance have:%d want:%d", balance.Uint64(), order.Volume)
	}

	mapping, err := adapter.buildOrderMapping(order)
	if err != nil {
		return err
	}

	req := httpadapter.OpenOrderRequest{
		Address:               adapter.trader.Address().String(),
		OrderFragmentMappings: []httpadapter.OrderFragmentMapping{mapping},
	}

	data, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(data)

	fmt.Println("Waiting for the post response")
	resp, err := http.DefaultClient.Post(fmt.Sprintf("%s/orders", adapter.httpAddress), "application/json", buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Recieved the post response")

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	type Response struct {
		Signature string `json:"signature"`
	}

	respose := Response{}
	if err := json.Unmarshal(respBytes, respose); err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return fmt.Errorf("Unexpected status code %d", resp.StatusCode)
	}

	sigBytes, err := base64.StdEncoding.DecodeString(respose.Signature)
	if err != nil {
		return err
	}
	tx, err := adapter.trader.SendTx(func() (client.Client, *types.Transaction, error) {
		tx, err := adapter.orderbookContract.OpenOrder(adapter.trader.TransactOpts(), 1, sigBytes, order.ID)
		return adapter.client, tx, err
	})
	if err != nil {
		return err
	}

	if _, err := adapter.client.WaitTillMined(context.Background(), tx); err != nil {
		return err
	}

	return adapter.store.AppendOrder(order)
}

func (adapter *adapter) RequestCancelOrder(orderID order.ID) error {
	tx, err := adapter.trader.SendTx(func() (client.Client, *types.Transaction, error) {
		tx, err := adapter.orderbookContract.CancelOrder(adapter.trader.TransactOpts(), orderID)
		return adapter.client, tx, err
	})

	if err != nil {
		return err
	}

	if _, err := adapter.client.WaitTillMined(context.Background(), tx); err != nil {
		return err
	}
	return adapter.store.DeleteOrder(orderID)
}

func (adapter *adapter) ListOrders() ([]order.ID, []string, []uint8, error) {
	numOrdersBig, err := adapter.orderbookContract.OrdersCount(&bind.CallOpts{})
	if err != nil {
		return nil, nil, nil, err
	}
	numOrders := numOrdersBig.Int64()
	orderIDs := make([]order.ID, 0, numOrders)
	addresses := make([]string, 0, numOrders)
	statuses := make([]uint8, 0, numOrders)

	// Get the first 500 orders
	start := big.NewInt(0)
	orderIDValues, addressValues, statusValues, err := adapter.orderbookContract.GetOrders(&bind.CallOpts{}, start, big.NewInt(500))

	// Loop until all darknode have been loaded
	for {
		if err != nil {
			return nil, nil, nil, err
		}
		for i := range orderIDValues {
			if bytes.Equal(addressValues[i].Bytes(), []byte{}) {
				// We are finished when a nil address is returned
				return orderIDs, addresses, statuses, nil
			}
			orderIDs = append(orderIDs, order.ID(orderIDValues[i]))
			addresses = append(addresses, addressValues[i].Hex())
			statuses = append(statuses, statusValues[i])
		}
		start = start.Add(start, big.NewInt(500))
		orderIDValues, addressValues, statusValues, err = adapter.orderbookContract.GetOrders(&bind.CallOpts{}, start, big.NewInt(500))
		if err != nil {
			return nil, nil, nil, err
		}
	}
}

func (adapter *adapter) Sign(data []byte) ([]byte, error) {
	return adapter.trader.Sign(data)
}

func (adapter *adapter) Address() []byte {
	return adapter.trader.Address().Bytes()
}

func (adapter *adapter) Status(id order.ID) (order.Status, error) {
	status, err := adapter.orderbookContract.OrderState(&bind.CallOpts{}, id)
	return order.Status(status), err
}

func (adapter *adapter) Settled(id order.ID) (bool, error) {
	det, err := adapter.renexSettlementContract.GetMatchDetails(&bind.CallOpts{}, id)
	if err != nil {
		return false, err
	}
	return det.Settled, nil
}

func (adapter *adapter) buildOrderMapping(order order.Order) (httpadapter.OrderFragmentMapping, error) {
	fmt.Println("Starting to build order mapping")
	pods, err := adapter.pods()
	if err != nil {
		return nil, err
	}

	orderFragmentMapping := httpadapter.OrderFragmentMapping{}

	for _, pod := range pods {
		fmt.Println("For pod:", pod.Position)
		n := int64(len(pod.Darknodes))
		k := int64(2 * (len(pod.Darknodes) + 1) / 3)
		hash := base64.StdEncoding.EncodeToString(pod.Hash[:])
		ordFragments, err := order.Split(n, k)
		if err != nil {
			return nil, err
		}
		orderFragmentMapping[hash] = []httpadapter.OrderFragment{}
		for i, ordFragment := range ordFragments {
			fmt.Println("Order Fragment:", i)
			marshaledOrdFragment := httpadapter.OrderFragment{
				Index: int64(i + 1),
			}

			pubKeyBytes, err := adapter.darknodeRegistryContract.GetDarknodePublicKey(&bind.CallOpts{}, common.BytesToAddress(pod.Darknodes[i].Hash()))
			if err != nil {
				return nil, err
			}
			pubKey, err := crypto.RsaPublicKeyFromBytes(pubKeyBytes)
			if err != nil {
				return nil, err
			}

			encryptedFragment, err := ordFragment.Encrypt(pubKey)
			marshaledOrdFragment.ID = base64.StdEncoding.EncodeToString(encryptedFragment.ID[:])
			marshaledOrdFragment.OrderID = base64.StdEncoding.EncodeToString(encryptedFragment.OrderID[:])
			marshaledOrdFragment.OrderParity = encryptedFragment.OrderParity
			marshaledOrdFragment.OrderSettlement = encryptedFragment.OrderSettlement
			marshaledOrdFragment.OrderType = encryptedFragment.OrderType
			marshaledOrdFragment.OrderExpiry = encryptedFragment.OrderExpiry.Unix()
			marshaledOrdFragment.Tokens = base64.StdEncoding.EncodeToString(encryptedFragment.Tokens)
			marshaledOrdFragment.Price = []string{
				base64.StdEncoding.EncodeToString(encryptedFragment.Price.Co),
				base64.StdEncoding.EncodeToString(encryptedFragment.Price.Exp),
			}
			marshaledOrdFragment.Volume = []string{
				base64.StdEncoding.EncodeToString(encryptedFragment.Volume.Co),
				base64.StdEncoding.EncodeToString(encryptedFragment.Volume.Exp),
			}
			marshaledOrdFragment.MinimumVolume = []string{
				base64.StdEncoding.EncodeToString(encryptedFragment.MinimumVolume.Co),
				base64.StdEncoding.EncodeToString(encryptedFragment.MinimumVolume.Exp),
			}
			marshaledOrdFragment.Nonce = base64.StdEncoding.EncodeToString(encryptedFragment.Nonce)
			orderFragmentMapping[hash] = append(orderFragmentMapping[hash], marshaledOrdFragment)
		}
	}

	return orderFragmentMapping, nil
}

func (adapter *adapter) pods() ([]registry.Pod, error) {
	epoch, err := adapter.darknodeRegistryContract.CurrentEpoch(&bind.CallOpts{})
	if err != nil {
		return []registry.Pod{}, err
	}

	darknodeAddrs, err := adapter.darknodes()
	if err != nil {
		return []registry.Pod{}, err
	}

	numberOfNodesInPod, err := adapter.darknodeRegistryContract.MinimumPodSize(&bind.CallOpts{})
	if err != nil {
		return []registry.Pod{}, err
	}
	if len(darknodeAddrs) < int(numberOfNodesInPod.Int64()) {
		return []registry.Pod{}, fmt.Errorf("degraded pod: expected at least %v addresses, got %v", int(numberOfNodesInPod.Int64()), len(darknodeAddrs))
	}

	numberOfDarknodes := big.NewInt(int64(len(darknodeAddrs)))
	x := big.NewInt(0).Mod(epoch.Epochhash, numberOfDarknodes)
	positionInOcean := make([]int, len(darknodeAddrs))
	for i := 0; i < len(darknodeAddrs); i++ {
		positionInOcean[i] = -1
	}
	pods := make([]registry.Pod, (len(darknodeAddrs) / int(numberOfNodesInPod.Int64())))
	for i := 0; i < len(darknodeAddrs); i++ {
		isRegistered, err := adapter.darknodeRegistryContract.IsRegistered(&bind.CallOpts{}, common.BytesToAddress(darknodeAddrs[x.Int64()].Hash()))
		if err != nil {
			return []registry.Pod{}, err
		}
		fmt.Println("Waiting for node registration")
		for !isRegistered || positionInOcean[x.Int64()] != -1 {
			fmt.Println("Is", hex.EncodeToString(darknodeAddrs[x.Int64()].Hash()), "registered:", isRegistered, "position", positionInOcean[x.Int64()])
			x.Add(x, big.NewInt(1))
			x.Mod(x, numberOfDarknodes)
			isRegistered, err = adapter.darknodeRegistryContract.IsRegistered(&bind.CallOpts{}, common.BytesToAddress(darknodeAddrs[x.Int64()].Hash()))
			if err != nil {
				return []registry.Pod{}, err
			}
		}
		fmt.Println("Node registrations complete")
		positionInOcean[x.Int64()] = i
		podID := i % (len(darknodeAddrs) / int(numberOfNodesInPod.Int64()))
		pods[podID].Darknodes = append(pods[podID].Darknodes, darknodeAddrs[x.Int64()])
		x.Mod(x.Add(x, epoch.Epochhash), numberOfDarknodes)
	}

	for i := range pods {
		hashData := [][]byte{}
		for _, darknodeAddr := range pods[i].Darknodes {
			hashData = append(hashData, darknodeAddr.ID())
		}
		copy(pods[i].Hash[:], crypto.Keccak256(hashData...))
		pods[i].Position = i
	}
	return pods, nil
}

func (adapter *adapter) darknodes() (identity.Addresses, error) {
	numDarknodesBig, err := adapter.darknodeRegistryContract.NumDarknodes(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	numDarknodes := numDarknodesBig.Int64()
	darknodes := make(identity.Addresses, 0, numDarknodes)

	// Get the first 20 pods worth of darknodes
	nilValue := common.HexToAddress("0x0000000000000000000000000000000000000000")
	values, err := adapter.darknodeRegistryContract.GetDarknodes(&bind.CallOpts{}, nilValue, big.NewInt(480))

	// Loop until all darknode have been loaded
	for {
		if err != nil {
			return nil, err
		}
		for _, value := range values {
			if bytes.Equal(value.Bytes(), nilValue.Bytes()) {
				// We are finished when a nil address is returned
				return darknodes, nil
			}
			darknodes = append(darknodes, identity.Address(value.Bytes()))
		}
		lastValue := values[len(values)-1]
		values, err = adapter.darknodeRegistryContract.GetDarknodes(&bind.CallOpts{}, lastValue, big.NewInt(480))
		if err != nil {
			return nil, err
		}
		// Skip the first value returned so that we do not duplicate values
		values = values[1:]
	}
}

func getTokenCode(ord order.Order) order.Token {
	if ord.Parity == 0 {
		return ord.Tokens.PriorityToken()
	}
	return ord.Tokens.NonPriorityToken()
}
