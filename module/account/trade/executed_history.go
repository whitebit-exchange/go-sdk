package trade

import (
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/module/account"
)

const executedHistoryEndpointUrl = "/api/v4/trade-account/executed-history"

type ExecutedHistory struct {
	Id            int64   `json:"id"`
	ClientOrderId string  `json:"clientOrderId"`
	Time          float64 `json:"time"`
	Side          string  `json:"side"`
	Role          int     `json:"role"`
	Amount        string  `json:"amount"`
	Price         string  `json:"price"`
	Deal          string  `json:"deal"`
	Fee           string  `json:"fee"`
	OrderId       int64   `json:"orderId"`
}

type executedHistoryEndpoint struct {
	whitebit.AuthParams
	account.ParamsForFiltration
}

func newExecutedHistoryEndpoint(market string, limit int, offset int, clientOrderId string) *executedHistoryEndpoint {
	optionsForFiltration := account.NewParamsForFiltration(market, limit, offset, 0, clientOrderId)
	return &executedHistoryEndpoint{ParamsForFiltration: optionsForFiltration, AuthParams: whitebit.NewAuthParams(executedHistoryEndpointUrl)}
}
