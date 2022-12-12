package main

import (
	"fmt"
	"github.com/whitebit-exchange/whitebit"
	"github.com/whitebit-exchange/whitebit/module/server"
)

func main() {

	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	fmt.Println("========================= Ping ========================= ")
	// Create a server service
	service := server.NewService(client)

	// Call SDK function Ping
	response, err := service.Ping()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

	fmt.Println("========================= GetTime ========================= ")
	// Call SDK function GetTime
	resp, err := service.GetTime()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", resp)
}
