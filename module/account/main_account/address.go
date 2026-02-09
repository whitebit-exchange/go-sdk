package main_account

import "github.com/whitebit-exchange/go-sdk"

const (
	addressEndpointURL     = "/api/v4/main-account/address"
	createNewAddressURL    = "/api/v4/main-account/create-new-address"
	fiatDepositURLEndpoint = "/api/v4/main-account/fiat-deposit-url"
)

type addressEndpoint struct {
	whitebit.AuthParams
	Ticker  string `json:"ticker"`
	Network string `json:"network,omitempty"`
}

func newAddressEndpoint(ticker, network string) *addressEndpoint {
	return &addressEndpoint{
		AuthParams: whitebit.NewAuthParams(addressEndpointURL),
		Ticker:     ticker,
		Network:    network,
	}
}

type createNewAddressEndpoint struct {
	whitebit.AuthParams
	Ticker  string `json:"ticker"`
	Network string `json:"network,omitempty"`
	Type    string `json:"type,omitempty"`
}

func newCreateNewAddressEndpoint(ticker, network, addressType string) *createNewAddressEndpoint {
	return &createNewAddressEndpoint{
		AuthParams: whitebit.NewAuthParams(createNewAddressURL),
		Ticker:     ticker,
		Network:    network,
		Type:       addressType,
	}
}

type Customer struct {
	FirstName string           `json:"firstName,omitempty"`
	LastName  string           `json:"lastName,omitempty"`
	Email     string           `json:"email,omitempty"`
	BirthDate string           `json:"birthDate,omitempty"`
	Address   *CustomerAddress `json:"address,omitempty"`
}

type CustomerAddress struct {
	Line1       string `json:"line1,omitempty"`
	Line2       string `json:"line2,omitempty"`
	City        string `json:"city,omitempty"`
	ZipCode     string `json:"zipCode,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
}

type fiatDepositURLEndpointParams struct {
	whitebit.AuthParams
	Ticker      string    `json:"ticker"`
	Provider    string    `json:"provider"`
	Amount      string    `json:"amount"`
	UniqueID    string    `json:"uniqueId"`
	Customer    *Customer `json:"customer,omitempty"`
	SuccessLink string    `json:"successLink,omitempty"`
	FailureLink string    `json:"failureLink,omitempty"`
	ReturnLink  string    `json:"returnLink,omitempty"`
}

func newFiatDepositURLEndpoint(ticker, provider, amount, uniqueID string, customer *Customer, successLink, failureLink, returnLink string) *fiatDepositURLEndpointParams {
	return &fiatDepositURLEndpointParams{
		AuthParams:  whitebit.NewAuthParams(fiatDepositURLEndpoint),
		Ticker:      ticker,
		Provider:    provider,
		Amount:      amount,
		UniqueID:    uniqueID,
		Customer:    customer,
		SuccessLink: successLink,
		FailureLink: failureLink,
		ReturnLink:  returnLink,
	}
}

type AddressResponse struct {
	Account  AddressAccount  `json:"account"`
	Required AddressRequired `json:"required"`
}

type AddressAccount struct {
	Address string `json:"address"`
	Memo    string `json:"memo,omitempty"`
}

type AddressRequired struct {
	FixedFee  string   `json:"fixedFee"`
	FlexFee   *FlexFee `json:"flexFee,omitempty"`
	MaxAmount string   `json:"maxAmount"`
	MinAmount string   `json:"minAmount"`
}

type FlexFee struct {
	MaxFee  string `json:"maxFee"`
	MinFee  string `json:"minFee"`
	Percent string `json:"percent"`
}

type FiatDepositURLResponse struct {
	URL string `json:"url"`
}
