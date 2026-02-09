package smartlending

import "github.com/whitebit-exchange/go-sdk"

const (
	fixedPlansURL           = "/api/v4/main-account/smart/plans"
	fixedInvestmentURL      = "/api/v4/main-account/smart/investment"
	fixedInvestmentCloseURL = "/api/v4/main-account/smart/investment/close"
	fixedInvestmentsURL     = "/api/v4/main-account/smart/investments"
	fixedInterestHistoryURL = "/api/v4/main-account/smart/interest-payment-history"
)

type fixedPlansEndpoint struct {
	whitebit.AuthParams
	Ticker string `json:"ticker,omitempty"`
}

func newFixedPlansEndpoint(ticker string) *fixedPlansEndpoint {
	return &fixedPlansEndpoint{
		AuthParams: whitebit.NewAuthParams(fixedPlansURL),
		Ticker:     ticker,
	}
}

type fixedInvestmentEndpoint struct {
	whitebit.AuthParams
	PlanID string `json:"planId"`
	Amount string `json:"amount"`
}

func newFixedInvestmentEndpoint(planID, amount string) *fixedInvestmentEndpoint {
	return &fixedInvestmentEndpoint{
		AuthParams: whitebit.NewAuthParams(fixedInvestmentURL),
		PlanID:     planID,
		Amount:     amount,
	}
}

type fixedInvestmentCloseEndpoint struct {
	whitebit.AuthParams
	ID string `json:"id"`
}

func newFixedInvestmentCloseEndpoint(id string) *fixedInvestmentCloseEndpoint {
	return &fixedInvestmentCloseEndpoint{
		AuthParams: whitebit.NewAuthParams(fixedInvestmentCloseURL),
		ID:         id,
	}
}

type fixedInvestmentsEndpoint struct {
	whitebit.AuthParams
	ID     string `json:"id,omitempty"`
	Ticker string `json:"ticker,omitempty"`
	Status int    `json:"status,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
}

func newFixedInvestmentsEndpoint(id, ticker string, status, limit, offset int) *fixedInvestmentsEndpoint {
	return &fixedInvestmentsEndpoint{
		AuthParams: whitebit.NewAuthParams(fixedInvestmentsURL),
		ID:         id,
		Ticker:     ticker,
		Status:     status,
		Limit:      limit,
		Offset:     offset,
	}
}

type fixedInterestHistoryEndpoint struct {
	whitebit.AuthParams
	PlanID string `json:"planId,omitempty"`
	Ticker string `json:"ticker,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
}

func newFixedInterestHistoryEndpoint(planID, ticker string, limit, offset int) *fixedInterestHistoryEndpoint {
	return &fixedInterestHistoryEndpoint{
		AuthParams: whitebit.NewAuthParams(fixedInterestHistoryURL),
		PlanID:     planID,
		Ticker:     ticker,
		Limit:      limit,
		Offset:     offset,
	}
}
