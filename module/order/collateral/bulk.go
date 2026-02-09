package collateral

import "github.com/whitebit-exchange/go-sdk"

const bulkOrderEndpointURL = "/api/v4/order/collateral/bulk"

type BulkOrderParams struct {
	Market        string `json:"market"`
	Side          string `json:"side"`
	Amount        string `json:"amount"`
	Price         string `json:"price,omitempty"`
	Activation    string `json:"activation,omitempty"`
	ClientOrderID string `json:"clientOrderId,omitempty"`
}

type bulkOrderEndpoint struct {
	whitebit.AuthParams
	Orders []BulkOrderParams `json:"orders"`
}

func newBulkOrderEndpoint(orders []BulkOrderParams) *bulkOrderEndpoint {
	return &bulkOrderEndpoint{
		AuthParams: whitebit.NewAuthParams(bulkOrderEndpointURL),
		Orders:     orders,
	}
}

type BulkOrderResult struct {
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
	Price         string  `json:"price"`
}

type BulkOrderResponse struct {
	Result BulkOrderResult `json:"result,omitempty"`
	Error  *BulkOrderError `json:"error,omitempty"`
}

type BulkOrderError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
