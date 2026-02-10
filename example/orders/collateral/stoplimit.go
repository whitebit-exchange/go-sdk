//go:build example

package main

import (
	"fmt"

	col "github.com/whitebit-exchange/go-sdk/module/order/collateral"
	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

// WARNING: May create a real collateral stop-limit order with valid keys/params.
func main() {
	apiKey := ""
	apiSecret := ""

	s := sdkpkg.New(apiKey, apiSecret)

	params := col.StopLimitOrderParams{
		Market:          "BTC_USDT",
		Side:            "sell",
		Amount:          "0.0001",
		Price:           "9",
		ActivationPrice: "10",
		ClientOrderId:   "",
	}

	order, err := s.OrdersCollateral.CreateStopLimitOrder(params)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("order id=%d, market=%s, side=%s, activation=%s, price=%s\n",
		order.OrderID, order.Market, order.Side, order.ActivationPrice, order.Price)
}
