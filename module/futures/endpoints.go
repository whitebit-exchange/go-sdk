package futures

import "github.com/whitebit-exchange/whitebit"

const futuresEndpointUrl = "/api/v4/public/futures"

type endpoint struct {
	whitebit.NoAuth
}

func newFuturesMarketsEndpoint() *endpoint {
	return &endpoint{}
}

func (endpoint *endpoint) Url() string {
	return futuresEndpointUrl
}
