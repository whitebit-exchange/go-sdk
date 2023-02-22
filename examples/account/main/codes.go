package main

import (
	"fmt"
	whitebit "github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/account/main_account"
)

func main() {
	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create account main service
	service := main_account.NewService(client)

	fmt.Println("========================= Create Whitebit code ========================= ")

	// Call SDK function
	//Name			Type		Mandatory	Description
	//ticker		String		Yes			Currency's ticker. Example: BTC
	//amount	Numeric string	Yes			Amount to transfer. Max precision = 8, value should be greater than zero and your main balance.
	//passphrase	String		No			Passphrase that will be used for applying codes. Passphrase must contain only latin letters,
	//										numbers and symbols (like !@#$%^, no whitespaces). Max: 25 symbols.
	//description	String		No			Additional text description for code. Visible only for creator Max: 75 symbols.

	code, err := service.CreateCode("USDT", "10", "", "test")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", code)

	fmt.Println("========================= Apply Whitebit code ========================= ")

	//Name			Type	  Mandatory		Description
	//code			String		Yes			Code that will be applied.
	//passphrase	String		No			Should be provided if the code was created with passphrase Max: 25 symbols.

	applyCode, err := service.ApplyCode("WB52d23850-b98a-48d5-8956-2159f5bf8966USDT", "")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", applyCode)

	fmt.Println("========================= My Whitebit codes ========================= ")

	//Name	 Type	Mandatory	Description
	//limit	 Int	No		LIMIT is a special clause used to limit records a particular query can return.
	//						Default: 30, Min: 1, Max: 100
	//offset Int	No		If you want the request to return entries starting from a particular line,
	//						you can use OFFSET clause to tell it where it should start. Default: 0, Min: 0, Max: 10000

	myCodes, err := service.GetCodes(0, 0)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", myCodes)

	fmt.Println("========================= Whitebit codes history ========================= ")

	//Name	 Type	Mandatory	Description
	//limit	 Int	No		LIMIT is a special clause used to limit records a particular query can return.
	//						Default: 30, Min: 1, Max: 100
	//offset Int	No		If you want the request to return entries starting from a particular line,
	//						you can use OFFSET clause to tell it where it should start. Default: 0, Min: 0, Max: 10000

	historyCodes, err := service.GetCodesHistory(0, 0)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", historyCodes)
}
