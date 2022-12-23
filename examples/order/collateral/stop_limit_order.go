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

	fmt.Println("========================= CreateStopLimitOrder ========================= ")
	// Call SDK function CreateStopLimitOrder
	response, err := service.CreateStopLimitOrder(collateral.StopLimitOrderParams{
		Market:          "BTC_USDT",
		Amount:          "0.001",
		Side:            order.SideBuy,
		Price:           "3000",
		ActivationPrice: "3300",
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
