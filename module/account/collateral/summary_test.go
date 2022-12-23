package collateral

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *CollateralAccountTestSuite) TestSummaryWithResult() {
	expectedServerResponse := Summary{Equity: "1760331.93",
		Margin:            "371610.399413175322",
		FreeMargin:        "1375118.640520948068",
		UnrealizedFunding: "0",
		Pnl:               "-13602.89006587661",
		Leverage:          5}

	byteResponse := []byte(`{"equity":"1760331.93","margin":"371610.399413175322",
		"freeMargin":"1375118.640520948068","unrealizedFunding":"0","pnl":"-13602.89006587661","leverage":5}`)

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	expectedRequest := "/api/v4/collateral-account/summary"
	endpoint := newSummaryEndpoint()
	request, _ := whitebit.CreateRequest(endpoint.Url())

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetSummary()
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *CollateralAccountTestSuite) TestSummaryResponseError() {
	byteResponse := []byte(`{"success": True, "message": null}`)
	expectedRequest := "/api/v4/collateral-account/summary"
	expectedServerResponse := Summary{}
	expectedResponse, _ := json.Marshal(expectedServerResponse)
	endpoint := newSummaryEndpoint()
	request, _ := whitebit.CreateRequest(endpoint.Url())

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetSummary()
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal(string(expectedResponse), string(response))
	s.Equal(expectedError, err.Error())

}
