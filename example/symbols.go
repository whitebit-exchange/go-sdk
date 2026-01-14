//go:build example

package main

import (
	"fmt"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

// Runnable example: public endpoint, keys are not required.
func main() {
	s := sdkpkg.New("", "")

	res, err := s.Symbols.GetSymbols()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("success=%v, symbols=%d\n", res.Success, len(res.Result))
	if len(res.Result) > 0 {
		fmt.Println("first:", res.Result[0])
	}
}
