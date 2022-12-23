package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/account/collateral"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create collateral service
	service := collateral.NewService(client)

	fmt.Println("========================= GetOpenPositions ========================= ")
	// Call SDK function GetOpenPositions

	response, err := service.GetOpenPositions("BTC_USDT")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
