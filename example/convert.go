//go:build example

package main

import (
	"fmt"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

func main() {
	// Create SDK with your API credentials
	s := sdkpkg.New("your-api-key", "your-api-secret")

	// Get conversion estimate
	fmt.Println("=== Conversion Estimate ===")
	estimate, err := s.Convert.Estimate("BTC", "USDT", "0.1")
	if err != nil {
		fmt.Println("estimate error:", err)
	} else {
		fmt.Printf("Query ID: %s\n", estimate.QueryID)
		fmt.Printf("Give: %s %s\n", estimate.GiveAmount, estimate.From)
		fmt.Printf("Get: %s %s\n", estimate.GetAmount, estimate.To)
		fmt.Printf("Rate: %s\n", estimate.Rate)
		fmt.Printf("Fee: %s\n", estimate.Fee)
	}

	// Confirm conversion (using the query ID from estimate)
	fmt.Println("\n=== Confirm Conversion ===")
	if estimate.QueryID != "" {
		confirmation, err := s.Convert.Confirm(estimate.QueryID)
		if err != nil {
			fmt.Println("confirm error:", err)
		} else {
			fmt.Printf("Conversion ID: %d\n", confirmation.ID)
			fmt.Printf("Status: %d\n", confirmation.Status)
			fmt.Printf("From: %s -> To: %s\n", confirmation.From, confirmation.To)
		}
	}

	// Get conversion history
	fmt.Println("\n=== Conversion History ===")
	history, err := s.Convert.GetHistory(10, 0)
	if err != nil {
		fmt.Println("history error:", err)
	} else {
		fmt.Printf("Total records: %d\n", history.Total)
		for _, record := range history.Records {
			fmt.Printf("  ID: %d, %s %s -> %s %s, Status: %d\n",
				record.ID, record.GiveAmount, record.From,
				record.GetAmount, record.To, record.Status)
		}
	}
}
