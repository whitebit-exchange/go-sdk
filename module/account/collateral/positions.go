package collateral

import "github.com/whitebit-exchange/go-sdk"

const (
	positionsEndpointURL     = "/api/v4/collateral-account/positions"
	positionCloseEndpointURL = "/api/v4/collateral-account/position/close"
)

type positionsEndpoint struct {
	whitebit.AuthParams
	Market string `json:"market,omitempty"`
}

func newPositionsEndpoint(market string) *positionsEndpoint {
	return &positionsEndpoint{
		AuthParams: whitebit.NewAuthParams(positionsEndpointURL),
		Market:     market,
	}
}

type positionCloseEndpoint struct {
	whitebit.AuthParams
	Market       string `json:"market"`
	PositionSide string `json:"positionSide,omitempty"`
}

func newPositionCloseEndpoint(market, positionSide string) *positionCloseEndpoint {
	return &positionCloseEndpoint{
		AuthParams:   whitebit.NewAuthParams(positionCloseEndpointURL),
		Market:       market,
		PositionSide: positionSide,
	}
}

type Position struct {
	PositionID    int64   `json:"positionId"`
	Market        string  `json:"market"`
	Amount        string  `json:"amount"`
	BasePrice     string  `json:"basePrice"`
	LiqPrice      string  `json:"liqPrice"`
	PnL           string  `json:"pnl"`
	PnLPercent    string  `json:"pnlPercent"`
	Margin        string  `json:"margin"`
	FreeMargin    string  `json:"freeMargin"`
	Funding       string  `json:"funding"`
	UnrealizedPnL string  `json:"unrealizedPnl"`
	PositionSide  string  `json:"positionSide"`
	OpenDate      float64 `json:"openDate"`
	ModifyDate    float64 `json:"modifyDate"`
}
