package convert

import (
	"encoding/json"

	"github.com/whitebit-exchange/go-sdk"
)

type Estimate struct {
	QueryID    string `json:"queryId"`
	From       string `json:"from"`
	To         string `json:"to"`
	GiveAmount string `json:"giveAmount"`
	GetAmount  string `json:"getAmount"`
	Rate       string `json:"rate"`
	Fee        string `json:"fee"`
	ExpiresAt  int64  `json:"expiresAt"`
}

type Confirmation struct {
	ID         int64  `json:"id"`
	Status     int    `json:"status"`
	From       string `json:"from"`
	To         string `json:"to"`
	GiveAmount string `json:"giveAmount"`
	GetAmount  string `json:"getAmount"`
	Fee        string `json:"fee"`
}

type HistoryRecord struct {
	ID         int64  `json:"id"`
	From       string `json:"from"`
	To         string `json:"to"`
	GiveAmount string `json:"giveAmount"`
	GetAmount  string `json:"getAmount"`
	Rate       string `json:"rate"`
	Fee        string `json:"fee"`
	Status     int    `json:"status"`
	CreatedAt  int64  `json:"createdAt"`
}

type HistoryResponse struct {
	Total   int             `json:"total"`
	Records []HistoryRecord `json:"records"`
	Limit   int             `json:"limit"`
	Offset  int             `json:"offset"`
}

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

// Estimate returns a conversion estimate between two currencies.
func (s *Service) Estimate(from, to, amount string) (Estimate, error) {
	resp, err := s.client.SendRequest(newEstimateEndpoint(from, to, amount))
	if err != nil {
		return Estimate{}, err
	}

	var result Estimate
	if err := json.Unmarshal(resp, &result); err != nil {
		return Estimate{}, err
	}
	return result, nil
}

// Confirm confirms a conversion using the query ID from Estimate.
func (s *Service) Confirm(queryID string) (Confirmation, error) {
	resp, err := s.client.SendRequest(newConfirmEndpoint(queryID))
	if err != nil {
		return Confirmation{}, err
	}

	var result Confirmation
	if err := json.Unmarshal(resp, &result); err != nil {
		return Confirmation{}, err
	}
	return result, nil
}

// GetHistory returns the conversion history.
func (s *Service) GetHistory(limit, offset int) (HistoryResponse, error) {
	resp, err := s.client.SendRequest(newHistoryEndpoint(limit, offset))
	if err != nil {
		return HistoryResponse{}, err
	}

	var result HistoryResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return HistoryResponse{}, err
	}
	return result, nil
}
