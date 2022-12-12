package symbol

import (
	"encoding/json"
	"github.com/whitebit-exchange/whitebit"
)

type SymbolsResult struct {
	Success bool                `json:"success"`
	Message map[string][]string `json:"message"`
	Result  []string            `json:"result"`
}

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

func (service *Service) GetSymbols() (SymbolsResult, error) {
	response, err := service.client.SendRequest(newEndpoint())
	if err != nil {
		return SymbolsResult{}, err
	}

	var result SymbolsResult
	err = json.Unmarshal(response, &result)

	if err != nil {
		return SymbolsResult{}, err
	}
	return result, nil

}
