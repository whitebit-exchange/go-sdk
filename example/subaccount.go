//go:build example

package main

import (
	"fmt"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

func main() {
	// Create SDK with your API credentials
	s := sdkpkg.New("your-api-key", "your-api-secret")

	// List sub-accounts
	fmt.Println("=== List Sub-Accounts ===")
	list, err := s.SubAccount.List(10, 0)
	if err != nil {
		fmt.Println("list error:", err)
	} else {
		fmt.Printf("Total sub-accounts: %d\n", list.Total)
		for _, acc := range list.Records {
			fmt.Printf("  ID: %s, Alias: %s, Email: %s, Status: %s\n",
				acc.ID, acc.Alias, acc.Email, acc.Status)
		}
	}

	// Create a new sub-account
	fmt.Println("\n=== Create Sub-Account ===")
	newAcc, err := s.SubAccount.Create("my-sub-account", "sub@example.com")
	if err != nil {
		fmt.Println("create error:", err)
	} else {
		fmt.Printf("Created sub-account ID: %s, Alias: %s\n", newAcc.ID, newAcc.Alias)
	}

	// Get balances for a sub-account
	fmt.Println("\n=== Sub-Account Balances ===")
	if newAcc.ID != "" {
		balances, err := s.SubAccount.GetBalances(newAcc.ID, "")
		if err != nil {
			fmt.Println("balances error:", err)
		} else {
			for ticker, balance := range balances {
				fmt.Printf("  %s: %s\n", ticker, balance.MainBalance)
			}
		}
	}

	// Transfer funds between sub-accounts
	fmt.Println("\n=== Transfer Between Sub-Accounts ===")
	err = s.SubAccount.Transfer("", newAcc.ID, "USDT", "10")
	if err != nil {
		fmt.Println("transfer error:", err)
	} else {
		fmt.Println("Transfer successful")
	}

	// Get transfer history
	fmt.Println("\n=== Transfer History ===")
	transferHistory, err := s.SubAccount.GetTransferHistory("", 10, 0)
	if err != nil {
		fmt.Println("transfer history error:", err)
	} else {
		fmt.Printf("Total transfers: %d\n", transferHistory.Total)
		for _, record := range transferHistory.Records {
			fmt.Printf("  From: %s, To: %s, %s %s\n",
				record.From, record.To, record.Amount, record.Ticker)
		}
	}

	// Create API key for sub-account
	fmt.Println("\n=== Create API Key ===")
	if newAcc.ID != "" {
		apiKey, err := s.SubAccount.CreateAPIKey(newAcc.ID, "trading-key", []string{"read", "trade"})
		if err != nil {
			fmt.Println("create API key error:", err)
		} else {
			fmt.Printf("API Key: %s\n", apiKey.APIKey)
			fmt.Printf("Secret: %s\n", apiKey.SecretKey)
			fmt.Printf("Permissions: %v\n", apiKey.Permissions)
		}
	}

	// List API keys
	fmt.Println("\n=== List API Keys ===")
	if newAcc.ID != "" {
		keys, err := s.SubAccount.ListAPIKeys(newAcc.ID, 10, 0)
		if err != nil {
			fmt.Println("list API keys error:", err)
		} else {
			fmt.Printf("Total API keys: %d\n", keys.Total)
			for _, key := range keys.Records {
				fmt.Printf("  ID: %s, Label: %s, Permissions: %v\n",
					key.ID, key.Label, key.Permissions)
			}
		}
	}
}
