package main_account

import (
	"github.com/whitebit-exchange/go-sdk"
)

const historyEndpointUrl = "/api/v4/main-account/history"

type historyEndpoint struct {
	whitebit.AuthParams
	HistoryParams
}

type Details struct {
	Partial interface{} `json:"partial"`
}

type HistoryRecords struct {
	Address         string      `json:"address"`
	UniqueId        string      `json:"uniqueId"`
	CreatedAt       int         `json:"createdAt"`
	Currency        string      `json:"currency"`
	Ticker          string      `json:"ticker"`
	Method          int         `json:"method"`
	Amount          string      `json:"amount"`
	Description     interface{} `json:"description"`
	Memo            string      `json:"memo"`
	Fee             string      `json:"fee"`
	Status          int         `json:"status"`
	Network         string      `json:"network"`
	TransactionHash string      `json:"transactionHash"`
	Details         Details     `json:"details"`
	Centralized     bool        `json:"centralized,omitempty"`
}

type HistoryResult struct {
	Records []HistoryRecords `json:"records"`
	Offset  int              `json:"offset"`
	Limit   int              `json:"limit"`
	Total   int              `json:"total"`
}

type HistoryParams struct {
	TransactionMethod string   `json:"transactionMethod,omitempty"`
	Ticker            string   `json:"ticker,omitempty"`
	Addresses         []string `json:"address,omitempty"`
	UniqueId          string   `json:"uniqueId,omitempty"`
	Status            []int    `json:"status,omitempty"`
	Offset            int      `json:"offset"`
	Limit             int      `json:"limit"`
}

func newHistoryEndpoint(params HistoryParams) *historyEndpoint {
	return &historyEndpoint{HistoryParams: params, AuthParams: whitebit.NewAuthParams(historyEndpointUrl)}
}
