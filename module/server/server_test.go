package server

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/tests/mocks"
	"testing"
)

type ServerTestSuite struct {
	client       *mocks.Client
	pingEndpoint *pingEndpoint
	timeEndpoint *timeEndpoint
	service      *Service
	suite.Suite
}

func (s *ServerTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
	s.pingEndpoint = newPingEndpoint()
	s.timeEndpoint = newTimeEndpoint()

}

func (s *ServerTestSuite) TestPingWithResult() {
	expectedServerResponse := PingResponse{"pong"}

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.pingEndpoint.Url())
	expectedRequest := "/api/v4/public/ping"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`["pong"]`)
	s.client.On("SendRequest", s.pingEndpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.Ping()
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *ServerTestSuite) TestTimeWithResult() {
	expectedServerResponse := TimeResponse{Time: 1234567}

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.timeEndpoint.Url())
	expectedRequest := "/api/v4/public/time"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"time":1234567}`)
	s.client.On("SendRequest", s.timeEndpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetTime()
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *ServerTestSuite) TestInvalidResponseError() {
	expectedServerResponse := TimeResponse{Time: 0}
	expectedResponse, _ := json.Marshal(expectedServerResponse)

	request, _ := whitebit.CreateRequest(s.timeEndpoint.Url())
	expectedRequest := "/api/v4/public/time"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", s.timeEndpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetTime()
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal(string(expectedResponse), string(response))
	s.Equal(expectedError, err.Error())

}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
