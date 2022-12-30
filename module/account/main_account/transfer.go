package main_account

import (
	"github.com/whitebit-exchange/go-sdk"
)

const transferEndpointUrl = "/api/v4/main-account/transfer"

type transferEndpoint struct {
	whitebit.AuthParams
	TransferParams
}

type TransferParams struct {
	Ticker string `json:"ticker"`
	Amount string `json:"amount"`
	From   string `json:"from"`
	To     string `json:"to"`
}

const (
	Main       = "main"
	Trade      = "trade"
	Collateral = "collateral"
)

func newTransferEndpoint(params TransferParams) *transferEndpoint {
	return &transferEndpoint{TransferParams: params, AuthParams: whitebit.NewAuthParams(transferEndpointUrl)}
}
