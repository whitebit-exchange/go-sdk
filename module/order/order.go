package order

const (
	SideBuy  = "buy"
	SideSell = "sell"
)

type CreateParams struct {
	Market        string `json:"market"`
	Amount        string `json:"amount"`
	Side          string `json:"side"`
	ClientOrderId string `json:"clientOrderId,omitempty"`
}

func NewCreateParams(market string, amount string, side string, clientOrderId string) CreateParams {
	return CreateParams{Market: market, Amount: amount, Side: side, ClientOrderId: clientOrderId}
}
