package spot

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/whitebit"
)

func (s *OrderTestSuite) TestStopLimitOrderSuccess() {

	params := StopLimitOrderParams{
		Market:          "BTC_USDT",
		Amount:          "0.001",
		Side:            "buy",
		Price:           "3000",
		ActivationPrice: "3300",
		ClientOrderId:   "71",
	}

	endpoint := newStopLimitEndpoint(params)

	byteResponse := []byte(`{"orderId":3310726581,"clientOrderId":"71","market":"BTC_USDT","side":"buy",
		"type":"stop limit","timestamp":1669334361.97947,"dealMoney":"0","dealStock":"0","amount":"0.001",
		"takerFee":"0.001","makerFee":"0.001","left":"0.001","dealFee":"0","activation_price":"3300","price":"3000"}`)

	serverResponse := StopLimitOrder{StopMarketOrder: StopMarketOrder{MarketOrder: MarketOrder{OrderID: 3310726581,
		ClientOrderID: "71", Market: "BTC_USDT", Side: "buy", Type: "stop limit", Timestamp: 1.66933436197947e+09,
		DealMoney: "0", DealStock: "0", Amount: "0.001", TakerFee: "0.001", MakerFee: "0.001", Left: "0.001",
		DealFee: "0"}, ActivationPrice: "3300"}, Price: "3000"}

	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/order/stop_limit"

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.CreateStopLimitOrder(params)

	response, _ := json.Marshal(responseJson)
	expectedResponse, _ := json.Marshal(serverResponse)
	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}
