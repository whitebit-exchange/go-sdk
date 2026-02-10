//go:build example

package main

import (
	"fmt"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

func main() {
	s := sdkpkg.New("", "")
	res, err := s.Status.GetMaintenanceStatus()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%d", res.Status)
}
