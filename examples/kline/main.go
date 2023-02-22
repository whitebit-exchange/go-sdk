package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/kline"
)

func main() {
	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)
	/*
	   Create Options
	   Required params:
	       market   string
	   Optional params:
	       start    int
	       end      int
	       interval string
	       limit    int
	*/

	options := kline.Options{Market: "BTC_USDT", Start: 1667469600, End: 1667487600, Limit: 100}

	// Create a kline service
	service := kline.NewService(client)

	fmt.Println("========================= GetKline ========================= ")
	// Call SDK function GetKline with endpoints options
	response, err := service.GetKline(options)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)
}
