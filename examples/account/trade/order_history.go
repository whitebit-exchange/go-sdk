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

	fmt.Println("========================= GetHistory ========================= ")
	// Call SDK function history

	/*
		   Create Options
		   Optional params:
				market   	  string
				Limit         int
				Offset        int
				OrderId       int
				ClientOrderId string
	*/

	response, err := service.GetHistory(trade.HistoryOptions{Market: "BTC_USDT"})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
