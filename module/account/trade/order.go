package trade

import (
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/module/account"
)

const orderEndpointUrl = "/api/v4/trade-account/order"

type Order struct {
	Time          float64 `json:"time"`
	Fee           string  `json:"fee"`
	Price         string  `json:"price"`
	Amount        string  `json:"amount"`
	Id            int64   `json:"id"`
	ClientOrderId string  `json:"clientOrderId"`
	DealOrderId   int     `json:"dealOrderId"`
	Role          int     `json:"role"`
	Deal          string  `json:"deal"`
}

type orderEndpoint struct {
	whitebit.AuthParams
	account.LimitOffsetParams

	OrderId int64 `json:"orderId"`
}

func newOrderEndpoint(orderId int64, limit int, offset int) *orderEndpoint {
	return &orderEndpoint{
		AuthParams:        whitebit.NewAuthParams(orderEndpointUrl),
		LimitOffsetParams: account.NewLimitOffsetParams(limit, offset),
		OrderId:           orderId,
	}
}
