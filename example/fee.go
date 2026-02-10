//go:build example

package main

import (
	"fmt"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

// This endpoint may require authentication. Provide your own keys if needed.
func main() {
	apiKey := ""
	apiSecret := ""
	s := sdkpkg.New(apiKey, apiSecret)

	res, err := s.Fee.GetTradingFee()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("success=%v, makerFee=%s, takerFee=%s\n", res.Success, res.Result.MakerFee, res.Result.TakerFee)
}
