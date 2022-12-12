package trade

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/module/order/spot"
)

func (s *AccountTestSuite) TestOrdersWithResult() {
	expectedServerResponse := []Orders{{
		MarketOrder: spot.MarketOrder{OrderID: 3313899825, ClientOrderID: "", Market: "BTC_USDT", Side: "buy",
			Type: "margin limit", Timestamp: 1.669375581341061e+09, DealMoney: "0", DealStock: "0", Amount: "100",
			TakerFee: "0.001", MakerFee: "0.001", Left: "100", DealFee: "0"}, Price: "2000", ActivationPrice: "",
	}, {
		MarketOrder: spot.MarketOrder{OrderID: 3313899824, ClientOrderID: "", Market: "BTC_USDT", Side: "buy",
			Type: "margin stop limit", Timestamp: 1.669375581341061e+09, DealMoney: "0", DealStock: "0", Amount: "100",
			TakerFee: "0.001", MakerFee: "0.001", Left: "100", DealFee: "0"}, Price: "40000", ActivationPrice: "30300"},
	}

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	expectedRequest := "/api/v4/orders"
	endpoint := newOrdersEndpoint("BTC_USDT", 100, 0)
	request, _ := whitebit.CreateRequest(endpoint.Url())

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`[{"orderId":3313899825,"clientOrderId":"","market":"BTC_USDT","side":"buy",
	"type":"margin limit","timestamp":1669375581.341061,"dealMoney":"0","dealStock":"0","amount":"100",
	"takerFee":"0.001","makerFee":"0.001","left":"100","dealFee":"0","price":"2000"},
	{"orderId":3313899824,"clientOrderId":"","market":"BTC_USDT","side":"buy","type":"margin stop limit",
	"timestamp":1669375581.341061,"dealMoney":"0","dealStock":"0","amount":"100","takerFee":"0.001",
	"makerFee":"0.001","left":"100","dealFee":"0","price":"40000","activation_price":"30300"}]`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetOrders("BTC_USDT", 100, 0)
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *AccountTestSuite) TestOrdersInvalidResponseError() {
	endpoint := newOrdersEndpoint("BTC_USDT", 100, 0)
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/orders"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetOrders("BTC_USDT", 100, 0)
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal("null", string(response))
	s.Equal(expectedError, err.Error())

}
