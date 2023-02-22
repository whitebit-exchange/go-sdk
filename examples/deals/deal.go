package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/deal"
)

func main() {
	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	/*
	   Create options
	   Required params:
	       market string
	   Optional params:
	       type  string
	*/
	options := deal.Options{Market: "BTC_USDT", Type: "Sell"}

	// Create a deals service
	service := deal.NewService(client)

	fmt.Println("========================= GetDeals ========================= ")
	// Call SDK function GetDeals with endpoints options
	response, err := service.GetDeals(options)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)
}
