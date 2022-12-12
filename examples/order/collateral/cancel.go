package main

import (
	"fmt"
	"github.com/whitebit-exchange/whitebit"
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

	fmt.Println("========================= CancelOrder ========================= ")
	// Call SDK function CancelOrder

	response, err := service.CancelOrder("BTC_USDT", 3310300051)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
