package assets

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/tests/mocks"
	"testing"
)

type AssetTestSuite struct {
	client   *mocks.Client
	service  *Service
	endpoint *endpoint
	suite.Suite
}

func (s *AssetTestSuite) SetupTest() {
	s.client = &mocks.Client{}
	s.service = NewService(s.client)
	s.endpoint = newAssetsEndpoint()
}

func (s *AssetTestSuite) TestWithResult() {
	byteResponse := []byte(`{
"1INCH":{"name":"1inch","unified_cryptoasset_id":0,"can_withdraw":true,"can_deposit":true,
	"min_withdraw":"25","max_withdraw":"0","maker_fee":"0.1","taker_fee":"0.1","min_deposit":"15", 
	"max_deposit":"0","networks":{"deposits":["ERC20"],"withdraws":["ERC20"],"default":"ERC20"},
	"confirmations":{"ERC20":15},"providers":{},"limits":{"deposit":{"ERC20":{"min":"15"}},
	"withdraw":{"ERC20":{"min":"25"}}},"currency_precision":18,"is_memo":false},
"AAPL":	{"name":"Apple Inc","unified_cryptoasset_id":0,"can_withdraw":false,"can_deposit":false,
	"min_withdraw":"0","max_withdraw":"0","maker_fee":"0.1","taker_fee":"0.1","min_deposit":"0",
	"max_deposit":"0","networks":{"default":"ERC20"},"confirmations":{"ERC20":15},"providers":{},
	"limits":{"deposit":{"ERC20":{"min":"0"}}, "withdraw":{"ERC20":{"min":"0"}}}, "currency_precision":18, 
	"is_memo":false}}`)

	assetMap := map[string]Asset{
		"1INCH": {
			Name:        "1inch",
			UCID:        0,
			CanWithdraw: true,
			CanDeposit:  true,
			MinWithdraw: "25",
			MaxWithdraw: "0",
			MakerFee:    "0.1",
			TakerFee:    "0.1",
			MinDeposit:  "15",
			MaxDeposit:  "0",
			Networks: AssetNetworks{
				AssetProviders: AssetProviders{
					Deposits:  []string{"ERC20"},
					Withdraws: []string{"ERC20"}},
				Default: "ERC20"},
			Confirmations: map[string]int{"ERC20": 15},
			Providers:     AssetProviders{},
			Limits: AssetLimits{
				Deposit:  map[string]AssetLimitItem{"ERC20": {Min: "15"}},
				Withdraw: map[string]AssetLimitItem{"ERC20": {Min: "25"}}},
			CurrencyPrecision: 18, IsMemo: false},
		"AAPL": {
			Name:        "Apple Inc",
			UCID:        0,
			CanWithdraw: false,
			CanDeposit:  false,
			MinWithdraw: "0",
			MaxWithdraw: "0",
			MakerFee:    "0.1",
			TakerFee:    "0.1",
			MinDeposit:  "0",
			MaxDeposit:  "0",
			Networks: AssetNetworks{
				Default: "ERC20"},
			Confirmations: map[string]int{"ERC20": 15},
			Providers:     AssetProviders{},
			Limits: AssetLimits{
				Deposit:  map[string]AssetLimitItem{"ERC20": {Min: "0"}},
				Withdraw: map[string]AssetLimitItem{"ERC20": {Min: "0"}}},
			CurrencyPrecision: 18, IsMemo: false},
	}

	request, _ := whitebit.CreateRequest(s.endpoint.Url())
	expectedRequest := "/api/v4/public/assets"

	s.Equal(expectedRequest, request.URL.String())

	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetAssets()

	response, _ := json.Marshal(responseJson)
	expectedResponse, _ := json.Marshal(assetMap)
	s.Equal(string(expectedResponse), string(response))
	s.Equal(err, error(nil))

}

func (s *AssetTestSuite) TestInvalidResponseError() {
	byteResponse := []byte(`"": {"name":"1inch", "unified_cryptoasset_id":0}}`)

	s.client.On("SendRequest", s.endpoint).Return(byteResponse, nil).Once()

	responseJson, err := s.service.GetAssets()
	response, _ := json.Marshal(responseJson)

	expectedResponse := "null"
	s.Equal(expectedResponse, string(response))

	expectedError := "invalid character ':' after top-level value"
	s.Equal(expectedError, err.Error())

}

func TestAssetTestSuite(t *testing.T) {
	suite.Run(t, new(AssetTestSuite))
}
