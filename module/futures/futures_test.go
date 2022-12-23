package futures

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/tests/mocks"
	"testing"
)

type FuturesTestSuite struct {
	client   *mocks.Client
	endpoint *endpoint
	service  *Service
	suite.Suite
}

func (s *FuturesTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
	s.endpoint = newFuturesMarketsEndpoint()

}

func (s *FuturesTestSuite) TestWithResult() {
	expectedServerResponse := MarketList{
		Success: true,
		Message: nil,
		Result: []Market{{
			TickerId:                 "ADA_PERP",
			StockCurrency:            "ADA",
			MoneyCurrency:            "USDT",
			LastPrice:                "0.313598",
			StockVolume:              "2274",
			MoneyVolume:              "7096.0",
			Bid:                      "0.313493",
			Ask:                      "0.313546",
			High:                     "0.318935",
			Low:                      "0.298932",
			ProductType:              "Perpetual",
			OpenInterest:             "604500",
			IndexPrice:               "0.3137910470166666",
			IndexName:                "Cardano",
			IndexCurrency:            "ADA",
			FundingRate:              "0.0001",
			NextFundingRateTimestamp: "1668096000000",
		}, {
			TickerId:                 "BTC_PERP",
			StockCurrency:            "BTC",
			MoneyCurrency:            "USDT",
			LastPrice:                "16542",
			StockVolume:              "5401.017",
			MoneyVolume:              "87615043.1559",
			Bid:                      "16542",
			Ask:                      "16542.5",
			High:                     "16609.6",
			Low:                      "15717.3",
			ProductType:              "Perpetual",
			OpenInterest:             "1813.65",
			IndexPrice:               "16543.886666666665",
			IndexName:                "Bitcoin",
			IndexCurrency:            "BTC",
			FundingRate:              "-0.000585680717062522",
			NextFundingRateTimestamp: "1668096000000"}},
	}

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v4/public/futures"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success":true,"message":null,"result":[
		{"ticker_id":"ADA_PERP","stock_currency":"ADA","money_currency":"USDT","last_price":"0.313598",
	"stock_volume":"2274","money_volume":"7096.0","bid":"0.313493","ask":"0.313546",
	"high":"0.318935","low":"0.298932","product_type":"Perpetual","open_interest":"604500",
	"index_price":"0.3137910470166666","index_name":"Cardano","index_currency":"ADA","funding_rate":"0.0001",
	"next_funding_rate_timestamp":"1668096000000"},
		{"ticker_id":"BTC_PERP","stock_currency":"BTC","money_currency":"USDT","last_price":"16542",
	"stock_volume":"5401.017","money_volume":"87615043.1559","bid":"16542","ask":"16542.5","high":"16609.6",
	"low":"15717.3","product_type":"Perpetual","open_interest":"1813.65","index_price":"16543.886666666665",
	"index_name":"Bitcoin","index_currency":"BTC","funding_rate":"-0.000585680717062522",
	"next_funding_rate_timestamp":"1668096000000"}]}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetFuturesMarkets()
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *FuturesTestSuite) TestInvalidResponseError() {
	expectedServerResponse := MarketList{
		Success: false,
		Message: nil,
		Result:  nil,
	}
	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v4/public/futures"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetFuturesMarkets()
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal(string(expectedResponse), string(response))
	s.Equal(expectedError, err.Error())

}

func TestFuturesTestSuite(t *testing.T) {
	suite.Run(t, new(FuturesTestSuite))
}
