package tickers

import (
	"fmt"
	go_sdk "github.com/whitebit-exchange/whitebit"
	"net/url"
)

const (
	marketActivityEndpointUrl       = "/api/v2/public/ticker"
	singleMarketActivityEndpointUrl = "/api/v1/public/ticker"
	tickerEndpointUrl               = "/api/v4/public/ticker"
	tickersEndpointUrl              = "/api/v1/public/tickers"
)

type marketActivityEndpoint struct {
	go_sdk.NoAuth
}

func newMarketActivityEndpoint() *marketActivityEndpoint {
	return &marketActivityEndpoint{}
}

func (endpoint *marketActivityEndpoint) Url() string {
	return marketActivityEndpointUrl
}

type singleMarketActivityEndpoint struct {
	go_sdk.NoAuth

	Market string
}

func newSingleMarketActivityEndpoint(market string) *singleMarketActivityEndpoint {
	return &singleMarketActivityEndpoint{Market: market}
}

func (endpoint *singleMarketActivityEndpoint) Url() string {
	queryParams := url.Values{}
	queryParams.Add("market", endpoint.Market)

	return fmt.Sprintf("%s?%s", singleMarketActivityEndpointUrl, queryParams.Encode())
}

type tickerEndpoint struct {
	go_sdk.NoAuth
}

func newTickerEndpoint() *tickerEndpoint {
	return &tickerEndpoint{}
}

func (endpoint *tickerEndpoint) Url() string {
	return tickerEndpointUrl
}

type tickersEndpoint struct {
	go_sdk.NoAuth
}

func newTickersEndpoint() *tickersEndpoint {
	return &tickersEndpoint{}
}

func (endpoint *tickersEndpoint) Url() string {
	return tickersEndpointUrl
}
