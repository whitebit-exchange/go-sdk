package trade

import (
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/module/account"
	"github.com/whitebit-exchange/whitebit/module/order/spot"
)

const ordersEndpointUrl = "/api/v4/orders"

type Orders struct {
	spot.MarketOrder

	Price           string `json:"price,omitempty"`
	ActivationPrice string `json:"activation_price,omitempty"`
}

type ordersEndpoint struct {
	whitebit.AuthParams
	account.MarketWithPaginationParams
}

func newOrdersEndpoint(market string, limit int, offset int) *ordersEndpoint {
	return &ordersEndpoint{
		AuthParams:                 whitebit.NewAuthParams(ordersEndpointUrl),
		MarketWithPaginationParams: account.NewMarketWithPaginationParams(market, limit, offset),
	}
}
