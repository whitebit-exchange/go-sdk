package market

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/tests/mocks"
	"testing"
)

type CollateralMarketTestSuite struct {
	client   *mocks.Client
	endpoint *collateralEndpoint
	service  *Service
	suite.Suite
}

func (s *CollateralMarketTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
	s.endpoint = newCollateralMarketsEndpoint()

}

func (s *CollateralMarketTestSuite) TestWithResult() {
	expectedServerResponse := CollateralMarketsResult{
		Success: true,
		Message: nil,
		Result: []string{"ADA_BTC", "ADA_PERP", "ADA_USDT", "APE_USDT", "AVAX_USDT", "BCH_USDT", "BTC_PERP",
			"BTC_USDT", "DOGE_USDT", "DOT_USDT", "EOS_USDT", "ETH_BTC", "ETH_PERP", "ETH_USDT", "LINK_USDT",
			"LTC_BTC", "LTC_USDT", "MATIC_USDT", "NEAR_USDT", "SHIB_USDT", "SOL_USDT", "TRX_USDT", "USDC_USDT",
			"XLM_USDT", "XRP_BTC", "XRP_PERP", "XRP_USDT"}}

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v4/public/collateral/markets"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success":true,"message":null,"result":["ADA_BTC","ADA_PERP","ADA_USDT","APE_USDT",
		"AVAX_USDT","BCH_USDT","BTC_PERP","BTC_USDT","DOGE_USDT","DOT_USDT","EOS_USDT","ETH_BTC","ETH_PERP","ETH_USDT",
		"LINK_USDT","LTC_BTC","LTC_USDT","MATIC_USDT","NEAR_USDT","SHIB_USDT","SOL_USDT","TRX_USDT","USDC_USDT",
		"XLM_USDT","XRP_BTC","XRP_PERP","XRP_USDT"]}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetCollateralMarkets()
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *CollateralMarketTestSuite) TestInvalidResponseError() {
	expectedServerResponse := CollateralMarketsResult{
		Success: false,
		Message: nil,
		Result:  nil,
	}
	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v4/public/collateral/markets"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetCollateralMarkets()
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal(string(expectedResponse), string(response))
	s.Equal(expectedError, err.Error())

}

func TestCollateralMarketTestSuite(t *testing.T) {
	suite.Run(t, new(CollateralMarketTestSuite))
}
