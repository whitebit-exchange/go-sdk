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

	// Create a tickers service
	service := tickers.NewService(client)

	fmt.Println("========================= GetSingleMarketActivity ========================= ")
	// Call SDK function GetSingleMarketActivity
	/*
	   Required params:
	       market string
	*/
	response, err := service.GetSingleMarketActivity("BTC_USDT")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)
}
