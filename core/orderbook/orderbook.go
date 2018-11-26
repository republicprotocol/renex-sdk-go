package orderbook

import (
	"math/big"

	"github.com/republicprotocol/republic-go/order"
)

type service struct {
	Adapter
}

type OrderMatch struct {
	Settled         bool
	OrderIsBuy      bool
	MatchedID       [32]byte
	PriorityVolume  *big.Int
	SecondaryVolume *big.Int
	PriorityFee     *big.Int
	SecondaryFee    *big.Int
	PriorityToken   uint32
	SecondaryToken  uint32
}

type Adapter interface {
	Status(order.ID) (order.Status, error)
	Settled(order.ID) (bool, error)
	RequestOpenOrder(order order.Order) error
	RequestCancelOrder(orderID order.ID) error
	MatchDetails(orderID order.ID) (OrderMatch, error)
	ListOrders() ([]order.ID, []order.Status, []string, error)
}
type Orderbook interface {
	Status(order.ID) (order.Status, error)
	Settled(order.ID) (bool, error)
	OpenOrder(order order.Order) error
	CancelOrder(orderID order.ID) error
	MatchDetails(orderID order.ID) (OrderMatch, error)
	ListOrdersByTrader(address string) ([]order.ID, error)
	ListOrdersByStatus(status order.Status) ([]order.ID, error)
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
	orderIds, _, addresses, err := service.ListOrders()
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

func (service *service) ListOrdersByStatus(status order.Status) ([]order.ID, error) {
	orderIds, statuses, _, err := service.ListOrders()
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
