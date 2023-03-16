package main_account

import (
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *MainAccountTestSuite) TestCustomFee() {
	expectedServerResponse := CustomFee{
		Error:     nil,
		Taker:     "0.001",
		Maker:     "0.002",
		CustomFee: map[string][]string{"BTC_USDT": {"0.001", "0.001"}},
	}

	endpoint := newCustomFeeEndpoint()
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/market/fee"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{
    "error": null,
    "taker": "0.001",
    "maker": "0.002",
    "custom_fee": {
        "BTC_USDT": [
            "0.001",
            "0.001"
        ]
    }
}
`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetCustomFee()

	s.Equal(expectedServerResponse, responseJson)
	s.Equal(err, error(nil))

}

func (s *MainAccountTestSuite) TestMyFeeBYMarketWithResult() {
	expectedServerResponse := MyFeeByMarket{Error: nil, Taker: "0.002", Maker: "0.001"}

	endpoint := newMyFeeByMarketEndpoint("BTC_USDT")
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/market/fee/single"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"error":null,"taker":"0.002","maker":"0.001"}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetMyFeeByMarket("BTC_USDT")

	s.Equal(expectedServerResponse, responseJson)
	s.Equal(err, error(nil))

}
