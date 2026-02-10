package stream

import (
	"encoding/json"
)

// Order type constants.
const (
	TypeLimit                   = 1
	TypeMarket                  = 2
	TypeStopLimit               = 3
	TypeStopMarket              = 4
	TypeConditionalLimit        = 5
	TypeConditionalMarket       = 6
	TypeMarginMarket            = 8
	TypeMarginStopLimit         = 9
	TypeMarginTriggerStopMarket = 10
	TypeCollateralNormalization = 14
	TypeMarketStock             = 202
)

// WebSocket method constants.
const (
	PingRequest = "ping"
	TimeRequest = "time"

	KlineRequest     = "candles_request"
	KlineSubscribe   = "candles_subscribe"
	KlineUpdate      = "candles_update"
	KlineUnsubscribe = "candles_unsubscribe"

	DepthRequest     = "depth_request"
	DepthSubscribe   = "depth_subscribe"
	DepthUpdate      = "depth_update"
	DepthUnsubscribe = "depth_unsubscribe"

	LastPriceRequest     = "lastprice_request"
	LastPriceSubscribe   = "lastprice_subscribe"
	LastPriceUpdate      = "lastprice_update"
	LastPriceUnsubscribe = "lastprice_unsubscribe"

	MarketStatRequest     = "market_request"
	MarketStatSubscribe   = "market_subscribe"
	MarketStatUpdate      = "market_update"
	MarketStatUnsubscribe = "market_unsubscribe"

	MarketStatTodayRequest     = "marketToday_query"
	MarketStatTodaySubscribe   = "marketToday_subscribe"
	MarketStatTodayUpdate      = "marketToday_update"
	MarketStatTodayUnsubscribe = "marketToday_unsubscribe"

	TradesRequest     = "trades_request"
	TradesSubscribe   = "trades_subscribe"
	TradesUpdate      = "trades_update"
	TradesUnsubscribe = "trades_unsubscribe"

	OrderPendingRequest      = "ordersPending_request"
	OrdersPendingSubscribe   = "ordersPending_subscribe"
	OrdersPendingUpdate      = "ordersPending_update"
	OrdersPendingUnsubscribe = "ordersPending_unsubscribe"

	DealsRequest     = "deals_request"
	DealsSubscribe   = "deals_subscribe"
	DealsUpdate      = "deals_update"
	DealsUnsubscribe = "deals_unsubscribe"

	SpotBalanceRequest     = "balanceSpot_request"
	SpotBalanceSubscribe   = "balanceSpot_subscribe"
	SpotBalanceUpdate      = "balanceSpot_update"
	SpotBalanceUnsubscribe = "balanceSpot_unsubscribe"

	OrdersExecutedRequest     = "ordersExecuted_request"
	OrdersExecutedSubscribe   = "ordersExecuted_subscribe"
	OrdersExecutedUpdate      = "ordersExecuted_update"
	OrdersExecutedUnsubscribe = "ordersExecuted_unsubscribe"

	MarginBalanceRequest     = "balanceMargin_request"
	MarginBalanceSubscribe   = "balanceMargin_subscribe"
	MarginBalanceUpdate      = "balanceMargin_update"
	MarginBalanceUnsubscribe = "balanceMargin_unsubscribe"

	WsTypeQuery     = true
	WsTypeSubscribe = false
)

// Command represents a WebSocket command to send to the server.
type Command struct {
	ID      int64  `json:"id"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	IsQuery bool   `json:"-"`
}

// CommandReply represents a response to a Command from the server.
type CommandReply struct {
	ID     int `json:"id"`
	Error  any `json:"error"`
	Result any `json:"result"`
}

// Subscription represents a WebSocket subscription with an event handler.
type Subscription struct {
	Command
	EventMethod       string
	OnEvent           func(event Event)
	UnsubscribeMethod Command
	onError           func(err error)
}

// Response represents a basic WebSocket response.
type Response struct {
	ID    int `json:"id"`
	Error any `json:"error"`
}

// ResponsePagination contains pagination info from paginated responses.
type ResponsePagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

// TransformEvent converts a generic Event into a typed event struct.
func TransformEvent[E any](event Event) (E, error) {
	var updateEvent E
	response, _ := json.Marshal(event)
	err := json.Unmarshal(response, &updateEvent)
	if err != nil {
		return *new(E), err
	}

	return updateEvent, err
}

// NewAuthorizeCommand creates an authorization command with the given token.
func NewAuthorizeCommand(token string) Command {
	return Command{
		ID:      1,
		Method:  "authorize",
		Params:  []any{token, "go-sdk"},
		IsQuery: WsTypeQuery,
	}
}

// NewStreamCommand creates a generic stream command.
func NewStreamCommand(id int64, method string, params []any) Command {
	return Command{ID: id, Method: method, Params: params}
}

// NewUnsubscribeCommand creates a generic unsubscribe command.
func NewUnsubscribeCommand(method string) Command {
	return Command{ID: 0, Method: method, Params: []any{}}
}

// AssetBalance represents the balance of an asset.
type AssetBalance struct {
	Available string `json:"available"`
	Freeze    string `json:"freeze"`
}

// SpotBalance represents a spot balance response.
type SpotBalance struct {
	Response
	Result map[string]AssetBalance
}

// SpotBalanceUpdateEvent represents a spot balance update event.
type SpotBalanceUpdateEvent struct {
	Balances []map[string]AssetBalance `json:"params"`
}

// NewSpotBalanceSubscription creates a subscription for spot balance updates.
func NewSpotBalanceSubscription(handle func(event SpotBalanceUpdateEvent), assets ...string) *Subscription {
	params := make([]any, len(assets))
	for i := range assets {
		params[i] = assets[i]
	}

	sub := &Subscription{
		Command:           Command{0, SpotBalanceSubscribe, params, WsTypeSubscribe},
		EventMethod:       SpotBalanceUpdate,
		UnsubscribeMethod: NewSpotBalanceUnsubscribe(),
	}

	sub.OnEvent = func(event Event) {
		updateEvent, err := TransformEvent[SpotBalanceUpdateEvent](event)
		if err != nil {
			if sub.onError != nil {
				sub.onError(err)
			}
			return
		}
		handle(updateEvent)
	}

	return sub
}

// NewSpotBalanceUnsubscribe creates a spot balance unsubscribe command.
func NewSpotBalanceUnsubscribe() Command {
	return Command{0, SpotBalanceUnsubscribe, []any{}, WsTypeQuery}
}

// NewSpotBalanceCommand creates a spot balance query command.
func NewSpotBalanceCommand(assets []string) Command {
	params := make([]any, len(assets))
	for i := range assets {
		params[i] = assets[i]
	}
	return Command{1, SpotBalanceRequest, params, WsTypeQuery}
}

// MarginBalanceResponse represents a margin balance response.
type MarginBalanceResponse struct {
	Response
	Result map[string]string
}

// MarginBalanceUpdateEvent represents a margin balance update event.
type MarginBalanceUpdateEvent struct {
	Balances []MarginBalanceEventInfo `json:"params"`
}

// MarginBalanceEventInfo contains margin balance event details.
type MarginBalanceEventInfo struct {
	Balance                string `json:"B"`
	Asset                  string `json:"a"`
	AvailableWithBorrow    string `json:"ab"`
	AvailableWithoutBorrow string `json:"av"`
	Borrow                 string `json:"b"`
}

// MarginBalanceQueryInfo contains margin balance query details.
type MarginBalanceQueryInfo struct {
	AvailableWithBorrow    string `json:"available_with_borrow"`
	AvailableWithoutBorrow string `json:"available_without_borrow"`
	Balance                string `json:"balance"`
	Borrow                 string `json:"borrow"`
}

// MarginBalance is a map of asset name to margin balance info.
type MarginBalance map[string]MarginBalanceQueryInfo

// NewMarginBalanceSubscription creates a subscription for margin balance updates.
func NewMarginBalanceSubscription(handle func(event MarginBalanceUpdateEvent), assets ...string) *Subscription {
	params := make([]any, len(assets))
	for i := range assets {
		params[i] = assets[i]
	}

	sub := &Subscription{
		Command:           Command{0, MarginBalanceSubscribe, params, WsTypeSubscribe},
		EventMethod:       MarginBalanceUpdate,
		UnsubscribeMethod: NewMarginBalanceUnsubscribe(),
	}

	sub.OnEvent = func(event Event) {
		updateEvent, err := TransformEvent[MarginBalanceUpdateEvent](event)
		if err != nil {
			if sub.onError != nil {
				sub.onError(err)
			}
			return
		}
		handle(updateEvent)
	}

	return sub
}

// NewMarginBalanceUnsubscribe creates a margin balance unsubscribe command.
func NewMarginBalanceUnsubscribe() Command {
	return Command{0, MarginBalanceUnsubscribe, []any{}, WsTypeQuery}
}

// NewMarginBalanceCommand creates a margin balance query command.
func NewMarginBalanceCommand(assets []string) Command {
	params := make([]any, len(assets))
	for i := range assets {
		params[i] = assets[i]
	}
	return Command{1, MarginBalanceRequest, params, WsTypeQuery}
}

// OrdersRecords contains order details.
type OrdersRecords struct {
	ID            int     `json:"id"`
	Market        string  `json:"market"`
	Type          int     `json:"type"`
	Side          int     `json:"side"`
	PostOnly      bool    `json:"post_only"`
	IOC           bool    `json:"ioc"`
	Ctime         float64 `json:"ctime"`
	Mtime         float64 `json:"mtime"`
	Price         string  `json:"price"`
	Amount        string  `json:"amount"`
	Left          string  `json:"left"`
	DealStock     string  `json:"deal_stock"`
	DealMoney     string  `json:"deal_money"`
	DealFee       string  `json:"deal_fee"`
	ClientOrderID string  `json:"client_order_id"`
}

// Orders contains a paginated list of order records.
type Orders struct {
	ResponsePagination
	Records []OrdersRecords `json:"records"`
}

// OrdersExecuted contains executed order records with pagination.
type OrdersExecuted struct {
	Limit   int             `json:"limit"`
	Offset  int             `json:"offset"`
	Records []OrdersRecords `json:"records"`
}

// PendingOrders represents a pending orders response.
type PendingOrders struct {
	Response
	Result Orders `json:"result"`
}

// ExecutedOrders represents an executed orders response.
type ExecutedOrders struct {
	Response
	Result Orders `json:"result"`
}

// Pending represents a pending order.
type Pending struct {
	ID                   int64
	PendingOrdersRecords OrdersRecords
}

// PendingOrdersUpdateEvent represents a pending orders update event.
type PendingOrdersUpdateEvent struct {
	Balances []any `json:"params"`
}

// NewPendingOrdersSubscription creates a subscription for pending order updates.
func NewPendingOrdersSubscription(handle func(event PendingOrdersUpdateEvent), market ...string) *Subscription {
	params := make([]any, len(market))
	for i := range market {
		params[i] = market[i]
	}

	sub := &Subscription{
		Command:           Command{0, OrdersPendingSubscribe, params, WsTypeSubscribe},
		EventMethod:       OrdersPendingUpdate,
		UnsubscribeMethod: NewOrdersPendingUnsubscribe(),
	}

	sub.OnEvent = func(event Event) {
		updateEvent, err := TransformEvent[PendingOrdersUpdateEvent](event)
		if err != nil {
			if sub.onError != nil {
				sub.onError(err)
			}
			return
		}
		handle(updateEvent)
	}

	return sub
}

// NewOrdersPendingUnsubscribe creates a pending orders unsubscribe command.
func NewOrdersPendingUnsubscribe() Command {
	return Command{0, OrdersPendingUnsubscribe, []any{}, WsTypeQuery}
}

// NewPendingOrdersCommand creates a pending orders query command.
func NewPendingOrdersCommand(market string, offset int64, limit int64) Command {
	return Command{
		ID:      1,
		Method:  OrderPendingRequest,
		Params:  []any{market, offset, limit},
		IsQuery: WsTypeQuery,
	}
}

// OrderExecutedUpdateEvent represents an executed order update event.
type OrderExecutedUpdateEvent struct {
	Params []OrdersRecords `json:"params"`
}

// NewOrderExecutedSubscription creates a subscription for executed order updates.
func NewOrderExecutedSubscription(handle func(event OrderExecutedUpdateEvent), market []string, filter int64) *Subscription {
	sub := &Subscription{
		Command:           Command{0, OrdersExecutedSubscribe, []any{market, filter}, WsTypeSubscribe},
		EventMethod:       OrdersExecutedUpdate,
		UnsubscribeMethod: NewOrdersExecutedUnsubscribe(),
	}

	sub.OnEvent = func(event Event) {
		updateEvent, err := TransformEvent[OrderExecutedUpdateEvent](event)
		if err != nil {
			if sub.onError != nil {
				sub.onError(err)
			}
			return
		}
		handle(updateEvent)
	}

	return sub
}

// NewOrdersExecutedUnsubscribe creates an executed orders unsubscribe command.
func NewOrdersExecutedUnsubscribe() Command {
	return Command{0, OrdersExecutedUnsubscribe, []any{}, WsTypeQuery}
}

// NewOrdersExecutedCommand creates an executed orders query command.
func NewOrdersExecutedCommand(market string, orderTypes []int64, limit int64, offset int64) Command {
	return Command{
		ID:     1,
		Method: OrdersExecutedRequest,
		Params: []any{
			map[string]any{"market": market, "order_types": orderTypes},
			offset,
			limit},
		IsQuery: WsTypeQuery,
	}
}

// Deal represents a single deal/trade.
type Deal struct {
	Time        float64 `json:"time"`
	ID          int     `json:"id"`
	Side        int     `json:"side"`
	Role        int     `json:"role"`
	Price       string  `json:"price"`
	Amount      string  `json:"amount"`
	Deal        string  `json:"deal"`
	Fee         string  `json:"fee"`
	Market      string  `json:"market"`
	DealOrderID int64   `json:"deal_order_id"`
}

// DealsResult contains a paginated list of deals.
type DealsResult struct {
	ResponsePagination
	Records []Deal `json:"records"`
}

// Deals represents a deals response.
type Deals struct {
	Response
	Result DealsResult `json:"result"`
}

// DealsUpdateEvent represents a deals update event.
type DealsUpdateEvent struct {
	Params []any `json:"params"`
}

// NewDealsSubscription creates a subscription for deal updates.
func NewDealsSubscription(handle func(event DealsUpdateEvent), market []string) *Subscription {
	sub := &Subscription{
		Command:           Command{0, DealsSubscribe, []any{market}, WsTypeSubscribe},
		EventMethod:       DealsUpdate,
		UnsubscribeMethod: NewDealsUnsubscribe(),
	}

	sub.OnEvent = func(event Event) {
		updateEvent, err := TransformEvent[DealsUpdateEvent](event)
		if err != nil {
			if sub.onError != nil {
				sub.onError(err)
			}
			return
		}
		handle(updateEvent)
	}

	return sub
}

// NewDealsUnsubscribe creates a deals unsubscribe command.
func NewDealsUnsubscribe() Command {
	return Command{0, DealsUnsubscribe, []any{}, WsTypeQuery}
}

// NewDealsCommand creates a deals query command.
func NewDealsCommand(market string, offset int64, limit int64) Command {
	return Command{
		ID:      1,
		Method:  DealsRequest,
		Params:  []any{market, offset, limit},
		IsQuery: WsTypeQuery,
	}
}

// Ping represents a ping response.
type Ping struct {
	Response
	Result string `json:"result"`
}

// NewPingCommand creates a ping command.
func NewPingCommand() Command {
	return Command{
		ID:      1,
		Method:  PingRequest,
		Params:  make([]any, 0),
		IsQuery: WsTypeQuery,
	}
}

// Time represents a time response.
type Time struct {
	Response
	Result int64 `json:"result"`
}

// NewTimeCommand creates a server time query command.
func NewTimeCommand() Command {
	return Command{
		ID:      1,
		Method:  TimeRequest,
		Params:  make([]any, 0),
		IsQuery: WsTypeQuery,
	}
}

// KlineResult represents a single kline/candlestick data point.
type KlineResult struct {
	Time        int64
	Open        string
	Close       string
	Highest     string
	Lowest      string
	StockVolume string
	DealVolume  string
	Market      string
}

// Kline represents a kline response.
type Kline struct {
	Response
	Result []KlineResult `json:"result"`
}

// KlineUpdateEvent represents a kline update event.
type KlineUpdateEvent struct {
	Params []any `json:"params"`
}

// NewKlineSubscription creates a subscription for kline/candlestick updates.
func NewKlineSubscription(handle func(event KlineUpdateEvent), market string, interval int64) *Subscription {
	sub := &Subscription{
		Command:           Command{0, KlineSubscribe, []any{market, interval}, WsTypeSubscribe},
		EventMethod:       KlineUpdate,
		UnsubscribeMethod: NewKlineUnsubscribe(),
	}

	sub.OnEvent = func(event Event) {
		updateEvent, err := TransformEvent[KlineUpdateEvent](event)
		if err != nil {
			if sub.onError != nil {
				sub.onError(err)
			}
			return
		}
		handle(updateEvent)
	}

	return sub
}

// NewKlineUnsubscribe creates a kline unsubscribe command.
func NewKlineUnsubscribe() Command {
	return Command{0, KlineUnsubscribe, []any{}, WsTypeQuery}
}

// NewKlineCommand creates a kline query command.
func NewKlineCommand(market string, startTime int64, endTime int64, interval int64) Command {
	return Command{
		ID:      1,
		Method:  KlineRequest,
		Params:  []any{market, startTime, endTime, interval},
		IsQuery: WsTypeQuery,
	}
}

// LastPrice represents a last price response.
type LastPrice struct {
	Response
	Result string `json:"result"`
}

// LastPriceUpdateEvent represents a last price update event.
type LastPriceUpdateEvent struct {
	Params any `json:"params"`
}

// NewLastPriceSubscription creates a subscription for last price updates.
func NewLastPriceSubscription(handle func(event LastPriceUpdateEvent), market []string) *Subscription {
	params := make([]any, len(market))
	for i := range market {
		params[i] = market[i]
	}

	sub := &Subscription{
		Command:           Command{0, LastPriceSubscribe, params, WsTypeSubscribe},
		EventMethod:       LastPriceUpdate,
		UnsubscribeMethod: NewLastPriceUnsubscribe(),
	}

	sub.OnEvent = func(event Event) {
		updateEvent, err := TransformEvent[LastPriceUpdateEvent](event)
		if err != nil {
			if sub.onError != nil {
				sub.onError(err)
			}
			return
		}
		handle(updateEvent)
	}

	return sub
}

// NewLastPriceUnsubscribe creates a last price unsubscribe command.
func NewLastPriceUnsubscribe() Command {
	return Command{0, LastPriceUnsubscribe, []any{}, WsTypeQuery}
}

// NewLastPriceCommand creates a last price query command.
func NewLastPriceCommand(market string) Command {
	return Command{
		ID:      1,
		Method:  LastPriceRequest,
		Params:  []any{market},
		IsQuery: WsTypeQuery,
	}
}

// MarketStatResult contains market statistics data.
type MarketStatResult struct {
	Period int    `json:"period"`
	Last   string `json:"last"`
	Open   string `json:"open"`
	Close  string `json:"close"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Volume string `json:"volume"`
	Deal   string `json:"deal"`
}

// MarketStat represents a market statistics response.
type MarketStat struct {
	Response
	Result MarketStatResult `json:"result"`
}

// MarketStatUpdateEvent represents a market statistics update event.
type MarketStatUpdateEvent struct {
	Params []any `json:"params"`
}

// NewMarketStatSubscription creates a subscription for market statistics updates.
func NewMarketStatSubscription(handle func(MarketStatUpdateEvent), market []string) *Subscription {
	params := make([]any, len(market))
	for i := range market {
		params[i] = market[i]
	}

	sub := &Subscription{
		Command:           Command{0, MarketStatSubscribe, params, WsTypeSubscribe},
		EventMethod:       MarketStatUpdate,
		UnsubscribeMethod: NewMarketStatUnsubscribe(),
	}

	sub.OnEvent = func(event Event) {
		updateEvent, err := TransformEvent[MarketStatUpdateEvent](event)
		if err != nil {
			if sub.onError != nil {
				sub.onError(err)
			}
			return
		}
		handle(updateEvent)
	}

	return sub
}

// NewMarketStatUnsubscribe creates a market statistics unsubscribe command.
func NewMarketStatUnsubscribe() Command {
	return Command{0, MarketStatUnsubscribe, []any{}, WsTypeQuery}
}

// NewMarketStatCommand creates a market statistics query command.
func NewMarketStatCommand(market string, period int64) Command {
	return Command{
		ID:      1,
		Method:  MarketStatRequest,
		Params:  []any{market, period},
		IsQuery: WsTypeQuery,
	}
}

// MarketStatTodayUpdateEvent represents a market statistics today update event.
type MarketStatTodayUpdateEvent struct {
	Params []any `json:"params"`
}

// NewMarketStatTodaySubscription creates a subscription for today's market statistics updates.
func NewMarketStatTodaySubscription(handle func(MarketStatTodayUpdateEvent), market []string) *Subscription {
	params := make([]any, len(market))
	for i := range market {
		params[i] = market[i]
	}

	sub := &Subscription{
		Command:           Command{0, MarketStatTodaySubscribe, params, WsTypeSubscribe},
		EventMethod:       MarketStatTodayUpdate,
		UnsubscribeMethod: NewMarketStatTodayUnsubscribe(),
	}

	sub.OnEvent = func(event Event) {
		updateEvent, err := TransformEvent[MarketStatTodayUpdateEvent](event)
		if err != nil {
			if sub.onError != nil {
				sub.onError(err)
			}
			return
		}
		handle(updateEvent)
	}

	return sub
}

// NewMarketStatTodayUnsubscribe creates a market statistics today unsubscribe command.
func NewMarketStatTodayUnsubscribe() Command {
	return Command{0, MarketStatTodayUnsubscribe, []any{}, WsTypeQuery}
}

// NewMarketStatTodayCommand creates a market statistics today query command.
func NewMarketStatTodayCommand(market string) Command {
	return Command{
		ID:      1,
		Method:  MarketStatTodayRequest,
		Params:  []any{market},
		IsQuery: WsTypeQuery,
	}
}

// MarketTradesResult contains a single market trade.
type MarketTradesResult struct {
	ID     int     `json:"id"`
	Time   float64 `json:"time"`
	Price  string  `json:"price"`
	Amount string  `json:"amount"`
	Type   string  `json:"type"`
}

// MarketTrades represents a market trades response.
type MarketTrades struct {
	Response
	Result MarketTradesResult `json:"result"`
}

// MarketTradesUpdateEvent represents a market trades update event.
type MarketTradesUpdateEvent struct {
	Params []any `json:"params"`
}

// NewMarketTradesSubscription creates a subscription for market trade updates.
func NewMarketTradesSubscription(handle func(MarketTradesUpdateEvent), market []string) *Subscription {
	params := make([]any, len(market))
	for i := range market {
		params[i] = market[i]
	}

	sub := &Subscription{
		Command:           Command{0, TradesSubscribe, params, WsTypeSubscribe},
		EventMethod:       TradesUpdate,
		UnsubscribeMethod: NewMarketTradesUnsubscribe(),
	}

	sub.OnEvent = func(event Event) {
		updateEvent, err := TransformEvent[MarketTradesUpdateEvent](event)
		if err != nil {
			if sub.onError != nil {
				sub.onError(err)
			}
			return
		}
		handle(updateEvent)
	}

	return sub
}

// NewMarketTradesUnsubscribe creates a market trades unsubscribe command.
func NewMarketTradesUnsubscribe() Command {
	return Command{0, TradesUnsubscribe, []any{}, WsTypeQuery}
}

// NewMarketTradesCommand creates a market trades query command.
func NewMarketTradesCommand(market string, limit int64, startTradeID int64) Command {
	return Command{
		ID:      1,
		Method:  TradesRequest,
		Params:  []any{market, limit, startTradeID},
		IsQuery: WsTypeQuery,
	}
}

// Pair represents a price-amount pair in the order book.
type Pair [2]string

// AsksAndBids contains asks and bids from the order book.
type AsksAndBids struct {
	Asks []Pair `json:"asks"`
	Bids []Pair `json:"bids"`
}

// MarketDepth represents a market depth response.
type MarketDepth struct {
	Response
	Result []AsksAndBids `json:"result"`
}

// MarketDepthUpdateEvent represents a market depth update event.
type MarketDepthUpdateEvent struct {
	Params []any `json:"params"`
}

// NewMarketDepthSubscription creates a subscription for market depth updates.
func NewMarketDepthSubscription(handle func(MarketDepthUpdateEvent), market string, limit int64, priceInterval string, multiSubscription bool) *Subscription {
	sub := &Subscription{
		Command:           Command{0, DepthSubscribe, []any{market, limit, priceInterval, multiSubscription}, WsTypeSubscribe},
		EventMethod:       DepthUpdate,
		UnsubscribeMethod: NewDepthUnsubscribe(),
	}

	sub.OnEvent = func(event Event) {
		updateEvent, err := TransformEvent[MarketDepthUpdateEvent](event)
		if err != nil {
			if sub.onError != nil {
				sub.onError(err)
			}
			return
		}
		handle(updateEvent)
	}

	return sub
}

// NewDepthUnsubscribe creates a market depth unsubscribe command.
func NewDepthUnsubscribe() Command {
	return Command{0, DepthUnsubscribe, []any{}, WsTypeQuery}
}

// NewMarketDepthCommand creates a market depth query command.
func NewMarketDepthCommand(market string, limit int64, priceInterval string) Command {
	return Command{
		ID:      1,
		Method:  DepthRequest,
		Params:  []any{market, limit, priceInterval},
		IsQuery: WsTypeQuery,
	}
}
