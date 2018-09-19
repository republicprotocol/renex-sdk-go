package swapper

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	netHttp "net/http"

	"github.com/republicprotocol/renex-swapper-go/adapter/http"
	"github.com/republicprotocol/renex-swapper-go/domain/swap"
	"github.com/republicprotocol/republic-go/order"
)

type swapper struct {
	httpAddress string
}
type Swapper interface {
	PostOrder(order.ID) error
	Status(order.ID) swap.Status
	Balance(order.Token) (uint64, error)
}

func New(httpAddress string) Swapper {
	return &swapper{
		httpAddress: httpAddress,
	}
}

func (swapper *swapper) PostOrder(orderID order.ID) error {
	req := http.PostOrderRequest{
		OrderID: hex.EncodeToString(orderID[:]),
	}

	data, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(data)

	resp, err := netHttp.DefaultClient.Post(fmt.Sprintf("%s/orders", swapper.httpAddress), "application/json", buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode == netHttp.StatusCreated) {
		return fmt.Errorf("Unexpected status code %d", resp.StatusCode)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	response := http.PostOrderResponse{}
	if err := json.Unmarshal(respBytes, &response); err != nil {
		return err
	}

	// TODO: Verify response signatiure
	// sigBytes, err := hex.DecodeString(response.Signature)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (swapper *swapper) Status(orderID order.ID) swap.Status {
	resp, err := netHttp.DefaultClient.Get(fmt.Sprintf("%s/status/%s", swapper.httpAddress, hex.EncodeToString(orderID[:])))
	if err != nil {
		return swap.StatusUnknown
	}

	if !(resp.StatusCode == netHttp.StatusOK) {
		return swap.StatusUnknown
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return swap.StatusUnknown
	}

	response := http.Status{}
	if err := json.Unmarshal(respBytes, &response); err != nil {
		return swap.StatusUnknown
	}

	return swap.Status(response.Status)
}

func (swapper *swapper) Balance(token order.Token) (uint64, error) {
	resp, err := netHttp.DefaultClient.Get(fmt.Sprintf("%s/balances", swapper.httpAddress))
	if err != nil {
		return 0, err
	}

	if !(resp.StatusCode == netHttp.StatusOK) {
		return 0, fmt.Errorf("Unexpected status code %d", resp.StatusCode)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	response := http.Balances{}
	if err := json.Unmarshal(respBytes, &response); err != nil {
		return 0, err
	}

	if token == order.TokenETH {
		return response.Ethereum.Amount, nil
	}

	if token == order.TokenBTC {
		return response.Bitcoin.Amount, nil
	}

	return 0, fmt.Errorf("Unknown token")
}
