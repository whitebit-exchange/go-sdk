package spot

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *OrderTestSuite) TestMarketOrderSuccess() {
	params := MarketOrderParams{
		Market:        "BTC_USDT",
		Amount:        "100",
		Side:          "buy",
		ClientOrderId: "31",
	}
	endpoint := newMarketOrderEndpoint(params)

	byteResponse := []byte(`{"orderId":3310597927,"clientOrderId":"31","market":"BTC_USDT","side":"buy",
		"type":"market","timestamp":1669331945.235558,"dealMoney":"99.89994743","dealStock":"0.006023",
		"amount":"99.9","takerFee":"0.001","makerFee":"0","left":"0.00005257","dealFee":"0.09989994743"}`)

	serverResponse := MarketOrder{OrderID: 3310597927, ClientOrderID: "31", Market: "BTC_USDT", Side: "buy",
		Type: "market", Timestamp: 1.669331945235558e+09, DealMoney: "99.89994743", DealStock: "0.006023", Amount: "99.9",
		TakerFee: "0.001", MakerFee: "0", Left: "0.00005257", DealFee: "0.09989994743"}

	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/order/market"

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.CreateMarketOrder(params)

	response, _ := json.Marshal(responseJson)
	expectedResponse, _ := json.Marshal(serverResponse)
	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}
