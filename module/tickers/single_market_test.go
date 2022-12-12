package tickers

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/tests/mocks"
	"testing"
)

type SingleMarketTestSuite struct {
	client   *mocks.Client
	endpoint *singleMarketActivityEndpoint
	service  *Service
	suite.Suite
}

func (s *SingleMarketTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)

}

func (s *SingleMarketTestSuite) TestWithResult() {
	expectedServerResponse := SingleMarketActivityResult{
		Success: true,
		Message: nil,
		Result: SingleMarketActivity{
			Open: "16173.59", Bid: "16488.42", Ask: "16490.9", Low: "16048.17", High: "16649.43", Last: "16489.66",
			Vol: "21251.812818", Deal: "347448310.17071561", Change: "1.95"},
	}
	s.endpoint = newSingleMarketActivityEndpoint("BTC_USDT")

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v1/public/ticker?market=BTC_USDT"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success":true,"message":null,"result":{"open":"16173.59","bid":"16488.42",
		"ask":"16490.9","low":"16048.17","high":"16649.43","last":"16489.66",
		"volume":"21251.812818","deal":"347448310.17071561","change":"1.95"}}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetSingleMarketActivity("BTC_USDT")
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *SingleMarketTestSuite) TestInvalidResponseError() {
	expectedServerResponse := SingleMarketActivityResult{
		Success: false,
		Message: nil,
		Result:  SingleMarketActivity{},
	}
	s.endpoint = newSingleMarketActivityEndpoint("BTC_USDT")
	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v1/public/ticker?market=BTC_USDT"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetSingleMarketActivity("BTC_USDT")
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal(string(expectedResponse), string(response))
	s.Equal(expectedError, err.Error())

}

func TestSingleMarketTestSuite(t *testing.T) {
	suite.Run(t, new(SingleMarketTestSuite))
}
