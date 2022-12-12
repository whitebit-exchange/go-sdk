package depth

import (
	"fmt"
	"github.com/spf13/cast"
	go_sdk "github.com/whitebit-exchange/whitebit"
	"net/url"
)

type depthEndpoint struct {
	go_sdk.NoAuth

	Market     string
	RequestUrl string
}

func newDepthEndpoint(market string) *depthEndpoint {
	return &depthEndpoint{Market: market, RequestUrl: "/api/v2/public/depth/"}
}

func (d *depthEndpoint) Url() string {
	return fmt.Sprintf("%s%s", d.RequestUrl, d.Market)
}

type orderBookEndpoint struct {
	go_sdk.NoAuth

	Market      string
	RequestUrl  string
	queryParams map[string]string
}

func newOrderBookEndpoint(market string) *orderBookEndpoint {
	return &orderBookEndpoint{
		Market:      market,
		RequestUrl:  "/api/v4/public/orderbook/",
		queryParams: make(map[string]string),
	}
}

func (d *orderBookEndpoint) SetLimit(limit int) {
	d.queryParams["limit"] = cast.ToString(limit)
}

func (d *orderBookEndpoint) SetLevel(level int) {
	d.queryParams["level"] = cast.ToString(level)
}

func (d *orderBookEndpoint) Url() string {
	queryParams := url.Values{}

	for key, value := range d.queryParams {
		queryParams.Add(key, value)
	}

	return fmt.Sprintf("%s%s?%s", d.RequestUrl, d.Market, queryParams.Encode())
}
