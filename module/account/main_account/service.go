package main_account

import (
	"encoding/json"
	"github.com/whitebit-exchange/go-sdk"
)

type PositionsHistoryOptions struct {
	Market     string
	PositionId int64
	StartDate  int64
	EndDate    int64
	Limit      string
	Offset     string
}

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

func (service *Service) GetMainBalance() (BalanceResult, error) {
	endpoint := newBalanceEndpoint("")
	response := make(BalanceResult)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return response, err
	}

	err = json.Unmarshal(result, &response)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (service *Service) GetMainBalanceTicker(ticker string) (MainBalanceResult, error) {
	endpoint := newBalanceEndpoint(ticker)
	var response MainBalanceResult
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return response, err
	}

	err = json.Unmarshal(result, &response)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (service *Service) GetFee() ([]FeeResult, error) {
	endpoint := newFeeEndpoint()
	result, err := service.client.SendRequest(endpoint)
	var fee []FeeResult

	if err != nil {
		return fee, err
	}

	err = json.Unmarshal(result, &fee)

	if err != nil {
		return fee, err
	}

	return fee, nil
}

func (service *Service) GetHistory(params HistoryParams) (HistoryResult, error) {
	endpoint := newHistoryEndpoint(params)
	var HistoryData HistoryResult
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return HistoryData, err
	}

	err = json.Unmarshal(result, &HistoryData)

	if err != nil {
		return HistoryData, err
	}

	return HistoryData, nil
}

func (service *Service) Transfer(params TransferParams) error {
	endpoint := newTransferEndpoint(params)
	_, err := service.client.SendRequest(endpoint)

	if err != nil {
		return err
	}

	return nil
}
