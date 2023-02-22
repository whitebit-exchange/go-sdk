package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/futures"
)

func main() {
	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create a futures service
	service := futures.NewService(client)

	fmt.Println("========================= GetFuturesMarkets ========================= ")
	// Call SDK function GetFuturesMarkets
	response, err := service.GetFuturesMarkets()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)
}
