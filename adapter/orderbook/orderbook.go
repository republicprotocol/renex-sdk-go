package orderbook

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/republicprotocol/renex-ingress-go/httpadapter"
	"github.com/republicprotocol/renex-sdk-go/adapter/bindings"
	"github.com/republicprotocol/renex-sdk-go/adapter/client"
	"github.com/republicprotocol/renex-sdk-go/adapter/store"
	"github.com/republicprotocol/renex-sdk-go/adapter/trader"
	"github.com/republicprotocol/renex-sdk-go/core/funds"
	"github.com/republicprotocol/renex-sdk-go/core/orderbook"
	"github.com/republicprotocol/republic-go/contract"
	"github.com/republicprotocol/republic-go/order"
)

type adapter struct {
	httpAddress             string
	republicBinder          contract.Binder
	renexSettlementContract *bindings.RenExSettlement
	trader                  trader.Trader
	client                  client.Client
	funds                   funds.Funds
	store                   store.Store
}

func NewAdapter(httpAddress string, client client.Client, trader trader.Trader, funds funds.Funds, store store.Store) (orderbook.Adapter, error) {
	conn, err := contract.Connect(contract.Config{Network: contract.NetworkNightly})
	if err != nil {
		return nil, err
	}

	republicBinder, err := contract.NewBinder(trader.TransactOpts(), conn)

	renexSettlement, err := bindings.NewRenExSettlement(client.RenExSettlementAddress(), bind.ContractBackend(client.Client()))
	if err != nil {
		return nil, err
	}
	return &adapter{
		republicBinder:          republicBinder,
		renexSettlementContract: renexSettlement,
		httpAddress:             httpAddress,
		trader:                  trader,
		client:                  client,
		funds:                   funds,
		store:                   store,
	}, nil
}

func (adapter *adapter) RequestOpenOrder(order order.Order) error {
	if err := adapter.BalanceCheck(order); err != nil {
		return err
	}

	mapping, err := adapter.buildOrderMapping(order)
	if err != nil {
		return err
	}

	req := httpadapter.OpenOrderRequest{
		Address:               adapter.trader.Address().String()[2:],
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

	// type Response struct {
	// 	Signature string `json:"signature"`
	// }

	// respose := Response{}
	// if err := json.Unmarshal(respBytes, respose); err != nil {
	// 	return err
	// }

	if resp.StatusCode != 201 {
		return fmt.Errorf("Unexpected status code %d", resp.StatusCode)
	}

	sigBytes, err := base64.StdEncoding.DecodeString(string(respBytes))
	if err != nil {
		return err
	}

	sig, err := toBytes65(sigBytes)
	if err != nil {
		return err
	}

	if err := adapter.republicBinder.OpenOrder(1, sig, order.ID); err != nil {
		return err
	}

	return adapter.store.AppendOrder(order)
}

func (adapter *adapter) RequestCancelOrder(orderID order.ID) error {
	if err := adapter.republicBinder.CancelOrder(orderID); err != nil {
		return err
	}
	return adapter.store.DeleteOrder(orderID)
}

func (adapter *adapter) ListOrders() ([]order.ID, []order.Status, []string, error) {
	numOrders, err := adapter.republicBinder.OrderCounts()
	if err != nil {
		return nil, nil, nil, err
	}
	orderCount := int(numOrders)
	orderIDs := make([]order.ID, 0, orderCount)
	addresses := make([]string, 0, orderCount)
	statuses := make([]order.Status, 0, orderCount)

	start := 0
	if orderCount < 500 {
		return adapter.republicBinder.Orders(0, orderCount)
	}

	for {
		if orderCount-start < 0 {
			return orderIDs, statuses, addresses, nil
		}
		orderIDValues, statusValues, addressValues, err := adapter.republicBinder.Orders(0, orderCount)
		if err != nil {
			return nil, nil, nil, err
		}
		orderIDs = append(orderIDs, orderIDValues...)
		addresses = append(addresses, addressValues...)
		statuses = append(statuses, statusValues...)
		start = start + 500
	}
}

func (adapter *adapter) Sign(data []byte) ([]byte, error) {
	return adapter.trader.Sign(data)
}

func (adapter *adapter) Address() []byte {
	return adapter.trader.Address().Bytes()
}

func (adapter *adapter) Status(id order.ID) (order.Status, error) {
	return adapter.republicBinder.Status(id)
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
	pods, err := adapter.republicBinder.Pods()
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

			pubKey, err := adapter.republicBinder.PublicKey(pod.Darknodes[i].ID().Address())
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

func (adapter *adapter) BalanceCheck(order order.Order) error {
	token := getTokenCode(order)
	balance, err := adapter.funds.UsableRenExBalance(token)
	if err != nil {
		return err
	}
	decodedVolume := decodeVolume(token, order.Volume)
	if balance.Cmp(decodedVolume) < 0 {
		return fmt.Errorf("[%v] Order volume exceeded usable balance have:%v want:%v", token, balance, decodedVolume)
	}
	return nil
}

func getTokenCode(ord order.Order) order.Token {
	if ord.Parity == order.ParityBuy {
		return ord.Tokens.PriorityToken()
	}
	return ord.Tokens.NonPriorityToken()
}

func toBytes65(b []byte) ([65]byte, error) {
	bytes65 := [65]byte{}
	if len(b) != 65 {
		return bytes65, errors.New("Length mismatch")
	}
	for i := range b {
		bytes65[i] = b[i]
	}
	return bytes65, nil
}

func decodeVolume(token order.Token, volume uint64) *big.Int {
	switch token {
	case order.TokenABC:
		return big.NewInt(int64(volume))
	case order.TokenDGX:
		return big.NewInt(int64(volume) / 1000)
	case order.TokenBTC:
		return big.NewInt(int64(volume) / 10000)
	default:
		return big.NewInt(int64(volume)).Mul(big.NewInt(int64(volume)), big.NewInt(1000000))
	}
}
