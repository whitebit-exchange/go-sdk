package main

import (
	"fmt"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/module/order"
	"github.com/whitebit-exchange/whitebit/module/order/collateral"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create collateral service
	service := collateral.NewService(client)

	fmt.Println("========================= CollateralLimit ========================= ")
	// Call SDK function CreateLimitOrder
	response, err := service.CreateLimitOrder(collateral.LimitOrderParams{
		Market: "BTC_USDT",
		Amount: "0.001",
		Side:   order.SideBuy,
		Price:  "3000",
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
