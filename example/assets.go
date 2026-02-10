//go:build example

package main

import (
	"fmt"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

func main() {
	s := sdkpkg.New("", "")
	res, err := s.Assets.GetAssets()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("assets=%d\n", len(res))
	if a, ok := res["BTC"]; ok {
		fmt.Println("BTC asset name:", a.Name)
	}
}
