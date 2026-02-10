//go:build example

package main

import (
	"fmt"

	"github.com/whitebit-exchange/go-sdk/module/kline"
	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

func main() {
	s := sdkpkg.New("", "")
	opts := kline.Options{Market: "BTC_USDT", Interval: "1m", Limit: 5}
	res, err := s.Kline.GetKline(opts)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("success=%v, klines=%d\n", res.Success, len(res.Result))
}
