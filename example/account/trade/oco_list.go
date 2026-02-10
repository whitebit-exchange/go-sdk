//go:build example

package main

import (
	"fmt"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

// Lists OCO orders for a market (read-only). May require authentication.
func main() {
	apiKey := ""
	apiSecret := ""

	s := sdkpkg.New(apiKey, apiSecret)

	market := "BTC_USDT"
	limit := 10
	offset := 0

	res, err := s.AccountTrade.GetOcoOrders(market, limit, offset)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("oco orders: %d\n", len(res))
	if len(res) > 0 {
		fmt.Printf("first oco id=%d\n", res[0].ID)
	}
}
