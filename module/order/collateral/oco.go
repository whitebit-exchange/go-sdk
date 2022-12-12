package collateral

import (
	"github.com/whitebit-exchange/whitebit"
)

const ocoEndpointUrl = "/api/v4/order/collateral/oco"

type OcoOrder struct {
	ID         int64 `json:"id"`
	StopLoss   `json:"stop_loss"`
	TakeProfit `json:"take_profit"`
}

type OcoOrderParams struct {
	Market          string `json:"market"`
	Amount          string `json:"amount"`
	Side            string `json:"side"`
	Price           string `json:"price"`
	ActivationPrice string `json:"activation_price"`
	StopLimitPrice  string `json:"stop_limit_price"`
	ClientOrderId   string `json:"clientOrderId,omitempty"`
}

type ocoEndpoint struct {
	whitebit.AuthParams
	OcoOrderParams
}

func newOcoEndpoint(params OcoOrderParams) *ocoEndpoint {
	return &ocoEndpoint{
		AuthParams:     whitebit.NewAuthParams(ocoEndpointUrl),
		OcoOrderParams: params,
	}
}
