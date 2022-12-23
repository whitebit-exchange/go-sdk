package futures

import (
	"encoding/json"
	"github.com/whitebit-exchange/go-sdk"
)

type Market struct {
	TickerId                 string `json:"ticker_id"`
	StockCurrency            string `json:"stock_currency"`
	MoneyCurrency            string `json:"money_currency"`
	LastPrice                string `json:"last_price"`
	StockVolume              string `json:"stock_volume"`
	MoneyVolume              string `json:"money_volume"`
	Bid                      string `json:"bid"`
	Ask                      string `json:"ask"`
	High                     string `json:"high"`
	Low                      string `json:"low"`
	ProductType              string `json:"product_type"`
	OpenInterest             string `json:"open_interest"`
	IndexPrice               string `json:"index_price"`
	IndexName                string `json:"index_name"`
	IndexCurrency            string `json:"index_currency"`
	FundingRate              string `json:"funding_rate"`
	NextFundingRateTimestamp string `json:"next_funding_rate_timestamp"`
}

type MarketList struct {
	Success bool                `json:"success"`
	Message map[string][]string `json:"message"`
	Result  []Market            `json:"result"`
}

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

func (service *Service) GetFuturesMarkets() (MarketList, error) {
	response, err := service.client.SendRequest(newFuturesMarketsEndpoint())
	if err != nil {
		return MarketList{}, err
	}

	var result MarketList
	err = json.Unmarshal(response, &result)

	if err != nil {
		return MarketList{}, err
	}

	return result, nil

}
