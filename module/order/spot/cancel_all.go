package spot

import "github.com/whitebit-exchange/go-sdk"

const cancelAllOrdersEndpointURL = "/api/v4/order/cancel/all"

type cancelAllOrdersEndpoint struct {
	whitebit.AuthParams
	Market string   `json:"market,omitempty"`
	Type   []string `json:"type,omitempty"`
}

func newCancelAllOrdersEndpoint(market string, types []string) *cancelAllOrdersEndpoint {
	return &cancelAllOrdersEndpoint{
		AuthParams: whitebit.NewAuthParams(cancelAllOrdersEndpointURL),
		Market:     market,
		Type:       types,
	}
}
