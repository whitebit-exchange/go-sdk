package collateral

import "github.com/whitebit-exchange/go-sdk"

const fundingHistoryEndpointURL = "/api/v4/collateral-account/funding-history"

type fundingHistoryEndpoint struct {
	whitebit.AuthParams
	Market string `json:"market,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
}

func newFundingHistoryEndpoint(market string, limit, offset int) *fundingHistoryEndpoint {
	return &fundingHistoryEndpoint{
		AuthParams: whitebit.NewAuthParams(fundingHistoryEndpointURL),
		Market:     market,
		Limit:      limit,
		Offset:     offset,
	}
}

type FundingRecord struct {
	ID          int64  `json:"id"`
	Market      string `json:"market"`
	FundingRate string `json:"fundingRate"`
	Amount      string `json:"amount"`
	Timestamp   int64  `json:"timestamp"`
}

type FundingHistoryResponse struct {
	Total   int             `json:"total"`
	Records []FundingRecord `json:"records"`
	Limit   int             `json:"limit"`
	Offset  int             `json:"offset"`
}
