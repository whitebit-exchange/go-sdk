package main

import (
	"fmt"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/module/fee"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create a fee service
	service := fee.NewService(client)

	fmt.Println("========================= GetTradingFee ========================= ")
	// Call SDK function GetTradingFee
	response, err := service.GetTradingFee()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
