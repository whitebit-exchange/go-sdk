package collateral

import (
	"github.com/whitebit-exchange/whitebit"
)

const balanceEndpointUrl = "/api/v4/collateral-account/balance"

type balanceEndpoint struct {
	whitebit.AuthParams

	Ticker string `json:"ticker,omitempty"`
}

func newBalanceEndpoint(ticker string) *balanceEndpoint {
	return &balanceEndpoint{Ticker: ticker, AuthParams: whitebit.NewAuthParams(balanceEndpointUrl)}
}
