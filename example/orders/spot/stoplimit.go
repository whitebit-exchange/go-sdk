//go:build example

package main

import (
	"fmt"

	spot "github.com/whitebit-exchange/go-sdk/module/order/spot"
	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

// WARNING: May place a real stop-limit order with valid keys/params.
func main() {
	apiKey := ""    // your API key
	apiSecret := "" // your API secret

	s := sdkpkg.New(apiKey, apiSecret)

	params := spot.StopLimitOrderParams{
		Market:          "BTC_USDT",
		Side:            "buy",
		Amount:          "0.0001",
		Price:           "11", // limit price
		ActivationPrice: "10", // far activation price for safety
		ClientOrderId:   "",
	}

	order, err := s.OrdersSpot.CreateStopLimitOrder(params)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("order id=%d, market=%s, side=%s, activation=%s, price=%s\n",
		order.OrderID, order.Market, order.Side, order.ActivationPrice, order.Price)
}
