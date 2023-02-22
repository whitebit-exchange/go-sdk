package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/depth"
)

func main() {
	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create a depth service
	service := depth.NewService(client)

	fmt.Println("========================= GetDepth ========================= ")
	// Call SDK function GetDepth with endpoints options
	/*
	   Required params:
	       market string
	*/
	response, err := service.GetDepth("BTC_USDT")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)
}
