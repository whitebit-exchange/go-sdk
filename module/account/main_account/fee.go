package main_account

import (
	"github.com/whitebit-exchange/go-sdk"
)

const feeEndpointUrl = "/api/v4/main-account/fee"

type Deposit struct {
	MinFlex     string `json:"minFlex"`
	MaxFlex     string `json:"maxFlex"`
	PercentFlex string `json:"percentFlex"`
	Fixed       string `json:"fixed"`
	MinAmount   string `json:"minAmount"`
	MaxAmount   string `json:"maxAmount"`
}

type Withdraw struct {
	MinFlex     string `json:"minFlex"`
	MaxFlex     string `json:"maxFlex"`
	PercentFlex string `json:"percentFlex"`
	Fixed       string `json:"fixed"`
	MinAmount   string `json:"minAmount"`
	MaxAmount   string `json:"maxAmount"`
}

type Fee struct {
	Ticker      string `json:"ticker"`
	Name        string `json:"name"`
	CanDeposit  bool   `json:"can_deposit"`
	CanWithdraw bool   `json:"can_withdraw"`
	Deposit     `json:"deposit"`
	Withdraw    `json:"withdraw"`
}

type feeEndpoint struct {
	whitebit.AuthParams
}

func newFeeEndpoint() *feeEndpoint {
	return &feeEndpoint{AuthParams: whitebit.NewAuthParams(feeEndpointUrl)}
}
