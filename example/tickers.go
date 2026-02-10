//go:build example

package main

import (
	"fmt"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

func main() {
	s := sdkpkg.New("", "")
	res, err := s.Tickers.GetTickers()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("success=%v, markets=%d\n", res.Success, len(res.Result))
	for mkt := range res.Result {
		fmt.Println("market:", mkt)
		break
	}
}
