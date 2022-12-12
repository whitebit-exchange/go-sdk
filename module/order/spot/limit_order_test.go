package spot

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/tests/mocks"
	"testing"
)

type OrderTestSuite struct {
	client  *mocks.Client
	service *Service
	suite.Suite
}

func (s *OrderTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
}

func (s *OrderTestSuite) TestLimitOrderSuccess() {
	endpoint := newLimitEndpoint(LimitOrderParams{
		Market:        "BTC_USDT",
		Amount:        "1",
		Side:          "buy",
		Price:         "3000",
		ClientOrderId: "23",
	})

	byteResponse := []byte(`{"orderId":3310300051,"clientOrderId":"25","market":"BTC_USDT","side":"buy",
		"type":"limit","timestamp":1669326096.937214,"dealMoney":"0","dealStock":"0","amount":"0.001",
		"takerFee":"0.001","makerFee":"0.001","left":"0.001","dealFee":"0","price":"3000"}`)

	serverResponse := LimitOrder{MarketOrder: MarketOrder{OrderID: 3310300051,
		ClientOrderID: "25", Market: "BTC_USDT", Side: "buy", Type: "limit", Timestamp: 1.669326096937214e+09,
		DealMoney: "0", DealStock: "0", Amount: "0.001", TakerFee: "0.001",
		MakerFee: "0.001", Left: "0.001", DealFee: "0"}, Price: "3000"}

	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/order/new"

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.CreateLimitOrder(LimitOrderParams{
		Market:        "BTC_USDT",
		Amount:        "1",
		Side:          "buy",
		Price:         "3000",
		ClientOrderId: "23",
	})

	response, _ := json.Marshal(responseJson)
	expectedResponse, _ := json.Marshal(serverResponse)
	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func TestOrderTestSuite(t *testing.T) {
	suite.Run(t, new(OrderTestSuite))
}
