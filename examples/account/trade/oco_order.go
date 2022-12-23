package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/account/trade"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create trade service
	service := trade.NewService(client)

	fmt.Println("========================= GetOcoOrders ========================= ")
	// Call SDK function GetOcoOrders
	response, err := service.GetOcoOrders("BTC_USDT", 100, 0)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
