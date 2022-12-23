package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/order"
	"github.com/whitebit-exchange/go-sdk/module/order/spot"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create spot service
	service := spot.NewService(client)

	fmt.Println("========================= CreateMarketStock ========================= ")
	// Call SDK function CreateMarketStock
	response, err := service.CreateMarketStock(spot.MarketOrderParams{
		Market: "BTC_USDT",
		Amount: "0.001",
		Side:   order.SideBuy,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
