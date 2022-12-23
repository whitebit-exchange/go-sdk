package collateral

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *OrderTestSuite) TestStopMarketOrderSuccess() {
	endpoint := newStopMarketEndpoint(StopMarketOrderParams{
		Market:          "BTC_USDT",
		Amount:          "100",
		Side:            "buy",
		ActivationPrice: "3000",
		ClientOrderId:   "181",
	})

	byteResponse := []byte(`{"orderId":3310974885,"clientOrderId":"182","market":"BTC_USDT","side":"buy",
		"type":"trigger margin market","timestamp":1669337921.653628,"dealMoney":"0", "dealStock":"0",
		"amount":"100","takerFee":"0.001","makerFee":"0","left":"100","dealFee":"0","activation_price":"3000"}`)

	serverResponse := StopMarketOrder{MarketOrder: MarketOrder{OrderID: 3310974885,
		ClientOrderID: "182", Market: "BTC_USDT", Side: "buy", Type: "trigger margin market",
		Timestamp: 1.669337921653628e+09, DealMoney: "0", DealStock: "0", Amount: "100", TakerFee: "0.001",
		MakerFee: "0", Left: "100", DealFee: "0"}, ActivationPrice: "3000"}

	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/order/collateral/trigger-market"

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.CreateStopMarketOrder(StopMarketOrderParams{
		Market:          "BTC_USDT",
		Amount:          "100",
		Side:            "buy",
		ActivationPrice: "3000",
		ClientOrderId:   "181",
	})

	response, _ := json.Marshal(responseJson)
	expectedResponse, _ := json.Marshal(serverResponse)
	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}
