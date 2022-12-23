package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/account/collateral"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create collateral service
	service := collateral.NewService(client)

	fmt.Println("========================= GetPositionsHistory ========================= ")
	// Call SDK function GetPositionsHistory

	response, err := service.GetPositionsHistory(collateral.PositionsHistoryOptions{Market: "BTC_USDT", Limit: "2"})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
