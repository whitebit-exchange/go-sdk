package main

import (
	"fmt"
	"github.com/whitebit-exchange/go-sdk"
	"github.com/whitebit-exchange/go-sdk/module/account/main_account"
)

func main() {
	// Create a client with your own apiKey and apiSecret
	client := whitebit.NewClient(
		"",
		"",
	)

	// Create main account service
	service := main_account.NewService(client)

	fmt.Println("========================= GetHistory ========================= ")

	// This function retrieves the history of deposits and withdraws
	//Name		Type	Mandatory	Description
	//trMethod	Number	No	Method. Example: 1 to display deposits / 2 to display withdraws. Do not send
	//								this parameter in order to receive both deposits and withdraws.
	//ticker	String	No	Currency's ticker. Example: BTC
	//address	String	No	Can be used for filtering transactions by specific address or memo.
	//addresses	Array	No	Can be used for filtering transactions by specific addresses or memos (max: 20).
	//uniqueId	String	No	Can be used for filtering transactions by specific unique id
	//limit	Int	Yes	LIMIT is a special clause used to limit records a particular query can return. Default: 50, Min: 1, Max: 100

	//Deposit status codes:
	//	Successful - 3, 7
	//	Canceled - 4, 9
	//	Unconfirmed by user - 5
	//	Uncredited - 22
	//	Pending - 15
	//Withdraw status codes:
	//	Pending - 1, 2, 6, 10, 11, 12, 13, 14, 15, 16, 17
	//	Successful - 3, 7
	//	Canceled - 4
	//	Unconfirmed by user - 5
	//	Partially successful - 18

	response, err := service.GetHistory(main_account.HistoryParams{
		Status: []int{3},
		Offset: 0,
		Limit:  0,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", response)

}
