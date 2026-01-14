//go:build example

package main

import (
	"fmt"

	col "github.com/whitebit-exchange/go-sdk/module/order/collateral"
	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

// WARNING: May create a real collateral stop-market order with valid keys/params.
func main() {
	apiKey := ""
	apiSecret := ""

	s := sdkpkg.New(apiKey, apiSecret)

	params := col.StopMarketOrderParams{
		Market:          "BTC_USDT",
		Side:            "sell",
		Amount:          "0.0001",
		ActivationPrice: "10000",
		ClientOrderId:   "",
	}

	order, err := s.OrdersCollateral.CreateStopMarketOrder(params)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("order id=%d, market=%s, side=%s, activation=%s\n",
		order.OrderID, order.Market, order.Side, order.ActivationPrice)
}
