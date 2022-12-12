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

	fmt.Println("========================= GetExecutedHistory ========================= ")
	// Call SDK function ExecutedHistory
	/*
		   Create Options
		   Optional params:
				market   	  string
				Limit         int
				Offset        int
				ClientOrderId string
	*/

	response, err := service.GetExecutedHistory(trade.ExecutedHistoryOptions{Market: "BTC_USDT", Limit: 2})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
