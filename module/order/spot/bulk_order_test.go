package spot

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *OrderTestSuite) TestBulkOrderSuccess() {
	endpoint := newBulkEndpoint(
		[]LimitOrderParams{{"sell", "0.01", "1000", "BTC_USDT", false, false, ""},
			{"buy", "0.01", "1000", "BTC_USDT", true, true, "someId"},
		})

	byteResponse := []byte(`[{"result":null,"error":{"code":10,"message":"Inner validation failed",
		"errors":{"amount":["Not enough balance."]}}},{"result":{"orderId":4330889006,"clientOrderId":"25",
		"market":"BTC_USDT","side":"sell","type":"limit","timestamp":1685025867.619288,"dealMoney":"0",
		"dealStock":"0","amount":"0.002","takerFee":"0.002","makerFee":"0.02","left":"0.002","dealFee":"0",
		"ioc":false,"postOnly":false,"price":"41000"},"error":null}]`)

	serverResponse := []BulkOrderResponseRecord{
		{
			Error: &whitebit.Error{Code: 10,
				Message: "Inner validation failed",
				Errors:  map[string][]string{"amount": {"Not enough balance."}}},
		},
		{
			Result: &LimitOrder{MarketOrder: MarketOrder{OrderID: 4330889006, ClientOrderID: "25", Market: "BTC_USDT",
				Side: "sell", Type: "limit", Timestamp: 1685025867.619288, DealMoney: "0", DealStock: "0",
				Amount: "0.002", TakerFee: "0.002", MakerFee: "0.02", Left: "0.002", DealFee: "0"}, Price: "41000"},
		},
	}

	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/order/bulk"

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.CreateBulkOrder([]LimitOrderParams{
		{"sell", "0.01", "1000", "BTC_USDT", false, false, ""},
		{"buy", "0.01", "1000", "BTC_USDT", true, true, "someId"},
	})

	response, _ := json.Marshal(responseJson)
	expectedResponse, _ := json.Marshal(serverResponse)
	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}
