package collateral

import (
	"github.com/whitebit-exchange/go-sdk"
)

const openPositionsEndpointUrl = "/api/v4/collateral-account/positions/open"

type OpenPosition struct {
	PositionID        int64   `json:"positionId"`
	Market            string  `json:"market"`
	OpenDate          float64 `json:"openDate"`
	ModifyDate        float64 `json:"modifyDate"`
	Amount            string  `json:"amount"`
	BasePrice         string  `json:"basePrice"`
	LiquidationPrice  string  `json:"liquidationPrice"`
	Leverage          string  `json:"leverage"`
	Pnl               string  `json:"pnl"`
	PnlPercent        string  `json:"pnlPercent"`
	Margin            string  `json:"margin"`
	FreeMargin        string  `json:"freeMargin"`
	Funding           string  `json:"funding"`
	UnrealizedFunding string  `json:"unrealizedFunding"`
	LiquidationState  string  `json:"liquidationState"`
}

type openPositionsEndpoint struct {
	whitebit.AuthParams

	Market string
}

func newOpenPositionsEndpoint(market string) *openPositionsEndpoint {
	return &openPositionsEndpoint{Market: market, AuthParams: whitebit.NewAuthParams(openPositionsEndpointUrl)}
}
