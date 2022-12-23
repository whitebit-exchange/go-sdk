package collateral

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *CollateralAccountTestSuite) TestExecutedWithResult() {
	expectedServerResponse := map[string]int{"leverage": 5}
	expectedResponse, _ := json.Marshal(expectedServerResponse)

	expectedRequest := "/api/v4/collateral-account/leverage"
	endpoint := newLeverageEndpoint("5")
	request, _ := whitebit.CreateRequest(endpoint.Url())

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"leverage":5}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.SetLeverage("5")
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *CollateralAccountTestSuite) TestExecutedInvalidResponseError() {
	endpoint := newLeverageEndpoint("5")
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/collateral-account/leverage"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.SetLeverage("5")
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal("null", string(response))
	s.Equal(expectedError, err.Error())

}
