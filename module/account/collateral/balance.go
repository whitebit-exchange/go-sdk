package collateral

import (
	"github.com/whitebit-exchange/go-sdk"
)

const balanceEndpointUrl = "/api/v4/collateral-account/balance"
const balanceSummaryEndpointUrl = "/api/v4/collateral-account/balance-summary"

type balanceEndpoint struct {
	whitebit.AuthParams

	Ticker string `json:"ticker,omitempty"`
}

type BalanceSummary struct {
	Asset                  string `json:"asset"`
	Balance                string `json:"balance"`
	Borrow                 string `json:"borrow"`
	AvailableWithoutBorrow string `json:"availableWithoutBorrow"`
	AvailableWithBorrow    string `json:"availableWithBorrow"`
}

func newBalanceEndpoint(ticker string) *balanceEndpoint {
	return &balanceEndpoint{Ticker: ticker, AuthParams: whitebit.NewAuthParams(balanceEndpointUrl)}
}

func newBalanceSummaryEndpoint(ticker string) *balanceEndpoint {
	return &balanceEndpoint{Ticker: ticker, AuthParams: whitebit.NewAuthParams(balanceSummaryEndpointUrl)}
}
