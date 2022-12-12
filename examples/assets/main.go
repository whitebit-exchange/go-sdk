package main

import (
	"fmt"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/module/assets"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create assets service
	service := assets.NewService(client)

	fmt.Println("========================= GetAssets ========================= ")
	// Call SDK function GetAssets
	response, err := service.GetAssets()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
