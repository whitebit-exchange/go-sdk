package fee

import (
	"encoding/json"
	"github.com/whitebit-exchange/whitebit"
)

type List map[string]interface{}

type Fee struct {
	MakerFee string `json:"makerFee"`
	TakerFee string `json:"takerFee"`
}

type Result struct {
	Success bool                `json:"success"`
	Message map[string][]string `json:"message"`
	Result  Fee                 `json:"result"`
}

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

func (service *Service) GetTradingFee() (Result, error) {
	response, err := service.client.SendRequest(newTradingFeeEndpoint())
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

func (service *Service) GetTradingFeesList() (List, error) {
	response, err := service.client.SendRequest(newListEndpoint())
	if err != nil {
		return List{}, err
	}

	result := make(List)
	err = json.Unmarshal(response, &result)

	if err != nil {
		return List{}, err
	}

	return result, nil

}
