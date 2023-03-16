package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/account/main_account"
)

func main() {
	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create main account service
	service := main_account.NewService(client)

	fmt.Println("========================= GetFee ========================= ")

	// Returns an object containing your trading fee for market.
	// Zero value in amount fields means that the setting is disabled.
	response, err := service.GetMyFeeByMarket("BTC_USDT")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

	// Returns an object containing your personal trading fee and markets with custom fee if there are any
	// Zero value in amount fields means that the setting is disabled.
	myFee, err := service.GetCustomFee()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", myFee)

}
