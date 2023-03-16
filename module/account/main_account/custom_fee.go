package main_account

import (
	"github.com/whitebit-exchange/go-sdk"
)

const customFeeEndpointUrl = "/api/v4/market/fee"
const customFeeByMarketEndpointUrl = "/api/v4/market/fee/single"

type MyFeeByMarket struct {
	Error interface{} `json:"error"`
	Taker string      `json:"taker"`
	Maker string      `json:"maker"`
}

type myFeeByMarketEndpoint struct {
	whitebit.AuthParams
	Market string `json:"market"`
}

type CustomFee struct {
	Error     interface{}         `json:"error,omitempty"`
	Taker     string              `json:"taker,omitempty"`
	Maker     string              `json:"maker,omitempty"`
	CustomFee map[string][]string `json:"custom_fee,omitempty"`
}

type customFeeEndpoint struct {
	whitebit.AuthParams
}

func newCustomFeeEndpoint() *customFeeEndpoint {
	return &customFeeEndpoint{AuthParams: whitebit.NewAuthParams(customFeeEndpointUrl)}
}

func newMyFeeByMarketEndpoint(market string) *myFeeByMarketEndpoint {
	return &myFeeByMarketEndpoint{AuthParams: whitebit.NewAuthParams(customFeeByMarketEndpointUrl), Market: market}
}
