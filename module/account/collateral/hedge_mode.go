package collateral

import "github.com/whitebit-exchange/go-sdk"

const (
	hedgeModeEndpointURL       = "/api/v4/collateral-account/hedge-mode"
	hedgeModeUpdateEndpointURL = "/api/v4/collateral-account/hedge-mode/update"
)

type hedgeModeEndpoint struct {
	whitebit.AuthParams
}

func newHedgeModeEndpoint() *hedgeModeEndpoint {
	return &hedgeModeEndpoint{
		AuthParams: whitebit.NewAuthParams(hedgeModeEndpointURL),
	}
}

type hedgeModeUpdateEndpoint struct {
	whitebit.AuthParams
	HedgeMode bool `json:"hedgeMode"`
}

func newHedgeModeUpdateEndpoint(hedgeMode bool) *hedgeModeUpdateEndpoint {
	return &hedgeModeUpdateEndpoint{
		AuthParams: whitebit.NewAuthParams(hedgeModeUpdateEndpointURL),
		HedgeMode:  hedgeMode,
	}
}

type HedgeModeResponse struct {
	HedgeMode bool `json:"hedgeMode"`
}
