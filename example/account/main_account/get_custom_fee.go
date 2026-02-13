//go:build example

package main

import (
	"fmt"

	whitebit "github.com/whitebit-exchange/go-sdk"
	mainAccount "github.com/whitebit-exchange/go-sdk/module/account/main_account"
)

func main() {
	// Create a client with your API credentials
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create main account service
	service := mainAccount.NewService(client)

	fmt.Println("========================= GetCustomFee =========================")

	// Get custom fee information for your account
	// This returns your maker/taker fees and any custom fees for specific markets
	customFee, err := service.GetCustomFee()

	if err != nil {
		fmt.Printf("Error getting custom fee: %v\n", err)
		return
	}

	// Print general fees
	fmt.Printf("Account Fee Information:\n")
	fmt.Printf("  Taker Fee: %s%%\n", customFee.Taker)
	fmt.Printf("  Maker Fee: %s%%\n", customFee.Maker)

	// Check for errors in response
	if customFee.Error != nil {
		fmt.Printf("\nAPI Error: %v\n", customFee.Error)
	}

	fmt.Println("\n========================= Complete =========================")
}
