package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/fee"
)

func main() {
	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create a fee service
	service := fee.NewService(client)

	fmt.Println("========================= GetTradingFeesList ========================= ")
	// Call SDK function GetTradingFeesList
	response, err := service.GetTradingFeesList()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
