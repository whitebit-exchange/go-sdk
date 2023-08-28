package spot

import (
	"github.com/whitebit-exchange/go-sdk"
)

const (
	killSwitchEndpointUrl       = "/api/v4/order/kill-switch"
	killSwitchStatusEndpointUrl = "/api/v4/order/kill-switch/status"
	OrderTypeSpot               = "spot"
	OrderTypeMargin             = "margin"
	OrderTypeFutures            = "futures"
)

type KillSwitchResponse struct {
	Market           string   `json:"market"`
	StartTime        int      `json:"startTime"`
	CancellationTime int      `json:"cancellationTime"`
	Types            []string `json:"types"`
}

type KillSwitchParams struct {
	Market  string   `json:"market"`
	Timeout string   `json:"timeout"`
	Types   []string `json:"types,omitempty"`
}

type KillSwitchEndpoint struct {
	whitebit.AuthParams
	KillSwitchParams
}

func newKillSwitchEndpoint(params KillSwitchParams) *KillSwitchEndpoint {
	return &KillSwitchEndpoint{
		AuthParams:       whitebit.NewAuthParams(killSwitchEndpointUrl),
		KillSwitchParams: params,
	}
}

type KillSwitchStatusResponse struct {
	Market           string   `json:"market"`
	StartTime        int      `json:"startTime"`
	CancellationTime int      `json:"cancellationTime"`
	Types            []string `json:"types"`
}

type KillSwitchStatusParams struct {
	Market string `json:"market,omitempty"`
}

type KillSwitchStatusEndpoint struct {
	whitebit.AuthParams
	KillSwitchStatusParams
}

func newKillSwitchStatusEndpoint(params KillSwitchStatusParams) *KillSwitchStatusEndpoint {
	return &KillSwitchStatusEndpoint{
		AuthParams:             whitebit.NewAuthParams(killSwitchStatusEndpointUrl),
		KillSwitchStatusParams: params,
	}
}
