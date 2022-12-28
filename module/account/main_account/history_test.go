package main_account

import (
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *MainAccountTestSuite) TestHistoryWithStatus() {
	expectedServerResponse := HistoryResult{
		Records: []HistoryRecords{{
			Address:         "",
			UniqueId:        "",
			CreatedAt:       1634999285,
			Currency:        "Tether US",
			Ticker:          "USDT",
			Method:          1,
			Amount:          "1000000",
			Description:     nil,
			Memo:            "",
			Fee:             "0",
			Status:          3,
			Network:         "ERC20",
			TransactionHash: "",
			Details:         Details{nil},
			Centralized:     false,
		}, {
			Address:         "",
			UniqueId:        "",
			CreatedAt:       1634999274,
			Currency:        "Bitcoin",
			Ticker:          "BTC",
			Method:          1,
			Amount:          "1000",
			Description:     nil,
			Memo:            "",
			Fee:             "0",
			Status:          3,
			Network:         "",
			TransactionHash: "",
			Details:         Details{},
			Centralized:     false,
		}},
		Offset: 0,
		Limit:  100,
		Total:  2,
	}

	endpoint := newHistoryEndpoint(HistoryParams{
		Status: []int{3},
		Offset: 0,
		Limit:  100,
	})
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/main-account/history"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{
    "records": [
        {
            "address": null,
            "uniqueId": null,
            "createdAt": 1634999285,
            "currency": "Tether US",
            "ticker": "USDT",
            "method": 1,
            "amount": "1000000",
            "description": null,
            "memo": "",
            "fee": "0",
            "status": 3,
            "network": "ERC20",
            "transactionHash": null,
            "details": {
                "partial": null
            }
        },
        {
            "address": null,
            "uniqueId": null,
            "createdAt": 1634999274,
            "currency": "Bitcoin",
            "ticker": "BTC",
            "method": 1,
            "amount": "1000",
            "description": null,
            "memo": "",
            "fee": "0",
            "status": 3,
            "network": null,
            "transactionHash": null,
            "details": {
                "partial": null
            }
        }
    ],
    "offset": 0,
    "limit": 100,
    "total": 2
	}`)

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetHistory(HistoryParams{
		Status: []int{3},
		Offset: 0,
		Limit:  100,
	})

	s.Equal(expectedServerResponse, responseJson)
	s.Equal(err, error(nil))

}

func (s *MainAccountTestSuite) TestHistoryWithTicker() {
	expectedServerResponse := HistoryResult{
		Records: []HistoryRecords{{
			Address:         "",
			UniqueId:        "",
			CreatedAt:       1634999285,
			Currency:        "Tether US",
			Ticker:          "USDT",
			Method:          1,
			Amount:          "1000000",
			Description:     nil,
			Memo:            "",
			Fee:             "0",
			Status:          3,
			Network:         "ERC20",
			TransactionHash: "",
			Details:         Details{nil},
			Centralized:     false,
		}},
		Offset: 0,
		Limit:  100,
		Total:  2,
	}

	endpoint := newHistoryEndpoint(HistoryParams{
		Status: []int{3},
		Offset: 0,
		Limit:  100,
	})
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/main-account/history"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{
    "records": [
        {
            "address": null,
            "uniqueId": null,
            "createdAt": 1634999285,
            "currency": "Tether US",
            "ticker": "USDT",
            "method": 1,
            "amount": "1000000",
            "description": null,
            "memo": "",
            "fee": "0",
            "status": 3,
            "network": "ERC20",
            "transactionHash": null,
            "details": {
                "partial": null
            }
        }
    ],
    "offset": 0,
    "limit": 100,
    "total": 2
	}`)

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetHistory(HistoryParams{
		Ticker: "USDT",
		Status: []int{3},
		Offset: 0,
		Limit:  100,
	})

	s.Equal(expectedServerResponse, responseJson)
	s.Equal(err, error(nil))

}
