package deal

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/tests/mocks"
	"testing"
)

type HistoryTestSuite struct {
	client   *mocks.Client
	service  *Service
	endpoint *tradeHistoryEndpoint
	suite.Suite
}

func (s *HistoryTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
}

func (s *HistoryTestSuite) TestWithResult() {
	expectedResult := TradeHistoryResult{
		Success: true,
		Message: nil,
		Result: []Deals{{Id: 2378100024, Time: 1668974839.944864, Price: "16572.65", Amount: "0.132096", Type: "sell"},
			{Id: 2378100023, Time: 1668974839.938267, Price: "16572.65", Amount: "0.058624", Type: "sell"},
			{Id: 2378100016, Time: 1668974839.653658, Price: "16572.65", Amount: "0.094976", Type: "sell"}}}

	expectedResponse, _ := json.Marshal(expectedResult)

	s.endpoint = newTradeHistoryEndpoint("BTC_USDT", 1)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v1/public/history?lastId=1&market=BTC_USDT"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{
		"success":true,
		"message":null,
		"result":[{"id":2378100024,"time":1668974839.944864,"price":"16572.65","amount":"0.132096","type":"sell"},
		{"id":2378100023,"time":1668974839.938267,"price":"16572.65","amount":"0.058624","type":"sell"},
		{"id":2378100016,"time":1668974839.653658,"price":"16572.65","amount":"0.094976","type":"sell"}]}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetTradeHistory(TradeHistoryOptions{Market: "BTC_USDT", LastId: 1})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *HistoryTestSuite) TestWithLimitResult() {
	expectedResult := TradeHistoryResult{
		Success: true,
		Message: nil,
		Result:  []Deals{{Id: 2378100024, Time: 1668974839.944864, Price: "16572.65", Amount: "0.132096", Type: "sell"}}}

	expectedResponse, _ := json.Marshal(expectedResult)

	s.endpoint = newTradeHistoryEndpoint("BTC_USDT", 1)
	s.endpoint.SetLimit(1)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v1/public/history?lastId=1&limit=1&market=BTC_USDT"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{
		"success":true,
		"message":null,
		"result":[{"id":2378100024,"time":1668974839.944864,"price":"16572.65","amount":"0.132096","type":"sell"}]}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetTradeHistory(TradeHistoryOptions{Market: "BTC_USDT", LastId: 1, Limit: 1})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *HistoryTestSuite) TestWithSellResult() {
	expectedResult := TradeHistoryResult{
		Success: true,
		Message: nil,
		Result: []Deals{{Id: 2378100024, Time: 1668974839.944864, Price: "16572.65", Amount: "0.132096", Type: "sell"},
			{Id: 2378100023, Time: 1668974839.938267, Price: "16572.65", Amount: "0.058624", Type: "sell"},
			{Id: 2378100016, Time: 1668974839.653658, Price: "16572.65", Amount: "0.094976", Type: "sell"}}}

	expectedResponse, _ := json.Marshal(expectedResult)

	s.endpoint = newTradeHistoryEndpoint("BTC_USDT", 1)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v1/public/history?lastId=1&market=BTC_USDT"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{
		"success":true,
		"message":null,
		"result":[{"id":2378100024,"time":1668974839.944864,"price":"16572.65","amount":"0.132096","type":"sell"},
		{"id":2378100023,"time":1668974839.938267,"price":"16572.65","amount":"0.058624","type":"sell"},
		{"id":2378100016,"time":1668974839.653658,"price":"16572.65","amount":"0.094976","type":"sell"}]}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetTradeHistory(TradeHistoryOptions{Market: "BTC_USDT", LastId: 1})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *HistoryTestSuite) TestWithBadLastIdResult() {
	expectedResult := TradeHistoryResult{
		Success: false,
		Message: nil,
		Result:  nil}

	s.endpoint = newTradeHistoryEndpoint("BTC_USDT", -1)
	expectedResponse, _ := json.Marshal(expectedResult)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v1/public/history?lastId=-1&market=BTC_USDT"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success":false,"message":null,"result":null}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetTradeHistory(TradeHistoryOptions{Market: "BTC_USDT", LastId: -1})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(error(nil), err)

}

func (s *HistoryTestSuite) TestBadMarketError() {
	expectedResult := TradeHistoryResult{
		Success: false,
		Message: map[string][]string{"market": {"Market is not available."}},
		Result:  nil}

	s.endpoint = newTradeHistoryEndpoint("BTC_US_DT", 1)
	expectedResponse, _ := json.Marshal(expectedResult)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v1/public/history?lastId=1&market=BTC_US_DT"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success":false,"message":{"market":["Market is not available."]},"result":null}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetTradeHistory(TradeHistoryOptions{Market: "BTC_US_DT", LastId: 1})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(error(nil), err)

}

func TestHistoryTestSuite(t *testing.T) {
	suite.Run(t, new(HistoryTestSuite))
}
