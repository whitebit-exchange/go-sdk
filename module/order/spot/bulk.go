package spot

import (
	"github.com/whitebit-exchange/go-sdk"
)

const bulkEndpointUrl = "/api/v4/order/bulk"

type BulkOrderResponseRecord struct {
	Result *LimitOrder     `json:"result"`
	Error  *whitebit.Error `json:"error"`
}


type BulkOrderEndpoint struct {
	whitebit.AuthParams
	Orders []LimitOrderParams `json:"orders"`
}

func newBulkEndpoint(params []LimitOrderParams) *BulkOrderEndpoint {
	return &BulkOrderEndpoint{
		AuthParams: whitebit.NewAuthParams(bulkEndpointUrl),
		Orders:     params,
	}
}
