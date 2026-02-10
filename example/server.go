//go:build example

package main

import (
	"fmt"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

func main() {
	s := sdkpkg.New("", "")

	// Ping
	ping, err := s.Server.Ping()
	if err != nil {
		fmt.Println("ping error:", err)
	} else {
		fmt.Println("ping:", ping)
	}

	// Time
	tm, err := s.Server.GetTime()
	if err != nil {
		fmt.Println("time error:", err)
	} else {
		fmt.Println("server time:", tm.Time)
	}
}
