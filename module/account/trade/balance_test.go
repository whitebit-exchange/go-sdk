package trade

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/tests/mocks"
	"testing"
)

type AccountTestSuite struct {
	client  *mocks.Client
	service *Service
	suite.Suite
}

func (s *AccountTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
}

func (s *AccountTestSuite) TestBalanceWithResult() {
	expectedServerResponse := map[string]interface{}{"available": "128999.065109", "freeze": "0"}

	expectedResponse, _ := json.Marshal(expectedServerResponse)
	endpoint := newBalanceEndpoint("BTC")
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/trade-account/balance"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"available":"128999.065109","freeze":"0"}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetBalance("BTC")
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *AccountTestSuite) TestBalanceInvalidResponseError() {
	endpoint := newBalanceEndpoint("BTC")
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/trade-account/balance"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetBalance("BTC")
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal("null", string(response))
	s.Equal(expectedError, err.Error())

}

func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}
