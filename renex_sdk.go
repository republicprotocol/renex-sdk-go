package renex

import (
	"github.com/republicprotocol/renex-sdk-go/core/funds"
	"github.com/republicprotocol/renex-sdk-go/core/orderbook"
)

type renEx struct {
}

type RenEx interface {
	orderbook.Input
	funds.Input
}

func NewRenEx() RenEx {
	return nil
}
