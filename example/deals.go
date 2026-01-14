//go:build example

package main

import (
	"fmt"

	"github.com/whitebit-exchange/go-sdk/module/deal"
	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

func main() {
	s := sdkpkg.New("", "")
	opts := deal.Options{Market: "BTC_USDT"}
	res, err := s.Deals.GetDeals(opts)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("deals=%d\n", len(res))
	if len(res) > 0 {
		fmt.Println("first tradeID:", res[0].TradeID)
	}
}
