package main_account

import "github.com/whitebit-exchange/go-sdk"

const (
	withdrawEndpointURL    = "/api/v4/main-account/withdraw"
	withdrawPayEndpointURL = "/api/v4/main-account/withdraw-pay"
)

type Beneficiary struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	TIN       int    `json:"tin,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Email     string `json:"email,omitempty"`
	BirthDate string `json:"birthDate,omitempty"`
}

type TravelRule struct {
	Type    string `json:"type"`
	VASP    string `json:"vasp"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type WithdrawParams struct {
	Ticker             string       `json:"ticker"`
	Amount             string       `json:"amount"`
	Address            string       `json:"address"`
	UniqueID           string       `json:"uniqueId"`
	Memo               string       `json:"memo,omitempty"`
	Provider           string       `json:"provider,omitempty"`
	Network            string       `json:"network,omitempty"`
	PartialEnable      bool         `json:"partialEnable,omitempty"`
	Beneficiary        *Beneficiary `json:"beneficiary,omitempty"`
	TravelRule         *TravelRule  `json:"travelRule,omitempty"`
	PaymentDescription string       `json:"paymentDescription,omitempty"`
}

type withdrawEndpoint struct {
	whitebit.AuthParams
	WithdrawParams
}

func newWithdrawEndpoint(params WithdrawParams) *withdrawEndpoint {
	return &withdrawEndpoint{
		AuthParams:     whitebit.NewAuthParams(withdrawEndpointURL),
		WithdrawParams: params,
	}
}

type withdrawPayEndpoint struct {
	whitebit.AuthParams
	WithdrawParams
}

func newWithdrawPayEndpoint(params WithdrawParams) *withdrawPayEndpoint {
	return &withdrawPayEndpoint{
		AuthParams:     whitebit.NewAuthParams(withdrawPayEndpointURL),
		WithdrawParams: params,
	}
}
