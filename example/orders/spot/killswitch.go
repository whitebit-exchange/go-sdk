//go:build example

package main

import (
	"fmt"

	spot "github.com/whitebit-exchange/go-sdk/module/order/spot"
	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

func main() {
	apiKey := ""
	apiSecret := ""

	s := sdkpkg.New(apiKey, apiSecret)

	params := spot.KillSwitchStatusParams{}
	res, err := s.OrdersSpot.GetKillSwitchStatus(params)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("kill switch records: %d\n", len(res))
	for i, r := range res {
		if i > 2 {
			break
		}
		fmt.Printf("%d) market=%s, start=%d, cancelAt=%d, types=%v\n", i+1, r.Market, r.StartTime, r.CancellationTime, r.Types)
	}
}
