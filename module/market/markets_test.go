package market

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/tests/mocks"
	"testing"
)

type MarketTestSuite struct {
	client   *mocks.Client
	endpoint *infoEndpoint
	service  *Service
	suite.Suite
}

func (s *MarketTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
	s.endpoint = newMarketsInfoEndpoint()

}

func (s *MarketTestSuite) TestWithResult() {
	expectedServerResponse := Result{
		Success: true,
		Message: nil,
		Result: []Market{
			{Name: "1INCH_BTC", Stock: "1INCH", Money: "BTC", StockPrec: "0", MoneyPrec: "8", FeePrec: "6",
				MakerFee: "0.1", TakerFee: "0.1", MinAmount: "2", MinTotal: "0.0001", TradesEnabled: true},
			{Name: "1INCH_UAH", Stock: "1INCH", Money: "UAH", StockPrec: "2", MoneyPrec: "6", FeePrec: "6",
				MakerFee: "0.1", TakerFee: "0.1", MinAmount: "2", MinTotal: "140", TradesEnabled: true}},
	}

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v2/public/markets"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success":true,"message":null,"result":[
		{"name":"1INCH_BTC","stock":"1INCH","money":"BTC","stockPrec":"0","moneyPrec":"8","feePrec":"6",
		"makerFee":"0.1","takerFee":"0.1","minAmount":"2","minTotal":"0.0001","tradesEnabled":true},
		{"name":"1INCH_UAH","stock":"1INCH","money":"UAH","stockPrec":"2","moneyPrec":"6","feePrec":"6",
		"makerFee":"0.1","takerFee":"0.1","minAmount":"2","minTotal":"140","tradesEnabled":true}]}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetMarkets()
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *MarketTestSuite) TestInvalidResponseError() {
	expectedServerResponse := Result{
		Success: false,
		Message: nil,
		Result:  nil,
	}
	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v2/public/markets"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetMarkets()
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal(string(expectedResponse), string(response))
	s.Equal(expectedError, err.Error())

}

func TestMarketTestSuite(t *testing.T) {
	suite.Run(t, new(MarketTestSuite))
}
