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

	fmt.Println("========================= CreateKillSwitch ========================= ")
	// Call SDK function CreateKillSwitch

	response, err := service.CreateKillSwitch(spot.KillSwitchParams{Market: "BTC_USDT", Timeout: "120", Types: []string{spot.OrderTypeSpot}})
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
