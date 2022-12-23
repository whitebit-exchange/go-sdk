package symbol

import "github.com/whitebit-exchange/go-sdk"

const symbolsEndpointUrl = "/api/v1/public/symbols"

type endpoint struct {
	whitebit.NoAuth
}

func newEndpoint() *endpoint {
	return &endpoint{}
}

func (endpoint *endpoint) Url() string {
	return symbolsEndpointUrl
}
