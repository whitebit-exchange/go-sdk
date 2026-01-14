package status

import (
	//"encoding/json"

	"encoding/json"

	"github.com/whitebit-exchange/go-sdk"
)

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

type State struct {
	Status int `json:"status"` // 1 - system operational, 0 - system maintenance
}

func (service *Service) GetMaintenanceStatus() (State, error) {
	s := State{Status: 0}
	response, err := service.client.SendRequest(newMaintenanceStatusEndpoint())
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(response, &s)

	if err != nil {
		return s, err
	}

	return s, nil

}
