//go:build example

package main

import (
	"fmt"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

// WARNING: Attempts to cancel a collateral OCO order. Use a real OCO order ID only if you are sure.
func main() {
	apiKey := ""
	apiSecret := ""

	s := sdkpkg.New(apiKey, apiSecret)

	market := "BTC_USDT"
	orderID := int64(0) // change to a real OCO order ID

	res, err := s.OrdersCollateral.CancelOcoOrder(market, orderID)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("canceled OCO id=%d, market=%s\n", res.OrderID, market)
}
