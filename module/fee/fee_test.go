package fee

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/tests/mocks"
	"testing"
)

type FeeTestSuite struct {
	client   *mocks.Client
	endpoint *tradingFeeEndpoint
	service  *Service
	suite.Suite
}

func (s *FeeTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
	s.endpoint = newTradingFeeEndpoint()

}

func (s *FeeTestSuite) TestWithResult() {
	expectedServerResponse := Result{
		Success: true,
		Message: nil,
		Result:  Fee{MakerFee: "0.1", TakerFee: "0.1"},
	}

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v2/public/fee"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success":true,"message":null,"result":{"makerFee":"0.1","takerFee":"0.1"}}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetTradingFee()
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *FeeTestSuite) TestInvalidResponseError() {
	expectedServerResponse := Result{
		Success: false,
		Message: nil,
		Result:  Fee{},
	}
	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v2/public/fee"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetTradingFee()
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal(string(expectedResponse), string(response))
	s.Equal(expectedError, err.Error())

}

func TestFeeTestSuite(t *testing.T) {
	suite.Run(t, new(FeeTestSuite))
}
