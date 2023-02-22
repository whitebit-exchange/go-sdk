package main_account

import (
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *MainAccountTestSuite) TestCreateCodeWithResult() {
	expectedServerResponse := Code{
		Code:       "WB136e7b54-d887-45c3-97f0-a916ef1be821USDT",
		Message:    "Code was successfully created",
		ExternalId: "00000"}

	endpoint := newCodeEndpoint("USDT", "10", "", "test")
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/main-account/codes"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"message":"Code was successfully created","code":"WB136e7b54-d887-45c3-97f0-a916ef1be821USDT","external_id":"00000"}`)

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.CreateCode("USDT", "10", "", "test")

	s.Equal(expectedServerResponse, responseJson)
	s.Equal(err, error(nil))

}

func (s *MainAccountTestSuite) TestApplyCodeWithResult() {
	expectedServerResponse := CodeApply{Message: "Code was successfully applied",
		Ticker:     "USDT",
		Amount:     "10",
		ExternalId: "0000"}

	endpoint := newCodeApplyEndpoint("WB136e7b54-d887-45c3-97f0-a916ef1be821USDT", "")
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/main-account/codes/apply"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"message":"Code was successfully applied","ticker":"USDT","amount":"10","external_id":"0000"}`)

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.ApplyCode("WB136e7b54-d887-45c3-97f0-a916ef1be821USDT", "")

	s.Equal(expectedServerResponse, responseJson)
	s.Equal(err, error(nil))

}

func (s *MainAccountTestSuite) TestMyCodesWithResult() {
	expectedServerResponse := CodeMy{
		Total: 2, Data: Data{
			{Amount: "10", Code: "WB4dd98602-3489-47cc-b34c-f3db6a0a4467USDT", Date: 1674683426, Status: "New",
				Ticker: "USDT", ExternalId: "da73291c-ae70-4ad1-9d31-f84ce3341433"},
			{Amount: "714.17", Code: "WB99db7e8c-e2c8-4389-9dbe-ef3eb2f98ce8USDT", Date: 1673375353, Status: "Activated",
				Ticker: "USDT", ExternalId: "902ba1ae-4c3b-46c8-9b96-108d0e046b86"}},
		Limit:  100,
		Offset: 0,
	}

	endpoint := newCodeMyEndpoint(0, 0)
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/main-account/codes/my"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"total":2,"data":[{"amount":"10","code":"WB4dd98602-3489-47cc-b34c-f3db6a0a4467USDT",
		"date":1674683426,"status":"New","ticker":"USDT","external_id":"da73291c-ae70-4ad1-9d31-f84ce3341433"},
		{"amount":"714.17","code":"WB99db7e8c-e2c8-4389-9dbe-ef3eb2f98ce8USDT","date":1673375353,"status":"Activated",
		"ticker":"USDT","external_id":"902ba1ae-4c3b-46c8-9b96-108d0e046b86"}],"limit":100,"offset":0}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetCodes(0, 0)

	s.Equal(expectedServerResponse, responseJson)
	s.Equal(err, error(nil))

}

func (s *MainAccountTestSuite) TestHistoryCodesWithResult() {
	expectedServerResponse := CodeHistory{
		Total: 2,
		Data: Data{{Amount: "10", Code: "WB4dd98602-3489-47cc-b34c-f3db6a0a4467USDT", Date: 1674683426, Status: "New",
			Ticker: "USDT", ExternalId: "00"},
			{Amount: "714.17", Code: "WB99db7e8c-e2c8-4389-9dbe-ef3eb2f98ce8USDT", Date: 1673375353,
				Status: "Activated", Ticker: "USDT", ExternalId: "00"}},
		Limit:  100,
		Offset: 0,
	}

	endpoint := newCodeHistoryEndpoint(0, 0)
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/main-account/codes/history"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"total":2,"data":[{"amount":"10","code":"WB4dd98602-3489-47cc-b34c-f3db6a0a4467USDT",
		"date":1674683426,"status":"New","ticker":"USDT","external_id":"00"},
		{"amount":"714.17","code":"WB99db7e8c-e2c8-4389-9dbe-ef3eb2f98ce8USDT","date":1673375353,"status":"Activated",
		"ticker":"USDT","external_id":"00"}],"limit":100,"offset":0}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetCodesHistory(0, 0)

	s.Equal(expectedServerResponse, responseJson)
	s.Equal(err, error(nil))

}
