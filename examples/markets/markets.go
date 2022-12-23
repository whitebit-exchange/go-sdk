package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/market"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create a markets service
	service := market.NewService(client)

	fmt.Println("========================= GetMarkets ========================= ")
	// Call SDK function GetMarkets
	response, err := service.GetMarkets()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)
}
