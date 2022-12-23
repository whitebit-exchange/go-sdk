package depth

import (
	"encoding/json"
	"github.com/whitebit-exchange/go-sdk"
)

type OrderBookOptions struct {
	Market string
	Limit  int
	Level  int
}

type Pair [2]string

type AsksAndBids struct {
	Asks []Pair `json:"asks"`
	Bids []Pair `json:"bids"`
}

type Options struct {
	Name  string
	Value string
}

type Depth struct {
	Time string `json:"lastUpdateTimestamp"`
	Asks []Pair `json:"asks"`
	Bids []Pair `json:"bids"`
}

type Result struct {
	Success bool                `json:"success"`
	Message map[string][]string `json:"message"`
	Result  Depth               `json:"result"`
}

type OrderBook struct {
	Time int64 `json:"timestamp"`
	AsksAndBids
}

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

func (service *Service) GetDepth(market string) (Result, error) {
	endpoint := newDepthEndpoint(market)
	response, err := service.client.SendRequest(endpoint)
	if err != nil {
		return Result{}, err
	}

	var result Result
	err = json.Unmarshal(response, &result)

	if err != nil {
		return Result{}, err
	}

	return result, nil

}

func (service *Service) GetOrderBook(options OrderBookOptions) (OrderBook, error) {
	endpoint := newOrderBookEndpoint(options.Market)

	if options.Limit != 0 {
		endpoint.SetLimit(options.Limit)
	}

	if options.Level != 0 {
		endpoint.SetLevel(options.Level)
	}

	response, err := service.client.SendRequest(endpoint)
	if err != nil {
		return OrderBook{}, err
	}

	var result OrderBook
	err = json.Unmarshal(response, &result)

	if err != nil {
		return OrderBook{}, err
	}

	return result, nil

}
