package collateral

import (
	"github.com/whitebit-exchange/whitebit"
)

const leverageEndpointUrl = "/api/v4/collateral-account/leverage"

type leverageEndpoint struct {
	whitebit.AuthParams

	Leverage string `json:"leverage,omitempty"`
}

func newLeverageEndpoint(leverage string) *leverageEndpoint {
	return &leverageEndpoint{Leverage: leverage, AuthParams: whitebit.NewAuthParams(leverageEndpointUrl)}
}
