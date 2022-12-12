package depth

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/tests/mocks"
	"testing"
)

type OrderBookTestSuite struct {
	client   *mocks.Client
	service  *Service
	endpoint *orderBookEndpoint
	suite.Suite
}

func (s *OrderBookTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
}

func (s *OrderBookTestSuite) TestWithResult() {
	expectedResult := OrderBook{
		Time: 1668982941,
		AsksAndBids: AsksAndBids{
			Asks: []Pair{{"16325.8", "0.051914"}, {"16327", "0.134383"}, {"16328.2", "0.294826"}},
			Bids: []Pair{{"16324.5", "0.06358"}, {"16323.2", "0.581937"}, {"16322", "0.601753"}}},
	}

	expectedResponse, _ := json.Marshal(expectedResult)

	s.endpoint = newOrderBookEndpoint("BTC_USDT")

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v4/public/orderbook/BTC_USDT?"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"timestamp":1668982941,
		"asks":[["16325.8","0.051914"],["16327","0.134383"],["16328.2","0.294826"]],
		"bids":[["16324.5","0.06358"],["16323.2","0.581937"],["16322","0.601753"]]}
	`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetOrderBook(OrderBookOptions{Market: "BTC_USDT"})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *OrderBookTestSuite) TestWithLimitResult() {
	expectedResult := OrderBook{
		Time: 1668982941,
		AsksAndBids: AsksAndBids{
			Asks: []Pair{{"16325.8", "0.051914"}},
			Bids: []Pair{{"16324.5", "0.06358"}}},
	}

	expectedResponse, _ := json.Marshal(expectedResult)

	s.endpoint = newOrderBookEndpoint("BTC_USDT")
	s.endpoint.SetLimit(1)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v4/public/orderbook/BTC_USDT?limit=1"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"timestamp":1668982941,
		"asks":[["16325.8","0.051914"]],
		"bids":[["16324.5","0.06358"]]}
	`)

	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetOrderBook(OrderBookOptions{Market: "BTC_USDT", Limit: 1})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *OrderBookTestSuite) TestWithLevelResult() {
	expectedResult := OrderBook{
		Time: 1669021356,
		AsksAndBids: AsksAndBids{
			Asks: []Pair{{"16050", "7.66301"}},
			Bids: []Pair{{"16016", "0.29048"}}},
	}

	expectedResponse, _ := json.Marshal(expectedResult)

	s.endpoint = newOrderBookEndpoint("BTC_USDT")
	s.endpoint.SetLimit(1)
	s.endpoint.SetLevel(2)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v4/public/orderbook/BTC_USDT?level=2&limit=1"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"timestamp":1669021356,
		"asks":[["16050","7.66301"]],
		"bids":[["16016","0.29048"]]}
	`)

	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetOrderBook(OrderBookOptions{Market: "BTC_USDT", Limit: 1, Level: 2})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *OrderBookTestSuite) TestWithBadLimitResult() {
	expectedResult := OrderBook{
		Time: 1669021356,
		AsksAndBids: AsksAndBids{
			Asks: []Pair{},
			Bids: []Pair{}},
	}

	expectedResponse, _ := json.Marshal(expectedResult)

	s.endpoint = newOrderBookEndpoint("BTC_USDT")
	s.endpoint.SetLimit(-1)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v4/public/orderbook/BTC_USDT?limit=-1"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"timestamp":1669021356,"asks":[],"bids":[]}`)

	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetOrderBook(OrderBookOptions{Market: "BTC_USDT", Limit: -1})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *OrderBookTestSuite) TestWithBadMarket() {
	expectedResult := OrderBook{
		Time: 0,
		AsksAndBids: AsksAndBids{
			Asks: nil,
			Bids: nil},
	}

	s.endpoint = newOrderBookEndpoint("BTC_USDT")
	expectedResponse, _ := json.Marshal(expectedResult)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v4/public/orderbook/BTC_USDT?"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"timestamp":0,"asks":null,"bids":null}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetOrderBook(OrderBookOptions{Market: "BTC_USDT"})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(error(nil), err)

}

func (s *OrderBookTestSuite) TestInvalidResponseError() {
	expectedResult := OrderBook{
		Time: 0,
		AsksAndBids: AsksAndBids{
			Asks: nil,
			Bids: nil},
	}

	s.endpoint = newOrderBookEndpoint("BTC_USDT")
	expectedResponse, _ := json.Marshal(expectedResult)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v4/public/orderbook/BTC_USDT?"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"timestamp":0,"asks":None,"bids":null}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetOrderBook(OrderBookOptions{Market: "BTC_USDT"})
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'N' looking for beginning of value"
	s.Equal(string(expectedResponse), string(response))
	s.Equal(expectedError, err.Error())

}

func TestOrderBookTestSuite(t *testing.T) {
	suite.Run(t, new(OrderBookTestSuite))
}
