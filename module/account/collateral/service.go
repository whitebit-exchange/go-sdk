package collateral

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

func (service *Service) GetCollateralBalance(ticker string) (map[string]string, error) {
	endpoint := newBalanceEndpoint(ticker)
	response := make(map[string]string, 0)
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

func (service *Service) GetCollateralSummaryBalance(ticker string) ([]BalanceSummary, error) {
	endpoint := newBalanceSummaryEndpoint(ticker)
	response := make([]BalanceSummary, 0)
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

func (service *Service) GetSummary() (Summary, error) {
	endpoint := newSummaryEndpoint()
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return Summary{}, err
	}

	var summary Summary

	err = json.Unmarshal(result, &summary)

	if err != nil {
		return Summary{}, err
	}

	return summary, nil
}

func (service *Service) SetLeverage(leverage string) (map[string]int, error) {
	endpoint := newLeverageEndpoint(leverage)
	leverageData := make(map[string]int, 0)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &leverageData)

	if err != nil {
		return nil, err
	}

	return leverageData, nil
}

func (service *Service) GetPositionsHistory(options PositionsHistoryOptions) ([]PositionHistory, error) {
	endpoint := newPositionsHistoryEndpoint(options.Market, options.PositionId, options.StartDate, options.EndDate,
		options.Limit, options.Offset)
	response := make([]PositionHistory, 0)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Service) GetOpenPositions(market string) ([]OpenPosition, error) {
	endpoint := newOpenPositionsEndpoint(market)
	response := make([]OpenPosition, 0)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetFundingHistory returns funding payments history.
func (service *Service) GetFundingHistory(market string, limit, offset int) (FundingHistoryResponse, error) {
	endpoint := newFundingHistoryEndpoint(market, limit, offset)
	var response FundingHistoryResponse
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

// GetHedgeMode returns the current hedge mode status.
func (service *Service) GetHedgeMode() (HedgeModeResponse, error) {
	endpoint := newHedgeModeEndpoint()
	var response HedgeModeResponse
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

// UpdateHedgeMode enables or disables hedge mode.
func (service *Service) UpdateHedgeMode(hedgeMode bool) (HedgeModeResponse, error) {
	endpoint := newHedgeModeUpdateEndpoint(hedgeMode)
	var response HedgeModeResponse
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

// GetPositions returns all open positions.
func (service *Service) GetPositions(market string) ([]Position, error) {
	endpoint := newPositionsEndpoint(market)
	response := make([]Position, 0)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// ClosePosition closes a position by market order.
// positionSide: "long" or "short" (for hedge mode)
func (service *Service) ClosePosition(market, positionSide string) error {
	endpoint := newPositionCloseEndpoint(market, positionSide)
	_, err := service.client.SendRequest(endpoint)
	return err
}
