package spot

import (
	"encoding/json"
	"github.com/whitebit-exchange/go-sdk"
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
	endpoint := newMarketOrderEndpoint(params)
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

func (service *Service) CreateMarketStock(params MarketOrderParams) (*StockMarketOrder, error) {
	endpoint := newStockMarketEndpoint(params)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	var order StockMarketOrder
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
