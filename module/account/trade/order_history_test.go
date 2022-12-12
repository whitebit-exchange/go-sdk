package trade

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/whitebit"
)

func (s *AccountTestSuite) TestHistoryWithResult() {
	expectedServerResponse := map[string][]OrderHistory{"BTC_USDT": {
		{Id: 3313805042, ClientOrderId: "", StartTime: 1.669374684477233e+09, EndTime: 1.669374926437726e+09, Side: "buy",
			Amount: "100", Price: "2000", Type: "margin limit", TakerFee: "0.001", MakerFee: "0.001", DealFee: "0",
			DealStock: "0", DealMoney: "0"},
		{Id: 3311012525, ClientOrderId: "231", StartTime: 1.66933863129033e+09, EndTime: 1.66933863129033e+09, Side: "buy",
			Amount: "100", Price: "0", Type: "margin market", TakerFee: "0.001", MakerFee: "0", DealFee: "1654.19779926661",
			DealStock: "100", DealMoney: "1654197.79926661"},
		{Id: 3310999504, ClientOrderId: "", StartTime: 1.669338402712101e+09, EndTime: 1.669366263596253e+09, Side: "buy",
			Amount: "100", Price: "2000", Type: "margin limit", TakerFee: "0.001", MakerFee: "0.001", DealFee: "0",
			DealStock: "0", DealMoney: "0"},
	}}

	byteResponse := []byte(`{"BTC_USDT":[{"id":3313805042,"clientOrderId":"","ctime":1669374684.477233,
	"ftime":1669374926.437726,"side":"buy","amount":"100","price":"2000","type":"margin limit","takerFee":"0.001",
	"makerFee":"0.001","dealFee":"0","dealStock":"0","dealMoney":"0"},{"id":3311012525,"clientOrderId":"231",
	"ctime":1669338631.29033,"ftime":1669338631.29033,"side":"buy","amount":"100","price":"0","type":"margin market",
	"takerFee":"0.001","makerFee":"0","dealFee":"1654.19779926661","dealStock":"100","dealMoney":"1654197.79926661"},
	{"id":3310999504,"clientOrderId":"","ctime":1669338402.712101,"ftime":1669366263.596253,"side":"buy",
	"amount":"100","price":"2000","type":"margin limit","takerFee":"0.001","makerFee":"0.001","dealFee":"0",
	"dealStock":"0","dealMoney":"0"}]}`)

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	expectedRequest := "/api/v4/trade-account/order/history"
	endpoint := newHistoryEndpoint("BTC_USDT", 100, 0, 0, "")
	request, _ := whitebit.CreateRequest(endpoint.Url())

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetHistory(HistoryOptions{Market: "BTC_USDT", Limit: 100})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *AccountTestSuite) TestHistoryInvalidResponseError() {
	endpoint := newHistoryEndpoint("BTC_USDT", 0, 0, 0, "")
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/trade-account/order/history"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetHistory(HistoryOptions{Market: "BTC_USDT"})
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal("null", string(response))
	s.Equal(expectedError, err.Error())

}
