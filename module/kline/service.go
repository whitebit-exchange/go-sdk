package kline

import (
	"encoding/json"
	"github.com/whitebit-exchange/whitebit"
)

type Kline []interface{}

type Result struct {
	Success bool                `json:"success"`
	Message map[string][]string `json:"message"`
	Result  []Kline             `json:"result"`
}

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

type Options struct {
	Market   string
	Start    int
	End      int
	Interval string
	Limit    int
}

func (service *Service) GetKline(options Options) (Result, error) {
	endpoint := newKlineEndpoint(options.Market)

	if options.Start != 0 {
		endpoint.SetStart(options.Start)
	}

	if options.End != 0 {
		endpoint.SetEnd(options.End)
	}

	if options.Interval != "" {
		endpoint.SetInterval(options.Interval)
	}
	if options.Limit != 0 {
		endpoint.SetLimit(options.Limit)
	}

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
