package deal

import (
	"fmt"
	"github.com/spf13/cast"
	go_sdk "github.com/whitebit-exchange/whitebit"
	"net/url"
)

type dealsEndpoint struct {
	go_sdk.NoAuth

	Market     string
	Type       string
	RequestUrl string
}

func newDealsEndpoint(market string) *dealsEndpoint {
	return &dealsEndpoint{Market: market, RequestUrl: "/api/v4/public/trades/"}
}

func (d *dealsEndpoint) SetType(types string) {
	d.Type = types
}

func (d *dealsEndpoint) Url() string {
	queryParams := url.Values{}
	if d.Type != "" {
		queryParams.Add("type", d.Type)
	}
	return fmt.Sprintf("%s%s?%s", d.RequestUrl, d.Market, queryParams.Encode())
}

type tradeHistoryEndpoint struct {
	go_sdk.NoAuth

	RequestUrl  string
	queryParams map[string]string
}

func newTradeHistoryEndpoint(market string, lastId int64) *tradeHistoryEndpoint {
	return &tradeHistoryEndpoint{
		RequestUrl: "/api/v1/public/history",
		queryParams: map[string]string{
			"market": market,
			"lastId": cast.ToString(lastId),
		},
	}
}

func (h *tradeHistoryEndpoint) SetLimit(limit int) {
	h.queryParams["limit"] = cast.ToString(limit)
}

func (h *tradeHistoryEndpoint) Url() string {
	queryParams := url.Values{}
	for key, value := range h.queryParams {
		queryParams.Add(key, value)
	}

	return fmt.Sprintf("%s?%s", h.RequestUrl, queryParams.Encode())
}
