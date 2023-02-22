package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/tickers"
)

func main() {
	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create a ticker service
	service := tickers.NewService(client)

	fmt.Println("========================= GetAvailableTickers ========================= ")
	// Call SDK function GetAvailableTickers
	response, err := service.GetAvailableTickers()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)
}
