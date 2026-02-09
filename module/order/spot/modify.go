package spot

import "github.com/whitebit-exchange/go-sdk"

const modifyOrderEndpointURL = "/api/v4/order/modify"

type ModifyOrderParams struct {
	OrderID       int64  `json:"orderId"`
	Market        string `json:"market"`
	Price         string `json:"price,omitempty"`
	Amount        string `json:"amount,omitempty"`
	ClientOrderID string `json:"clientOrderId,omitempty"`
}

type modifyOrderEndpoint struct {
	whitebit.AuthParams
	ModifyOrderParams
}

func newModifyOrderEndpoint(params ModifyOrderParams) *modifyOrderEndpoint {
	return &modifyOrderEndpoint{
		AuthParams:        whitebit.NewAuthParams(modifyOrderEndpointURL),
		ModifyOrderParams: params,
	}
}

type ModifiedOrder struct {
	OrderID       int64   `json:"orderId"`
	ClientOrderID string  `json:"clientOrderId"`
	Market        string  `json:"market"`
	Side          string  `json:"side"`
	Type          string  `json:"type"`
	Timestamp     float64 `json:"timestamp"`
	DealMoney     string  `json:"dealMoney"`
	DealStock     string  `json:"dealStock"`
	Amount        string  `json:"amount"`
	TakerFee      string  `json:"takerFee"`
	MakerFee      string  `json:"makerFee"`
	Left          string  `json:"left"`
	DealFee       string  `json:"dealFee"`
	PostOnly      bool    `json:"postOnly"`
	IOC           bool    `json:"ioc"`
	Price         string  `json:"price"`
}
