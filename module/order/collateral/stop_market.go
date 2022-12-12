package collateral

import (
	"github.com/whitebit-exchange/whitebit"
)

const stopMarketOrderEndpointUrl = "/api/v4/order/collateral/trigger-market"

type StopMarketOrder struct {
	MarketOrder
	ActivationPrice string `json:"activation_price"`
}

type StopMarketOrderParams struct {
	Market          string `json:"market"`
	Amount          string `json:"amount"`
	Side            string `json:"side"`
	ActivationPrice string `json:"activation_price"`
	ClientOrderId   string `json:"clientOrderId,omitempty"`
}

type stopMarketEndpoint struct {
	whitebit.AuthParams
	StopMarketOrderParams
}

func newStopMarketEndpoint(params StopMarketOrderParams) *stopMarketEndpoint {
	return &stopMarketEndpoint{
		AuthParams:            whitebit.NewAuthParams(stopMarketOrderEndpointUrl),
		StopMarketOrderParams: params,
	}
}
