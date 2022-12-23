package server

import (
	"encoding/json"
	"github.com/whitebit-exchange/go-sdk"
)

type TimeResponse struct {
	Time int64 `json:"time"`
}

type PingResponse []string

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

func (service *Service) Ping() (PingResponse, error) {
	endpoint := newPingEndpoint()

	response, err := service.client.SendRequest(endpoint)
	if err != nil {
		return PingResponse{}, err
	}

	var result PingResponse
	err = json.Unmarshal(response, &result)

	if err != nil {
		return PingResponse{}, err
	}

	return result, nil

}

func (service *Service) GetTime() (TimeResponse, error) {
	endpoint := newTimeEndpoint()

	response, err := service.client.SendRequest(endpoint)
	if err != nil {
		return TimeResponse{}, err
	}

	var result TimeResponse
	err = json.Unmarshal(response, &result)

	if err != nil {
		return TimeResponse{}, err
	}

	return result, nil

}
