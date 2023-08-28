package spot

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *CancelTestSuite) TestKillSwitchSuccess() {

	byteResponse := []byte(`{"market":"BTC_USDT","startTime":1000,"cancellationTime":1100,"types":["spot"]}`)

	serverResponse := KillSwitchResponse{Market: "BTC_USDT", StartTime: 1000, CancellationTime: 1100, Types: []string{OrderTypeSpot}}

	expectedRequest := killSwitchEndpointUrl

	endpoint := newKillSwitchEndpoint(KillSwitchParams{Market: "BTC_USDT", Timeout: "100", Types: []string{"spot"}})
	request, _ := whitebit.CreateRequest(endpoint.Url())

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.CreateKillSwitch(KillSwitchParams{Market: "BTC_USDT", Timeout: "100", Types: []string{OrderTypeSpot}})
	response, _ := json.Marshal(responseJson)
	expectedResponse, _ := json.Marshal(serverResponse)
	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *CancelTestSuite) TestKillSwitchStatusSuccess() {

	byteResponse := []byte(`[{"market":"BTC_USDT","startTime":1000,"cancellationTime":1100,"types":["spot"]}, {"market":"ETH_USDT","startTime":1000,"cancellationTime":1100,"types":["margin", "futures"]}]`)

	serverResponse := []KillSwitchStatusResponse{
		{Market: "BTC_USDT", StartTime: 1000, CancellationTime: 1100, Types: []string{OrderTypeSpot}},
		{Market: "ETH_USDT", StartTime: 1000, CancellationTime: 1100, Types: []string{OrderTypeMargin, OrderTypeFutures}}}

	expectedRequest := killSwitchStatusEndpointUrl

	endpoint := newKillSwitchStatusEndpoint(KillSwitchStatusParams{})
	request, _ := whitebit.CreateRequest(endpoint.Url())

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetKillSwitchStatus(KillSwitchStatusParams{})
	response, _ := json.Marshal(responseJson)
	expectedResponse, _ := json.Marshal(serverResponse)
	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}
