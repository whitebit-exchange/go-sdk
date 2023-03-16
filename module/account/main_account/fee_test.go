package main_account

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/tests/mocks"
	"testing"
)

type MainAccountTestSuite struct {
	client  *mocks.Client
	service *Service
	suite.Suite
}

func (s *MainAccountTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
}

func (s *MainAccountTestSuite) TestFeeWithResult() {
	expectedServerResponse := []Fee{{
		Ticker:      "BSV",
		Name:        "Bitcoin SV",
		CanDeposit:  true,
		CanWithdraw: true,
		Deposit: Deposit{
			MinFlex:     "0",
			MaxFlex:     "0",
			PercentFlex: "0",
			Fixed:       "0",
			MinAmount:   "0.03",
			MaxAmount:   "0",
		},
		Withdraw: Withdraw{
			MinFlex:     "0",
			MaxFlex:     "0",
			PercentFlex: "0",
			Fixed:       "0.006",
			MinAmount:   "0.06",
			MaxAmount:   "0",
		}}, {
		Ticker:      "BTC",
		Name:        "Bitcoin",
		CanDeposit:  true,
		CanWithdraw: true,
		Deposit: Deposit{
			MinFlex:     "0",
			MaxFlex:     "0",
			PercentFlex: "0",
			Fixed:       "0",
			MinAmount:   "0.01",
			MaxAmount:   "0",
		},
		Withdraw: Withdraw{
			MinFlex:     "0",
			MaxFlex:     "0",
			PercentFlex: "0",
			Fixed:       "0.0004",
			MinAmount:   "0.001",
			MaxAmount:   "3",
		}},
	}

	endpoint := newFeeEndpoint()
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/main-account/fee"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`[
    {
        "ticker": "BSV",
        "name": "Bitcoin SV",
        "can_deposit": true,
        "can_withdraw": true,
        "deposit": {
            "minFlex": "0",
            "maxFlex": "0",
            "percentFlex": "0",
            "fixed": "0",
            "minAmount": "0.03",
            "maxAmount": "0"
        },
        "withdraw": {
            "minFlex": "0",
            "maxFlex": "0",
            "percentFlex": "0",
            "fixed": "0.006",
            "minAmount": "0.06",
            "maxAmount": "0"
        }
    },
    {
        "ticker": "BTC",
        "name": "Bitcoin",
        "can_deposit": true,
        "can_withdraw": true,
        "deposit": {
            "minFlex": "0",
            "maxFlex": "0",
            "percentFlex": "0",
            "fixed": "0",
            "minAmount": "0.01",
            "maxAmount": "0"
        },
        "withdraw": {
            "minFlex": "0",
            "maxFlex": "0",
            "percentFlex": "0",
            "fixed": "0.0004",
            "minAmount": "0.001",
            "maxAmount": "3"
        }
    }]
`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetFee()

	s.Equal(expectedServerResponse, responseJson)
	s.Equal(err, error(nil))

}

func (s *MainAccountTestSuite) TestFeeInvalidResponseError() {
	endpoint := newFeeEndpoint()
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/main-account/fee"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetFee()
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal("null", string(response))
	s.Equal(expectedError, err.Error())

}

func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(MainAccountTestSuite))
}
