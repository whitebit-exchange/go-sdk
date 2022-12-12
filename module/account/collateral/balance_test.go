package collateral

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/tests/mocks"
	"testing"
)

type CollateralAccountTestSuite struct {
	client  *mocks.Client
	service *Service
	suite.Suite
}

func (s *CollateralAccountTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
}

func (s *CollateralAccountTestSuite) TestBalanceWithResult() {
	expectedServerResponse := map[string]string{"USDT": "100000"}

	expectedResponse, _ := json.Marshal(expectedServerResponse)
	endpoint := newBalanceEndpoint("BTC")
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/collateral-account/balance"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"USDT":"100000"}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetCollateralBalance("BTC")
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *CollateralAccountTestSuite) TestBalanceInvalidResponseError() {
	endpoint := newBalanceEndpoint("BTC")
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/collateral-account/balance"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetCollateralBalance("BTC")
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal("{}", string(response))
	s.Equal(expectedError, err.Error())

}

func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(CollateralAccountTestSuite))
}
