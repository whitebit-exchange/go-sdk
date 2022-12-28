package main_account

import (
	"github.com/whitebit-exchange/go-sdk"
)

const balanceEndpointUrl = "/api/v4/main-account/balance"

type balanceEndpoint struct {
	whitebit.AuthParams
	Ticker string `json:"ticker,omitempty"`
}

type MainBalanceResult struct {
	MainBalance string `json:"main_balance"`
}

type BalanceResult map[string]MainBalanceResult

func newBalanceEndpoint(ticker string) *balanceEndpoint {
	return &balanceEndpoint{Ticker: ticker, AuthParams: whitebit.NewAuthParams(balanceEndpointUrl)}
}
