package collateral

const cancelOcoEndpointUrl = "/api/v4/order/oco-cancel"

type OcoCancelOrder struct {
	ID         int64 `json:"id"`
	StopLoss   `json:"stop_loss"`
	TakeProfit `json:"take_profit"`
}

type StopLoss struct {
	MarketOrder
	Mtime               float64 `json:"mtime"`
	Price               string  `json:"price"`
	ActivationPrice     string  `json:"activation_price"`
	ActivationCondition string  `json:"activation_condition"`
	Activated           int     `json:"activated"`
}

type TakeProfit struct {
	MarketOrder
	Mtime float64 `json:"mtime"`
	Price string  `json:"price"`
}

type ocoCancelOrderEndpoint struct {
	cancelOrderEndpoint
}

func newOcoCancelOrderEndpoint(market string, orderId int64) *ocoCancelOrderEndpoint {
	cancelOrderOptions := newCancelOrderEndpoint(market, orderId)
	cancelOrderOptions.AuthParams.Request = cancelOcoEndpointUrl

	return &ocoCancelOrderEndpoint{cancelOrderEndpoint: *cancelOrderOptions}
}
