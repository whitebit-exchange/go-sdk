package main_account

import (
	"github.com/whitebit-exchange/go-sdk"
)

const balanceEndpointUrl = "/api/v4/main-account/balance"

type balanceEndpoint struct {
	whitebit.AuthParams
	Ticker string `json:"ticker,omitempty"`
}

type MainBalance struct {
	MainBalance string `json:"main_balance"`
}

type State map[string]MainBalance

func newBalanceEndpoint(ticker string) *balanceEndpoint {
	return &balanceEndpoint{Ticker: ticker, AuthParams: whitebit.NewAuthParams(balanceEndpointUrl)}
}
