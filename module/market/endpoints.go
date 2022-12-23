package market

import "github.com/whitebit-exchange/go-sdk"

const (
	marketsEndpointUrl           = "/api/v2/public/markets"
	collateralMarketsEndpointUrl = "/api/v4/public/collateral/markets"
)

type infoEndpoint struct {
	whitebit.NoAuth
}

func newMarketsInfoEndpoint() *infoEndpoint {
	return &infoEndpoint{}
}

func (endpoint *infoEndpoint) Url() string {
	return marketsEndpointUrl
}

func (endpoint *infoEndpoint) IsAuthed() bool {
	return false
}

type collateralEndpoint struct {
	whitebit.NoAuth
}

func newCollateralMarketsEndpoint() *collateralEndpoint {
	return &collateralEndpoint{}
}

func (endpoint *collateralEndpoint) Url() string {
	return collateralMarketsEndpointUrl
}
