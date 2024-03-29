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

	fmt.Println("========================= GetCollateralBalance ========================= ")
	// Call SDK function GetCollateralBalance

	response, err := service.GetCollateralBalance("USDT")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

	summary, err := service.GetCollateralSummaryBalance("USDT")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", summary)

}
