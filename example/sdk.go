//go:build example

package main

import (
	"fmt"
	"net/http"
	"time"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

func main() {
	// Basic SDK initialization with default settings
	fmt.Println("=== Basic SDK Usage ===")
	sdk := sdkpkg.New("your-api-key", "your-api-secret")

	// Test connection with ping
	ping, err := sdk.Server.Ping()
	if err != nil {
		fmt.Println("ping error:", err)
	} else {
		fmt.Println("Ping:", ping)
	}

	// Get server time
	serverTime, err := sdk.Server.GetTime()
	if err != nil {
		fmt.Println("time error:", err)
	} else {
		fmt.Printf("Server time: %d\n", serverTime.Time)
	}

	// === SDK with custom options ===
	fmt.Println("\n=== SDK with Custom Options ===")

	// SDK with custom base URL
	sdkCustomURL := sdkpkg.New(
		"your-api-key",
		"your-api-secret",
		sdkpkg.WithBaseURL("https://whitebit.com"),
	)
	fmt.Println("SDK with custom base URL created")
	_ = sdkCustomURL

	// SDK with custom timeout
	sdkCustomTimeout := sdkpkg.New(
		"your-api-key",
		"your-api-secret",
		sdkpkg.WithTimeout(30*time.Second),
	)
	fmt.Println("SDK with 30s timeout created")
	_ = sdkCustomTimeout

	// SDK with custom HTTP client
	customHTTPClient := &http.Client{
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 50,
			IdleConnTimeout:     120 * time.Second,
		},
	}
	sdkCustomClient := sdkpkg.New(
		"your-api-key",
		"your-api-secret",
		sdkpkg.WithHTTPClient(customHTTPClient),
	)
	fmt.Println("SDK with custom HTTP client created")
	_ = sdkCustomClient

	// SDK with multiple options combined
	sdkFull := sdkpkg.New(
		"your-api-key",
		"your-api-secret",
		sdkpkg.WithBaseURL("https://whitebit.com"),
		sdkpkg.WithTimeout(45*time.Second),
	)
	fmt.Println("SDK with multiple options created")

	// === Available Services ===
	fmt.Println("\n=== Available SDK Services ===")
	fmt.Println("Public API:")
	fmt.Println("  - sdk.Server    : Server status, time, ping")
	fmt.Println("  - sdk.Market    : Market info")
	fmt.Println("  - sdk.Depth     : Order book depth")
	fmt.Println("  - sdk.Tickers   : Ticker data")
	fmt.Println("  - sdk.Symbols   : Trading pairs")
	fmt.Println("  - sdk.Deals     : Recent trades")
	fmt.Println("  - sdk.Assets    : Asset info")
	fmt.Println("  - sdk.Fee       : Fee info")
	fmt.Println("  - sdk.Futures   : Futures markets")
	fmt.Println("  - sdk.Kline     : Candlestick data")
	fmt.Println("  - sdk.Status    : Maintenance status")

	fmt.Println("\nOrders:")
	fmt.Println("  - sdk.OrdersSpot       : Spot orders")
	fmt.Println("  - sdk.OrdersCollateral : Collateral orders")

	fmt.Println("\nAccount:")
	fmt.Println("  - sdk.AccountTrade      : Trade account")
	fmt.Println("  - sdk.AccountCollateral : Collateral account")
	fmt.Println("  - sdk.AccountMain       : Main account")

	fmt.Println("\nAdditional Services:")
	fmt.Println("  - sdk.Convert      : Currency conversion")
	fmt.Println("  - sdk.SubAccount   : Sub-account management")
	fmt.Println("  - sdk.SmartLending : Smart lending (fixed & flexible)")
	fmt.Println("  - sdk.Mining       : Mining rewards & hashrate")

	_ = sdkFull
}
