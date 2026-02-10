//go:build example

package main

import (
	"fmt"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

func main() {
	s := sdkpkg.New("", "")
	market := "BTC_USDT"
	res, err := s.Depth.GetDepth(market)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("success=%v, asks=%d, bids=%d\n", res.Success, len(res.Result.Asks), len(res.Result.Bids))
}
