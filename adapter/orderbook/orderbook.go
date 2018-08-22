package orderbook

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/republicprotocol/renex-ingress-go/httpadapter"
	"github.com/republicprotocol/renex-sdk-go/adapter/bindings"
	"github.com/republicprotocol/renex-sdk-go/adapter/client"
	"github.com/republicprotocol/renex-sdk-go/adapter/trader"
	"github.com/republicprotocol/republic-go/order"
)

type adapter struct {
	httpAddress              string
	orderbookContract        *bindings.Orderbook
	darknodeRegistryContract *bindings.DarknodeRegistry
	trader                   trader.Trader
}

type Adapter interface {
	RequestOpenOrder(order order.Order, signature string) error
	RequestCloseOrder(orderID order.ID, signature string) error

	ListOrders() ([]order.Order, []string, []uint8, error)

	Sign([]byte) ([]byte, error)
}

func NewAdapter(httpAddress string, client client.Client, trader trader.Trader) (Adapter, error) {
	orderbook, err := bindings.NewOrderbook(client.OrderbookAddress(), client.Conn())
	if err != nil {
		return nil, err
	}
	darknodeRegistry, err := bindings.NewDarknodeRegistry(client.DarknodeRegistryAddress(), client.Conn())
	if err != nil {
		return nil, err
	}
	return &adapter{
		orderbookContract:        orderbook,
		darknodeRegistryContract: darknodeRegistry,
		httpAddress:              httpAddress,
	}
}

func (adapter *adapter) RequestOpenOrder(order order.Order, signature string) error {
	mapping, err := adapter.buildOrderMapping(order, signature)
	if err != nil {
		return err
	}

	req := httpadapter.OpenOrderRequest{
		Signature:             signature,
		OrderFragmentMappings: []httpadapter.OrderFragmentMapping{mapping},
	}

	data, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(data)

	resp, err := http.DefaultClient.Post(fmt.Sprintf("%s/orders", adapter.httpAddress), "application/json", buf)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}

	return fmt.Errorf("Unexpected status code %d", status)
}

func (adapter *adapter) RequestCloseOrder(orderID order.ID, signature string) error {

	client := http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/orders?id=%s&signature=%s", adapter.httpAddress, orderID.String(), signature), nil)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}

	return fmt.Errorf("Unexpected status code %d", status)
}

func (adapter *adapter) ListOrders() ([]order.ID, []string, []uint8, error) {
	limit, error := adapter.orderbookContract.GetOrdersCount(*bin.CallOpts{})
	if err != nil {
		return nil, nil, nil, err
	}

	// TODO: GetOrders will be updated to expect settlement id as the first
	// parameter.
	orderIds, addresses, statuses, err := adapter.orderbookContract.GetOrders(*bind.CallOpts{}, big.NewInt(0), limit)
	if err != nil {
		return nil, nil, nil, err
	}

	addressStrings := []string{}
	for _, address := range addresses {
		addressStrings = append(addressStrings, address.String())
	}

	return []order.ID{orderIds}, addressStrings, statuses, nil
}

func (adapter *adapter) Sign(data []byte) ([]byte, error) {
	return adapter.trader.Sign(data)
}

func (adapter *adapter) buildOrderMapping(order order.Order, signature string) (httpadapter.OrderFragmentMapping, error) {
	pods, err := adapter.pods()
	if err != nil {
		return nil, err
	}
	orderFragmentMapping := httpadapter.OrderFragmentMapping{}

	for _, pod := range pods {
		n := int64(len(pod.Darknodes))
		k := int64(2 * (len(pod.Darknodes) + 1) / 3)
		hash := base64.StdEncoding.EncodeToString(pod.Hash[:])
		ordFragments, err := order.Split(n, k)
		if err != nil {
			return nil, err
		}
		orderFragmentMapping[hash] = []httpadapter.OrderFragment{}
		for i, ordFragment := range ordFragments {
			marshaledOrdFragment := httpadapter.OrderFragment{
				Index: int64(i + 1),
			}

			pubKey, err := adapter.PublicKey(pod.Darknodes[i])
			if err != nil {
				return nil, err
			}

			encryptedFragment, err := ordFragment.Encrypt(pubKey)
			marshaledOrdFragment.OrderSignature = signature
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
