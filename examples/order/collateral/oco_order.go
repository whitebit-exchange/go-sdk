package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/order"
	"github.com/whitebit-exchange/go-sdk/module/order/collateral"
)

func main() {
	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create collateral service
	service := collateral.NewService(client)

	fmt.Println("========================= CreateOcoOrder ========================= ")
	// Call SDK function CreateOcoOrder
	response, err := service.CreateOcoOrder(collateral.OcoOrderParams{
		Market:          "BTC_USDT",
		Amount:          "100",
		Side:            order.SideBuy,
		Price:           "2000",
		ActivationPrice: "30300",
		StopLimitPrice:  "40000",
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
