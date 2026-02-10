//go:build example

package main

import (
	"fmt"

	col "github.com/whitebit-exchange/go-sdk/module/order/collateral"
	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

// WARNING: May create a real collateral OCO order with valid keys/params.
func main() {
	apiKey := ""
	apiSecret := ""

	s := sdkpkg.New(apiKey, apiSecret)

	params := col.OcoOrderParams{
		Market:          "BTC_USDT",
		Side:            "sell",
		Amount:          "0.0001",
		ActivationPrice: "10000",
		StopLimitPrice:  "9999",
		ClientOrderId:   "",
	}

	order, err := s.OrdersCollateral.CreateOcoOrder(params)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("oco id=%d, market=%s\n", order.ID, params.Market)
}
