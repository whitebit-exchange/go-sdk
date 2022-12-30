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

	fmt.Println("========================= Transfer ========================= ")

	// Returns an array of objects containing deposit/withdrawal settings for the corresponding currencies.
	// Zero value in amount fields means that the setting is disabled.
	// ticker	String	Yes	Currency's ticker. Example: BTC
	// amount	Numeric string	Yes	Amount to transfer. Max precision = 8, value should be greater than zero and less or equal your available balance.
	// from	String	No if method is set	Balance FROM which funds will move to. Acceptable values: main, spot, collateral
	// to	String	No if method is set	Balance TO which funds will move to. Acceptable values: main, spot, collateral
	err := service.Transfer(main_account.TransferParams{
		Ticker: "BTC",
		Amount: "0.001",
		From:   main_account.Main,
		To:     main_account.Trade,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

}
