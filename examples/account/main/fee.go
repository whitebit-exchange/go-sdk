package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/account/main_account"
)

func main() {
	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create main account service
	service := main_account.NewService(client)

	fmt.Println("========================= GetFee ========================= ")

	// Returns an array of objects containing deposit/withdrawal settings for the corresponding currencies.
	// Zero value in amount fields means that the setting is disabled.
	response, err := service.GetFee()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
