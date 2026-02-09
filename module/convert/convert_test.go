package convert

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/tests/mocks"
)

type ConvertTestSuite struct {
	client  *mocks.Client
	service *Service
	suite.Suite
}

func (s *ConvertTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
}

func (s *ConvertTestSuite) TestEstimate() {
	endpoint := newEstimateEndpoint("BTC", "USDT", "0.1")

	byteResponse := []byte(`{
		"queryId": "test-query-id",
		"from": "BTC",
		"to": "USDT",
		"giveAmount": "0.1",
		"getAmount": "5000",
		"rate": "50000",
		"fee": "0.001",
		"expiresAt": 1700000000
	}`)

	expectedResponse := Estimate{
		QueryID:    "test-query-id",
		From:       "BTC",
		To:         "USDT",
		GiveAmount: "0.1",
		GetAmount:  "5000",
		Rate:       "50000",
		Fee:        "0.001",
		ExpiresAt:  1700000000,
	}

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/convert/estimate", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.Estimate("BTC", "USDT", "0.1")

	s.NoError(err)
	s.Equal(expectedResponse.QueryID, result.QueryID)
	s.Equal(expectedResponse.From, result.From)
	s.Equal(expectedResponse.To, result.To)
	s.Equal(expectedResponse.GiveAmount, result.GiveAmount)
	s.Equal(expectedResponse.GetAmount, result.GetAmount)
}

func (s *ConvertTestSuite) TestConfirm() {
	endpoint := newConfirmEndpoint("test-query-id")

	byteResponse := []byte(`{
		"id": 12345,
		"status": 1,
		"from": "BTC",
		"to": "USDT",
		"giveAmount": "0.1",
		"getAmount": "5000",
		"fee": "0.001"
	}`)

	expectedResponse := Confirmation{
		ID:         12345,
		Status:     1,
		From:       "BTC",
		To:         "USDT",
		GiveAmount: "0.1",
		GetAmount:  "5000",
		Fee:        "0.001",
	}

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/convert/confirm", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.Confirm("test-query-id")

	s.NoError(err)
	s.Equal(expectedResponse.ID, result.ID)
	s.Equal(expectedResponse.Status, result.Status)
	s.Equal(expectedResponse.From, result.From)
	s.Equal(expectedResponse.To, result.To)
}

func (s *ConvertTestSuite) TestGetHistory() {
	endpoint := newHistoryEndpoint(10, 0)

	byteResponse := []byte(`{
		"total": 1,
		"records": [{
			"id": 1,
			"from": "BTC",
			"to": "USDT",
			"giveAmount": "0.1",
			"getAmount": "5000",
			"rate": "50000",
			"fee": "0.001",
			"status": 1,
			"createdAt": 1700000000
		}],
		"limit": 10,
		"offset": 0
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/convert/history", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.GetHistory(10, 0)

	s.NoError(err)
	s.Equal(1, result.Total)
	s.Equal(1, len(result.Records))
	s.Equal("BTC", result.Records[0].From)
	s.Equal("USDT", result.Records[0].To)
}

func (s *ConvertTestSuite) TestEstimateInvalidResponse() {
	byteResponse := []byte(`invalid json`)

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	_, err := s.service.Estimate("BTC", "USDT", "0.1")

	s.Error(err)
	s.Contains(err.Error(), "invalid character")
}

func TestConvertTestSuite(t *testing.T) {
	suite.Run(t, new(ConvertTestSuite))
}
