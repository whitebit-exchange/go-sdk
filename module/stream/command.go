package stream

import (
	"encoding/json"
	"fmt"
)

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

type Command struct {
	Id      int64         `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	IsQuery bool          `json:"-"`
}

type CommandReply struct {
	Id     int         `json:"id"`
	Error  interface{} `json:"error"`
	Result any         `json:"result"`
}

type Subscription struct {
	Command
	EventMethod       string
	OnEvent           func(event Event)
	UnsubscribeMethod Command
}

type Response struct {
	Id    int         `json:"id"`
	Error interface{} `json:"error"`
}

type ResponsePagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

func TransformEvent[E any](event Event) (E, error) {
	var updateEvent E
	response, _ := json.Marshal(event)
	err := json.Unmarshal(response, &updateEvent)
	if err != nil {
		return *new(E), err
	}

	return updateEvent, err
}

func NewAuthorizeCommand(token string) Command {
	return Command{
		Id:      1,
		Method:  "authorize",
		Params:  []interface{}{token, "go-sdk"},
		IsQuery: WsTypeQuery,
	}
}

func NewStreamCommand(id int64, method string, params []interface{}) Command {
	return Command{Id: id, Method: method, Params: params}
}

func NewUnsubscribeCommand(method string) Command {
	return Command{Id: 0, Method: method, Params: []interface{}{}}
}

type AssetBalance struct {
	Available string `json:"available"`
	Freeze    string `json:"freeze"`
}

type SpotBalance struct {
	Response
	Result map[string]AssetBalance
}

type SpotBalanceUpdateEvent struct {
	Balances []map[string]AssetBalance `json:"params"`
}

func NewSpotBalanceSubscription(handle func(event SpotBalanceUpdateEvent), assets ...string) *Subscription {
	assetsAsInterface := make([]interface{}, len(assets))
	for i := range assets {
		assetsAsInterface[i] = assets[i]
	}

	onEvent := func(event Event) {
		updateEvent, err := TransformEvent[SpotBalanceUpdateEvent](event)
		if err != nil {
			fmt.Println(err)
		}

		handle(updateEvent)
	}

	return &Subscription{Command{0, SpotBalanceSubscribe, assetsAsInterface, WsTypeSubscribe},
		SpotBalanceUpdate, onEvent, NewSpotBalanceUnsubscribe()}
}

func NewSpotBalanceUnsubscribe() Command {
	return Command{0, SpotBalanceUnsubscribe, []any{}, WsTypeQuery}
}

func NewSpotBalanceCommand(assets []string) Command {
	assetsAsInterface := make([]interface{}, len(assets))
	for i := range assets {
		assetsAsInterface[i] = assets[i]
	}
	return Command{1, SpotBalanceRequest, assetsAsInterface, WsTypeQuery}
}

type MarginBalance struct {
	Response
	Result map[string]string
}

type MarginBalanceUpdateEvent struct {
	Balances []map[string]string `json:"params"`
}

func NewMarginBalanceSubscription(handle func(event MarginBalanceUpdateEvent), assets ...string) *Subscription {
	params := make([]interface{}, len(assets))
	for i := range assets {
		params[i] = assets[i]
	}

	onEvent := func(event Event) {
		updateEvent, err := TransformEvent[MarginBalanceUpdateEvent](event)
		if err != nil {
			fmt.Println(err)
		}
		handle(updateEvent)
	}

	return &Subscription{Command{0, MarginBalanceSubscribe, params, WsTypeSubscribe},
		MarginBalanceUpdate, onEvent, NewMarginBalanceUnsubscribe()}
}

func NewMarginBalanceUnsubscribe() Command {
	return Command{0, MarginBalanceUnsubscribe, []any{}, WsTypeQuery}
}

func NewMarginBalanceCommand(assets []string) Command {
	assetsAsInterface := make([]interface{}, len(assets))
	for i := range assets {
		assetsAsInterface[i] = assets[i]
	}
	return Command{1, MarginBalanceRequest, assetsAsInterface, WsTypeQuery}
}

type OrdersRecords struct {
	Id            int     `json:"id"`
	Market        string  `json:"market"`
	Type          int     `json:"type"`
	Side          int     `json:"side"`
	PostOnly      bool    `json:"post_only"`
	Ctime         float64 `json:"ctime"`
	Mtime         float64 `json:"mtime"`
	Price         string  `json:"price"`
	Amount        string  `json:"amount"`
	Left          string  `json:"left"`
	DealStock     string  `json:"deal_stock"`
	DealMoney     string  `json:"deal_money"`
	DealFee       string  `json:"deal_fee"`
	ClientOrderId string  `json:"client_order_id"`
}

type Orders struct {
	ResponsePagination
	Records []OrdersRecords `json:"records"`
}

type OrdersExecuted struct {
	Limit   int             `json:"limit"`
	Offset  int             `json:"offset"`
	Records []OrdersRecords `json:"records"`
}

type PendingOrders struct {
	Response
	Result Orders `json:"result"`
}

type ExecutedOrders struct {
	Response
	Result Orders `json:"result"`
}

type Pending struct {
	Id                   int64
	PendingOrdersRecords OrdersRecords
}

type PendingOrdersUpdateEvent struct {
	Balances []any `json:"params"`
}

func NewPendingOrdersSubscription(handle func(event PendingOrdersUpdateEvent), market ...string) *Subscription {
	params := make([]interface{}, len(market))
	for i := range market {
		params[i] = market[i]
	}

	onEvent := func(event Event) {
		updateEvent, err := TransformEvent[PendingOrdersUpdateEvent](event)
		if err != nil {
			fmt.Println(err)
		}

		handle(updateEvent)
	}

	return &Subscription{Command{0, OrdersPendingSubscribe, params, WsTypeSubscribe},
		OrdersPendingUpdate, onEvent, NewOrdersPendingUnsubscribe()}
}

func NewOrdersPendingUnsubscribe() Command {
	return Command{0, OrdersPendingUnsubscribe, []any{}, WsTypeQuery}
}

func NewPendingOrdersCommand(market string, offset int64, limit int64) Command {
	return Command{
		Id:      1,
		Method:  OrderPendingRequest,
		Params:  []interface{}{market, offset, limit},
		IsQuery: WsTypeQuery,
	}
}

type OrderExecutedUpdateEvent struct {
	Params []OrdersRecords `json:"params"`
}

func NewOrderExecutedSubscription(handle func(event OrderExecutedUpdateEvent), market []string, filter int64) *Subscription {
	onEvent := func(event Event) {
		updateEvent, err := TransformEvent[OrderExecutedUpdateEvent](event)
		if err != nil {
			fmt.Println(err)
		}

		handle(updateEvent)
	}

	return &Subscription{Command{0, OrdersExecutedSubscribe, []interface{}{market, filter}, WsTypeSubscribe},
		OrdersExecutedUpdate, onEvent, NewOrdersExecutedUnsubscribe()}
}

func NewOrdersExecutedUnsubscribe() Command {
	return Command{0, OrdersExecutedUnsubscribe, []any{}, WsTypeQuery}
}

func NewOrdersExecutedCommand(market string, orderTypes []int64, limit int64, offset int64) Command {
	return Command{
		Id:     1,
		Method: OrdersExecutedRequest,
		Params: []interface{}{
			map[string]interface{}{"market": market, "order_types": orderTypes},
			offset,
			limit},
		IsQuery: WsTypeQuery,
	}
}

type Deal struct {
	Time        float64 `json:"time"`
	Id          int     `json:"id"`
	Side        int     `json:"side"`
	Role        int     `json:"role"`
	Price       string  `json:"price"`
	Amount      string  `json:"amount"`
	Deal        string  `json:"deal"`
	Fee         string  `json:"fee"`
	Market      string  `json:"market"`
	DealOrderId int64   `json:"deal_order_id"`
}

type DealsResult struct {
	ResponsePagination
	Records []Deal `json:"records"`
}

type Deals struct {
	Response
	Result DealsResult `json:"result"`
}

type DealsUpdateEvent struct {
	Params []any `json:"params"`
}

func NewDealsSubscription(handle func(event DealsUpdateEvent), market []string) *Subscription {
	onEvent := func(event Event) {
		updateEvent, err := TransformEvent[DealsUpdateEvent](event)
		if err != nil {
			fmt.Println(err)
		}

		handle(updateEvent)
	}

	return &Subscription{Command{0, DealsSubscribe, []interface{}{market}, WsTypeSubscribe},
		DealsUpdate, onEvent, NewDealsUnsubscribe()}
}

func NewDealsUnsubscribe() Command {
	return Command{0, DealsUnsubscribe, []any{}, WsTypeQuery}
}

func NewDealsCommand(market string, offset int64, limit int64) Command {
	return Command{
		Id:      1,
		Method:  DealsRequest,
		Params:  []interface{}{market, offset, limit},
		IsQuery: WsTypeQuery,
	}
}

type Ping struct {
	Response
	Result string `json:"result"`
}

func NewPingCommand() Command {
	return Command{
		Id:      1,
		Method:  PingRequest,
		Params:  make([]interface{}, 0),
		IsQuery: WsTypeQuery,
	}
}

type Time struct {
	Response
	Result int64 `json:"result"`
}

func NewTimeCommand() Command {
	return Command{
		Id:      1,
		Method:  TimeRequest,
		Params:  make([]interface{}, 0),
		IsQuery: WsTypeQuery,
	}
}

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

type Kline struct {
	Response
	Result []KlineResult `json:"result"`
}

type KlineUpdateEvent struct {
	Params []any `json:"params"`
}

func NewKlineSubscription(handle func(event KlineUpdateEvent), market string, interval int64) *Subscription {
	onEvent := func(event Event) {
		updateEvent, err := TransformEvent[KlineUpdateEvent](event)
		if err != nil {
			fmt.Println(err)
		}

		handle(updateEvent)
	}

	return &Subscription{Command{0, KlineSubscribe, []interface{}{market, interval}, WsTypeSubscribe},
		KlineUpdate, onEvent, NewKlineUnsubscribe()}
}

func NewKlineUnsubscribe() Command {
	return Command{0, KlineUnsubscribe, []any{}, WsTypeQuery}
}

func NewKlineCommand(market string, startTime int64, endTime int64, interval int64) Command {
	return Command{
		Id:      1,
		Method:  KlineRequest,
		Params:  []interface{}{market, startTime, endTime, interval},
		IsQuery: WsTypeQuery,
	}
}

type LastPrice struct {
	Response
	Result string `json:"result"`
}

type LastPriceUpdateEvent struct {
	Params any `json:"params"`
}

func NewLastPriceSubscription(handle func(event LastPriceUpdateEvent), market []string) *Subscription {
	params := make([]interface{}, len(market))
	for i := range market {
		params[i] = market[i]
	}

	onEvent := func(event Event) {
		updateEvent, err := TransformEvent[LastPriceUpdateEvent](event)
		if err != nil {
			fmt.Println(err)
		}

		handle(updateEvent)
	}

	return &Subscription{Command{0, LastPriceSubscribe, params, WsTypeSubscribe},
		LastPriceUpdate, onEvent, NewLastPriceUnsubscribe()}
}

func NewLastPriceUnsubscribe() Command {
	return Command{0, LastPriceUnsubscribe, []any{}, WsTypeQuery}
}

func NewLastPriceCommand(market string) Command {
	return Command{
		Id:      1,
		Method:  LastPriceRequest,
		Params:  []interface{}{market},
		IsQuery: WsTypeQuery,
	}
}

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

type MarketStat struct {
	Response
	Result MarketStatResult `json:"result"`
}

type MarketStatUpdateEvent struct {
	Params []any `json:"params"`
}

func NewMarketStatSubscription(handle func(MarketStatUpdateEvent), market []string) *Subscription {
	params := make([]interface{}, len(market))
	for i := range market {
		params[i] = market[i]
	}

	onEvent := func(event Event) {
		updateEvent, err := TransformEvent[MarketStatUpdateEvent](event)
		if err != nil {
			fmt.Println(err)
		}

		handle(updateEvent)
	}

	return &Subscription{Command{0, MarketStatSubscribe, params, WsTypeSubscribe},
		MarketStatUpdate, onEvent, NewMarketStatUnsubscribe()}
}

func NewMarketStatUnsubscribe() Command {
	return Command{0, MarketStatUnsubscribe, []any{}, WsTypeQuery}
}

func NewMarketStatCommand(market string, period int64) Command {
	return Command{
		Id:      1,
		Method:  MarketStatRequest,
		Params:  []interface{}{market, period},
		IsQuery: WsTypeQuery,
	}
}

type MarketStatTodayUpdateEvent struct {
	Params []any `json:"params"`
}

func NewMarketStatTodaySubscription(handle func(MarketStatTodayUpdateEvent), market []string) *Subscription {
	params := make([]interface{}, len(market))
	for i := range market {
		params[i] = market[i]
	}

	onEvent := func(event Event) {
		updateEvent, err := TransformEvent[MarketStatTodayUpdateEvent](event)
		if err != nil {
			fmt.Println(err)
		}

		handle(updateEvent)
	}

	return &Subscription{Command{0, MarketStatTodaySubscribe, params, WsTypeSubscribe},
		MarketStatTodayUpdate, onEvent, NewMarketStatTodayUnsubscribe()}
}

func NewMarketStatTodayUnsubscribe() Command {
	return Command{0, MarketStatTodayUnsubscribe, []any{}, WsTypeQuery}
}

func NewMarketStatTodayCommand(market string) Command {
	return Command{
		Id:      1,
		Method:  MarketStatTodayRequest,
		Params:  []interface{}{market},
		IsQuery: WsTypeQuery,
	}
}

type MarketTradesResult struct {
	Id     int     `json:"id"`
	Time   float64 `json:"time"`
	Price  string  `json:"price"`
	Amount string  `json:"amount"`
	Type   string  `json:"type"`
}

type MarketTrades struct {
	Response
	Result MarketTradesResult `json:"result"`
}

type MarketTradesUpdateEvent struct {
	Params []any `json:"params"`
}

func NewMarketTradesSubscription(handle func(MarketTradesUpdateEvent), market []string) *Subscription {
	params := make([]interface{}, len(market))
	for i := range market {
		params[i] = market[i]
	}

	onEvent := func(event Event) {
		updateEvent, err := TransformEvent[MarketTradesUpdateEvent](event)
		if err != nil {
			fmt.Println(err)
		}

		handle(updateEvent)
	}

	return &Subscription{Command{0, TradesSubscribe, params, WsTypeSubscribe},
		TradesUpdate, onEvent, NewMarketTradesUnsubscribe()}
}

func NewMarketTradesUnsubscribe() Command {
	return Command{0, TradesUnsubscribe, []any{}, WsTypeQuery}
}

func NewMarketTradesCommand(market string, limit int64, StartTradeId int64) Command {
	return Command{
		Id:      1,
		Method:  TradesRequest,
		Params:  []interface{}{market, limit, StartTradeId},
		IsQuery: WsTypeQuery,
	}
}

type Pair [2]string

type AsksAndBids struct {
	Asks []Pair `json:"asks"`
	Bids []Pair `json:"bids"`
}

type MarketDepth struct {
	Response
	Result []AsksAndBids `json:"result"`
}

type MarketDepthUpdateEvent struct {
	Params []any `json:"params"`
}

func NewMarketDepthSubscription(handle func(MarketDepthUpdateEvent), market string, limit int64, priceInterval string, multiSubscription bool) *Subscription {
	onEvent := func(event Event) {
		updateEvent, err := TransformEvent[MarketDepthUpdateEvent](event)
		if err != nil {
			fmt.Println(err)
		}

		handle(updateEvent)
	}

	return &Subscription{Command{0, DepthSubscribe, []interface{}{market, limit, priceInterval, multiSubscription}, WsTypeSubscribe},
		DepthUpdate, onEvent, NewDepthUnsubscribeUnsubscribe()}
}

func NewDepthUnsubscribeUnsubscribe() Command {
	return Command{0, DepthUnsubscribe, []any{}, WsTypeQuery}
}

func NewMarketDepthCommand(market string, limit int64, priceInterval string) Command {
	return Command{
		Id:      1,
		Method:  DepthRequest,
		Params:  []interface{}{market, limit, priceInterval},
		IsQuery: WsTypeQuery,
	}
}
