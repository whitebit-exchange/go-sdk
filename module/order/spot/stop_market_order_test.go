package spot

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *OrderTestSuite) TestStopMarketOrderSuccess() {
	params := StopMarketOrderParams{
		Market:          "BTC_USDT",
		Amount:          "100",
		Side:            "buy",
		ActivationPrice: "3000",
		ClientOrderId:   "81",
	}

	endpoint := newStopMarketEndpoint(params)

	byteResponse := []byte(`{"orderId":3310741147,"clientOrderId":"81","market":"BTC_USDT","side":"buy",
		"type":"stop market","timestamp":1669334575.3439,"dealMoney":"0","dealStock":"0","amount":"99.9",
		"takerFee":"0.001","makerFee":"0","left":"99.9","dealFee":"0","activation_price":"3000"}`)

	serverResponse := StopMarketOrder{MarketOrder: MarketOrder{OrderID: 3310741147,
		ClientOrderID: "81", Market: "BTC_USDT", Side: "buy", Type: "stop market", Timestamp: 1.6693345753439e+09,
		DealMoney: "0", DealStock: "0", Amount: "99.9", TakerFee: "0.001", MakerFee: "0", Left: "99.9", DealFee: "0"},
		ActivationPrice: "3000"}

	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/order/stop_market"

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.CreateStopMarketOrder(params)

	response, _ := json.Marshal(responseJson)
	expectedResponse, _ := json.Marshal(serverResponse)
	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}
