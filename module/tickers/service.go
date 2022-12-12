package tickers

import (
	"encoding/json"
	"github.com/whitebit-exchange/whitebit"
)

type MarketActivity struct {
	LastUpdatedTimestamp string `json:"lastUpdateTimestamp"`
	TradingPairs         string `json:"tradingPairs"`
	LastPrice            string `json:"lastPrice"`
	LowestAsk            string `json:"lowestAsk"`
	HighestBid           string `json:"highestBid"`
	BaseVolume           string `json:"baseVolume24h"`
	QuoteVolume          string `json:"quoteVolume24h"`
	TradesEnabled        bool   `json:"tradesEnabled"`
}

type MarketActivityResult struct {
	Success bool                `json:"success"`
	Message map[string][]string `json:"message"`
	Result  []MarketActivity    `json:"result"`
}

type SingleMarketActivity struct {
	Open   string `json:"open"`
	Bid    string `json:"bid"`
	Ask    string `json:"ask"`
	Low    string `json:"low"`
	High   string `json:"high"`
	Last   string `json:"last"`
	Vol    string `json:"volume"`
	Deal   string `json:"deal"`
	Change string `json:"change"`
}

type SingleMarketActivityResult struct {
	Success bool                 `json:"success"`
	Message map[string][]string  `json:"message"`
	Result  SingleMarketActivity `json:"result"`
}

type Ticker struct {
	BaseID      int    `json:"base_id"`
	QuoteID     int    `json:"quote_id"`
	LastPrice   string `json:"last_price"`
	QuoteVolume string `json:"quote_volume"`
	BaseVolume  string `json:"base_volume"`
	IsFrozen    bool   `json:"isFrozen"`
	Change      string `json:"change"`
}

type TickerResult map[string]Ticker

type Tickers map[string]TickerStatus

type TickerStatus struct {
	SingleMarketActivity `json:"ticker"`
	At                   int64 `json:"at"`
}

type Result struct {
	Success bool                `json:"success"`
	Message map[string][]string `json:"message"`
	Result  Tickers             `json:"result"`
}

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

func (service *Service) GetSingleMarketActivity(market string) (SingleMarketActivityResult, error) {
	endpoint := newSingleMarketActivityEndpoint(market)
	response, err := service.client.SendRequest(endpoint)
	if err != nil {
		return SingleMarketActivityResult{}, err
	}

	var result SingleMarketActivityResult
	err = json.Unmarshal(response, &result)

	if err != nil {
		return SingleMarketActivityResult{}, err
	}

	return result, nil

}

func (service *Service) GetMarketActivity() (MarketActivityResult, error) {
	endpoint := newMarketActivityEndpoint()
	response, err := service.client.SendRequest(endpoint)
	if err != nil {
		return MarketActivityResult{}, err
	}

	var result MarketActivityResult
	err = json.Unmarshal(response, &result)

	if err != nil {
		return MarketActivityResult{}, err
	}

	return result, nil
}

func (service *Service) GetAvailableTickers() (TickerResult, error) {
	endpoint := newTickerEndpoint()
	response, err := service.client.SendRequest(endpoint)
	if err != nil {
		return TickerResult{}, err
	}

	tickers := new(TickerResult)
	err = json.Unmarshal(response, tickers)

	if err != nil {
		return TickerResult{}, err
	}

	return *tickers, nil
}

func (service *Service) GetTickers() (Result, error) {
	endpoint := newTickersEndpoint()

	response, err := service.client.SendRequest(endpoint)
	if err != nil {
		return Result{}, err
	}

	tickers := new(Result)
	err = json.Unmarshal(response, tickers)

	if err != nil {
		return Result{}, err
	}

	return *tickers, nil
}
