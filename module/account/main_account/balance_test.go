package main_account

import (
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *MainAccountTestSuite) TestMainBalanceWithResult() {
	expectedServerResponse := BalanceResult{
		"BTC": MainBalanceResult{MainBalance: "476.7151833"},
		"ETH": MainBalanceResult{MainBalance: "0"},
		"USD": MainBalanceResult{MainBalance: "0"}}

	endpoint := newBalanceEndpoint("")
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/main-account/balance"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{
    "BTC": {
        "main_balance": "476.7151833"
    },
    "ETH": {
        "main_balance": "0"
    },
    "USD": {
        "main_balance": "0"
    }}
`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetMainBalance()

	s.Equal(expectedServerResponse, responseJson)
	s.Equal(err, error(nil))

}

func (s *MainAccountTestSuite) TestMainBalanceTickerWithResult() {
	expectedServerResponse := MainBalanceResult{MainBalance: "476.7151833"}

	endpoint := newBalanceEndpoint("BTC")
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/main-account/balance"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"main_balance":"476.7151833"}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetMainBalanceTicker("BTC")

	s.Equal(expectedServerResponse, responseJson)
	s.Equal(err, error(nil))

}
