package kline

import (
	"fmt"
	"github.com/spf13/cast"
	go_sdk "github.com/whitebit-exchange/go-sdk"
	"net/url"
)

type endpoint struct {
	go_sdk.NoAuth

	RequestUrl  string
	queryParams map[string]string
}

type QueryParams struct {
	Market   string
	Start    int64
	End      int64
	Interval string
	Limit    int
}

func newKlineEndpoint(market string) *endpoint {
	return &endpoint{
		RequestUrl: "/api/v1/public/kline",
		queryParams: map[string]string{
			"market": market,
		},
	}
}

func (h *endpoint) SetStart(start int) {
	h.queryParams["start"] = cast.ToString(start)
}

func (h *endpoint) SetEnd(end int) {
	h.queryParams["end"] = cast.ToString(end)
}

func (h *endpoint) SetInterval(interval string) {
	h.queryParams["interval"] = interval
}

func (h *endpoint) SetLimit(limit int) {
	h.queryParams["limit"] = cast.ToString(limit)
}

func (h *endpoint) Url() string {
	queryParams := url.Values{}

	for key, value := range h.queryParams {
		queryParams.Add(key, value)
	}

	return fmt.Sprintf("%s?%s", h.RequestUrl, queryParams.Encode())
}
