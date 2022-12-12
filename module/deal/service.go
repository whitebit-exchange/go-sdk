package deal

import (
	"encoding/json"
	"github.com/whitebit-exchange/whitebit"
)

type Deal struct {
	TradeID int    `json:"tradeID"`
	Price   string `json:"price"`
	Amount  string `json:"quote_volume"`
	Volume  string `json:"base_volume"`
	Time    int    `json:"trade_timestamp"`
	Type    string `json:"type"`
}

type Deals struct {
	Id     int     `json:"id"`
	Time   float64 `json:"time"`
	Price  string  `json:"price"`
	Amount string  `json:"amount"`
	Type   string  `json:"type"`
}

type TradeHistoryResult struct {
	Success bool                `json:"success"`
	Message map[string][]string `json:"message"`
	Result  []Deals             `json:"result"`
}

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

type Options struct {
	Market string
	Type   string
}

func (service *Service) GetDeals(options Options) ([]Deal, error) {
	endpoint := newDealsEndpoint(options.Market)

	if options.Type != "" {
		endpoint.SetType(options.Type)
	}

	response, err := service.client.SendRequest(endpoint)
	if err != nil {
		return []Deal{}, err
	}

	dealsList := make([]Deal, 0)
	err = json.Unmarshal(response, &dealsList)

	if err != nil {
		return []Deal{}, err
	}

	return dealsList, nil

}

type TradeHistoryOptions struct {
	Market string
	LastId int64
	Limit  int
}

func (service *Service) GetTradeHistory(options TradeHistoryOptions) (TradeHistoryResult, error) {
	endpoint := newTradeHistoryEndpoint(options.Market, options.LastId)

	if options.Limit != 0 {
		endpoint.SetLimit(options.Limit)
	}

	response, err := service.client.SendRequest(endpoint)
	if err != nil {
		return TradeHistoryResult{}, err
	}

	var result TradeHistoryResult
	err = json.Unmarshal(response, &result)

	if err != nil {
		return TradeHistoryResult{}, err
	}

	return result, nil

}
