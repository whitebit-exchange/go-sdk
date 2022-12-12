package server

import "github.com/whitebit-exchange/whitebit"

const (
	timeEndpointUrl = "/api/v4/public/time"
	pingEndpointUrl = "/api/v4/public/ping"
)

type timeEndpoint struct {
	whitebit.NoAuth
}

func newTimeEndpoint() *timeEndpoint {
	return &timeEndpoint{}
}

func (endpoint *timeEndpoint) Url() string {
	return timeEndpointUrl
}

type pingEndpoint struct {
	whitebit.NoAuth
}

func newPingEndpoint() *pingEndpoint {
	return &pingEndpoint{}
}

func (endpoint *pingEndpoint) Url() string {
	return pingEndpointUrl
}
