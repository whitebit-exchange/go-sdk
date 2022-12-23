package collateral

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

func (s *CollateralAccountTestSuite) TestOpenPositionWithResult() {
	expectedServerResponse := []OpenPosition{
		{PositionID: 664, Market: "BTC_USDT", OpenDate: "1669338402.712101", ModifyDate: "1669375581.3410611",
			Amount: "100", BasePrice: "16558.5199706587661", LiquidationPrice: "8568.62", Leverage: "5",
			Pnl: "-11629.86", PnlPercent: "-0.70", Margin: "371610.4", FreeMargin: "1379086.41", Funding: "0",
			UnrealizedFunding: "0", LiquidationState: ""}}

	expectedResponse, _ := json.Marshal(expectedServerResponse)

	expectedRequest := "/api/v4/collateral-account/positions/open"
	endpoint := newOpenPositionsEndpoint("BTC_USDT")
	request, _ := whitebit.CreateRequest(endpoint.Url())

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`[{"positionId":664,"market":"BTC_USDT","openDate":"1669338402.712101",
	"modifyDate":"1669375581.3410611","amount":"100","basePrice":"16558.5199706587661","liquidationPrice":"8568.62",
	"leverage":"5","pnl":"-11629.86","pnlPercent":"-0.70","margin":"371610.4","freeMargin":"1379086.41",
	"funding":"0","unrealizedFunding":"0","liquidationState":""}]`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetOpenPositions("BTC_USDT")
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.NoError(err)

}

func (s *CollateralAccountTestSuite) TestOpenPositionResponseError() {
	byteResponse := []byte(`{"success": True, "message": null}`)
	expectedRequest := "/api/v4/collateral-account/positions/open"

	endpoint := newOpenPositionsEndpoint("BTC_USDT")
	request, _ := whitebit.CreateRequest(endpoint.Url())

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetOpenPositions("BTC_USDT")
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal("null", string(response))
	s.Equal(expectedError, err.Error())

}
