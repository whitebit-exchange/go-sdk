package spot

type OcoOrder struct {
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

type OcoOrderParams struct {
	Market          string `json:"market"`
	Amount          string `json:"amount"`
	Side            string `json:"side"`
	ActivationPrice string `json:"activation_price"`
	StopLimitPrice  string `json:"stop_limit_price"`
	ClientOrderId   string `json:"clientOrderId,omitempty"`
}
