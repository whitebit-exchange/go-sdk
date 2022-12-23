package collateral

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *OrderTestSuite) TestStopLimitOrderSuccess() {

	endpoint := newStopLimitEndpoint(StopLimitOrderParams{
		Market:          "BTC_USDT",
		Amount:          "0.001",
		Side:            "buy",
		Price:           "3000",
		ActivationPrice: "3300",
		ClientOrderId:   "71",
	})

	byteResponse := []byte(`{"orderId":3310980564,"clientOrderId":"","market":"BTC_USDT","side":"buy",
		"type":"margin stop limit","timestamp":1669338076.215123,"dealMoney":"0","dealStock":"0","amount":"0.001",
		"takerFee":"0.001","makerFee":"0.001","left":"0.001","dealFee":"0","activation_price":"3300","price":"3000"}`)

	serverResponse := StopLimitOrder{StopMarketOrder: StopMarketOrder{
		MarketOrder: MarketOrder{OrderID: 3310980564, ClientOrderID: "", Market: "BTC_USDT", Side: "buy",
			Type: "margin stop limit", Timestamp: 1.669338076215123e+09, DealMoney: "0", DealStock: "0", Amount: "0.001",
			TakerFee: "0.001", MakerFee: "0.001", Left: "0.001", DealFee: "0"}, ActivationPrice: "3300"}, Price: "3000"}

	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/order/collateral/stop-limit"

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.CreateStopLimitOrder(StopLimitOrderParams{
		Market:          "BTC_USDT",
		Amount:          "0.001",
		Side:            "buy",
		Price:           "3000",
		ActivationPrice: "3300",
		ClientOrderId:   "71",
	})

	response, _ := json.Marshal(responseJson)
	expectedResponse, _ := json.Marshal(serverResponse)
	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}
