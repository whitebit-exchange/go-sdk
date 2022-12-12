package collateral

import (
	"encoding/json"
	"github.com/whitebit-exchange/whitebit"
)

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

func (service *Service) CreateLimitOrder(params LimitOrderParams) (*LimitOrder, error) {
	endpoint := newLimitEndpoint(params)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	var order LimitOrder
	err = json.Unmarshal(result, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (service *Service) CreateMarketOrder(params MarketOrderParams) (*MarketOrder, error) {
	endpoint := newMarketEndpoint(params)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	var order MarketOrder
	err = json.Unmarshal(result, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (service *Service) CreateStopLimitOrder(params StopLimitOrderParams) (*StopLimitOrder, error) {
	endpoint := newStopLimitEndpoint(params)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	var order StopLimitOrder
	err = json.Unmarshal(result, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (service *Service) CreateStopMarketOrder(params StopMarketOrderParams) (*StopMarketOrder, error) {
	endpoint := newStopMarketEndpoint(params)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	var order StopMarketOrder
	err = json.Unmarshal(result, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (service *Service) CreateOcoOrder(params OcoOrderParams) (*OcoOrder, error) {
	endpoint := newOcoEndpoint(params)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	var order OcoOrder
	err = json.Unmarshal(result, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (service *Service) CancelOrder(market string, orderId int64) (*CancelOrder, error) {
	endpoint := newCancelOrderEndpoint(market, orderId)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	var order CancelOrder
	err = json.Unmarshal(result, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (service *Service) CancelOcoOrder(market string, orderId int64) (*OcoCancelOrder, error) {
	endpoint := newOcoCancelOrderEndpoint(market, orderId)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	var order OcoCancelOrder
	err = json.Unmarshal(result, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
