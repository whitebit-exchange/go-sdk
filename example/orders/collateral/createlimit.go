//go:build example

package main

import (
	"fmt"

	col "github.com/whitebit-exchange/go-sdk/module/order/collateral"
	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

// WARNING: May create a real collateral limit order with valid keys/params.
func main() {
	apiKey := ""    // your API key
	apiSecret := "" // your API secret

	s := sdkpkg.New(apiKey, apiSecret)

	params := col.LimitOrderParams{
		Market:        "BTC_USDT",
		Side:          "buy",
		Amount:        "0.0001",
		Price:         "10",
		ClientOrderId: "",
	}

	order, err := s.OrdersCollateral.CreateLimitOrder(params)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("order id=%d, market=%s, side=%s, price=%s\n",
		order.OrderID, order.Market, order.Side, order.Price)
}
