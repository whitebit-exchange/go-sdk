package main

import (
	"fmt"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/module/symbol"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create a symbols service
	service := symbol.NewService(client)

	fmt.Println("========================= GetSymbols ========================= ")
	// Call SDK function GetSymbols
	response, err := service.GetSymbols()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)
}
