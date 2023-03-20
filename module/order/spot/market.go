package spot

import (
	"github.com/whitebit-exchange/go-sdk"
)

const marketOrderEndpointUrl = "/api/v4/order/market"

type MarketOrder struct {
	OrderID       int64   `json:"orderId"`
	ClientOrderID string  `json:"clientOrderId"`
	Market        string  `json:"market"`
	Side          string  `json:"side"`
	Type          string  `json:"type"`
	Timestamp     float64 `json:"timestamp"`
	DealMoney     string  `json:"dealMoney"`
	DealStock     string  `json:"dealStock"`
	Amount        string  `json:"amount"`
	TakerFee      string  `json:"takerFee"`
	MakerFee      string  `json:"makerFee"`
	Left          string  `json:"left"`
	DealFee       string  `json:"dealFee"`
	PostOnly      bool    `json:"postOnly"`
	IOC           bool    `json:"ioc"`
}

type MarketOrderParams struct {
	Market        string `json:"market"`
	Amount        string `json:"amount"`
	Side          string `json:"side"`
	ClientOrderId string `json:"clientOrderId,omitempty"`
}

type marketEndpoint struct {
	whitebit.AuthParams
	MarketOrderParams
}

func newMarketOrderEndpoint(params MarketOrderParams) *marketEndpoint {
	return &marketEndpoint{
		AuthParams:        whitebit.NewAuthParams(marketOrderEndpointUrl),
		MarketOrderParams: params,
	}
}
