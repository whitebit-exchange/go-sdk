package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/order"
	"github.com/whitebit-exchange/go-sdk/module/order/spot"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create spot service
	service := spot.NewService(client)

	fmt.Println("========================= CreateStopMarketOrder ========================= ")
	// Call SDK function CreateStopMarketOrder
	response, err := service.CreateStopMarketOrder(spot.StopMarketOrderParams{
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
