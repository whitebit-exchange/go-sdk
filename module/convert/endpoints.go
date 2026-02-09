package convert

import "github.com/whitebit-exchange/go-sdk"

const (
	estimateURL = "/api/v4/convert/estimate"
	confirmURL  = "/api/v4/convert/confirm"
	historyURL  = "/api/v4/convert/history"
)

type estimateEndpoint struct {
	whitebit.AuthParams
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

func newEstimateEndpoint(from, to, amount string) *estimateEndpoint {
	return &estimateEndpoint{
		AuthParams: whitebit.NewAuthParams(estimateURL),
		From:       from,
		To:         to,
		Amount:     amount,
	}
}

type confirmEndpoint struct {
	whitebit.AuthParams
	QueryID string `json:"queryId"`
}

func newConfirmEndpoint(queryID string) *confirmEndpoint {
	return &confirmEndpoint{
		AuthParams: whitebit.NewAuthParams(confirmURL),
		QueryID:    queryID,
	}
}

type historyEndpoint struct {
	whitebit.AuthParams
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

func newHistoryEndpoint(limit, offset int) *historyEndpoint {
	return &historyEndpoint{
		AuthParams: whitebit.NewAuthParams(historyURL),
		Limit:      limit,
		Offset:     offset,
	}
}
