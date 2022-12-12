package main

import (
	"fmt"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/module/account/trade"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create trade service
	service := trade.NewService(client)

	fmt.Println("========================= GetOrders ========================= ")
	// Call SDK function orders

	response, err := service.GetOrders("BTC_USDT", 100, 0)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
