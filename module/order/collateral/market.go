package collateral

import (
	"github.com/whitebit-exchange/whitebit"
)

const marketOrderEndpointUrl = "/api/v4/order/collateral/market"

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

func newMarketEndpoint(params MarketOrderParams) *marketEndpoint {
	return &marketEndpoint{
		AuthParams:        whitebit.NewAuthParams(marketOrderEndpointUrl),
		MarketOrderParams: params,
	}
}
