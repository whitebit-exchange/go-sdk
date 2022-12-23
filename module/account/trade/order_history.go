package trade

import (
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/account"
)

const OrderHistoryEndpointUrl = "/api/v4/trade-account/order/history"

type OrderHistory struct {
	Id            int64   `json:"id"`
	ClientOrderId string  `json:"clientOrderId"`
	StartTime     float64 `json:"ctime"`
	EndTime       float64 `json:"ftime"`
	Side          string  `json:"side"`
	Amount        string  `json:"amount"`
	Price         string  `json:"price"`
	Type          string  `json:"type"`
	TakerFee      string  `json:"takerFee"`
	MakerFee      string  `json:"makerFee"`
	DealFee       string  `json:"dealFee"`
	DealStock     string  `json:"dealStock"`
	DealMoney     string  `json:"dealMoney"`
}

type historyEndpoint struct {
	whitebit.AuthParams
	account.ParamsForFiltration
}

func newHistoryEndpoint(market string, limit int, offset int, orderId int, clientOrderId string) *historyEndpoint {
	optionsForFiltration := account.NewParamsForFiltration(market, limit, offset, orderId, clientOrderId)
	return &historyEndpoint{
		AuthParams:          whitebit.NewAuthParams(OrderHistoryEndpointUrl),
		ParamsForFiltration: optionsForFiltration,
	}
}
