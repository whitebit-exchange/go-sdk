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

	fmt.Println("========================= CreateBulkOrder ========================= ")
	// Call SDK function bulk
	response, err := service.CreateBulkOrder([]spot.LimitOrderParams{
		{"sell", "0.01", "400000", "BTC_USDT", false, false, "idSome"},
		{"buy", "0.01", "5000", "BTC_USDT", true, true, "someId"},
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
