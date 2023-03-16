package collateral

import (
	"github.com/whitebit-exchange/go-sdk"
)

const summaryEndpointUrl = "/api/v4/collateral-account/summary"

type Summary struct {
	Equity            string `json:"equity"`
	Margin            string `json:"margin"`
	FreeMargin        string `json:"freeMargin"`
	UnrealizedFunding string `json:"unrealizedFunding"`
	Pnl               string `json:"pnl"`
	Leverage          int    `json:"leverage"`
	MarginFraction    string `json:"marginFraction"`
}

type summaryEndpoint struct {
	whitebit.AuthParams
}

func newSummaryEndpoint() *summaryEndpoint {
	return &summaryEndpoint{AuthParams: whitebit.NewAuthParams(summaryEndpointUrl)}
}
