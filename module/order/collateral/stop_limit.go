package collateral

import (
	"github.com/whitebit-exchange/whitebit"
)

const stopLimitEndpointUrl = "/api/v4/order/collateral/stop-limit"

type StopLimitOrder struct {
	StopMarketOrder
	Price string `json:"price"`
}

type StopLimitOrderParams struct {
	Market          string `json:"market"`
	Amount          string `json:"amount"`
	Side            string `json:"side"`
	Price           string `json:"price"`
	ActivationPrice string `json:"activation_price"`
	ClientOrderId   string `json:"clientOrderId,omitempty"`
}

type stopLimitEndpoint struct {
	whitebit.AuthParams
	StopLimitOrderParams
}

func newStopLimitEndpoint(params StopLimitOrderParams) *stopLimitEndpoint {
	return &stopLimitEndpoint{
		AuthParams:           whitebit.NewAuthParams(stopLimitEndpointUrl),
		StopLimitOrderParams: params,
	}
}
