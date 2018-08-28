package orderbook

import (
	"github.com/republicprotocol/republic-go/order"
)

type service struct {
	Adapter
}

type Adapter interface {
	Status(order.ID) (order.Status, error)
	Settled(order.ID) (bool, error)
	RequestOpenOrder(order order.Order) error
	RequestCancelOrder(orderID order.ID) error
	ListOrders() ([]order.ID, []string, []uint8, error)
}
type Orderbook interface {
	Status(order.ID) (order.Status, error)
	Settled(order.ID) (bool, error)
	OpenOrder(order order.Order) error
	CancelOrder(orderID order.ID) error
	ListOrdersByTrader(address string) ([]order.ID, error)
	ListOrdersByStatus(status uint8) ([]order.ID, error)
}

func NewService(adapter Adapter) Orderbook {
	return &service{
		adapter,
	}
}

func (service *service) OpenOrder(order order.Order) error {
	return service.RequestOpenOrder(order)
}

func (service *service) CancelOrder(orderID order.ID) error {
	return service.RequestCancelOrder(orderID)
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

func (service *service) ListOrdersByStatus(status uint8) ([]order.ID, error) {
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
