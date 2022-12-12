package spot

import (
	"github.com/whitebit-exchange/whitebit"
)

const cancelOrderEndpointUrl = "/api/v4/order/cancel"

type CancelOrder struct {
	OrderID         int64   `json:"orderId"`
	ClientOrderID   string  `json:"clientOrderId"`
	Market          string  `json:"market"`
	Side            string  `json:"side"`
	Type            string  `json:"type"`
	Timestamp       float64 `json:"timestamp"`
	DealMoney       string  `json:"dealMoney"`
	DealStock       string  `json:"dealStock"`
	Amount          string  `json:"amount"`
	TakerFee        string  `json:"takerFee"`
	MakerFee        string  `json:"makerFee"`
	Left            string  `json:"left"`
	DealFee         string  `json:"dealFee"`
	Price           string  `json:"price"`
	ActivationPrice string  `json:"activation_price"`
}

type cancelOrderEndpoint struct {
	whitebit.AuthParams

	Market  string `json:"market"`
	OrderId int64  `json:"orderId"`
}

func newCancelOrderEndpoint(market string, orderId int64) *cancelOrderEndpoint {
	return &cancelOrderEndpoint{
		AuthParams: whitebit.NewAuthParams(cancelOrderEndpointUrl),
		Market:     market,
		OrderId:    orderId,
	}
}
