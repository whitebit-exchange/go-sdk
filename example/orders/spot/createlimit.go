//go:build example

package main

import (
	"fmt"

	spot "github.com/whitebit-exchange/go-sdk/module/order/spot"
	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

// WARNING: Running this with real keys and valid parameters will CREATE A REAL LIMIT ORDER.
// Defaults are chosen to avoid accidental fills.
func main() {
	apiKey := ""    // your API key
	apiSecret := "" // your API secret

	s := sdkpkg.New(apiKey, apiSecret)

	params := spot.LimitOrderParams{
		Market:        "BTC_USDT",
		Side:          "buy",
		Amount:        "0.0001",
		Price:         "10",
		ClientOrderId: "",
	}

	order, err := s.OrdersSpot.CreateLimitOrder(params)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("order id=%d, market=%s, side=%s, type=%s, price=%s, left=%s\n",
		order.OrderID, order.Market, order.Side, order.Type, order.Price, order.Left)
}
