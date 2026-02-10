//go:build example

package main

import (
	"fmt"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

func main() {
	s := sdkpkg.New("", "")
	res, err := s.Market.GetMarkets()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("success=%v, markets=%d\n", res.Success, len(res.Result))
	if len(res.Result) > 0 {
		fmt.Println("first:", res.Result[0].Name)
	}
}
