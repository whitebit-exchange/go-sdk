package main

import (
	"fmt"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/module/order"
	"github.com/whitebit-exchange/whitebit/module/order/collateral"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create collateral service
	service := collateral.NewService(client)

	fmt.Println("========================= CreateStopMarketOrder ========================= ")
	// Call SDK function CreateStopMarketOrder
	response, err := service.CreateStopMarketOrder(collateral.StopMarketOrderParams{
		Market:          "BTC_USDT",
		Amount:          "100",
		Side:            order.SideBuy,
		ActivationPrice: "3000",
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
