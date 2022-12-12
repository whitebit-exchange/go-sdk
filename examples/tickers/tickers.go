package main

import (
	"fmt"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/module/tickers"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create a tickers service
	service := tickers.NewService(client)

	fmt.Println("========================= GetTickers ========================= ")
	// Call SDK function GetTickers
	response, err := service.GetTickers()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)
}
