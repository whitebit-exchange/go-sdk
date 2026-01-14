//go:build example

package main

import (
	"fmt"

	col "github.com/whitebit-exchange/go-sdk/module/order/collateral"
	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

// WARNING: May create a real collateral market order with valid keys/params.
func main() {
	apiKey := ""
	apiSecret := ""

	s := sdkpkg.New(apiKey, apiSecret)

	params := col.MarketOrderParams{
		Market:        "BTC_USDT",
		Side:          "buy",
		Amount:        "0.0001",
		ClientOrderId: "",
	}

	order, err := s.OrdersCollateral.CreateMarketOrder(params)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("order id=%d, market=%s, side=%s, amount=%s\n",
		order.OrderID, order.Market, order.Side, order.Amount)
}
