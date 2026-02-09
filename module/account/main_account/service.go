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

func (service *Service) GetMainBalance() (State, error) {
	endpoint := newBalanceEndpoint("")
	response := make(State)
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

func (service *Service) GetAssetBalance(ticker string) (MainBalance, error) {
	endpoint := newBalanceEndpoint(ticker)
	var response MainBalance
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

func (service *Service) GetFee() ([]Fee, error) {
	endpoint := newFeeEndpoint()
	result, err := service.client.SendRequest(endpoint)
	var fee []Fee

	if err != nil {
		return fee, err
	}

	err = json.Unmarshal(result, &fee)

	if err != nil {
		return fee, err
	}

	return fee, nil
}

func (service *Service) GetHistory(params HistoryParams) (History, error) {
	endpoint := newHistoryEndpoint(params)
	var HistoryData History
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

func (service *Service) CreateCode(ticker string, amount string, pass string, description string) (Code, error) {
	endpoint := newCodeEndpoint(ticker, amount, pass, description)
	var response Code
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

func (service *Service) ApplyCode(code string, pass string) (CodeApply, error) {
	endpoint := newCodeApplyEndpoint(code, pass)
	var response CodeApply
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

func (service *Service) GetCodes(limit int64, offset int64) (CodeMy, error) {
	endpoint := newCodeMyEndpoint(limit, offset)
	var response CodeMy
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

func (service *Service) GetCodesHistory(limit int64, offset int64) (CodeHistory, error) {
	endpoint := newCodeMyEndpoint(limit, offset)
	var response CodeHistory
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

func (service *Service) GetAccountFeeByMarket(market string) (AccountFeeByMarket, error) {
	endpoint := newAccountFeeByMarketEndpoint(market)
	var response AccountFeeByMarket
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

func (service *Service) GetCustomFee() (CustomFee, error) {
	endpoint := newCustomFeeEndpoint()
	var response CustomFee
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

// GetDepositAddress returns a deposit address for the specified ticker.
func (service *Service) GetDepositAddress(ticker, network string) (AddressResponse, error) {
	endpoint := newAddressEndpoint(ticker, network)
	var response AddressResponse
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

// CreateNewAddress creates a new deposit address for the specified ticker.
func (service *Service) CreateNewAddress(ticker, network, addressType string) (AddressResponse, error) {
	endpoint := newCreateNewAddressEndpoint(ticker, network, addressType)
	var response AddressResponse
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

// GetFiatDepositURL returns a fiat deposit URL.
func (service *Service) GetFiatDepositURL(ticker, provider, amount, uniqueID string, customer *Customer, successLink, failureLink, returnLink string) (FiatDepositURLResponse, error) {
	endpoint := newFiatDepositURLEndpoint(ticker, provider, amount, uniqueID, customer, successLink, failureLink, returnLink)
	var response FiatDepositURLResponse
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

// Withdraw creates a withdrawal request. Amount includes fee.
func (service *Service) Withdraw(params WithdrawParams) error {
	endpoint := newWithdrawEndpoint(params)
	_, err := service.client.SendRequest(endpoint)
	return err
}

// WithdrawPay creates a withdrawal request where the fee is added on top of the amount.
func (service *Service) WithdrawPay(params WithdrawParams) error {
	endpoint := newWithdrawPayEndpoint(params)
	_, err := service.client.SendRequest(endpoint)
	return err
}
