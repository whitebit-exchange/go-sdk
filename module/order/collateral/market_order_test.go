package collateral

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *OrderTestSuite) TestMarketOrderSuccess() {
	endpoint := newMarketEndpoint(MarketOrderParams{
		Market:        "BTC_USDT",
		Amount:        "100",
		Side:          "buy",
		ClientOrderId: "31",
	})

	byteResponse := []byte(`{"orderId":3311012525,"clientOrderId":"231","market":"BTC_USDT","side":"buy",
		"type":"margin market","timestamp":1669338631.29033,"dealMoney":"1654197.79926661","dealStock":"100",
		"amount":"100","takerFee":"0.001","makerFee":"0","left":"0","dealFee":"1654.19779926661"}`)

	serverResponse := MarketOrder{OrderID: 3311012525, ClientOrderID: "231", Market: "BTC_USDT", Side: "buy",
		Type: "margin market", Timestamp: 1.66933863129033e+09, DealMoney: "1654197.79926661", DealStock: "100",
		Amount: "100", TakerFee: "0.001", MakerFee: "0", Left: "0", DealFee: "1654.19779926661"}

	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/order/collateral/market"

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.CreateMarketOrder(MarketOrderParams{
		Market:        "BTC_USDT",
		Amount:        "100",
		Side:          "buy",
		ClientOrderId: "31",
	})

	response, _ := json.Marshal(responseJson)
	expectedResponse, _ := json.Marshal(serverResponse)
	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}
