package spot

import (
	"github.com/whitebit-exchange/whitebit"
)

const stockMarketOrderEndpointUrl = "/api/v4/order/stock_market"

type StockMarketOrder struct {
	MarketOrder
}

type stockMarketEndpoint struct {
	whitebit.AuthParams
	MarketOrderParams
}

func newStockMarketEndpoint(params MarketOrderParams) *stockMarketEndpoint {
	return &stockMarketEndpoint{
		AuthParams:        whitebit.NewAuthParams(stockMarketOrderEndpointUrl),
		MarketOrderParams: params,
	}
}
