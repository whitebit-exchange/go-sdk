package sdk

import (
	"net"
	"net/http"
	"time"

	whitebit "github.com/whitebit-exchange/go-sdk"
	accountCollateral "github.com/whitebit-exchange/go-sdk/module/account/collateral"
	accountMain "github.com/whitebit-exchange/go-sdk/module/account/main_account"
	accountTrade "github.com/whitebit-exchange/go-sdk/module/account/trade"
	"github.com/whitebit-exchange/go-sdk/module/assets"
	"github.com/whitebit-exchange/go-sdk/module/convert"
	"github.com/whitebit-exchange/go-sdk/module/deal"
	"github.com/whitebit-exchange/go-sdk/module/depth"
	"github.com/whitebit-exchange/go-sdk/module/fee"
	"github.com/whitebit-exchange/go-sdk/module/futures"
	"github.com/whitebit-exchange/go-sdk/module/kline"
	"github.com/whitebit-exchange/go-sdk/module/market"
	"github.com/whitebit-exchange/go-sdk/module/mining"
	orderCollateral "github.com/whitebit-exchange/go-sdk/module/order/collateral"
	orderSpot "github.com/whitebit-exchange/go-sdk/module/order/spot"
	"github.com/whitebit-exchange/go-sdk/module/server"
	"github.com/whitebit-exchange/go-sdk/module/smartlending"
	"github.com/whitebit-exchange/go-sdk/module/status"
	"github.com/whitebit-exchange/go-sdk/module/subaccount"
	"github.com/whitebit-exchange/go-sdk/module/symbol"
	"github.com/whitebit-exchange/go-sdk/module/tickers"
)

type SDK struct {
	Client *whitebit.Whitebit

	// Public API
	Server  *server.Service
	Market  *market.Service
	Depth   *depth.Service
	Tickers *tickers.Service
	Symbols *symbol.Service
	Deals   *deal.Service
	Assets  *assets.Service
	Fee     *fee.Service
	Futures *futures.Service
	Kline   *kline.Service
	Status  *status.Service

	// Orders
	OrdersSpot       *orderSpot.Service
	OrdersCollateral *orderCollateral.Service

	// Account
	AccountTrade      *accountTrade.Service
	AccountCollateral *accountCollateral.Service
	AccountMain       *accountMain.Service

	// Additional Services
	Convert      *convert.Service
	SubAccount   *subaccount.Service
	SmartLending *smartlending.Service
	Mining       *mining.Service
}

type config struct {
	BaseURL    string
	HTTPClient *http.Client
	Timeout    time.Duration
}

type Option func(*config)

func WithBaseURL(baseURL string) Option { return func(c *config) { c.BaseURL = baseURL } }

func WithHTTPClient(cli *http.Client) Option { return func(c *config) { c.HTTPClient = cli } }

func WithTimeout(d time.Duration) Option { return func(c *config) { c.Timeout = d } }

func New(apiKey, apiSecret string, opts ...Option) *SDK {
	cfg := config{
		BaseURL: "https://whitebit.com",
		Timeout: 15 * time.Second,
	}
	for _, o := range opts {
		o(&cfg)
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		transport := &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           (&net.Dialer{Timeout: 10 * time.Second, KeepAlive: 60 * time.Second}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          200,
			MaxIdleConnsPerHost:   100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
		httpClient = &http.Client{Timeout: cfg.Timeout, Transport: transport}
	}

	wb := whitebit.NewClientWithHTTPClient(apiKey, apiSecret, httpClient, cfg.BaseURL)

	sdk := &SDK{Client: wb}

	sdk.Server = server.NewService(wb)
	sdk.Market = market.NewService(wb)
	sdk.Depth = depth.NewService(wb)
	sdk.Tickers = tickers.NewService(wb)
	sdk.Symbols = symbol.NewService(wb)
	sdk.Deals = deal.NewService(wb)
	sdk.Assets = assets.NewService(wb)
	sdk.Fee = fee.NewService(wb)
	sdk.Futures = futures.NewService(wb)
	sdk.Kline = kline.NewService(wb)
	sdk.Status = status.NewService(wb)

	sdk.OrdersSpot = orderSpot.NewService(wb)
	sdk.OrdersCollateral = orderCollateral.NewService(wb)

	sdk.AccountTrade = accountTrade.NewService(wb)
	sdk.AccountCollateral = accountCollateral.NewService(wb)
	sdk.AccountMain = accountMain.NewService(wb)

	// Additional services
	sdk.Convert = convert.NewService(wb)
	sdk.SubAccount = subaccount.NewService(wb)
	sdk.SmartLending = smartlending.NewService(wb)
	sdk.Mining = mining.NewService(wb)

	return sdk
}
