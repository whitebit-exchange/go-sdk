package fee

import "github.com/whitebit-exchange/whitebit"

const (
	tradingFeeEndpointUrl = "/api/v2/public/fee"
	feeListEndpointUrl    = "/api/v4/public/fee"
)

type tradingFeeEndpoint struct {
	whitebit.NoAuth
}

func newTradingFeeEndpoint() *tradingFeeEndpoint {
	return &tradingFeeEndpoint{}
}

func (endpoint *tradingFeeEndpoint) Url() string {
	return tradingFeeEndpointUrl
}

func (endpoint *tradingFeeEndpoint) IsAuthed() bool {
	return false
}

type listEndpoint struct {
	whitebit.NoAuth
}

func newListEndpoint() *listEndpoint {
	return &listEndpoint{}
}

func (endpoint *listEndpoint) Url() string {
	return feeListEndpointUrl
}
