//go:build example

package main

import (
	"fmt"

	spot "github.com/whitebit-exchange/go-sdk/module/order/spot"
	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

// WARNING: May place a real stop-market order with valid keys/params.
func main() {
	apiKey := ""    // your API key
	apiSecret := "" // your API secret

	s := sdkpkg.New(apiKey, apiSecret)

	params := spot.StopMarketOrderParams{
		Market:          "BTC_USDT",
		Side:            "sell",
		Amount:          "0.0001",
		ActivationPrice: "10000", // far price for safety
		ClientOrderId:   "",
	}

	order, err := s.OrdersSpot.CreateStopMarketOrder(params)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("order id=%d, market=%s, side=%s, activation=%s\n",
		order.OrderID, order.Market, order.Side, order.ActivationPrice)
}
