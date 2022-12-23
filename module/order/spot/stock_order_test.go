package spot

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *OrderTestSuite) TestStockOrderSuccess() {
	byteResponse := []byte(`{"orderId":3310696165,"clientOrderId":"61","market":"BTC_USDT","side":"buy",
		"type":"stock market","timestamp":1669333716.366427,"dealMoney":"16.58143","dealStock":"0.001",
		"amount":"0.001","takerFee":"0.001","makerFee":"0","left":"0","dealFee":"0.01658143"}`)

	serverResponse := StockMarketOrder{MarketOrder: MarketOrder{OrderID: 3310696165,
		ClientOrderID: "61", Market: "BTC_USDT", Side: "buy", Type: "stock market",
		Timestamp: 1.669333716366427e+09, DealMoney: "16.58143", DealStock: "0.001",
		Amount: "0.001", TakerFee: "0.001", MakerFee: "0", Left: "0", DealFee: "0.01658143"}}

	params := MarketOrderParams{
		Market:        "BTC_USDT",
		Amount:        "0.001",
		Side:          "buy",
		ClientOrderId: "61",
	}

	endpoint := newStockMarketEndpoint(params)
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/order/stock_market"

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.CreateMarketStock(params)

	response, _ := json.Marshal(responseJson)
	expectedResponse, _ := json.Marshal(serverResponse)
	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}
