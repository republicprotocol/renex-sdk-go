package orderbook

import (
	"github.com/republicprotocol/renex-sdk-go/adapter/orderbook"
	"github.com/republicprotocol/republic-go/crypto"
	"github.com/republicprotocol/republic-go/order"
)

type service struct {
	orderbook.Adapter
}

type Service interface {
	OpenOrder(order order.Order) error
	CloseOrder(orderID order.ID) error
	ListOrdersByTrader(address string) ([]order.ID, error)
	ListOrdersByStatus(status uint8) ([]order.ID, error)
}

func NewService(adapter orderbook.Adapter) Service {
	return &service{
		adapter,
	}
}

func (service *service) OpenOrder(order order.Order) error {
	signatureData := crypto.Keccak256([]byte("Republic Protocol: open: "), order.ID[:])
	signatureData = crypto.Keccak256([]byte("\x19Ethereum Signed Message:\n32"), signatureData)
	signature, err := service.Sign(signatureData)
	if err != nil {
		return err
	}
	return service.RequestOpenOrder(order, trader)
}

func (service *service) CloseOrder(orderID order.ID) error {
	signatureData := crypto.Keccak256([]byte("Republic Protocol: close: "), orderID[:])
	signatureData = crypto.Keccak256([]byte("\x19Ethereum Signed Message:\n32"), signatureData)
	signature, err := service.Sign(signatureData)
	if err != nil {
		return err
	}
	return service.RequestCloseOrder(order, trader)
}

func (service *service) ListOrdersByTrader(traderAddress string) ([]order.ID, error) {
	orderIds, addresses, _, err := service.ListOrders()
	if err != nil {
		return nil, err
	}
	orderList := []order.ID{}
	for i, id := range orderIds {
		if traderAddress == addresses[i] {
			orderList = append(orderList, id)
		}
	}
	return orderList, nil
}

func (service *service) ListOrdersByStatus(status uint8) error {
	orderIds, _, statuses, err := service.ListOrders()
	if err != nil {
		return nil, err
	}
	orderList := []order.ID{}
	for i, id := range orderIds {
		if status == statuses[i] {
			orderList = append(orderList, id)
		}
	}
	return orderList, nil
}
