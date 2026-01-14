package main_account

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/go-sdk/tests/mocks"
)

type MainAccountServiceSuite struct {
	suite.Suite
	client  *mocks.Client
	service *Service
}

func (s *MainAccountServiceSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
}

// GetMainBalance
func (s *MainAccountServiceSuite) TestGetMainBalance_Success() {
	// arrange
	payload := []byte(`{"BTC":{"main_balance":"1.23"},"USDT":{"main_balance":"1000"}}`)
	s.client.On("SendRequest", mock.Anything).Return(payload, nil).Once()

	// act
	got, err := s.service.GetMainBalance()

	// assert
	s.Require().NoError(err)
	b, _ := json.Marshal(got)
	s.Equal(`{"BTC":{"main_balance":"1.23"},"USDT":{"main_balance":"1000"}}`, string(b))
}

func (s *MainAccountServiceSuite) TestGetMainBalance_APIError() {
	s.client.On("SendRequest", mock.Anything).Return(nil, errors.New("api error")).Once()
	got, err := s.service.GetMainBalance()
	s.Error(err)
	s.Equal(State{}, got)
}

func (s *MainAccountServiceSuite) TestGetMainBalance_InvalidJSON() {
	s.client.On("SendRequest", mock.Anything).Return([]byte(`{"BTC":{"main_balance":1.23}`), nil).Once()
	got, err := s.service.GetMainBalance()
	s.Error(err)
	s.Equal(State{}, got)
}

// GetFee
func (s *MainAccountServiceSuite) TestGetFee_Success() {
	payload := []byte(`[{"ticker":"BTC","name":"Bitcoin","can_deposit":true,"can_withdraw":true,
        "deposit":{"minFlex":"0","maxFlex":"0","percentFlex":"0","fixed":"0","minAmount":"0","maxAmount":"0"},
        "withdraw":{"minFlex":"0","maxFlex":"0","percentFlex":"0","fixed":"0","minAmount":"0","maxAmount":"0"}}]`)
	s.client.On("SendRequest", mock.Anything).Return(payload, nil).Once()
	fees, err := s.service.GetFee()
	s.Require().NoError(err)
	s.Require().Len(fees, 1)
	s.Equal("BTC", fees[0].Ticker)
	s.Equal("Bitcoin", fees[0].Name)
	s.True(fees[0].CanDeposit)
	s.True(fees[0].CanWithdraw)
}

func (s *MainAccountServiceSuite) TestGetFee_APIError() {
	s.client.On("SendRequest", mock.Anything).Return(nil, errors.New("api error")).Once()
	fees, err := s.service.GetFee()
	s.Error(err)
	s.Nil(fees)
}

func (s *MainAccountServiceSuite) TestGetFee_InvalidJSON() {
	s.client.On("SendRequest", mock.Anything).Return([]byte(`{"ticker_id":"BTC"}`), nil).Once()
	fees, err := s.service.GetFee()
	s.Error(err)
	s.Nil(fees)
}

// GetHistory (smoke scenarios)
func (s *MainAccountServiceSuite) TestGetHistory_Success() {
	payload := []byte(`{"records":[],"offset":0,"limit":10,"total":0}`)
	s.client.On("SendRequest", mock.Anything).Return(payload, nil).Once()
	res, err := s.service.GetHistory(HistoryParams{})
	s.NoError(err)
	s.Equal(0, res.Offset)
	s.Equal(10, res.Limit)
	s.Equal(0, res.Total)
	s.Len(res.Records, 0)
}

func (s *MainAccountServiceSuite) TestGetHistory_APIError() {
	s.client.On("SendRequest", mock.Anything).Return(nil, errors.New("api error")).Once()
	res, err := s.service.GetHistory(HistoryParams{})
	s.Error(err)
	// zero value struct
	s.Equal(0, res.Offset)
	s.Equal(0, res.Limit)
	s.Equal(0, res.Total)
	s.Nil(res.Records)
}

func (s *MainAccountServiceSuite) TestGetHistory_InvalidJSON() {
	s.client.On("SendRequest", mock.Anything).Return([]byte(`{"records":[`), nil).Once()
	res, err := s.service.GetHistory(HistoryParams{})
	s.Error(err)
	s.Equal(0, res.Offset)
	s.Equal(0, res.Limit)
	s.Equal(0, res.Total)
	s.Nil(res.Records)
}

func TestMainAccountServiceSuite(t *testing.T) { suite.Run(t, new(MainAccountServiceSuite)) }
