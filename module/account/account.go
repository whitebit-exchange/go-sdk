package account

type LimitOffsetParams struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

func NewLimitOffsetParams(limit int, offset int) LimitOffsetParams {
	return LimitOffsetParams{Limit: limit, Offset: offset}
}

type MarketWithPaginationParams struct {
	LimitOffsetParams

	Market string `json:"market,omitempty"`
}

func NewMarketWithPaginationParams(market string, limit int, offset int) MarketWithPaginationParams {
	return MarketWithPaginationParams{Market: market, LimitOffsetParams: NewLimitOffsetParams(limit, offset)}
}

type ParamsForFiltration struct {
	MarketWithPaginationParams

	OrderId       int    `json:"orderId,omitempty"`
	ClientOrderId string `json:"clientOrderId,omitempty"`
}

func NewParamsForFiltration(market string, limit int, offset int, orderId int, clientOrderId string) ParamsForFiltration {
	options := NewMarketWithPaginationParams(market, limit, offset)
	return ParamsForFiltration{MarketWithPaginationParams: options, OrderId: orderId, ClientOrderId: clientOrderId}
}
