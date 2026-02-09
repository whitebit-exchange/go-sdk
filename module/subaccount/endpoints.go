package subaccount

import "github.com/whitebit-exchange/go-sdk"

const (
	createURL          = "/api/v4/sub-account/create"
	deleteURL          = "/api/v4/sub-account/delete"
	editURL            = "/api/v4/sub-account/edit"
	listURL            = "/api/v4/sub-account/list"
	transferURL        = "/api/v4/sub-account/transfer"
	blockURL           = "/api/v4/sub-account/block"
	unblockURL         = "/api/v4/sub-account/unblock"
	balancesURL        = "/api/v4/sub-account/balances"
	transferHistoryURL = "/api/v4/sub-account/transfer/history"
)

type createEndpoint struct {
	whitebit.AuthParams
	Alias string `json:"alias"`
	Email string `json:"email,omitempty"`
}

func newCreateEndpoint(alias, email string) *createEndpoint {
	return &createEndpoint{
		AuthParams: whitebit.NewAuthParams(createURL),
		Alias:      alias,
		Email:      email,
	}
}

type deleteEndpoint struct {
	whitebit.AuthParams
	ID string `json:"id"`
}

func newDeleteEndpoint(id string) *deleteEndpoint {
	return &deleteEndpoint{
		AuthParams: whitebit.NewAuthParams(deleteURL),
		ID:         id,
	}
}

type editEndpoint struct {
	whitebit.AuthParams
	ID    string `json:"id"`
	Alias string `json:"alias"`
}

func newEditEndpoint(id, alias string) *editEndpoint {
	return &editEndpoint{
		AuthParams: whitebit.NewAuthParams(editURL),
		ID:         id,
		Alias:      alias,
	}
}

type listEndpoint struct {
	whitebit.AuthParams
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

func newListEndpoint(limit, offset int) *listEndpoint {
	return &listEndpoint{
		AuthParams: whitebit.NewAuthParams(listURL),
		Limit:      limit,
		Offset:     offset,
	}
}

type transferEndpoint struct {
	whitebit.AuthParams
	FromSubAccountID string `json:"fromSubAccountUid,omitempty"`
	ToSubAccountID   string `json:"toSubAccountUid,omitempty"`
	Ticker           string `json:"ticker"`
	Amount           string `json:"amount"`
}

func newTransferEndpoint(from, to, ticker, amount string) *transferEndpoint {
	return &transferEndpoint{
		AuthParams:       whitebit.NewAuthParams(transferURL),
		FromSubAccountID: from,
		ToSubAccountID:   to,
		Ticker:           ticker,
		Amount:           amount,
	}
}

type blockEndpoint struct {
	whitebit.AuthParams
	ID string `json:"id"`
}

func newBlockEndpoint(id string) *blockEndpoint {
	return &blockEndpoint{
		AuthParams: whitebit.NewAuthParams(blockURL),
		ID:         id,
	}
}

type unblockEndpoint struct {
	whitebit.AuthParams
	ID string `json:"id"`
}

func newUnblockEndpoint(id string) *unblockEndpoint {
	return &unblockEndpoint{
		AuthParams: whitebit.NewAuthParams(unblockURL),
		ID:         id,
	}
}

type balancesEndpoint struct {
	whitebit.AuthParams
	ID     string `json:"id"`
	Ticker string `json:"ticker,omitempty"`
}

func newBalancesEndpoint(id, ticker string) *balancesEndpoint {
	return &balancesEndpoint{
		AuthParams: whitebit.NewAuthParams(balancesURL),
		ID:         id,
		Ticker:     ticker,
	}
}

type transferHistoryEndpoint struct {
	whitebit.AuthParams
	SubAccountID string `json:"subAccountUid,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	Offset       int    `json:"offset,omitempty"`
}

func newTransferHistoryEndpoint(subAccountID string, limit, offset int) *transferHistoryEndpoint {
	return &transferHistoryEndpoint{
		AuthParams:   whitebit.NewAuthParams(transferHistoryURL),
		SubAccountID: subAccountID,
		Limit:        limit,
		Offset:       offset,
	}
}
