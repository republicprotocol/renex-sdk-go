package store

import (
	"encoding/json"
	"math/big"

	"github.com/republicprotocol/republic-go/order"
)

type store struct {
	StoreAdapter
}

type orders struct {
	ids []order.ID `json:"id"`
}

type StoreAdapter interface {
	Read(key []byte) ([]byte, error)
	Write(key []byte, value []byte) error
	Delete(key []byte) error
	Close() error
}

type Store interface {
	RequestLockedBalance(order.Token) (*big.Int, error)
	OpenOrdersExist(order.Token) (bool, error)
}

func NewStore(adapter StoreAdapter) Store {
	return &store{}
}

func (store *store) RequestLockedBalance(tokenCode order.Token) (*big.Int, error) {
	ords, err := store.openOrders(tokenCode)
	if err != nil {
		return nil, err
	}
	balance := uint64(0)
	for _, ord := range ords {
		balance += ord.Volume
	}
	return big.NewInt(int64(balance)), nil
}

func (store *store) OpenOrdersExist(tokenCode order.Token) (bool, error) {
	orders, err := store.openOrders(tokenCode)
	if err != nil {
		return false, err
	}
	return len(orders) == 0, nil
}

func (store *store) openOrders(tokenCode order.Token) ([]order.Order, error) {
	data, err := store.Read([]byte("ORDERS"))
	if err != nil {
		return nil, err
	}
	orderList := orders{}
	if err := json.Unmarshal(data, &orderList); err != nil {
		return nil, err
	}

	orders := []order.Order{}
	for _, id := range orderList.ids {
		ord, err := store.Order(id)
		if err != nil {
			return nil, err
		}
		if checkToken(ord, tokenCode) {
			orders = append(orders, ord)
		}
	}

	return orders, nil
}

func (store *store) Order(id order.ID) (order.Order, error) {
	data, err := store.Read(append([]byte("ORDER"), id[:]...))
	if err != nil {
		return order.Order{}, err
	}
	ord := order.Order{}
	if err := json.Unmarshal(data, &ord); err != nil {
		return order.Order{}, err
	}
	return ord, nil
}

func (store *store) AppendOrder(ord order.Order) error {
	data, err := json.Marshal(ord)
	if err != nil {
		return err
	}
	if err := store.Write(append([]byte("ORDER"), ord.ID[:]...), data); err != nil {
		return err
	}
	ordersData, err := store.Read([]byte("ORDERS"))
	if err != nil {
		return err
	}
	orderList := orders{}
	if err := json.Unmarshal(ordersData, &orderList); err != nil {
		return err
	}
	orderList.ids = append(orderList.ids, ord.ID)
	orderListData, err := json.Marshal(orderList)
	if err != nil {
		return err
	}
	return store.Write([]byte("ORDERS"), orderListData)
}

func (store *store) DeleteOrder(id order.ID) error {
	if err := store.Delete(append([]byte("ORDER"), id[:]...)); err != nil {
		return err
	}
	data, err := store.Read([]byte("ORDERS"))
	if err != nil {
		return err
	}
	orderList := orders{}
	if err := json.Unmarshal(data, &orderList); err != nil {
		return err
	}
	orderList.ids = deleteFromList(orderList.ids, id)
	orderListData, err := json.Marshal(orderList)
	if err != nil {
		return err
	}
	return store.Write([]byte("ORDERS"), orderListData)
}

func deleteFromList(ids []order.ID, deleteID order.ID) []order.ID {
	for i, id := range ids {
		if id == deleteID {
			return append(ids[:i], ids[i+1:]...)
		}
	}
	return ids
}

func checkToken(ord order.Order, tokenCode order.Token) bool {
	if ord.Parity == 0 {
		return (order.Token(ord.Tokens.PriorityToken()) == tokenCode)
	}
	return (order.Token(ord.Tokens.NonPriorityToken()) == tokenCode)
}
