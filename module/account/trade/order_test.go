package trade

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *AccountTestSuite) TestOrderWithResult() {
	expectedServerResponse := map[string]interface{}{"limit": 100, "offset": 0,
		"records": []interface{}{map[string]interface{}{
			"amount": "0.029791", "clientOrderId": "", "deal": "599.39521791", "dealOrderId": 3.26360797e+09,
			"fee": "0.59939521791", "id": 8.920261e+07, "price": "20120.01", "role": 2, "time": 1.668176415399715e+09}}}

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	expectedRequest := "/api/v4/trade-account/order"
	endpoint := newOrderEndpoint(3263845935, 100, 0)
	request, _ := whitebit.CreateRequest(endpoint.Url())

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"limit":100,"offset":0,"records":[{"amount":"0.029791","clientOrderId":"",
		"deal":"599.39521791","dealOrderId":3263607970,"fee":"0.59939521791","id":89202610,
		"price":"20120.01","role":2,"time":1668176415.399715}]}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetOrder(3263845935, 100, 0)
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *AccountTestSuite) TestOrderInvalidResponseError() {
	endpoint := newOrderEndpoint(3263845935, 100, 0)
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/trade-account/order"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetOrder(3263845935, 100, 0)
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal("null", string(response))
	s.Equal(expectedError, err.Error())

}
