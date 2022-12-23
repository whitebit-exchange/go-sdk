package trade

import (
	"github.com/whitebit-exchange/go-sdk"
)

const balanceEndpointUrl = "/api/v4/trade-account/balance"

type Balance struct {
	Available string `json:"available"`
	Freeze    string `json:"freeze"`
}

type balanceEndpoint struct {
	whitebit.AuthParams

	Ticker string `json:"ticker,omitempty"`
}

func newBalanceEndpoint(ticker string) *balanceEndpoint {
	return &balanceEndpoint{Ticker: ticker, AuthParams: whitebit.NewAuthParams(balanceEndpointUrl)}
}
