package market

import (
	"encoding/json"
	"github.com/whitebit-exchange/go-sdk"
)

type Market struct {
	Name          string `json:"name"`
	Stock         string `json:"stock"`
	Money         string `json:"money"`
	StockPrec     string `json:"stockPrec"`
	MoneyPrec     string `json:"moneyPrec"`
	FeePrec       string `json:"feePrec"`
	MakerFee      string `json:"makerFee"`
	TakerFee      string `json:"takerFee"`
	MinAmount     string `json:"minAmount"`
	MinTotal      string `json:"minTotal"`
	TradesEnabled bool   `json:"tradesEnabled"`
}

type Result struct {
	Success bool                `json:"success"`
	Message map[string][]string `json:"message"`
	Result  []Market            `json:"result"`
}

type CollateralMarketsResult struct {
	Success bool                `json:"success"`
	Message map[string][]string `json:"message"`
	Result  []string            `json:"result"`
}

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

func (service *Service) GetMarkets() (Result, error) {
	endpoint := newMarketsInfoEndpoint()

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

func (service *Service) GetCollateralMarkets() (CollateralMarketsResult, error) {
	endpoint := newCollateralMarketsEndpoint()

	response, err := service.client.SendRequest(endpoint)
	if err != nil {
		return CollateralMarketsResult{}, err
	}

	var result CollateralMarketsResult
	err = json.Unmarshal(response, &result)

	if err != nil {
		return CollateralMarketsResult{}, err
	}

	return result, nil

}
