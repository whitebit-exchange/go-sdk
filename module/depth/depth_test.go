package depth

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/tests/mocks"
	"testing"
)

type DepthTestSuite struct {
	client   *mocks.Client
	endpoint *depthEndpoint
	service  *Service
	suite.Suite
}

func (s *DepthTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
	s.endpoint = newDepthEndpoint("ETH_USDT")

}

func (s *DepthTestSuite) TestWithResult() {
	expectedResult := Depth{
		Time: "2022-11-15T14:21:31Z",
		Asks: []Pair{{"10", "100"}, {"11", "110"}, {"12", "120"}},
		Bids: []Pair{{"20", "200"}, {"21", "110"}, {"22", "120"}},
	}
	expectedServerResponse := Result{
		Success: true,
		Message: nil,
		Result:  expectedResult,
	}
	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v2/public/depth/ETH_USDT"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": true, "message": null,
							"result":{"lastUpdateTimestamp": "2022-11-15T14:21:31Z",
								"asks":[["10", "100"], ["11", "110"], ["12", "120"]],
								"bids": [["20", "200"], ["21", "110"], ["22", "120"]]}}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetDepth("ETH_USDT")
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *DepthTestSuite) TestWithMessage() {
	s.endpoint = newDepthEndpoint("XXX_XXX")
	expectedServerResponse := Result{
		Success: false,
		Message: map[string][]string{
			"market": {"Market is not available."},
		},
		Result: Depth{},
	}
	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v2/public/depth/XXX_XXX"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": false,
							"message": {
								"market": ["Market is not available."]
								},
							"result": null
							}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetDepth("XXX_XXX")
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *DepthTestSuite) TestInvalidResponseError() {
	expectedServerResponse := Result{
		Success: false,
		Message: nil,
		Result:  Depth{},
	}
	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v2/public/depth/ETH_USDT"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetDepth("ETH_USDT")
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal(string(expectedResponse), string(response))
	s.Equal(expectedError, err.Error())

}

func TestDepthTestSuite(t *testing.T) {
	suite.Run(t, new(DepthTestSuite))
}
