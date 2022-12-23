package symbol

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/tests/mocks"
	"testing"
)

type SymbolsTestSuite struct {
	client   *mocks.Client
	endpoint *endpoint
	service  *Service
	suite.Suite
}

func (s *SymbolsTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
	s.endpoint = newEndpoint()

}

func (s *SymbolsTestSuite) TestWithResult() {
	expectedServerResponse := SymbolsResult{
		Success: true,
		Message: nil,
		Result:  []string{"1INCH_BTC", "1INCH_UAH", "1INCH_USDT", "AAVE_BTC", "AAVE_USDT", "ADA_BTC"},
	}

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v1/public/symbols"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success":true,"message":null,"result":["1INCH_BTC","1INCH_UAH","1INCH_USDT","AAVE_BTC","AAVE_USDT","ADA_BTC"]}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetSymbols()
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *SymbolsTestSuite) TestInvalidResponseError() {
	expectedServerResponse := SymbolsResult{
		Success: false,
		Message: nil,
		Result:  nil,
	}
	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v1/public/symbols"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetSymbols()
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal(string(expectedResponse), string(response))
	s.Equal(expectedError, err.Error())

}

func TestSymbolsTestSuite(t *testing.T) {
	suite.Run(t, new(SymbolsTestSuite))
}
