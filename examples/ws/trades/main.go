package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/server"
	"github.com/whitebit-exchange/go-sdk/module/stream"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a client with your own apiKey and apiSecret (need for authorize websocket connecting)
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create new service
	service := server.NewService(client)

	fmt.Println("========================= GetWsToken ========================= ")
	// Call SDK function GetWsToken - you can use this token all time
	token, err := service.GetWsToken()
	if err != nil {
		log.Fatal(err)
	}

	// Create ctx and cancel for close websocket connect
	ctx, cancel := context.WithCancel(context.Background())

	// Create websocket connection,
	streamService, err := stream.NewStream(ctx, token, func(err error) {
		fmt.Println(err)
		switch err.(type) {
		case *net.OpError:
			{
				fmt.Println(err.Error())
				time.Sleep(time.Second)
			}
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	// function what close websocket connection after sleep
	go func() {
		time.Sleep(time.Second * 1000)
		cancel()
	}()

	// Create handler for processing websocket KlineUpdateEvent
	marketTradesHandler := func(event stream.MarketTradesUpdateEvent) {
		fmt.Printf("%#v\n", event)
	}

	// Subscribe on deals events
	err = streamService.Subscribe(stream.NewMarketTradesSubscription(marketTradesHandler, []string{"BTC_USDT"}))
	if err != nil {
		log.Fatal(err)
	}

	streamService.Query(stream.NewMarketTradesCommand("BTC_USDT", 10, 0), func(command stream.Command, response []byte) {
		var result []stream.MarketTradesResult
		err = json.Unmarshal(response, &result)
		if err != nil {
			print(err.Error())
			return
		}
		fmt.Println(command, result)
	})

	time.Sleep(time.Second * 10)
	// unsubscribe example if you need
	err = streamService.Unsubscribe(stream.NewMarketTradesUnsubscribe())
	if err != nil {
		log.Fatal(err)
	}

	gracefulShutdown := make(chan os.Signal, 1)
	// add any other syscalls that you want to be notified with
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-gracefulShutdown
}
