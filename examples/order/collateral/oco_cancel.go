package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
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

	fmt.Println("========================= OcoCancel ========================= ")
	// Call SDK function OcoCancel

	response, err := service.CancelOcoOrder("BTC_USDT", 3313805040)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
