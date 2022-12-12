package collateral

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/whitebit"
)

func (s *CollateralAccountTestSuite) TestPositionWithResult() {
	expectedServerResponse := []PositionHistory{
		{PositionID: 664, Market: "BTC_USDT", OpenDate: "1669338402.712101", ModifyDate: "1669375581.3410611",
			Amount: "100", BasePrice: "16558.5199706588", RealizedFunding: "0", LiquidationPrice: "8568.62",
			LiquidationState: "", OrderDetail: OrderDetail{ID: 3313899825,
				TradeAmount: "100", Price: "2000", TradeFee: "0", FundingFee: "0"}},
		{PositionID: 664, Market: "BTC_USDT", OpenDate: "1669338402.712101",
			ModifyDate: "1669374926.437654", Amount: "100", BasePrice: "16558.5199706588", RealizedFunding: "0",
			LiquidationPrice: "8568.62", LiquidationState: "", OrderDetail: OrderDetail{ID: 3313805042,
				TradeAmount: "0", Price: "16575.21", TradeFee: "0", FundingFee: "0"}}}

	byteResponse := []byte(`[{"positionId":664,"market":"BTC_USDT","openDate":1669338402.712101,
		"modifyDate":1669375581.3410611,"amount":"100","basePrice":"16558.5199706588","realizedFunding":"0",
		"liquidationPrice":"8568.62","liquidationState":"","orderDetail":{"id":3313899825,"tradeAmount":"100",
		"price":"2000","tradeFee":"0","fundingFee":"0"}},{"positionId":664,"market":"BTC_USDT",
		"openDate":1669338402.712101,"modifyDate":1669374926.437654,"amount":"100",
		"basePrice":"16558.5199706588","realizedFunding":"0","liquidationPrice":"8568.62",
		"liquidationState":"","orderDetail":{"id":3313805042,"tradeAmount":"0","price":"16575.21",
		"tradeFee":"0","fundingFee":"0"}}]`)
	expectedResponse, _ := json.Marshal(expectedServerResponse)

	expectedRequest := "/api/v4/collateral-account/positions/history"
	endpoint := newPositionsHistoryEndpoint("BTC_USDT", 0, 0, 0, "2", "")
	request, _ := whitebit.CreateRequest(endpoint.Url())

	s.Equal(expectedRequest, request.URL.String())
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetPositionsHistory(PositionsHistoryOptions{Market: "BTC_USDT", Limit: "2"})
	response, _ := json.Marshal(responseJson)

	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *CollateralAccountTestSuite) TestPositionResponseError() {
	endpoint := newPositionsHistoryEndpoint("BTC_USDT", 0, 0, 0, "", "")
	request, _ := whitebit.CreateRequest(endpoint.Url())
	expectedRequest := "/api/v4/collateral-account/positions/history"

	s.Equal(expectedRequest, request.URL.String())

	byteResponse := []byte(`{"success": True, "message": null}`)
	s.client.On("SendRequest", mock.Anything).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetPositionsHistory(PositionsHistoryOptions{Market: "BTC_USDT"})
	response, _ := json.Marshal(responseJson)

	expectedError := "invalid character 'T' looking for beginning of value"
	s.Equal("null", string(response))
	s.Equal(expectedError, err.Error())

}
