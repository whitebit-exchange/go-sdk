package server

import "github.com/whitebit-exchange/go-sdk"

const (
	timeEndpointUrl = "/api/v4/public/time"
	pingEndpointUrl = "/api/v4/public/ping"
	wsTokenUrl      = "/api/v4/profile/websocket_token"
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

type wsTokenEndpoint struct {
	whitebit.AuthParams
}

func newWsTokenEndpoint() *wsTokenEndpoint {
	return &wsTokenEndpoint{
		AuthParams: whitebit.NewAuthParams(wsTokenUrl)}
}

func (endpoint *wsTokenEndpoint) Url() string {
	return wsTokenUrl
}

func (endpoint *wsTokenEndpoint) IsAuthed() bool {
	return true
}
