package smartlending

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/tests/mocks"
)

type SmartLendingTestSuite struct {
	client  *mocks.Client
	service *Service
	suite.Suite
}

func (s *SmartLendingTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
}

// Fixed Lending Tests

func (s *SmartLendingTestSuite) TestGetFixedPlans() {
	endpoint := newFixedPlansEndpoint("USDT")

	byteResponse := []byte(`[
		{
			"id": "plan-1",
			"ticker": "USDT",
			"status": 1,
			"percent": "10",
			"duration": 30,
			"interestTicker": "USDT",
			"interestRatio": "1",
			"minInvestment": "100",
			"maxInvestment": "10000",
			"maxPossibleInvestment": "5000"
		}
	]`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/main-account/smart/plans", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.GetFixedPlans("USDT")

	s.NoError(err)
	s.Equal(1, len(result))
	s.Equal("plan-1", result[0].ID)
	s.Equal("USDT", result[0].Ticker)
	s.Equal("10", result[0].Percent)
}

func (s *SmartLendingTestSuite) TestCreateFixedInvestment() {
	endpoint := newFixedInvestmentEndpoint("plan-1", "500")

	byteResponse := []byte(`{
		"id": "inv-1",
		"plan": {"id": "plan-1", "ticker": "USDT", "status": 1, "percent": "10", "duration": 30},
		"status": 1,
		"created": 1700000000,
		"updated": 1700000000,
		"paymentTime": 1702592000,
		"amount": "500",
		"interestPaid": "0"
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/main-account/smart/investment", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.CreateFixedInvestment("plan-1", "500")

	s.NoError(err)
	s.Equal("inv-1", result.ID)
	s.Equal("500", result.Amount)
	s.Equal(1, result.Status)
}

func (s *SmartLendingTestSuite) TestCloseFixedInvestment() {
	endpoint := newFixedInvestmentCloseEndpoint("inv-1")

	byteResponse := []byte(`{}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/main-account/smart/investment/close", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	err := s.service.CloseFixedInvestment("inv-1")

	s.NoError(err)
}

func (s *SmartLendingTestSuite) TestGetFixedInvestments() {
	endpoint := newFixedInvestmentsEndpoint("", "USDT", 1, 10, 0)

	byteResponse := []byte(`{
		"offset": 0,
		"limit": 10,
		"records": [
			{
				"id": "inv-1",
				"plan": {"id": "plan-1", "ticker": "USDT", "status": 1, "percent": "10", "duration": 30},
				"status": 1,
				"created": 1700000000,
				"updated": 1700000000,
				"paymentTime": 1702592000,
				"amount": "500",
				"interestPaid": "50"
			}
		]
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/main-account/smart/investments", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.GetFixedInvestments("", "USDT", 1, 10, 0)

	s.NoError(err)
	s.Equal(1, len(result.Records))
	s.Equal("inv-1", result.Records[0].ID)
}

func (s *SmartLendingTestSuite) TestGetFixedInterestHistory() {
	endpoint := newFixedInterestHistoryEndpoint("plan-1", "USDT", 10, 0)

	byteResponse := []byte(`{
		"offset": 0,
		"limit": 10,
		"records": [
			{
				"planId": "plan-1",
				"investmentId": "inv-1",
				"amount": "5.5",
				"ticker": "USDT",
				"timestamp": 1700000000
			}
		]
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/main-account/smart/interest-payment-history", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.GetFixedInterestHistory("plan-1", "USDT", 10, 0)

	s.NoError(err)
	s.Equal(1, len(result.Records))
	s.Equal("5.5", result.Records[0].Amount)
}

// Flexible Lending Tests

func (s *SmartLendingTestSuite) TestGetFlexPlans() {
	endpoint := newFlexPlansEndpoint("USDT", 10, 0)

	byteResponse := []byte(`[
		{
			"id": "flex-plan-1",
			"ticker": "USDT",
			"minInvestment": "10",
			"maxInvestment": "100000",
			"maxRate": "5"
		}
	]`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/main-account/smart-flex/plans", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.GetFlexPlans("USDT", 10, 0)

	s.NoError(err)
	s.Equal(1, len(result))
	s.Equal("flex-plan-1", result[0].ID)
}

func (s *SmartLendingTestSuite) TestGetFlexInvestments() {
	endpoint := newFlexInvestmentsEndpoint("USDT", "", "", 1, 10, 0)

	byteResponse := []byte(`{
		"data": [
			{
				"id": "flex-inv-1",
				"planId": "flex-plan-1",
				"currency": "USDT",
				"invested": "1000",
				"withAutoReinvest": true,
				"status": 1,
				"createdAt": 1700000000,
				"updatedAt": 1700000000
			}
		],
		"limit": 10,
		"offset": 0
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/main-account/smart-flex/investments", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.GetFlexInvestments("USDT", "", "", 1, 10, 0)

	s.NoError(err)
	s.Equal(1, len(result.Data))
	s.Equal("flex-inv-1", result.Data[0].ID)
	s.True(result.Data[0].WithAutoInvest)
}

func (s *SmartLendingTestSuite) TestFlexInvest() {
	endpoint := newFlexInvestEndpoint("flex-plan-1", "500")

	byteResponse := []byte(`{
		"id": "flex-inv-1",
		"planId": "flex-plan-1",
		"currency": "USDT",
		"invested": "500",
		"withAutoReinvest": false,
		"status": 1,
		"createdAt": 1700000000,
		"updatedAt": 1700000000
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/main-account/smart-flex/investments/invest", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.FlexInvest("flex-plan-1", "500")

	s.NoError(err)
	s.Equal("flex-inv-1", result.ID)
	s.Equal("500", result.Invested)
}

func (s *SmartLendingTestSuite) TestFlexWithdraw() {
	endpoint := newFlexWithdrawEndpoint("flex-inv-1", "100")

	byteResponse := []byte(`{
		"id": "flex-inv-1",
		"planId": "flex-plan-1",
		"currency": "USDT",
		"invested": "400",
		"withAutoReinvest": false,
		"status": 1,
		"createdAt": 1700000000,
		"updatedAt": 1700000001
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/main-account/smart-flex/investments/withdraw", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.FlexWithdraw("flex-inv-1", "100")

	s.NoError(err)
	s.Equal("400", result.Invested)
}

func (s *SmartLendingTestSuite) TestFlexClose() {
	endpoint := newFlexCloseEndpoint("flex-inv-1")

	byteResponse := []byte(`{}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/main-account/smart-flex/investments/close", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	err := s.service.FlexClose("flex-inv-1")

	s.NoError(err)
}

func (s *SmartLendingTestSuite) TestFlexSetAutoInvest() {
	endpoint := newFlexAutoInvestEndpoint("flex-inv-1", true)

	byteResponse := []byte(`{
		"id": "flex-inv-1",
		"planId": "flex-plan-1",
		"currency": "USDT",
		"invested": "500",
		"withAutoReinvest": true,
		"status": 1,
		"createdAt": 1700000000,
		"updatedAt": 1700000001
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/main-account/smart-flex/investments/auto-invest", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.FlexSetAutoInvest("flex-inv-1", true)

	s.NoError(err)
	s.True(result.WithAutoInvest)
}

func (s *SmartLendingTestSuite) TestGetFlexHistory() {
	endpoint := newFlexHistoryEndpoint("flex-plan-1", "", "", 0, 0, []int{1, 2}, 10, 0)

	byteResponse := []byte(`{
		"data": [
			{
				"createdAt": 1700000000,
				"planId": "flex-plan-1",
				"investmentId": "flex-inv-1",
				"currency": "USDT",
				"amount": "500",
				"transactionId": "tx-1",
				"actionType": 1
			}
		],
		"limit": 10,
		"offset": 0
	}`)

	request, _ := whitebit.CreateRequest(endpoint.Url())
	s.Equal("/api/v4/main-account/smart-flex/investments/history", request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	result, err := s.service.GetFlexHistory("flex-plan-1", "", "", 0, 0, []int{1, 2}, 10, 0)

	s.NoError(err)
	s.Equal(1, len(result.Data))
	s.Equal(1, result.Data[0].ActionType)
}

func (s *SmartLendingTestSuite) TestInvalidResponse() {
	byteResponse := []byte(`invalid json`)

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	_, err := s.service.GetFixedPlans("USDT")

	s.Error(err)
	s.Contains(err.Error(), "invalid character")
}

func TestSmartLendingTestSuite(t *testing.T) {
	suite.Run(t, new(SmartLendingTestSuite))
}
