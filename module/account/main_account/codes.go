package main_account

import (
	"github.com/whitebit-exchange/go-sdk"
)

const (
	codesEndpointUrl        = "/api/v4/main-account/codes"
	codesApplyEndpointUrl   = "/api/v4/main-account/codes/apply"
	codesMyEndpointUrl      = "/api/v4/main-account/codes/my"
	codesHistoryEndpointUrl = "/api/v4/main-account/codes/history"
)

type codeEndpoint struct {
	whitebit.AuthParams
	Ticker      string `json:"ticker"`
	Amount      string `json:"amount"`
	Passphrase  string `json:"passphrase,omitempty"`
	Description string `json:"description,omitempty"`
}

type Code struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	ExternalId string `json:"external_id"`
}

func newCodeEndpoint(ticker string, amount string, pass string, description string) *codeEndpoint {
	return &codeEndpoint{Ticker: ticker, Amount: amount, Passphrase: pass, Description: description, AuthParams: whitebit.NewAuthParams(codesEndpointUrl)}
}

type codeApplyEndpoint struct {
	whitebit.AuthParams
	Code       string `json:"code"`
	Passphrase string `json:"passphrase,omitempty"`
}

type CodeApply struct {
	Message    string `json:"message"`
	Ticker     string `json:"ticker"`
	Amount     string `json:"amount"`
	ExternalId string `json:"external_id"`
}

func newCodeApplyEndpoint(code string, pass string) *codeApplyEndpoint {
	return &codeApplyEndpoint{Code: code, Passphrase: pass, AuthParams: whitebit.NewAuthParams(codesApplyEndpointUrl)}
}

type codeMyEndpoint struct {
	whitebit.AuthParams
	Limit  int64 `json:"limit,omitempty"`
	Offset int64 `json:"offset,omitempty"`
}

type Data []struct {
	Amount     string `json:"amount"`
	Code       string `json:"code"`
	Date       int    `json:"date"`
	Status     string `json:"status"`
	Ticker     string `json:"ticker"`
	ExternalId string `json:"external_id"`
}

type CodeMy struct {
	Total  int  `json:"total"`
	Data   Data `json:"data"`
	Limit  int  `json:"limit"`
	Offset int  `json:"offset"`
}

func newCodeMyEndpoint(limit int64, offset int64) *codeMyEndpoint {
	return &codeMyEndpoint{Limit: limit, Offset: offset, AuthParams: whitebit.NewAuthParams(codesMyEndpointUrl)}
}

type codeHistoryEndpoint struct {
	whitebit.AuthParams
	Limit  int64 `json:"limit,omitempty"`
	Offset int64 `json:"offset,omitempty"`
}

type CodeHistory struct {
	Total  int  `json:"total"`
	Data   Data `json:"data"`
	Limit  int  `json:"limit"`
	Offset int  `json:"offset"`
}

func newCodeHistoryEndpoint(limit int64, offset int64) *codeHistoryEndpoint {
	return &codeHistoryEndpoint{Limit: limit, Offset: offset, AuthParams: whitebit.NewAuthParams(codesHistoryEndpointUrl)}
}
