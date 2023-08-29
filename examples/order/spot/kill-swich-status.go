package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
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

	fmt.Println("========================= GetKillSwitchStatus ========================= ")
	// Call SDK function GetKillSwitchStatus

	response, err := service.GetKillSwitchStatus(spot.KillSwitchStatusParams{Market: "BTC_USDT"})
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
