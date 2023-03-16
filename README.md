### A Golang SDK for [whitebit](https://www.whitebit.com)

For best compatibility, please use Go >= 1.18

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Please read [whitebit API document](https://whitebit-exchange.github.io/api-docs/) before continuing.

### API List

- [Private API](https://whitebit-exchange.github.io/api-docs/private/http-trade-v4/)
- [Public API](https://whitebit-exchange.github.io/api-docs/public/http-v4/)
- [Public WebSocket API](https://whitebit-exchange.github.io/api-docs/private/websocket/)
- [Private WebSocket API](https://whitebit-exchange.github.io/api-docs/private/websocket/)

v4 is the preferred one to use

---
## Disclaimer
“You acknowledge that the software is provided “as is”. Author makes no representations or warranties with respect to
the software whether express or implied, including but not limited to, implied warranties of merchantability and fitness
for a particular purpose. author makes no representation or warranty that: (i) the use and distribution of the software
will be uninterrupted or error free, and (ii) any use and distribution of the software is free from infringement of any
third party intellectual property rights. It shall be your sole responsibility to make such determination before the use
of software. Author disclaims any liability in case any such use and distribution infringe any third party’s
intellectual property rights. Author hereby disclaims any warranty and liability whatsoever for any development created
by or for you with respect to your customers. You acknowledge that you have relied on no warranties and that no
warranties are made by author or granted by law whenever it is permitted by law.”
---
### Installation
```shell
go get github.com/whitebit-exchange/go-sdk
```
---
### REST API

#### Setup

Init client for API services. Get APIKey/SecretKey from your whitebit account.

```golang
client := whitebit.NewClient(
    "", //your api key
    "", //your secret key
)
```
Following are some simple examples. 

See the **examples** folder for full references.

#### Create Spot Limit Order

```golang
// Create order/spot service
service := spot.NewService(client)

//Create OrderLimit params
// Call SDK function CreateLimitOrder
response, err := service.CreateLimitOrder(spot.LimitOrderParams{
    Market: "BTC_USDT",
    Amount: "0.001",
    Side:   order.SideBuy,
    Price:  "12000",
})

if err != nil {
    fmt.Println(err.Error())
}

fmt.Printf("%#v\n", response)
```

#### Get Order Info

```golang
// Create a client with your own apiKey and apiSecret
client := whitebit.NewClient(
    "",
    "",
)

// Create assets trade
service := trade.NewService(client)

fmt.Println("========================= GetOrderInfo ========================= ")
// Call SDK function GetOrder

response, err := service.GetOrder(3263845935, 100, 0)

if err != nil {
    fmt.Println(err.Error())
}

fmt.Printf("%#v\n", response)
```

#### Get Futures Markets
You don't need the APIKey and SecretKey to use public API
```golang
// Create a client with your own apiKey and apiSecret
client := whitebit.NewClient(
    "",
    "",
)

// Create a futures service
service := futures.NewService(client)

fmt.Println("========================= GetFuturesMarkets ========================= ")
// Call SDK function GetFuturesMarkets
response, err := service.GetFuturesMarkets()

if err != nil {
    fmt.Println(err.Error())
}

fmt.Printf("%#v\n", response)
```

#### Get Server Time and Ping
You don't need the APIKey and SecretKey to use public API
```golang
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
```
---
### Websocket
Get websocket token from your whitebit account.
```golang
// Create a client with your own apiKey and apiSecret (need for authorize websocket connecting)
client := whitebit.NewClient("publicKey", "secretKey")

// Create new service
service := server.NewService(client)

// Call SDK function GetWsToken - you can use this token all time
token, err := service.GetWsToken()
if err != nil {
	log.Fatal(err)
}
```
Init websocket stream with error handler.
```golang
// Create ctx and cancel for close websocket connect
ctx, cancel := context.WithCancel(context.Background())

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
```
Subscribe example
```golang
// Create handler for processing websocket LastPriceUpdateEvent
lastPriceHandler := func(event stream.LastPriceUpdateEvent) {
	fmt.Printf("%#v\n", event)
}

// Subscribe on market last price events
err = streamService.Subscribe(stream.NewLastPriceSubscription(lastPriceHandler, []string{"BTC_USDT"}))

time.Sleep(time.Second * 10)
// unsubscribe example if you need
err = streamService.Unsubscribe(stream.NewLastPriceUnsubscribe())
if err != nil {
    log.Fatal(err)
}
```
Query example
```golang
streamService.Query(stream.NewLastPriceCommand("BTC_USDT"), func(command stream.Command, response []byte) {
	var result string
	err = json.Unmarshal(response, &result)
	if err != nil {
		print(err.Error())
		return
	}
	fmt.Println(command, result)
})

```

