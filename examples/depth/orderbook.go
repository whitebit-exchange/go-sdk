package main

import (
	"fmt"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/module/depth"
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
	       limit  int
	       level  int
	*/
	options := depth.OrderBookOptions{
		Market: "BTC_USDT",
		Limit:  100,
		Level:  1,
	}
	// Create a depth service
	service := depth.NewService(client)

	fmt.Println("========================= GetOrderBook ========================= ")
	// Call SDK function GetOrderBook with endpoints options
	response, err := service.GetOrderBook(options)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)
}
