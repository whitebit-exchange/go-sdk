package fee

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/tests/mocks"
	"testing"
)

type FeeListTestSuite struct {
	client   *mocks.Client
	endpoint *listEndpoint
	service  *Service
	suite.Suite
}

func (s *FeeListTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
	s.endpoint = newListEndpoint()

}

func (s *FeeListTestSuite) TestWithResult() {
	expectedServerResponse := List{
		"ADA": map[string]interface{}{"deposit": map[string]interface{}{"fixed": interface{}(nil),
			"flex": interface{}(nil), "max_amount": "0", "min_amount": "8"},
			"is_depositable": true,
			"is_withdrawal":  true,
			"name":           "Cardano",
			"providers":      []interface{}{},
			"ticker":         "ADA",
			"withdraw": map[string]interface{}{"fixed": "2",
				"flex": interface{}(nil), "max_amount": "0", "min_amount": "16"}},
		"ADK": map[string]interface{}{"deposit": map[string]interface{}{"fixed": interface{}(nil),
			"flex": interface{}(nil), "max_amount": "0", "min_amount": "30"},
			"is_depositable": true,
			"is_withdrawal":  true,
			"name":           "Aidos Kuneen",
			"providers":      []interface{}{},
			"ticker":         "ADK",
			"withdraw": map[string]interface{}{"fixed": "5",
				"flex": interface{}(nil), "max_amount": "0", "min_amount": "50"}},
	}

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v4/public/fee"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"ADA":{"is_depositable":true,"is_withdrawal":true,"ticker":"ADA","name":"Cardano",
		"providers":[],"withdraw":{"max_amount":"0","min_amount":"16","fixed":"2","flex":null},
		"deposit":{"max_amount":"0","min_amount":"8","fixed":null,"flex":null}},
		"ADK":{"is_depositable":true,"is_withdrawal":true,"ticker":"ADK","name":"Aidos Kuneen","providers":[],
		"withdraw":{"max_amount":"0","min_amount":"50","fixed":"5","flex":null},"deposit":
		{"max_amount":"0","min_amount":"30","fixed":null,"flex":null}}}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetTradingFeesList()
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *FeeListTestSuite) TestInvalidResponseError() {
	expectedServerResponse := List{}
	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v4/public/fee"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetTradingFeesList()
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal(string(expectedResponse), string(response))
	s.Equal(expectedError, err.Error())

}

func TestFeeListTestSuite(t *testing.T) {
	suite.Run(t, new(FeeListTestSuite))
}
