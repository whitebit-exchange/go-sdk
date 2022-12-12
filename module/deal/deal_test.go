package deal

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/whitebit/tests/mocks"
	"testing"
)

type DealTestSuite struct {
	client   *mocks.Client
	service  *Service
	endpoint *dealsEndpoint
	suite.Suite
}

func (s *DealTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
	s.endpoint = newDealsEndpoint("BTC_USDT")
}

func (s *DealTestSuite) TestWithResult() {
	expectedResult := []Deal{{
		TradeID: 2370420409,
		Price:   "16771.58",
		Amount:  "0.008732",
		Volume:  "146.44",
		Time:    1668775571,
		Type:    "buy",
	}, {
		TradeID: 2372573653,
		Price:   "16782.98",
		Amount:  "0.065792",
		Volume:  "1104.18",
		Time:    1668769708,
		Type:    "sell",
	}, {
		TradeID: 2372573652,
		Price:   "16782.98",
		Amount:  "0.039349",
		Volume:  "660.39",
		Time:    1668769708,
		Type:    "sell",
	}}

	expectedResponse, _ := json.Marshal(expectedResult)

	byteResponse := []byte(`[{"tradeID":2370420409,"price":"16771.58","quote_volume":"0.008732",
	"base_volume":"146.44","trade_timestamp":1668775571,"type":"buy"},
	{"tradeID":2372573653,"price":"16782.98","quote_volume":"0.065792","base_volume":"1104.18",
	"trade_timestamp":1668769708,"type":"sell"},{"tradeID":2372573652,"price":"16782.98","quote_volume":"0.039349",
	"base_volume":"660.39","trade_timestamp":1668769708,"type":"sell"}]`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	options := Options{Market: "BTC_USDT"}
	responseJson, err := s.service.GetDeals(options)
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *DealTestSuite) TestWithSellResult() {
	expectedResult := []Deal{{
		TradeID: 2372573653,
		Price:   "16782.98",
		Amount:  "0.065792",
		Volume:  "1104.18",
		Time:    1668769708,
		Type:    "sell",
	}, {
		TradeID: 2372573652,
		Price:   "16782.98",
		Amount:  "0.039349",
		Volume:  "660.39",
		Time:    1668769708,
		Type:    "sell",
	}}

	expectedResponse, _ := json.Marshal(expectedResult)

	byteResponse := []byte(`[{"tradeID":2372573653,"price":"16782.98","quote_volume":"0.065792","base_volume":"1104.18",
	"trade_timestamp":1668769708,"type":"sell"},{"tradeID":2372573652,"price":"16782.98","quote_volume":"0.039349",
	"base_volume":"660.39","trade_timestamp":1668769708,"type":"sell"}]`)
	s.endpoint.SetType("sell")
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetDeals(Options{Market: "BTC_USDT", Type: "sell"})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *DealTestSuite) TestWithBuyResult() {
	expectedResult := []Deal{{
		TradeID: 2370420409,
		Price:   "16771.58",
		Amount:  "0.008732",
		Volume:  "146.44",
		Time:    1668775571,
		Type:    "buy",
	}}

	expectedResponse, _ := json.Marshal(expectedResult)

	byteResponse := []byte(`[{"tradeID":2370420409,"price":"16771.58","quote_volume":"0.008732",
	"base_volume":"146.44","trade_timestamp":1668775571,"type":"buy"}]`)
	s.endpoint.SetType("buy")
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetDeals(Options{Market: "BTC_USDT", Type: "buy"})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *DealTestSuite) TestInvalidResponseError() {
	expectedResponse := "[]"

	byteResponse := []byte(`[{"tradeID":2370420409 "price":"16771.58","quote_volume":"0.008732",
	"base_volume":"146.44","trade_timestamp":1668775571,"type":"buy"}]`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetDeals(Options{Market: "BTC_USDT"})
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character '\"' after object key:value pair"
	s.Equal(expectedResponse, string(response))
	s.Equal(expectedError, err.Error())

}

func TestDealTestSuite(t *testing.T) {
	suite.Run(t, new(DealTestSuite))
}
