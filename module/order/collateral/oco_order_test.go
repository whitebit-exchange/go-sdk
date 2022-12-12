package collateral

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/whitebit"
)

func (s *OrderTestSuite) TestOcoOrderSuccess() {
	endpoint := newOcoEndpoint(OcoOrderParams{
		Market:          "BTC_USDT",
		Amount:          "100",
		Side:            "buy",
		Price:           "2000",
		ActivationPrice: "30300",
		StopLimitPrice:  "40000",
		ClientOrderId:   "152",
	})

	byteResponse := []byte(`{"id":3310999502,"stop_loss":{"orderId":3310999503,"clientOrderId":"",
		"market":"BTC_USDT","side":"buy","type":"margin stop limit","timestamp":1669338402.712101,
		"dealMoney":"0","dealStock":"0","amount":"100","takerFee":"0.001","makerFee":"0.001","left":"100",
		"dealFee":"0","post_only":false,"mtime":1669338402.712101,"price":"40000","activation_price":"30300",
		"activation_condition":"gte","activated":0},"take_profit":{"orderId":3310999504,"clientOrderId":"",
		"market":"BTC_USDT","side":"buy","type":"margin limit","timestamp":1669338402.712101,"dealMoney":"0",
		"dealStock":"0","amount":"100","takerFee":"0.001","makerFee":"0.001","left":"100","dealFee":"0",
		"post_only":false,"mtime":1669338402.712101,"price":"2000"}}`)

	serverResponse := OcoOrder{
		ID: 3310999502,
		StopLoss: StopLoss{
			MarketOrder: MarketOrder{
				OrderID:       3310999503,
				ClientOrderID: "",
				Market:        "BTC_USDT",
				Side:          "buy",
				Type:          "margin stop limit",
				Timestamp:     1.669338402712101e+09,
				DealMoney:     "0",
				DealStock:     "0",
				Amount:        "100",
				TakerFee:      "0.001",
				MakerFee:      "0.001",
				Left:          "100",
				DealFee:       "0",
			},
			PostOnly:            false,
			Mtime:               1.669338402712101e+09,
			Price:               "40000",
			ActivationPrice:     "30300",
			ActivationCondition: "gte",
			Activated:           0,
		},
		TakeProfit: TakeProfit{
			MarketOrder: MarketOrder{
				OrderID:       3310999504,
				ClientOrderID: "",
				Market:        "BTC_USDT",
				Side:          "buy",
				Type:          "margin limit",
				Timestamp:     1.669338402712101e+09,
				DealMoney:     "0",
				DealStock:     "0",
				Amount:        "100",
				TakerFee:      "0.001",
				MakerFee:      "0.001",
				Left:          "100",
				DealFee:       "0",
			}, PostOnly: false,
			Mtime: 1.669338402712101e+09,
			Price: "2000",
		},
	}

	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/order/collateral/oco"

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.CreateOcoOrder(OcoOrderParams{
		Market:          "BTC_USDT",
		Amount:          "100",
		Side:            "buy",
		Price:           "2000",
		ActivationPrice: "30300",
		StopLimitPrice:  "40000",
		ClientOrderId:   "152",
	})

	response, _ := json.Marshal(responseJson)
	expectedResponse, _ := json.Marshal(serverResponse)
	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}
