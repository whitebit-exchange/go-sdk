package collateral

import (
	"encoding/json"
	"github.com/whitebit-exchange/go-sdk"
)

const positionsHistoryEndpointUrl = "/api/v4/collateral-account/positions/history"

type PositionHistory struct {
	PositionID       int64       `json:"positionId"`
	Market           string      `json:"market"`
	OpenDate         json.Number `json:"openDate"`
	ModifyDate       json.Number `json:"modifyDate"`
	Amount           string      `json:"amount"`
	BasePrice        string      `json:"basePrice"`
	RealizedFunding  string      `json:"realizedFunding"`
	LiquidationPrice string      `json:"liquidationPrice"`
	LiquidationState string      `json:"liquidationState"`
	OrderDetail      OrderDetail `json:"orderDetail"`
}

type OrderDetail struct {
	ID          int64  `json:"id"`
	TradeAmount string `json:"tradeAmount"`
	Price       string `json:"price"`
	TradeFee    string `json:"tradeFee"`
	FundingFee  string `json:"fundingFee"`
}

type positionHistoryEndpoint struct {
	whitebit.AuthParams

	Market     string `json:"market,omitempty"`
	PositionId int64  `json:"positionId,omitempty"`
	StartDate  int64  `json:"startDate,omitempty"`
	EndDate    int64  `json:"endDate,omitempty"`
	Limit      string `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
}

func newPositionsHistoryEndpoint(market string, positionId int64, startDate int64, endDate int64, limit string, offset string) *positionHistoryEndpoint {
	return &positionHistoryEndpoint{
		Market:     market,
		PositionId: positionId,
		StartDate:  startDate,
		EndDate:    endDate,
		Limit:      limit,
		Offset:     offset,
		AuthParams: whitebit.NewAuthParams(positionsHistoryEndpointUrl),
	}
}
