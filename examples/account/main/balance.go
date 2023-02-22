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

	// Create account main service
	service := main_account.NewService(client)

	fmt.Println("========================= GetBalance ========================= ")
	// Call SDK function GetBalance for main account

	allBalance, err := service.GetMainBalance()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", allBalance)

	fmt.Println("========================= GetBalance By Ticker ========================= ")
	// Call SDK function GetBalance by ticker

	tickerBalance, err := service.GetAssetBalance("BTC")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", tickerBalance)
}
