package whitebit

import (
	"time"
)

type Endpoint interface {
	Url() string
	IsAuthed() bool
}

type NoAuth struct{}

func (endpoint NoAuth) IsAuthed() bool {
	return false
}

func (endpoint NoAuth) Url() string {
	return ""
}

type AuthParams struct {
	Request     string `json:"request"`
	Nonce       int64  `json:"nonce"`
	NonceWindow bool   `json:"nonceWindow"`
}

func NewAuthParams(url string) AuthParams {
	return AuthParams{Nonce: time.Now().UnixMilli(), NonceWindow: true, Request: url}
}

func (params AuthParams) IsAuthed() bool {
	return true
}

func (params AuthParams) Url() string {
	return params.Request
}
