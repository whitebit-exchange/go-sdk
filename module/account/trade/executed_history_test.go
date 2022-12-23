package trade

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *AccountTestSuite) TestExecutedWithResult() {
	expectedServerResponse := []ExecutedHistory{{
		Id: 90452546, ClientOrderId: "231", Time: 1.66933863129033e+09, Side: "buy", Role: 2, Amount: "3.097524",
		Price: "16546.62", Deal: "51253.55256888", Fee: "51.25355256888", OrderId: 3311012525},
		{Id: 90452545, ClientOrderId: "231", Time: 1.66933863129033e+09, Side: "buy", Role: 2, Amount: "6.472764",
			Price: "16546.62", Deal: "107102.36625768", Fee: "107.10236625768", OrderId: 3311012525}}

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	expectedRequest := "/api/v4/trade-account/executed-history"
	endpoint := newExecutedHistoryEndpoint("BTC_USDT", 2, 0, "")
	request, _ := whitebit.CreateRequest(endpoint.Url())

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`[{"id":90452546,"clientOrderId":"231","time":1669338631.29033,"side":"buy","role":2,"amount":"3.097524","price":"16546.62","deal":"51253.55256888","fee":"51.25355256888","orderId":3311012525},{"id":90452545,"clientOrderId":"231","time":1669338631.29033,"side":"buy","role":2,"amount":"6.472764","price":"16546.62","deal":"107102.36625768","fee":"107.10236625768","orderId":3311012525}]`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetExecutedHistory(ExecutedHistoryOptions{Market: "BTC_USDT", Limit: 2})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *AccountTestSuite) TestExecutedInvalidResponseError() {
	endpoint := newExecutedHistoryEndpoint("BTC_USDT", 2, 0, "")
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/trade-account/executed-history"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetExecutedHistory(ExecutedHistoryOptions{Market: "BTC_USDT", Limit: 2})
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal("null", string(response))
	s.Equal(expectedError, err.Error())

}
