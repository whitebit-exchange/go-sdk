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

	//Send Ping query with handler for result processing
	streamService.Query(stream.NewPingCommand(), func(command stream.Command, response []byte) {
		var result string
		err = json.Unmarshal(response, &result)
		if err != nil {
			print(err.Error())
			return
		}
		fmt.Println(command, result)
	})

	//Send Time query with handler for result processing
	streamService.Query(stream.NewTimeCommand(), func(command stream.Command, response []byte) {
		var result int64
		err = json.Unmarshal(response, &result)
		if err != nil {
			print(err.Error())
			return
		}
		fmt.Println(command, result)
	})

	gracefulShutdown := make(chan os.Signal, 1)
	// add any other syscalls that you want to be notified with
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-gracefulShutdown
}
