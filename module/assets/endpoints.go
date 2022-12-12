package assets

import "github.com/whitebit-exchange/whitebit"

const assetsEndpointUrl = "/api/v4/public/assets"

type AssetLimitItem struct {
	Max string `json:"max,omitempty"`
	Min string `json:"min,omitempty"`
}

type endpoint struct {
	whitebit.NoAuth
}

func newAssetsEndpoint() *endpoint {
	return &endpoint{}
}

func (endpoint *endpoint) Url() string {
	return assetsEndpointUrl
}
