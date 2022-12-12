package main

import (
	"fmt"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/module/order"
	"github.com/whitebit-exchange/whitebit/module/order/spot"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create spot service
	service := spot.NewService(client)

	fmt.Println("========================= CreateStopLimitOrder ========================= ")
	// Call SDK function CreateStopLimitOrder
	response, err := service.CreateStopLimitOrder(spot.StopLimitOrderParams{
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
