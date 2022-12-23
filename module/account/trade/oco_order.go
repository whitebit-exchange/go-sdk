package trade

import (
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/account"
)

const ocoOrdersEndpointUrl = "/api/v4/oco-orders"

type ocoListEndpoint struct {
	account.MarketWithPaginationParams
	whitebit.AuthParams
}

func newOcoListEndpoint(market string, limit int, offset int) *ocoListEndpoint {
	return &ocoListEndpoint{
		MarketWithPaginationParams: account.NewMarketWithPaginationParams(market, limit, offset),
		AuthParams:                 whitebit.NewAuthParams(ocoOrdersEndpointUrl),
	}
}
