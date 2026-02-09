//go:build example

package main

import (
	"fmt"
	"time"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

func main() {
	// Create SDK with your API credentials
	s := sdkpkg.New("your-api-key", "your-api-secret")

	// Get mining rewards
	fmt.Println("=== Mining Rewards ===")
	rewards, err := s.Mining.GetRewards(10, 0)
	if err != nil {
		fmt.Println("rewards error:", err)
	} else {
		fmt.Printf("Total rewards: %d\n", rewards.Total)
		for _, reward := range rewards.Records {
			t := time.Unix(reward.Timestamp, 0)
			fmt.Printf("  ID: %d, Amount: %s %s, Time: %s\n",
				reward.ID, reward.Amount, reward.Ticker, t.Format(time.RFC3339))
		}
	}

	// Get hashrate history
	fmt.Println("\n=== Hashrate History ===")
	hashrate, err := s.Mining.GetHashrate(10, 0)
	if err != nil {
		fmt.Println("hashrate error:", err)
	} else {
		fmt.Printf("Total hashrate records: %d\n", hashrate.Total)
		for _, record := range hashrate.Records {
			t := time.Unix(record.Timestamp, 0)
			fmt.Printf("  Hashrate: %s, Time: %s\n",
				record.Hashrate, t.Format(time.RFC3339))
		}
	}
}
