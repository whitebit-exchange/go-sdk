package smartlending

import (
	"encoding/json"

	"github.com/whitebit-exchange/go-sdk"
)

// Fixed Lending Types

type FixedPlan struct {
	ID                    string `json:"id"`
	Ticker                string `json:"ticker"`
	Status                int    `json:"status"`
	Percent               string `json:"percent"`
	Duration              int    `json:"duration"`
	InterestTicker        string `json:"interestTicker"`
	InterestRatio         string `json:"interestRatio"`
	MinInvestment         string `json:"minInvestment"`
	MaxInvestment         string `json:"maxInvestment"`
	MaxPossibleInvestment string `json:"maxPossibleInvestment"`
}

type FixedInvestment struct {
	ID           string    `json:"id"`
	Plan         FixedPlan `json:"plan"`
	Status       int       `json:"status"`
	Created      int64     `json:"created"`
	Updated      int64     `json:"updated"`
	PaymentTime  int64     `json:"paymentTime"`
	Amount       string    `json:"amount"`
	InterestPaid string    `json:"interestPaid"`
}

type FixedInvestmentsResponse struct {
	Offset  int               `json:"offset"`
	Limit   int               `json:"limit"`
	Records []FixedInvestment `json:"records"`
}

type InterestPayment struct {
	PlanID       string `json:"planId"`
	InvestmentID string `json:"investmentId"`
	Amount       string `json:"amount"`
	Ticker       string `json:"ticker"`
	Timestamp    int64  `json:"timestamp"`
}

type InterestHistoryResponse struct {
	Offset  int               `json:"offset"`
	Limit   int               `json:"limit"`
	Records []InterestPayment `json:"records"`
}

// Flexible Lending Types

type FlexPlan struct {
	ID            string `json:"id"`
	Ticker        string `json:"ticker"`
	MinInvestment string `json:"minInvestment"`
	MaxInvestment string `json:"maxInvestment"`
	MaxRate       string `json:"maxRate"`
}

type FlexInvestment struct {
	ID             string `json:"id"`
	PlanID         string `json:"planId"`
	Currency       string `json:"currency"`
	Invested       string `json:"invested"`
	WithAutoInvest bool   `json:"withAutoReinvest"`
	Status         int    `json:"status"`
	CreatedAt      int64  `json:"createdAt"`
	UpdatedAt      int64  `json:"updatedAt"`
}

type FlexInvestmentsResponse struct {
	Data   []FlexInvestment `json:"data"`
	Limit  int              `json:"limit"`
	Offset int              `json:"offset"`
}

type FlexHistoryRecord struct {
	CreatedAt     int64  `json:"createdAt"`
	PlanID        string `json:"planId"`
	InvestmentID  string `json:"investmentId"`
	Currency      string `json:"currency"`
	Amount        string `json:"amount"`
	TransactionID string `json:"transactionId"`
	ActionType    int    `json:"actionType"`
}

type FlexHistoryResponse struct {
	Data   []FlexHistoryRecord `json:"data"`
	Limit  int                 `json:"limit"`
	Offset int                 `json:"offset"`
}

type Service struct {
	client whitebit.Client
}

func NewService(client whitebit.Client) *Service {
	return &Service{client: client}
}

// Fixed Lending Methods

// GetFixedPlans returns available fixed lending plans.
func (s *Service) GetFixedPlans(ticker string) ([]FixedPlan, error) {
	resp, err := s.client.SendRequest(newFixedPlansEndpoint(ticker))
	if err != nil {
		return nil, err
	}

	var result []FixedPlan
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateFixedInvestment creates a new fixed investment.
func (s *Service) CreateFixedInvestment(planID, amount string) (FixedInvestment, error) {
	resp, err := s.client.SendRequest(newFixedInvestmentEndpoint(planID, amount))
	if err != nil {
		return FixedInvestment{}, err
	}

	var result FixedInvestment
	if err := json.Unmarshal(resp, &result); err != nil {
		return FixedInvestment{}, err
	}
	return result, nil
}

// CloseFixedInvestment closes an active fixed investment.
func (s *Service) CloseFixedInvestment(id string) error {
	_, err := s.client.SendRequest(newFixedInvestmentCloseEndpoint(id))
	return err
}

// GetFixedInvestments returns fixed investments history.
// Status: 1 = active, 2 = closed
func (s *Service) GetFixedInvestments(id, ticker string, status, limit, offset int) (FixedInvestmentsResponse, error) {
	resp, err := s.client.SendRequest(newFixedInvestmentsEndpoint(id, ticker, status, limit, offset))
	if err != nil {
		return FixedInvestmentsResponse{}, err
	}

	var result FixedInvestmentsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return FixedInvestmentsResponse{}, err
	}
	return result, nil
}

// GetFixedInterestHistory returns interest payment history.
func (s *Service) GetFixedInterestHistory(planID, ticker string, limit, offset int) (InterestHistoryResponse, error) {
	resp, err := s.client.SendRequest(newFixedInterestHistoryEndpoint(planID, ticker, limit, offset))
	if err != nil {
		return InterestHistoryResponse{}, err
	}

	var result InterestHistoryResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return InterestHistoryResponse{}, err
	}
	return result, nil
}

// Flexible Lending Methods

// GetFlexPlans returns available flexible lending plans.
func (s *Service) GetFlexPlans(ticker string, limit, offset int) ([]FlexPlan, error) {
	resp, err := s.client.SendRequest(newFlexPlansEndpoint(ticker, limit, offset))
	if err != nil {
		return nil, err
	}

	var result []FlexPlan
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetFlexInvestments returns user's flexible investments.
// Status: 0 = closed, 1 = active
func (s *Service) GetFlexInvestments(ticker, plan, investment string, status, limit, offset int) (FlexInvestmentsResponse, error) {
	resp, err := s.client.SendRequest(newFlexInvestmentsEndpoint(ticker, plan, investment, status, limit, offset))
	if err != nil {
		return FlexInvestmentsResponse{}, err
	}

	var result FlexInvestmentsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return FlexInvestmentsResponse{}, err
	}
	return result, nil
}

// GetFlexHistory returns flexible investment history.
// ActionTypes: 1=INVEST, 2=REINVEST, 3=WITHDRAW, 4=DAILY_EARNING, 5=CLOSE, 6=OPEN
func (s *Service) GetFlexHistory(plan, investment, transaction string, dateFrom, dateTo int64, actionTypes []int, limit, offset int) (FlexHistoryResponse, error) {
	resp, err := s.client.SendRequest(newFlexHistoryEndpoint(plan, investment, transaction, dateFrom, dateTo, actionTypes, limit, offset))
	if err != nil {
		return FlexHistoryResponse{}, err
	}

	var result FlexHistoryResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return FlexHistoryResponse{}, err
	}
	return result, nil
}

// GetFlexPaymentHistory returns earnings history (DAILY_EARNING only).
func (s *Service) GetFlexPaymentHistory(plan, investment string, dateFrom, dateTo int64, limit, offset int) (FlexHistoryResponse, error) {
	resp, err := s.client.SendRequest(newFlexPaymentHistoryEndpoint(plan, investment, dateFrom, dateTo, limit, offset))
	if err != nil {
		return FlexHistoryResponse{}, err
	}

	var result FlexHistoryResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return FlexHistoryResponse{}, err
	}
	return result, nil
}

// FlexInvest creates a new flexible investment.
func (s *Service) FlexInvest(planID, amount string) (FlexInvestment, error) {
	resp, err := s.client.SendRequest(newFlexInvestEndpoint(planID, amount))
	if err != nil {
		return FlexInvestment{}, err
	}

	var result FlexInvestment
	if err := json.Unmarshal(resp, &result); err != nil {
		return FlexInvestment{}, err
	}
	return result, nil
}

// FlexWithdraw withdraws from a flexible investment.
func (s *Service) FlexWithdraw(investmentID, amount string) (FlexInvestment, error) {
	resp, err := s.client.SendRequest(newFlexWithdrawEndpoint(investmentID, amount))
	if err != nil {
		return FlexInvestment{}, err
	}

	var result FlexInvestment
	if err := json.Unmarshal(resp, &result); err != nil {
		return FlexInvestment{}, err
	}
	return result, nil
}

// FlexClose closes a flexible investment.
func (s *Service) FlexClose(investmentID string) error {
	_, err := s.client.SendRequest(newFlexCloseEndpoint(investmentID))
	return err
}

// FlexSetAutoInvest enables or disables auto-reinvest.
func (s *Service) FlexSetAutoInvest(investmentID string, autoInvest bool) (FlexInvestment, error) {
	resp, err := s.client.SendRequest(newFlexAutoInvestEndpoint(investmentID, autoInvest))
	if err != nil {
		return FlexInvestment{}, err
	}

	var result FlexInvestment
	if err := json.Unmarshal(resp, &result); err != nil {
		return FlexInvestment{}, err
	}
	return result, nil
}
