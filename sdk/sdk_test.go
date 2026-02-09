package sdk

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type SDKTestSuite struct {
	suite.Suite
}

func (s *SDKTestSuite) TestNewSDKWithDefaults() {
	sdk := New("test-api-key", "test-api-secret")

	s.NotNil(sdk)
	s.NotNil(sdk.Client)

	// Public API services
	s.NotNil(sdk.Server)
	s.NotNil(sdk.Market)
	s.NotNil(sdk.Depth)
	s.NotNil(sdk.Tickers)
	s.NotNil(sdk.Symbols)
	s.NotNil(sdk.Deals)
	s.NotNil(sdk.Assets)
	s.NotNil(sdk.Fee)
	s.NotNil(sdk.Futures)
	s.NotNil(sdk.Kline)
	s.NotNil(sdk.Status)

	// Order services
	s.NotNil(sdk.OrdersSpot)
	s.NotNil(sdk.OrdersCollateral)

	// Account services
	s.NotNil(sdk.AccountTrade)
	s.NotNil(sdk.AccountCollateral)
	s.NotNil(sdk.AccountMain)

	// Additional services
	s.NotNil(sdk.Convert)
	s.NotNil(sdk.SubAccount)
	s.NotNil(sdk.SmartLending)
	s.NotNil(sdk.Mining)
}

func (s *SDKTestSuite) TestNewSDKWithCustomBaseURL() {
	customURL := "https://custom.whitebit.com"
	sdk := New("test-api-key", "test-api-secret", WithBaseURL(customURL))

	s.NotNil(sdk)
	s.NotNil(sdk.Client)
}

func (s *SDKTestSuite) TestNewSDKWithCustomTimeout() {
	customTimeout := 30 * time.Second
	sdk := New("test-api-key", "test-api-secret", WithTimeout(customTimeout))

	s.NotNil(sdk)
	s.NotNil(sdk.Client)
}

func (s *SDKTestSuite) TestNewSDKWithCustomHTTPClient() {
	customClient := &http.Client{
		Timeout: 60 * time.Second,
	}
	sdk := New("test-api-key", "test-api-secret", WithHTTPClient(customClient))

	s.NotNil(sdk)
	s.NotNil(sdk.Client)
}

func (s *SDKTestSuite) TestNewSDKWithAllOptions() {
	customURL := "https://custom.whitebit.com"
	customTimeout := 30 * time.Second
	customClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	sdk := New(
		"test-api-key",
		"test-api-secret",
		WithBaseURL(customURL),
		WithTimeout(customTimeout),
		WithHTTPClient(customClient),
	)

	s.NotNil(sdk)
	s.NotNil(sdk.Client)

	// Verify all services are initialized
	s.NotNil(sdk.Server)
	s.NotNil(sdk.Convert)
	s.NotNil(sdk.SubAccount)
	s.NotNil(sdk.SmartLending)
	s.NotNil(sdk.Mining)
}

func (s *SDKTestSuite) TestWithBaseURLOption() {
	cfg := &config{}
	opt := WithBaseURL("https://test.com")
	opt(cfg)
	s.Equal("https://test.com", cfg.BaseURL)
}

func (s *SDKTestSuite) TestWithTimeoutOption() {
	cfg := &config{}
	opt := WithTimeout(45 * time.Second)
	opt(cfg)
	s.Equal(45*time.Second, cfg.Timeout)
}

func (s *SDKTestSuite) TestWithHTTPClientOption() {
	cfg := &config{}
	customClient := &http.Client{Timeout: 100 * time.Second}
	opt := WithHTTPClient(customClient)
	opt(cfg)
	s.Equal(customClient, cfg.HTTPClient)
}

func TestSDKTestSuite(t *testing.T) {
	suite.Run(t, new(SDKTestSuite))
}
