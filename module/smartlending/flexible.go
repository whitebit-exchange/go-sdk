package smartlending

import "github.com/whitebit-exchange/go-sdk"

const (
	flexPlansURL          = "/api/v4/main-account/smart-flex/plans"
	flexInvestmentsURL    = "/api/v4/main-account/smart-flex/investments"
	flexHistoryURL        = "/api/v4/main-account/smart-flex/investments/history"
	flexPaymentHistoryURL = "/api/v4/main-account/smart-flex/investments/payment-history"
	flexInvestURL         = "/api/v4/main-account/smart-flex/investments/invest"
	flexWithdrawURL       = "/api/v4/main-account/smart-flex/investments/withdraw"
	flexCloseURL          = "/api/v4/main-account/smart-flex/investments/close"
	flexAutoInvestURL     = "/api/v4/main-account/smart-flex/investments/auto-invest"
)

type flexPlansEndpoint struct {
	whitebit.AuthParams
	Ticker string `json:"ticker,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
}

func newFlexPlansEndpoint(ticker string, limit, offset int) *flexPlansEndpoint {
	return &flexPlansEndpoint{
		AuthParams: whitebit.NewAuthParams(flexPlansURL),
		Ticker:     ticker,
		Limit:      limit,
		Offset:     offset,
	}
}

type flexInvestmentsEndpoint struct {
	whitebit.AuthParams
	Ticker           string `json:"ticker,omitempty"`
	Plan             string `json:"plan,omitempty"`
	Investment       string `json:"investment,omitempty"`
	InvestmentStatus int    `json:"investmentStatus,omitempty"`
	Limit            int    `json:"limit,omitempty"`
	Offset           int    `json:"offset,omitempty"`
}

func newFlexInvestmentsEndpoint(ticker, plan, investment string, status, limit, offset int) *flexInvestmentsEndpoint {
	return &flexInvestmentsEndpoint{
		AuthParams:       whitebit.NewAuthParams(flexInvestmentsURL),
		Ticker:           ticker,
		Plan:             plan,
		Investment:       investment,
		InvestmentStatus: status,
		Limit:            limit,
		Offset:           offset,
	}
}

type flexHistoryEndpoint struct {
	whitebit.AuthParams
	Plan        string `json:"plan,omitempty"`
	Investment  string `json:"investment,omitempty"`
	Transaction string `json:"transaction,omitempty"`
	DateFrom    int64  `json:"dateFrom,omitempty"`
	DateTo      int64  `json:"dateTo,omitempty"`
	ActionTypes []int  `json:"actionTypes,omitempty"`
	Limit       int    `json:"limit,omitempty"`
	Offset      int    `json:"offset,omitempty"`
}

func newFlexHistoryEndpoint(plan, investment, transaction string, dateFrom, dateTo int64, actionTypes []int, limit, offset int) *flexHistoryEndpoint {
	return &flexHistoryEndpoint{
		AuthParams:  whitebit.NewAuthParams(flexHistoryURL),
		Plan:        plan,
		Investment:  investment,
		Transaction: transaction,
		DateFrom:    dateFrom,
		DateTo:      dateTo,
		ActionTypes: actionTypes,
		Limit:       limit,
		Offset:      offset,
	}
}

type flexPaymentHistoryEndpoint struct {
	whitebit.AuthParams
	Plan       string `json:"plan,omitempty"`
	Investment string `json:"investment,omitempty"`
	DateFrom   int64  `json:"dateFrom,omitempty"`
	DateTo     int64  `json:"dateTo,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Offset     int    `json:"offset,omitempty"`
}

func newFlexPaymentHistoryEndpoint(plan, investment string, dateFrom, dateTo int64, limit, offset int) *flexPaymentHistoryEndpoint {
	return &flexPaymentHistoryEndpoint{
		AuthParams: whitebit.NewAuthParams(flexPaymentHistoryURL),
		Plan:       plan,
		Investment: investment,
		DateFrom:   dateFrom,
		DateTo:     dateTo,
		Limit:      limit,
		Offset:     offset,
	}
}

type flexInvestEndpoint struct {
	whitebit.AuthParams
	PlanID string `json:"planId"`
	Amount string `json:"amount"`
}

func newFlexInvestEndpoint(planID, amount string) *flexInvestEndpoint {
	return &flexInvestEndpoint{
		AuthParams: whitebit.NewAuthParams(flexInvestURL),
		PlanID:     planID,
		Amount:     amount,
	}
}

type flexWithdrawEndpoint struct {
	whitebit.AuthParams
	InvestmentID string `json:"investmentId"`
	Amount       string `json:"amount"`
}

func newFlexWithdrawEndpoint(investmentID, amount string) *flexWithdrawEndpoint {
	return &flexWithdrawEndpoint{
		AuthParams:   whitebit.NewAuthParams(flexWithdrawURL),
		InvestmentID: investmentID,
		Amount:       amount,
	}
}

type flexCloseEndpoint struct {
	whitebit.AuthParams
	InvestmentID string `json:"investmentId"`
}

func newFlexCloseEndpoint(investmentID string) *flexCloseEndpoint {
	return &flexCloseEndpoint{
		AuthParams:   whitebit.NewAuthParams(flexCloseURL),
		InvestmentID: investmentID,
	}
}

type flexAutoInvestEndpoint struct {
	whitebit.AuthParams
	InvestmentID string `json:"investmentId"`
	AutoInvest   bool   `json:"autoInvest"`
}

func newFlexAutoInvestEndpoint(investmentID string, autoInvest bool) *flexAutoInvestEndpoint {
	return &flexAutoInvestEndpoint{
		AuthParams:   whitebit.NewAuthParams(flexAutoInvestURL),
		InvestmentID: investmentID,
		AutoInvest:   autoInvest,
	}
}
