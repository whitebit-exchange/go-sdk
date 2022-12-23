package trade

import (
	"encoding/json"
	"github.com/whitebit-exchange/go-sdk"
	order "github.com/whitebit-exchange/go-sdk/module/order/spot"
)

type ExecutedHistoryOptions struct {
	Market        string
	Limit         int
	Offset        int
	ClientOrderId string
}

type HistoryOptions struct {
	Market        string
	Limit         int
	Offset        int
	OrderId       int
	ClientOrderId string
}

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

func (service *Service) GetHistory(options HistoryOptions) (map[string][]OrderHistory, error) {
	endpoint := newHistoryEndpoint(options.Market, options.Limit, options.Offset, options.OrderId, options.ClientOrderId)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	history := make(map[string][]OrderHistory)

	err = json.Unmarshal(result, &history)

	if err != nil {
		return nil, err
	}

	return history, nil
}

func (service *Service) GetExecutedHistory(options ExecutedHistoryOptions) ([]ExecutedHistory, error) {
	endpoint := newExecutedHistoryEndpoint(options.Market, options.Limit, options.Offset, options.ClientOrderId)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	var executedHistory []ExecutedHistory

	err = json.Unmarshal(result, &executedHistory)

	if err != nil {
		return nil, err
	}

	return executedHistory, nil
}

func (service *Service) GetOrders(market string, limit int, offset int) ([]Orders, error) {
	endpoint := newOrdersEndpoint(market, limit, offset)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	var orders []Orders

	err = json.Unmarshal(result, &orders)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (service *Service) GetBalance(ticker string) (map[string]any, error) {
	endpoint := newBalanceEndpoint(ticker)
	balances := make(map[string]any)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &balances)

	if err != nil {
		return nil, err
	}

	return balances, nil
}

func (service *Service) GetOrder(orderId int64, limit int, offset int) (map[string]interface{}, error) {
	endpoint := newOrderEndpoint(orderId, limit, offset)
	response := make(map[string]interface{}, 0)
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

func (service *Service) GetOcoOrders(market string, limit int, offset int) ([]order.OcoOrder, error) {
	endpoint := newOcoListEndpoint(market, limit, offset)
	orderRecords := make([]order.OcoOrder, 0)
	result, err := service.client.SendRequest(endpoint)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &orderRecords)

	if err != nil {
		return nil, err
	}

	return orderRecords, nil
}
