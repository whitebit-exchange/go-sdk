package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/deal"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	/*
	   Create tradeHistoryEndpoint endpoint
	   Required params:
	       market string
	       lastId int
	   Optional params:
	       limit  int
	*/
	options := deal.TradeHistoryOptions{Market: "BTC_USDT", LastId: 1, Limit: 50}

	// Create a deals service
	service := deal.NewService(client)

	fmt.Println("========================= GetTradeHistory ========================= ")
	// Call SDK function GetTradeHistory with endpoints options
	response, err := service.GetTradeHistory(options)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)
}
