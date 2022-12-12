package tickers

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/tests/mocks"
	"testing"
)

type TickersTestSuite struct {
	client   *mocks.Client
	endpoint *tickersEndpoint
	service  *Service
	suite.Suite
}

func (s *TickersTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
	s.endpoint = newTickersEndpoint()
}

func (s *TickersTestSuite) TestWithResult() {
	expectedServerResponse := Result{
		Success: true,
		Message: map[string][]string(nil),
		Result: Tickers{
			"1INCH_BTC": TickerStatus{SingleMarketActivity: SingleMarketActivity{
				Open: "", Bid: "0.0000326", Ask: "0.00003263", Low: "0.00003211", High: "0.0000332", Last: "0.00003263",
				Vol: "", Deal: "0.61087781", Change: "0.80"}, At: 1669220500},
			"1INCH_UAH": TickerStatus{SingleMarketActivity: SingleMarketActivity{
				Open: "", Bid: "21.173251", Ask: "22.09536", Low: "21.000189", High: "21.951546", Last: "21.620607",
				Vol: "", Deal: "242274.98124506", Change: "2.82"}, At: 1669220500}}}

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v1/public/tickers"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success":true,"message":null,"result":{"1INCH_BTC":{
		"ticker":{"bid":"0.0000326","ask":"0.00003263","low":"0.00003211","high":"0.0000332",
		"last":"0.00003263","vol":"18691","deal":"0.61087781","change":"0.80"},"at":1669220500},
		"1INCH_UAH":{"ticker":{"bid":"21.173251","ask":"22.09536","low":"21.000189","high":"21.951546",
		"last":"21.620607","vol":"11221.47","deal":"242274.98124506","change":"2.82"},"at":1669220500}}}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetTickers()
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *TickersTestSuite) TestInvalidResponseError() {
	expectedServerResponse := Result{
		Success: false,
		Message: map[string][]string(nil),
		Result:  nil,
	}
	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v1/public/tickers"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetTickers()
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal(string(expectedResponse), string(response))
	s.Equal(expectedError, err.Error())
}

func TestTickersTestSuite(t *testing.T) {
	suite.Run(t, new(TickersTestSuite))
}
