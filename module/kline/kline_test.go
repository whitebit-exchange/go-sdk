package kline

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/tests/mocks"
	"testing"
)

type KlineTestSuite struct {
	client   *mocks.Client
	endpoint *endpoint
	service  *Service
	suite.Suite
}

func (s *KlineTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)

}

func (s *KlineTestSuite) TestWithResult() {
	expectedServerResponse := Result{
		Success: true,
		Message: map[string][]string(nil),
		Result: []Kline{
			{1667469600, "20319.48", "20267.69", "20337.79", "20260.09", "603.187972", "12239132.92483017"},
			{1667473200, "20267.69", "20129.73", "20278.34", "20058.85", "896.024204", "18080673.60376878"},
			{1667476800, "20129.73", "20107.08", "20178.14", "20049.82", "893.027185", "17958510.75884111"},
			{1667480400, "20108.59", "20108.59", "20209.97", "20096.52", "956.102104", "19265622.69211793"},
			{1667484000, "20107.08", "20316.43", "20337.79", "20088.98", "1170.173097", "23683142.02525501"}}}

	s.endpoint = newKlineEndpoint("BTC_USDT")
	s.endpoint.SetStart(1667469600)
	s.endpoint.SetEnd(1667487600)
	s.endpoint.SetLimit(5)

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v1/public/kline?end=1667487600&limit=5&market=BTC_USDT&start=1667469600"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success":true,"message":null,"result":[[1667469600,"20319.48","20267.69","20337.79",
		"20260.09","603.187972","12239132.92483017"],[1667473200,"20267.69","20129.73","20278.34","20058.85",
	"896.024204","18080673.60376878"],[1667476800,"20129.73","20107.08","20178.14","20049.82","893.027185",
	"17958510.75884111"],[1667480400,"20108.59","20108.59","20209.97","20096.52","956.102104","19265622.69211793"],
	[1667484000,"20107.08","20316.43","20337.79","20088.98","1170.173097","23683142.02525501"]]}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetKline(Options{Market: "BTC_USDT", Start: 1667469600, End: 1667487600, Limit: 5})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *KlineTestSuite) TestInvalidResponseError() {
	expectedServerResponse := Result{
		Success: false,
		Message: nil,
		Result:  nil,
	}
	expectedResponse, _ := json.Marshal(expectedServerResponse)
	s.endpoint = newKlineEndpoint("ETH_USDT")

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v1/public/kline?market=ETH_USDT"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetKline(Options{Market: "ETH_USDT"})
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal(string(expectedResponse), string(response))
	s.Equal(expectedError, err.Error())

}

func TestKlineTestSuite(t *testing.T) {
	suite.Run(t, new(KlineTestSuite))
}
