//go:build example

package main

import (
	"fmt"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

// WARNING: Attempts to cancel a collateral order. Use a real order ID only if you are sure.
func main() {
	apiKey := ""
	apiSecret := ""

	s := sdkpkg.New(apiKey, apiSecret)

	market := "BTC_USDT"
	orderID := int64(0) // change to a real order ID

	res, err := s.OrdersCollateral.CancelOrder(market, orderID)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("canceled order id=%d, market=%s, side=%s\n", res.OrderID, res.Market, res.Side)
}
