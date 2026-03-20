# Reference
<details><summary><code>client.ConvertEstimate(request) -> *gosdk.ConvertEstimateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a quote for converting one currency to another. Quote lifetime is 10 seconds, then quote will be expired.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ConvertEstimateRequest{
        From: "BTC",
        To: "USDT",
        Direction: gosdk.ConvertEstimateRequestDirectionTo,
        Amount: "35,103.1",
    }
client.ConvertEstimate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `string` — From currency. Example: BTC
    
</dd>
</dl>

<dl>
<dd>

**to:** `string` — To currency. Example: USDT
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*gosdk.ConvertEstimateRequestDirection` — Convert amount direction, defines in which currency corresponding "amount" field is populated. Use "to" in case amount is in "to" currency, use "from" if amount is in "from" currency
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Amount to convert or receive.
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*int` — Nonce for request
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` — Request path
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConvertConfirm(request) -> *gosdk.ConvertConfirmResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint confirms an estimated quote.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ConvertConfirmRequest{
        QuoteID: "4050",
    }
client.ConvertConfirm(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**quoteID:** `string` — Quote ID
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*int` — Nonce for request
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` — Request path
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConvertHistory(request) -> *gosdk.ConvertHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns convert history.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>
**Note:** The endpoint can retrieve data not older than 6 months from current month. For older data, use the Report on the History page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ConvertHistoryRequest{
        FromTicker: gosdk.String(
            "BTC",
        ),
    }
client.ConvertHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fromTicker:** `*string` — From currency. Example: BTC
    
</dd>
</dl>

<dl>
<dd>

**toTicker:** `*string` — To currency. Example: USDT
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — From time filter. Example: 1699260637. Default: now()
    
</dd>
</dl>

<dl>
<dd>

**to:** `*string` — To time filter. Example: 1699260637. Default: now() +
    
</dd>
</dl>

<dl>
<dd>

**quoteID:** `*string` — Quote Id. Example: 4050
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*string` — How many records to receive. Default: 100
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*string` — Amount to convert or receive. Default 0
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*int` — Nonce for request
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` — Request path
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Authentication
<details><summary><code>client.Authentication.OAuth20Authorization() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint initiates the OAuth 2.0 authorization flow for user authentication and obtaining an authorization code.

**Using the State Parameter (Best Practice)**

The `state` parameter is crucial for security in OAuth flows:

- Generate a cryptographically secure random string
- Store it in the session before redirecting
- Validate it matches when handling the callback
- This prevents CSRF attacks
<Note>
**Note:** OAuth scopes are predefined during client application setup and cannot be modified during the authorization request. The access token will include all scopes that were approved during client creation.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetAuthLoginRequest{
        ClientID: "YOUR_CLIENT_ID",
        State: gosdk.String(
            "SECURE_RANDOM_STATE",
        ),
    }
client.Authentication.OAuth20Authorization(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` — The application's client ID
    
</dd>
</dl>

<dl>
<dd>

**state:** `*string` — A secure random string used to maintain state between the request and callback and prevent CSRF attacks (Recommended)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Authentication.GetAccessToken(request) -> *gosdk.PostOauth2TokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint activates an access token by exchanging an authorization code.

<Warning>
**Important Notes:**

- Access token duration is 300 seconds
- The IP of the client must be added to WB Allowlist
</Warning>

**Request Headers:**
- Content-Type: application/x-www-form-urlencoded
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.PostOauth2TokenRequest{
        ClientID: "YOUR_CLIENT_ID",
        ClientSecret: "YOUR_CLIENT_SECRET",
        Code: "AUTHORIZATION_CODE",
    }
client.Authentication.GetAccessToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` — The application's client ID
    
</dd>
</dl>

<dl>
<dd>

**clientSecret:** `string` — The application's client secret
    
</dd>
</dl>

<dl>
<dd>

**code:** `string` — The authorization code received from the authorization endpoint
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Authentication.RefreshToken(request) -> *gosdk.PostOauth2RefreshTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint creates a new access token using a refresh token.

**Request Headers:**
- Content-Type: application/x-www-form-urlencoded

<Warning>
**Important Notes:**

- Refresh token duration is 600 seconds
- Rate limit: 1 request per second
- The IP of the client must be added to WB Allowlist
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.PostOauth2RefreshTokenRequest{
        ClientID: "YOUR_CLIENT_ID",
        ClientSecret: "YOUR_CLIENT_SECRET",
        Token: "REFRESH_TOKEN",
    }
client.Authentication.RefreshToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` — The application's client ID
    
</dd>
</dl>

<dl>
<dd>

**clientSecret:** `string` — The application's client secret
    
</dd>
</dl>

<dl>
<dd>

**token:** `string` — The refresh token received from the token endpoint
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AccountEndpoints
<details><summary><code>client.AccountEndpoints.GetAccountTransactions(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint retrieves a paginated list of account transactions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]any{
        "key": "value",
    }
client.AccountEndpoints.GetAccountTransactions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AccountEndpoints.GetCurrencyConversions(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint retrieves the history of currency conversions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]any{
        "key": "value",
    }
client.AccountEndpoints.GetCurrencyConversions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AccountEndpoints.GetOrdersHistory(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint retrieves the history of trading orders.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]any{
        "key": "value",
    }
client.AccountEndpoints.GetOrdersHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AccountEndpoints.GetExecutedDeals(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint retrieves the history of executed deals.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]any{
        "key": "value",
    }
client.AccountEndpoints.GetExecutedDeals(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AccountEndpoints.GetMainAccountBalance(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint retrieves the main account balance information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]any{
        "key": "value",
    }
client.AccountEndpoints.GetMainAccountBalance(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AccountEndpoints.GetSpotAccountBalance(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint retrieves the spot trading account balance information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]any{
        "key": "value",
    }
client.AccountEndpoints.GetSpotAccountBalance(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## PublicAPIV4
<details><summary><code>client.PublicAPIV4.MaintenanceStatus() -> *gosdk.GetAPIV4PublicPlatformStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves maintenance status
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.PublicAPIV4.MaintenanceStatus(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PublicAPIV4.MarketInfo() -> []*gosdk.GetAPIV4PublicMarketsResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves all information about available spot and futures markets.

<Note>
The API caches the response for 1 second
</Note>

<Warning>
Rate limit 2000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.PublicAPIV4.MarketInfo(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PublicAPIV4.MarketActivity() -> map[string]*gosdk.GetAPIV4PublicTickerResponseValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves a 24-hour pricing and volume summary for each market pair available on the exchange.

<Note>
The API caches the response for 1 second
</Note>

<Warning>
Rate limit: 2000 requests/10 sec. See [Public API V4 overview](/public/http-v4/index) for rate limit details.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.PublicAPIV4.MarketActivity(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PublicAPIV4.AssetStatusList() -> map[string]*gosdk.Asset</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the assets status.

<Note>
The API caches the response for 1 second
</Note>

<Warning>
Rate limit: 2000 requests/10 sec. See [Public API V4 overview](/public/http-v4/index) for rate limit details.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.PublicAPIV4.AssetStatusList(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PublicAPIV4.Orderbook(Market) -> *gosdk.OrderbookResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the current [order book](/glossary#order-book) as two arrays ([bids](/glossary#bid) / [asks](/glossary#ask)) with additional parameters.

<Note>
The API caches the response for 100 ms
</Note>

<Warning>
Rate limit 600 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetAPIV4PublicOrderbookMarketRequest{
        Market: "BTC_USDT",
        Limit: gosdk.Int(
            100,
        ),
        Level: gosdk.Int(
            2,
        ),
    }
client.PublicAPIV4.Orderbook(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` — Market pair name
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Orders depth quantity: 0 - 100. Not defined or 0 will return 100 entries.
    
</dd>
</dl>

<dl>
<dd>

**level:** `*int` — Optional parameter that controls the aggregation level. Level 0 – default, no aggregation. Levels 1–5 provide increasing aggregation of the order book.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PublicAPIV4.Depth(Market) -> *gosdk.OrderbookResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves depth price levels within ±2% of the market last price. Use when lightweight order book data is needed for a narrow price band around the current market price. The ±2% constraint limits the response to price levels near the last traded price, reducing payload size compared to the full order book.

<Note>
The API caches the response for 1 sec
</Note>

<Warning>
Rate limit: 2000 requests/10 sec. See [Public API V4 overview](/public/http-v4/index) for rate limit details.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetAPIV4PublicOrderbookDepthMarketRequest{
        Market: "BTC_USDT",
    }
client.PublicAPIV4.Depth(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` — Market pair name
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PublicAPIV4.RecentTrades(Market) -> []*gosdk.GetAPIV4PublicTradesMarketResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the [trades](/glossary#deal-trade) that have been executed recently on the requested [market](/glossary#market).

<Note>
The API caches the response for 1 second
</Note>

<Note>
  Public trade data can include executions that originate from RPI orders. Public order book feeds (`depth`, `bookTicker`) exclude RPI orders. RPI orders appear only in private active orders and in the exchange UI order book (web and mobile).
</Note>


<Warning>
Rate limit 2000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetAPIV4PublicTradesMarketRequest{
        Market: "BTC_USDT",
    }
client.PublicAPIV4.RecentTrades(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` — Market pair name
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*gosdk.GetAPIV4PublicTradesMarketRequestType` — Can be buy or sell
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PublicAPIV4.Fee() -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the list of [fees](/glossary#fee) and min/max amounts for deposits and withdrawals

<Note>
The API caches the response for 1 second
</Note>

<Warning>
Rate limit: 2000 requests/10 sec. See [Public API V4 overview](/public/http-v4/index) for rate limit details.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.PublicAPIV4.Fee(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PublicAPIV4.ServerTime() -> *gosdk.GetAPIV4PublicTimeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the current server time.

<Note>
The API caches the response for 1 second
</Note>

<Warning>
Rate limit 2000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.PublicAPIV4.ServerTime(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PublicAPIV4.ServerStatus() -> []string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the current API life-state.

<Note>
The API caches the response for 1 second
</Note>

<Warning>
Rate limit 2000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.PublicAPIV4.ServerStatus(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PublicAPIV4.CollateralMarketsList() -> *gosdk.GetAPIV4PublicCollateralMarketsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns the list of [markets](/glossary#market) that are available for [collateral](/glossary#collateral) trading

<Note>
The API caches the response for 1 second
</Note>

<Warning>
Rate limit 2000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.PublicAPIV4.CollateralMarketsList(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PublicAPIV4.AvailableFuturesMarketsList() -> *gosdk.GetAPIV4PublicFuturesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns the list of available futures markets.

<Note>
The API caches the response for 1 second
</Note>

<Warning>
Rate limit 2000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.PublicAPIV4.AvailableFuturesMarketsList(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PublicAPIV4.FundingHistory(Market) -> []*gosdk.GetAPIV4PublicFundingHistoryMarketResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns the funding rate history for a specified futures market.

<Warning>
Rate limit 2000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetAPIV4PublicFundingHistoryMarketRequest{
        Market: "BTC_PERP",
        StartDate: gosdk.Int(
            1752480000,
        ),
        EndDate: gosdk.Int(
            1752537600,
        ),
        Limit: gosdk.Int(
            100,
        ),
        Offset: gosdk.Int(
            0,
        ),
    }
client.PublicAPIV4.FundingHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` — Market name (e.g., BTC_PERP)
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*int` — Start timestamp in seconds
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*int` — End timestamp in seconds
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of records to return. Default: 100, Maximum: 1000
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Number of records to skip
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PublicAPIV4.MiningPoolOverview() -> *gosdk.GetAPIV4PublicMiningPoolResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns overall information about the current mining pool state.

Hash rate is expressed in H units.

<Warning>
Rate limit 1000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.PublicAPIV4.MiningPoolOverview(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Main Account
<details><summary><code>client.MainAccount.GetMainBalance(request) -> map[string]*gosdk.GetMainBalanceResponseValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the [main balance](/glossary#balance-main) by currency [ticker](/glossary#ticker) or all balances.

<Warning>
  Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetMainBalanceRequest{
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.MainAccount.GetMainBalance(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ticker:** `*string` — Currency's ticker.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MainAccount.GetDepositWithdrawHistory(request) -> *gosdk.GetDepositWithdrawHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the history of deposits and withdraws.

**Deposit status codes:**
- `Successful` - 3, 7
- `Canceled` - 4, 9
- `Unconfirmed by user` - 5
- `Additional data required` - 21
- `Uncredited` - 22
- `Pending` - 15

**Travel Rule Deposit check status codes:**
- `Awaiting verification` - 27: The transaction has been frozen due to the lack of data required under the Travel Rule. The user is required to provide this data manually through the exchange interface.
- `Confirmation in progress` - 28: The Travel Rule data provided by the user is currently being verified by WhiteBIT.

⚠️ Due to regulatory requirements in Turkey and [EU](/glossary#european-economic-area-eea), the system places every inbound crypto deposit on hold (frozen) until confirming the transaction's origin. The sender must provide certain details if the transaction is from another Virtual Asset Service Provider (VASP) or verify the address if from a self-hosted wallet. The system credits deposited funds to the account only after successful verification.

**Withdraw status codes:**
- `Pending` - 1, 2, 6, 10, 11, 12, 13, 14, 15, 16, 17
- `Successful` - 3, 7
- `Canceled` - 4
- `Unconfirmed by user` - 5
- `Additional data required` - 21
- `Partially successful` - 18

<Warning>
Rate limit: 200 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetDepositWithdrawHistoryRequest{
        TransactionMethod: gosdk.Int(
            1,
        ),
        Ticker: gosdk.String(
            "BTC",
        ),
        Limit: gosdk.Int(
            100,
        ),
        Offset: gosdk.Int(
            0,
        ),
        Status: []int{
            3,
            7,
        },
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.MainAccount.GetDepositWithdrawHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**transactionMethod:** `*int` — Method. Example: **1** to display deposits / **2** to display withdraws. Do not send this parameter in order to receive both deposits and withdraws.
    
</dd>
</dl>

<dl>
<dd>

**ticker:** `*string` — Currency's [ticker](/glossary#ticker). Example: BTC
    
</dd>
</dl>

<dl>
<dd>

**address:** `*string` — Can be used for filtering transactions by specific address.
    
</dd>
</dl>

<dl>
<dd>

**memo:** `*string` — Can be used for filtering transactions by specific [memo](/glossary#memodestination-tag)
    
</dd>
</dl>

<dl>
<dd>

**addresses:** `[]string` — Can be used for filtering transactions by specific array of addresses.
    
</dd>
</dl>

<dl>
<dd>

**uniqueID:** `*string` — Can be used for filtering transactions by specific unique id
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — LIMIT is a special clause used to limit records a particular query can return.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Use the OFFSET clause to return entries starting from a particular line.
    
</dd>
</dl>

<dl>
<dd>

**status:** `[]int` 

Can be used for filtering transactions by status codes.

⚠️ Caution: Use this parameter with the appropriate `transactionMethod` and valid status codes for that method. See the endpoint description above for valid codes. Example: `"status": [3,7]`
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Deposit
<details><summary><code>client.Deposit.GetDepositAddress(request) -> *gosdk.GetDepositAddressResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves a deposit address of the cryptocurrency.

<Accordion title="Errors">
```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "ticker": ["The selected ticker is invalid."]
  }
}
```

```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "network": ["The selected network is invalid."]
  }
}
```

```json
{
  "code": 1,
  "message": "Inner validation failed",
  "errors": {
    "ticker": ["Currency is not depositable"]
  }
}
```
</Accordion>

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetDepositAddressRequest{
        Ticker: "BTC",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.Deposit.GetDepositAddress(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ticker:** `string` — Currencies ticker. Example: BTC ⚠️ Currency [ticker](/glossary#ticker) should not be [fiat](/glossary#fiat) and it’s “can_deposit” status must be “true”. See [Asset Status endpoint](/public/http-v4/asset-status-list) response for the status.
    
</dd>
</dl>

<dl>
<dd>

**network:** `*string` — Cryptocurrency network. ⚠️ If currency has multiple networks like USDT, specify the network to use. See [ticker](/glossary#ticker) networks list in “networks” field from response [Asset Status endpoint](/public/http-v4/asset-status-list).
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Deposit.GetFiatDepositURL(request) -> *gosdk.GetFiatDepositURLResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves a deposit url of the [fiat](/glossary#fiat) invoice.

The endpoint works on demand. Contact WhiteBIT support and provide the API key to get access to the functionality.

<Accordion title="Errors">
```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "amount": ["Amount is too little for deposit"]
  }
}
```

```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "provider": ["Cannot find currency for specified provider"]
  }
}
```

```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "uniqueId": ["The unique id has already been taken."]
  }
}
```

```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "amount": ["The amount must be a number."],
    "provider": ["The selected provider is invalid."],
    "ticker": ["The selected ticker is invalid."]
  }
}
```

```json
{
  "code": 10,
  "message": "Failed to generate deposit url"
}
```

```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "amount": ["The amount field is required."],
    "provider": ["The provider field is required."],
    "ticker": ["The ticker field is required."],
    "uniqueId": ["The unique id field is required."]
  }
}
```

```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "successLink": [
      "Your domain is incorrect. Please contact support for more details"
    ],
    "failureLink": [
      "Your domain is incorrect. Please contact support for more details"
    ]
  }
}
```

```json
{
  "success": false,
  "message": "You don't have permission to use this endpoint. Please contact support for more details",
  "code": 0
}
```

```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "successLink": ["Uri domain must have only https scheme"],
    "failureLink": ["Uri domain must have only https scheme"]
  }
}
```
</Accordion>

<Warning>
When using VISAMASTER as [provider](/glossary#provider), pass the [referer header](https://developer.mozilla.org/ru/docs/Web/HTTP/Headers/Referer) when opening the invoice link (e.g. when opening https://someaddress.com). If the header is missing (e.g. when opening the link from a Telegram message), the browser is redirected to the WhiteBIT homepage
</Warning>

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetFiatDepositURLRequest{
        Ticker: "UAH",
        Provider: "VISAMASTER",
        Amount: "100",
        UniqueID: gosdk.String(
            "{{generateID}}",
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.Deposit.GetFiatDepositURL(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ticker:** `string` — Currency's [ticker](/glossary#ticker) ([fiat](/glossary#fiat)). ⚠️ Currencies ticker should be fiat and has "can_deposit" status must be "true". Use [Asset Status endpoint](/public/http-v4/asset-status-list) to know more about currency.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `string` — [Fiat](/glossary#fiat) currency [provider](/glossary#provider). ⚠️ Currency provider should be taken from [Asset Status endpoint](/public/http-v4/asset-status-list) response.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Deposit amount
    
</dd>
</dl>

<dl>
<dd>

**uniqueID:** `*string` — Unique transaction identifier on client's side
    
</dd>
</dl>

<dl>
<dd>

**customer:** `*gosdk.GetFiatDepositURLRequestCustomer` — Customer information (required for USD/EUR with VISAMASTER [provider](/glossary#provider))
    
</dd>
</dl>

<dl>
<dd>

**successLink:** `*string` — Customer will be redirected to this URL by acquiring [provider](/glossary#provider) after success deposit. To activate this feature, please contact support
    
</dd>
</dl>

<dl>
<dd>

**failureLink:** `*string` — Customer will be redirected to this URL in case of fail or rejection on acquiring provider side. To activate this feature, please contact support
    
</dd>
</dl>

<dl>
<dd>

**returnLink:** `*string` — Customer will be redirected to the URL defined if selects 'back' option after from the payment success or failure page. To activate this feature, define desired link. If not populated, option 'back' won't be displayed
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Deposit.IssueCardToken(request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint issues a [card token](/glossary#card-token) for the specified card number.

**Request Headers:**
- `Authorization`: JWT token from the [Issue JWT token](/private/http-main-v4/issue-jwt-token) response field `data.token` (**Required**)
- `Content-Type`: `application/json` (**Required**)

Use the `data.token` value returned by [Issue JWT token](/private/http-main-v4/issue-jwt-token) as the `Authorization` header when calling this Fiat Gateway endpoint.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.IssueCardTokenRequest{
        CardNumber: "4111111111111111",
    }
client.Deposit.IssueCardToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cardNumber:** `string` — Target card number
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Deposit.RefundDeposit(request) -> *gosdk.RefundDepositResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Refund a deposit.

The system processes refunds only for deposits with status `canceled`.

Obtain `transactionId` from the `deposit.canceled` webhook (`uniqueId`) or from the deposit/withdraw history in the WhiteBIT interface after support confirms refund availability.

The refund address must support the same network and asset as the original deposit. The refund address must not be a WhiteBIT address. The refund address can differ from the original deposit sender address.

<Accordion title="Errors">
```json
{
  "code": 0,
  "message": "Transaction not refundable"
}
```

```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "transactionId": ["The transaction id field is required."],
    "address": ["The address field is required."]
  }
}
```
</Accordion>

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>

<Note>
Refund processing does not complete instantly. Use the [refund.successful](/platform/webhook#whitebit-refund-successful) and [refund.failed](/platform/webhook#whitebit-refund-failed) webhooks to receive refund status updates.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.RefundDepositRequest{
        TransactionID: gosdk.String(
            "54bffeb7-7a8f-43f8-bcd8-f14ec10fee85",
        ),
        Address: "0x1234567890abcdef1234567890abcdef12345678",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.Deposit.RefundDeposit(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**transactionID:** `*string` — Transaction UUID of the deposit. Obtain from the [deposit.canceled](/platform/webhook) webhook (`uniqueId` field) or from the deposit/withdraw history in the WhiteBIT interface.
    
</dd>
</dl>

<dl>
<dd>

**address:** `string` — Destination wallet address for the refund. The address must support the same network and asset as the original deposit. Cannot be a WhiteBIT address. Does not have to match the original deposit address.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Base64-encoded request body. See the [authentication guide](/private/http-auth) for signature generation details.
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — A unique identifier for the request. Use a monotonically increasing value such as a Unix timestamp in milliseconds.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Deposit.CreateNewAddress(request) -> *gosdk.CreateNewAddressResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a new address even when the last created address is not used. The endpoint is not available by default, contact support@whitebit.com to get permissions to use the endpoint.

**Address types:**

| Currency | Types               | Default |
|----------|---------------------|---------|
| BTC      | p2sh-segwit, bech32 | bech32  |
| LTC      | p2sh-segwit, bech32 | bech32  |

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateNewAddressRequest{
        Ticker: "XLM",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.Deposit.CreateNewAddress(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ticker:** `string` — Currency's ticker.
    
</dd>
</dl>

<dl>
<dd>

**network:** `*string` — Currency's network (for multinetwork currencies). Example: OMNI or TRC20 or ERC20. For USDT default network is ERC20(ETH).
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*gosdk.CreateNewAddressRequestType` — Address type, available for specific currencies list (see address types table in endpoint description)
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## JWT
<details><summary><code>client.Jwt.IssueJwtToken(request) -> *gosdk.IssueJwtTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint issues a JWT token for the Fiat Gateway service.
The token is used to authenticate requests to the Fiat Gateway API.

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.IssueJwtTokenRequest{
        Request: "{{request}}",
        NonceWindow: gosdk.Bool(
            false,
        ),
        Nonce: "{{nonce}}",
    }
client.Jwt.IssueJwtToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonceWindow:** `*bool` — Nonce window setting
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jwt.GetWebSocketToken(request) -> *gosdk.GetWebSocketTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The V4 endpoint can be used to retrieve the WebSocket token for user.
The token is required to authorize WebSocket connections for private API access.

<Warning>
Rate limit: 10 requests/60 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetWebSocketTokenRequest{
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.Jwt.GetWebSocketToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Withdraw
<details><summary><code>client.Withdraw.CreateWithdraw(request) -> []any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates withdraw for the specified ticker.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>

<Note>
Also, fiat currencies can't be withdrawn without KYC verification.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateWithdrawRequest{
        Ticker: "ETH",
        Amount: "0.9",
        Address: "0x0964A6B8F794A4B8d61b62652dB27ddC9844FB4c",
        UniqueID: gosdk.String(
            "24529041",
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.Withdraw.CreateWithdraw(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ticker:** `string` 

Currency's [ticker](/glossary#ticker). Example: BTC

⚠️ Currencies ticker must have "can_deposit" status equal to "true". Use [Asset Status endpoint](/public/http-v4/asset-status-list) to know more about currency.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Withdraw amount (including [fee](/glossary#fee)). To add the fee to the specified amount, use the /main-account/withdraw-pay request.
    
</dd>
</dl>

<dl>
<dd>

**address:** `string` — Target address (wallet address for cryptocurrencies, identifier/[card token](/glossary#card-token) for [fiat](/glossary#fiat) currencies)
    
</dd>
</dl>

<dl>
<dd>

**memo:** `*string` 

[Memo](/glossary#memodestination-tag).

⚠️ Required if currency is memoable.
    
</dd>
</dl>

<dl>
<dd>

**uniqueID:** `*string` 

Unique transaction identifier.

⚠️ Generate a new unique ID for each withdrawal request.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` 

[Fiat](/glossary#fiat) currency [provider](/glossary#provider). Example: VISAMASTER

⚠️ Required for fiat currencies. Currency provider should be taken from [Asset Status endpoint](/public/http-v4/asset-status-list) response.
    
</dd>
</dl>

<dl>
<dd>

**network:** `*string` 

Cryptocurrency network. Available for multi network currencies. Example: OMNI

⚠️ Currency network should be taken from [Asset Status endpoint](/public/http-v4/asset-status-list) response. Default for USDT is ERC20
    
</dd>
</dl>

<dl>
<dd>

**partialEnable:** `*bool` — Optional parameter for [FIAT](/glossary#fiat) withdrawals with increased Maximum Limit if set as "true". To use this parameter, the application must support "Partially successful" withdrawal status and latest updates in deposit/withdrawal history.
    
</dd>
</dl>

<dl>
<dd>

**beneficiary:** `*gosdk.CreateWithdrawRequestBeneficiary` 

Beneficiary information data array.

⚠️ Required if currency [ticker](/glossary#ticker) is one of: UAH_IBAN, USD_VISAMASTER, EUR_VISAMASTER, USD, EUR
    
</dd>
</dl>

<dl>
<dd>

**travelRule:** `*gosdk.CreateWithdrawRequestTravelRule` 

Travel Rule information data array.

⚠️ Required if currency is crypto and the account is from [EEA](/glossary#european-economic-area-eea)
    
</dd>
</dl>

<dl>
<dd>

**paymentDescription:** `*string` 

Description of withdrawal destination

⚠️ Required if currency is crypto and withdrawal from whitebit-tr.com
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Withdraw.CreateWithdrawPay(request) -> []any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint has the similar logic as /main-account/withdraw, but with the only difference: amount that is specified will not include [fee](/glossary#fee) (it will be calculated to make target withdraw amount equal to the specified amount).

**Example:**
- When creating a base withdraw with amount = 100 USD, the receiver receives 100 USD minus the [fee](/glossary#fee), and the balance decreases by 100 USD.
- When using this endpoint with amount = 100 USD, the receiver receives 100 USD, and the balance decreases by 100 USD plus the [fee](/glossary#fee).

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.WithdrawRequest{
        Ticker: "ETH",
        Amount: "0.9",
        Address: "0x0964A6B8F794A4B8d61b62652dB27ddC9844FB4c",
        UniqueID: gosdk.String(
            "24529041",
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.Withdraw.CreateWithdrawPay(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ticker:** `string` — Currencies [ticker](/glossary#ticker). Example: BTC ⚠️ Currencies ticker must have "can_deposit" status equal to "true". Use [Asset Status endpoint](/public/http-v4/asset-status-list) to know more about currency.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Withdraw amount (including [fee](/glossary#fee)). To add the fee to the specified amount, use the /main-account/withdraw-pay request
    
</dd>
</dl>

<dl>
<dd>

**address:** `string` — Target address (wallet address for cryptocurrencies, identifier/[card token](/glossary#card-token) for [fiat](/glossary#fiat) currencies)
    
</dd>
</dl>

<dl>
<dd>

**memo:** `*string` — Required if currency is memoable. See [memo](/glossary#memodestination-tag) for details.
    
</dd>
</dl>

<dl>
<dd>

**uniqueID:** `*string` — Unique transaction identifier. ⚠️ Generate a new unique ID for each withdrawal request.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*string` — [Fiat](/glossary#fiat) currency [provider](/glossary#provider). Example: VISAMASTER ⚠️ Currency provider should be taken from [Asset Status endpoint](/public/http-v4/asset-status-list) response. Required if currency is fiat.
    
</dd>
</dl>

<dl>
<dd>

**network:** `*string` — Cryptocurrency network. Available for [multinetwork](/glossary#multinetwork) currencies. Example: OMNI ⚠️ Currency network should be taken from [Asset Status endpoint](/public/http-v4/asset-status-list) response. Default for USDT is ERC20
    
</dd>
</dl>

<dl>
<dd>

**partialEnable:** `*bool` — Optional parameter for [FIAT](/glossary#fiat) withdrawals with increased Maximum Limit if set as "true". To use this parameter, the application must support "Partially successful" withdrawal status and latest updates in deposit/withdrawal history.
    
</dd>
</dl>

<dl>
<dd>

**beneficiary:** `map[string]any` — Beneficiary information data. Required if currency [ticker](/glossary#ticker) is one of: UAH_IBAN, USD_VISAMASTER, EUR_VISAMASTER, USD, EUR
    
</dd>
</dl>

<dl>
<dd>

**travelRule:** `map[string]any` — Travel Rule information data. Required if currency is crypto and the account is from [EEA](/glossary#european-economic-area-eea)
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Transfer
<details><summary><code>client.Transfer.BetweenBalances(request) -> []any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint transfers the specified amount between [main](/glossary#balance-main), [trade](/glossary#balance-spotbalance-trade) and [collateral](/glossary#balance-collateral) balances.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>

<Note>
Also, fiat currencies can't be transferred without KYC verification.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.TransferBetweenBalancesRequest{
        Method: gosdk.TransferBetweenBalancesRequestMethodDeposit.Ptr(),
        Ticker: "XLM",
        Amount: "0.9",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.Transfer.BetweenBalances(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**method:** `*gosdk.TransferBetweenBalancesRequestMethod` 

Transfer method.

⚠️ We highly recommend to use **from** and **to** fields, which provides more flexibility. This way will be deprecated in future.

Example: **deposit** to transfer from [main](/glossary#balance-main) to [trade](/glossary#balance-spotbalance-trade) / **withdraw** to transfer from [trade](/glossary#balance-spotbalance-trade) balance to [main](/glossary#balance-main). For [collateral balances](/glossary#balance-collateral) use **collateral-deposit** to transfer from main to collateral and **collateral-withdraw** to transfer from collateral to main

**Not required** if **from** and **to** are set.
    
</dd>
</dl>

<dl>
<dd>

**from:** `*gosdk.TransferBetweenBalancesRequestFrom` 

Balance FROM which funds will move to. Acceptable values: [**main**](/glossary#balance-main), [**spot**](/glossary#balance-spotbalance-trade), [**collateral**](/glossary#balance-collateral)

**Not required** if **method** is set.
    
</dd>
</dl>

<dl>
<dd>

**to:** `*gosdk.TransferBetweenBalancesRequestTo` 

Balance TO which funds will move to. Acceptable values: [**main**](/glossary#balance-main), [**spot**](/glossary#balance-spotbalance-trade), [**collateral**](/glossary#balance-collateral)

**Not required** if **method** is set.
    
</dd>
</dl>

<dl>
<dd>

**ticker:** `string` — Currency's [ticker](/glossary#ticker). Example: BTC
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Amount to transfer. Max [precision](/glossary#precision) = 8, value must be greater than zero and less than or equal to the available balance.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Codes
<details><summary><code>client.Codes.CreateCode(request) -> *gosdk.CreateCodeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates [WhiteBIT code](/glossary#whitebit-codes).

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateCodeRequest{
        Ticker: "ETH",
        Amount: "0.002",
        Passphrase: gosdk.String(
            "some passphrase",
        ),
        Description: gosdk.String(
            "some description",
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.Codes.CreateCode(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ticker:** `string` — Currency's [ticker](/glossary#ticker). Example: BTC
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Amount to transfer. Max [precision](/glossary#precision) = 8, value must be greater than zero and less than or equal to the [main balance](/glossary#balance-main).
    
</dd>
</dl>

<dl>
<dd>

**passphrase:** `*string` — Passphrase for applying [WhiteBIT codes](/glossary#whitebit-codes). Passphrase must contain only latin letters, numbers and symbols (like !@#$%^, no whitespaces). Max: 25 symbols.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — Additional text description for [code](/glossary#whitebit-codes). Visible only for creator. Max: 75 symbols.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Codes.ApplyCode(request) -> *gosdk.ApplyCodeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint applies [WhiteBIT code](/glossary#whitebit-codes).

<Warning>
Rate limit: 60 requests/1 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ApplyCodeRequest{
        Code: "WBe11f4fce-2a53-4edc-b195-66b693bd77e3ETH",
        Passphrase: gosdk.String(
            "some passphrase",
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.Codes.ApplyCode(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**code:** `string` — [Code](/glossary#whitebit-codes) that will be applied.
    
</dd>
</dl>

<dl>
<dd>

**passphrase:** `*string` — Should be provided if the [code](/glossary#whitebit-codes) was created with passphrase. Max: 25 symbols.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Codes.GetMyCodes(request) -> *gosdk.GetMyCodesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the list of [WhiteBIT codes](/glossary#whitebit-codes) created by my account.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetMyCodesRequest{
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.Codes.GetMyCodes(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — LIMIT is a special clause used to limit records a particular query can return.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Use the OFFSET clause to return entries starting from a particular line.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Codes.GetCodesHistory(request) -> *gosdk.GetCodesHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the whole [codes](/glossary#whitebit-codes) history for the account.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetCodesHistoryRequest{
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.Codes.GetCodesHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — LIMIT is a special clause used to limit records a particular query can return.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Use the OFFSET clause to return entries starting from a particular line.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crypto Lending - Fixed
<details><summary><code>client.CryptoLendingFixed.GetFixedPlans(request) -> []*gosdk.FixedPlan</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves all active [plans](/glossary#crypto-lending).

<Note>
These endpoints are available only for B2B partner services. Fill the institutional services form to get permissions to use these endpoints.
</Note>

**Note:** When target currency is different from source currency, interest amount in target currency will be calculated using `interestRatio` value.

**Examples:**
- When source currency = USDT, target currency = BTC and interest ratio = 40000, interest is received in BTC and equals the USDT interest amount divided by the interest ratio (e.g. 0.000025 BTC per 1 USDT of interest).
- When source currency equals target currency, interest ratio equals 1.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetFixedPlansRequest{
        Ticker: gosdk.String(
            "USDT",
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CryptoLendingFixed.GetFixedPlans(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ticker:** `*string` — [Invest plan](/glossary#crypto-lending) source currency's [ticker](/glossary#ticker). Example: BTC
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CryptoLendingFixed.CreateFixedInvestment(request) -> *gosdk.CreateFixedInvestmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a new investment to the specified [invest plan](/glossary#crypto-lending).

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateFixedInvestmentRequest{
        PlanID: "8e667b4a-0b71-4988-8af5-9474dbfaeb51",
        Amount: "100",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CryptoLendingFixed.CreateFixedInvestment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planID:** `string` — [Invest plan](/glossary#crypto-lending) identifier
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Investment amount
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CryptoLendingFixed.CloseFixedInvestment(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint closes active investment.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CloseFixedInvestmentRequest{
        ID: "0d7b66ff-1909-4938-ab7a-d16d9a64dcd5",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CryptoLendingFixed.CloseFixedInvestment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Investment identifier
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CryptoLendingFixed.GetFixedInvestmentsHistory(request) -> *gosdk.GetFixedInvestmentsHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves an investments history.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetFixedInvestmentsHistoryRequest{
        ID: gosdk.String(
            "0d7b66ff-1909-4938-ab7a-d16d9a64dcd5",
        ),
        Ticker: gosdk.String(
            "USDT",
        ),
        Status: gosdk.Int(
            1,
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CryptoLendingFixed.GetFixedInvestmentsHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `*string` — Investment identifier
    
</dd>
</dl>

<dl>
<dd>

**ticker:** `*string` — [Invest plan](/glossary#crypto-lending) source currency's [ticker](/glossary#ticker)
    
</dd>
</dl>

<dl>
<dd>

**status:** `*int` — Investment status (1 - active, 2 - closed)
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — LIMIT is a special clause used to limit records a particular query can return.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Use the OFFSET clause to return entries starting from a particular line.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CryptoLendingFixed.GetInterestPaymentHistory(request) -> *gosdk.GetInterestPaymentHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the history of interest payments.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetInterestPaymentHistoryRequest{
        PlanID: gosdk.String(
            "8e667b4a-0b71-4988-8af5-9474dbfaeb51",
        ),
        Ticker: gosdk.String(
            "USDT",
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CryptoLendingFixed.GetInterestPaymentHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**planID:** `*string` — [Invest plan](/glossary#crypto-lending) identifier
    
</dd>
</dl>

<dl>
<dd>

**ticker:** `*string` — [Invest plan](/glossary#crypto-lending) target currency's [ticker](/glossary#ticker)
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — LIMIT is a special clause used to limit records a particular query can return.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Use the OFFSET clause to return entries starting from a particular line.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crypto Lending - Flex
<details><summary><code>client.CryptoLendingFlex.GetFlexPlans(request) -> []*gosdk.FlexPlan</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve list of active [Flex Plans](/glossary#crypto-lending).

Available after September 22, 2025.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetFlexPlansRequest{
        Limit: gosdk.Int(
            50,
        ),
        Offset: gosdk.Int(
            0,
        ),
        Ticker: gosdk.String(
            "USDT",
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CryptoLendingFlex.GetFlexPlans(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Pagination limit.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Pagination offset.
    
</dd>
</dl>

<dl>
<dd>

**ticker:** `*string` — Filter by currency [ticker](/glossary#ticker). Example: USDT
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CryptoLendingFlex.GetUserFlexInvestments(request) -> *gosdk.GetUserFlexInvestmentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve user's investment portfolio with optional filtering.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetUserFlexInvestmentsRequest{
        Limit: gosdk.Int(
            100,
        ),
        Offset: gosdk.Int(
            0,
        ),
        Ticker: gosdk.String(
            "USDT",
        ),
        Plan: gosdk.String(
            "8f2e9d3c-1a4b-4c2d-9e5f-6a7b8c9d0e1f",
        ),
        Investment: gosdk.String(
            "invest_id_123",
        ),
        InvestmentStatus: gosdk.Int(
            1,
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CryptoLendingFlex.GetUserFlexInvestments(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Pagination limit. Default: 100.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Pagination offset. Default: 0.
    
</dd>
</dl>

<dl>
<dd>

**ticker:** `*string` — Filter by currency [ticker](/glossary#ticker). Example: USDT.
    
</dd>
</dl>

<dl>
<dd>

**plan:** `*string` — Filter by plan ID (UUID).
    
</dd>
</dl>

<dl>
<dd>

**investment:** `*string` — Filter by investment ID.
    
</dd>
</dl>

<dl>
<dd>

**investmentStatus:** `*int` — Filter by status (1=ACTIVE, 0=CLOSED).
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CryptoLendingFlex.GetFlexInvestmentHistory(request) -> *gosdk.GetFlexInvestmentHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve complete investment operations history with advanced filtering.

**Available Action Types:**
- 1: INVEST - Investment creation
- 2: REINVEST - Automatic reinvestment
- 3: WITHDRAW_FROM_INVESTMENT - Partial withdrawal
- 4: DAILY_EARNING - Daily earnings
- 5: CLOSE_INVESTMENT - Investment closure
- 6: OPEN_INVESTMENT - Investment opening

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetFlexInvestmentHistoryRequest{
        Limit: gosdk.Int(
            50,
        ),
        Offset: gosdk.Int(
            0,
        ),
        Plan: gosdk.String(
            "8f2e9d3c-1a4b-4c2d-9e5f-6a7b8c9d0e1f",
        ),
        Investment: gosdk.String(
            "inv_123",
        ),
        Transaction: gosdk.String(
            "tx_456",
        ),
        DateFrom: gosdk.Int(
            1640995200,
        ),
        DateTo: gosdk.Int(
            1641081600,
        ),
        ActionTypes: []int{
            1,
            2,
            4,
        },
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CryptoLendingFlex.GetFlexInvestmentHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Pagination limit.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Pagination offset.
    
</dd>
</dl>

<dl>
<dd>

**plan:** `*string` — Filter by plan ID (UUID).
    
</dd>
</dl>

<dl>
<dd>

**investment:** `*string` — Filter by investment ID.
    
</dd>
</dl>

<dl>
<dd>

**transaction:** `*string` — Filter by transaction ID.
    
</dd>
</dl>

<dl>
<dd>

**dateFrom:** `*int` — Filter from date (timestamp).
    
</dd>
</dl>

<dl>
<dd>

**dateTo:** `*int` — Filter to date (timestamp).
    
</dd>
</dl>

<dl>
<dd>

**actionTypes:** `[]int` — Array of operation type IDs. See table below.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CryptoLendingFlex.GetFlexPaymentHistory(request) -> *gosdk.GetFlexPaymentHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve investment earnings history (ONLY DAILY_EARNING operations).

**Note:** The endpoint automatically filters to show ONLY DAILY_EARNING operations (type 4).

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetFlexPaymentHistoryRequest{
        Limit: gosdk.Int(
            50,
        ),
        Offset: gosdk.Int(
            0,
        ),
        Plan: gosdk.String(
            "8f2e9d3c-1a4b-4c2d-9e5f-6a7b8c9d0e1f",
        ),
        Investment: gosdk.String(
            "inv_123",
        ),
        Transaction: gosdk.String(
            "tx_456",
        ),
        DateFrom: gosdk.Int(
            1640995200,
        ),
        DateTo: gosdk.Int(
            1641081600,
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CryptoLendingFlex.GetFlexPaymentHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Pagination limit.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Pagination offset.
    
</dd>
</dl>

<dl>
<dd>

**plan:** `*string` — Filter by plan ID (UUID).
    
</dd>
</dl>

<dl>
<dd>

**investment:** `*string` — Filter by investment ID.
    
</dd>
</dl>

<dl>
<dd>

**transaction:** `*string` — Filter by transaction ID.
    
</dd>
</dl>

<dl>
<dd>

**dateFrom:** `*int` — Filter from date (timestamp).
    
</dd>
</dl>

<dl>
<dd>

**dateTo:** `*int` — Filter to date (timestamp).
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CryptoLendingFlex.CreateFlexInvestment(request) -> *gosdk.CreateFlexInvestmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create new investment in a Flex plan.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateFlexInvestmentRequest{
        Plan: "8f2e9d3c-1a4b-4c2d-9e5f-6a7b8c9d0e1f",
        Amount: "1000.500000",
        WithReinvest: gosdk.Bool(
            true,
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CryptoLendingFlex.CreateFlexInvestment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**plan:** `string` — Plan external ID (UUID).
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Investment amount.
    
</dd>
</dl>

<dl>
<dd>

**withReinvest:** `*bool` — Enable auto-reinvestment.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CryptoLendingFlex.WithdrawFromFlexInvestment(request) -> *gosdk.WithdrawFromFlexInvestmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Withdraw specified amount from user's investment.

**Note:** Plan must be active and accessible to user.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.WithdrawFromFlexInvestmentRequest{
        Plan: "8f2e9d3c-1a4b-4c2d-9e5f-6a7b8c9d0e1f",
        Amount: "500.250000",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CryptoLendingFlex.WithdrawFromFlexInvestment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**plan:** `string` — Plan external ID (UUID).
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Withdrawal amount.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CryptoLendingFlex.CloseFlexInvestment(request) -> *gosdk.CloseFlexInvestmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Completely close investment and withdraw all funds.

**Validation Rules:**
- plan: required, string, UUID format, must exist
- Investment must be ACTIVE

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CloseFlexInvestmentRequest{
        Plan: "8f2e9d3c-1a4b-4c2d-9e5f-6a7b8c9d0e1f",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CryptoLendingFlex.CloseFlexInvestment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**plan:** `string` — Plan external ID (UUID).
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CryptoLendingFlex.UpdateFlexAutoReinvestment(request) -> *gosdk.UpdateFlexAutoReinvestmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Enable/disable automatic reinvestment for user's investment.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.UpdateFlexAutoReinvestmentRequest{
        Plan: "8f2e9d3c-1a4b-4c2d-9e5f-6a7b8c9d0e1f",
        Enabled: gosdk.Bool(
            true,
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CryptoLendingFlex.UpdateFlexAutoReinvestment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**plan:** `string` — Plan external ID (UUID).
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Enable or disable auto-reinvestment.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Fees
<details><summary><code>client.Fees.GetFees(request) -> []*gosdk.FeeInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an array of objects containing deposit/withdrawal [fees](/glossary#fee) for the corresponding currencies.
Zero value in amount fields means that the setting is disabled.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetFeesRequest{
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.Fees.GetFees(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sub-Account
<details><summary><code>client.SubAccount.CreateSubAccount(request) -> *gosdk.SubAccount</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates new [sub-account](/glossary#sub-account).

<Note>
The `email` field requirement depends on the `shareKyc` parameter:
- When `shareKyc` is `false` or not provided: `email` is **required**
- When `shareKyc` is `true`: `email` is **optional**
</Note>

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateSubAccountRequest{
        Alias: "trading_bot",
        Permissions: &gosdk.CreateSubAccountRequestPermissions{
            SpotEnabled: true,
            CollateralEnabled: false,
        },
    }
client.SubAccount.CreateSubAccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**alias:** `string` — Name for sub-account
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — Sub-account email (required when shareKyc is false)
    
</dd>
</dl>

<dl>
<dd>

**shareKyc:** `*bool` — If KYC shared with main account
    
</dd>
</dl>

<dl>
<dd>

**permissions:** `*gosdk.CreateSubAccountRequestPermissions` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccount.DeleteSubAccount(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint deletes [sub-account](/glossary#sub-account).

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeleteSubAccountRequest{
        ID: "8e667b4a-0b71-4988-8af5-9474dbfaeb51",
    }
client.SubAccount.DeleteSubAccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Sub-account id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccount.EditSubAccount(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint edits [sub-account](/glossary#sub-account).

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.EditSubAccountRequest{
        ID: "8e667b4a-0b71-4988-8af5-9474dbfaeb51",
        Alias: "training",
        Permissions: &gosdk.EditSubAccountRequestPermissions{
            SpotEnabled: true,
            CollateralEnabled: false,
        },
    }
client.SubAccount.EditSubAccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Sub-account id
    
</dd>
</dl>

<dl>
<dd>

**alias:** `string` — Name for sub-account
    
</dd>
</dl>

<dl>
<dd>

**permissions:** `*gosdk.EditSubAccountRequestPermissions` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccount.ListSubAccounts(request) -> *gosdk.ListSubAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns list of current user [sub-accounts](/glossary#sub-account).

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListSubAccountsRequest{}
client.SubAccount.ListSubAccounts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**search:** `*string` — Search term
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccount.Transfer(request) -> *gosdk.SubAccountTransferResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates transfer from main account to [sub-account](/glossary#sub-account) or vice versa.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.SubAccountTransferRequest{
        ID: "8e667b4a-0b71-4988-8af5-9474dbfaeb51",
        Direction: gosdk.SubAccountTransferRequestDirectionMainToSub,
        Amount: "0.5",
        Ticker: "ETH",
    }
client.SubAccount.Transfer(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Sub-account id
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*gosdk.SubAccountTransferRequestDirection` — Transfer direction
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Transfer amount (min 0.00000001)
    
</dd>
</dl>

<dl>
<dd>

**ticker:** `string` — Currency's [ticker](/glossary#ticker)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccount.BlockSubAccount(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint blocks [sub-account](/glossary#sub-account).

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.BlockSubAccountRequest{
        ID: "8e667b4a-0b71-4988-8af5-9474dbfaeb51",
    }
client.SubAccount.BlockSubAccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Sub-account id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccount.UnblockSubAccount(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint unblocks [sub-account](/glossary#sub-account).

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.UnblockSubAccountRequest{
        ID: "8e667b4a-0b71-4988-8af5-9474dbfaeb51",
    }
client.SubAccount.UnblockSubAccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Sub-account id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccount.GetSubAccountBalances(request) -> map[string][]*gosdk.GetSubAccountBalancesResponseValueItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns [sub-account](/glossary#sub-account) balances.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetSubAccountBalancesRequest{
        ID: "8e667b4a-0b71-4988-8af5-9474dbfaeb51",
        Ticker: gosdk.String(
            "USDC",
        ),
    }
client.SubAccount.GetSubAccountBalances(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Sub-account id
    
</dd>
</dl>

<dl>
<dd>

**ticker:** `*string` — Currency's ticker (if not provided, returns data by all currencies)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccount.GetSubAccountTransferHistory(request) -> *gosdk.GetSubAccountTransferHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns history of transfers between main account and [sub-account](/glossary#sub-account).

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetSubAccountTransferHistoryRequest{
        ID: "8e667b4a-0b71-4988-8af5-9474dbfaeb51",
    }
client.SubAccount.GetSubAccountTransferHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Sub-account id
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*gosdk.GetSubAccountTransferHistoryRequestDirection` — Transfer direction (optional)
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sub-Account API Keys
<details><summary><code>client.SubAccountAPIKeys.CreateSubAccountAPIKey(request) -> *gosdk.SubAccountAPIKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a new API key for a [sub-account](/glossary#sub-account).

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateSubAccountAPIKeyRequest{
        Type: 1,
        SubAccountID: "8e667b4a-0b71-4988-8af5-9474dbfaeb51",
    }
client.SubAccountAPIKeys.CreateSubAccountAPIKey(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**type_:** `int` — Type of API key (1 - info and trading; 2 - info, trading, deposits, withdraws)
    
</dd>
</dl>

<dl>
<dd>

**subAccountID:** `string` — ID of the sub-account to create the API key for
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` — Custom title/name for the API key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountAPIKeys.EditSubAccountAPIKey(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint updates an existing [sub-account](/glossary#sub-account) API key.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.EditSubAccountAPIKeyRequest{
        APIKeyID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
        Title: "Trading Bot Key",
        URLs: []*gosdk.EditSubAccountAPIKeyRequestURLsItem{
            &gosdk.EditSubAccountAPIKeyRequestURLsItem{
                URL: gosdk.String(
                    "/api/v4/main-account/withdraw",
                ),
                Enable: gosdk.Bool(
                    false,
                ),
            },
            &gosdk.EditSubAccountAPIKeyRequestURLsItem{
                URL: gosdk.String(
                    "/api/v4/main-account/balance",
                ),
                Enable: gosdk.Bool(
                    true,
                ),
            },
        },
    }
client.SubAccountAPIKeys.EditSubAccountAPIKey(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKeyID:** `string` — ID of the API key to update
    
</dd>
</dl>

<dl>
<dd>

**title:** `string` — New title for the API key
    
</dd>
</dl>

<dl>
<dd>

**urls:** `[]*gosdk.EditSubAccountAPIKeyRequestURLsItem` — Array of URL objects for API key restrictions
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountAPIKeys.DeleteSubAccountAPIKey(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint deletes a [sub-account](/glossary#sub-account) API key.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeleteSubAccountAPIKeyRequest{
        APIKeyID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    }
client.SubAccountAPIKeys.DeleteSubAccountAPIKey(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKeyID:** `string` — ID of the API key to delete
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountAPIKeys.ListSubAccountAPIKeys(request) -> *gosdk.ListSubAccountAPIKeysResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves a list of API keys for a [sub-account](/glossary#sub-account).
Note: For security reasons, the apiSecret field returns an empty string.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListSubAccountAPIKeysRequest{}
client.SubAccountAPIKeys.ListSubAccountAPIKeys(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subAccountID:** `*string` — ID of the sub-account to list API keys for
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountAPIKeys.ResetSubAccountAPIKey(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint resets (regenerates) an existing [sub-account](/glossary#sub-account) API key.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ResetSubAccountAPIKeyRequest{
        APIKeyID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    }
client.SubAccountAPIKeys.ResetSubAccountAPIKey(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKeyID:** `string` — ID of the API key to reset
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountAPIKeys.ListSubAccountAPIKeyIPAddresses(request) -> *gosdk.ListSubAccountAPIKeyIPAddressesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the list of IP addresses allowed for a [sub-account](/glossary#sub-account) API key.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListSubAccountAPIKeyIPAddressesRequest{
        APIKeyID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    }
client.SubAccountAPIKeys.ListSubAccountAPIKeyIPAddresses(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKeyID:** `string` — ID of the API key to list IP addresses for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountAPIKeys.CreateSubAccountAPIKeyIPAddress(request) -> *gosdk.CreateSubAccountAPIKeyIPAddressResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint adds a new IP address to the allowed list for a [sub-account](/glossary#sub-account) API key.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateSubAccountAPIKeyIPAddressRequest{
        APIKeyID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
        IP: "192.168.1.100",
    }
client.SubAccountAPIKeys.CreateSubAccountAPIKeyIPAddress(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKeyID:** `string` — ID of the API key to add IP address to
    
</dd>
</dl>

<dl>
<dd>

**ip:** `string` — IP address to add to allowed list
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccountAPIKeys.DeleteSubAccountAPIKeyIPAddress(request) -> *gosdk.DeleteSubAccountAPIKeyIPAddressResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint removes an IP address from the allowed list for a [sub-account](/glossary#sub-account) API key.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeleteSubAccountAPIKeyIPAddressRequest{
        APIKeyID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
        IP: "192.168.1.100",
    }
client.SubAccountAPIKeys.DeleteSubAccountAPIKeyIPAddress(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKeyID:** `string` — ID of the API key to remove IP address from
    
</dd>
</dl>

<dl>
<dd>

**ip:** `string` — IP address to remove from allowed list
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Mining Pool
<details><summary><code>client.MiningPool.GetMiningRewards(request) -> *gosdk.GetMiningRewardsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns rewards received from mining.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetMiningRewardsRequest{}
client.MiningPool.GetMiningRewards(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**account:** `*string` — Mining pool account
    
</dd>
</dl>

<dl>
<dd>

**from:** `*int` — Date timestamp starting from which rewards are received
    
</dd>
</dl>

<dl>
<dd>

**to:** `*int` — Date timestamp until which rewards are received
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MiningPool.GetMiningHashrate(request) -> *gosdk.GetMiningHashrateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns hashrate of mining pool account.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetMiningHashrateRequest{
        Account: "miner123",
    }
client.MiningPool.GetMiningHashrate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**account:** `string` — Mining pool account
    
</dd>
</dl>

<dl>
<dd>

**from:** `*int` — Unix timestamp of starting point
    
</dd>
</dl>

<dl>
<dd>

**to:** `*int` — Unix timestamp of final point
    
</dd>
</dl>

<dl>
<dd>

**interval:** `*gosdk.GetMiningHashrateRequestInterval` — Timestamp interval
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MiningPool.GetMiningPayoutDestination(request) -> *gosdk.GetMiningPayoutDestinationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the current payout destination setting for a specific mining account belonging to the authenticated user.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetMiningPayoutDestinationRequest{
        AccountName: "my_miner_01",
    }
client.MiningPool.GetMiningPayoutDestination(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountName:** `string` — Mining pool account name
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MiningPool.SetMiningPayoutDestination(request) -> *gosdk.SetMiningPayoutDestinationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the payout destination for a specific mining account belonging to the authenticated user. Can be set to main balance or an external BTC address.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.SetMiningPayoutDestinationRequest{
        AccountName: "my_miner_01",
        Destination: gosdk.SetMiningPayoutDestinationRequestDestinationMainBalance,
    }
client.MiningPool.SetMiningPayoutDestination(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountName:** `string` — Mining pool account name
    
</dd>
</dl>

<dl>
<dd>

**destination:** `*gosdk.SetMiningPayoutDestinationRequestDestination` — Payout destination type
    
</dd>
</dl>

<dl>
<dd>

**address:** `*string` — External BTC address. Required when destination is external_address. Supports all standard Bitcoin address formats.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MiningPool.GetMiningMinerInfo(request) -> *gosdk.GetMiningMinerInfoResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns fee information and stratum connection details with worker counts for a specific mining account.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetMiningMinerInfoRequest{
        Account: "my_miner_01",
    }
client.MiningPool.GetMiningMinerInfo(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**account:** `string` — Mining pool account name
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MiningPool.GetMiningWorkerNames(request) -> *gosdk.GetMiningWorkerNamesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of online worker names for a specific mining account.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetMiningWorkerNamesRequest{
        Account: "my_miner_01",
    }
client.MiningPool.GetMiningWorkerNames(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**account:** `string` — Mining pool account name
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Pagination offset
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Pagination limit
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MiningPool.GetMiningWorkerHashrate(request) -> *gosdk.GetMiningWorkerHashrateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns hashrate performance history for a specific worker on a mining account.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetMiningWorkerHashrateRequest{
        Account: "my_miner_01",
        Worker: "worker_001",
    }
client.MiningPool.GetMiningWorkerHashrate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**account:** `string` — Mining pool account name
    
</dd>
</dl>

<dl>
<dd>

**worker:** `string` — Worker name
    
</dd>
</dl>

<dl>
<dd>

**interval:** `*gosdk.GetMiningWorkerHashrateRequestInterval` — Time frame granularity
    
</dd>
</dl>

<dl>
<dd>

**from:** `*int` — Start timestamp in Unix seconds. Must be <= now
    
</dd>
</dl>

<dl>
<dd>

**to:** `*int` — End timestamp in Unix seconds. Must be <= now
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MiningPool.CreateMiningWatcherLink(request) -> *gosdk.CreateMiningWatcherLinkResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new watcher link for one or more mining accounts, granting specific permissions with a configurable expiration.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateMiningWatcherLinkRequest{
        Accounts: []string{
            "my_miner_01",
            "my_miner_02",
        },
        Name: "monitoring_link",
        Permissions: []gosdk.CreateMiningWatcherLinkRequestPermissionsItem{
            gosdk.CreateMiningWatcherLinkRequestPermissionsItemDashboard,
            gosdk.CreateMiningWatcherLinkRequestPermissionsItemWorkers,
        },
        LiveUntil: gosdk.CreateMiningWatcherLinkRequestLiveUntilOneH,
    }
client.MiningPool.CreateMiningWatcherLink(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accounts:** `[]string` — Array of mining account names
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — Link name (alphanumeric and underscores only)
    
</dd>
</dl>

<dl>
<dd>

**permissions:** `[]*gosdk.CreateMiningWatcherLinkRequestPermissionsItem` — Array of permissions
    
</dd>
</dl>

<dl>
<dd>

**liveUntil:** `*gosdk.CreateMiningWatcherLinkRequestLiveUntil` — Expiration period
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MiningPool.ListMiningWatcherLinks(request) -> *gosdk.ListMiningWatcherLinksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns all active watcher links for a specific mining account.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListMiningWatcherLinksRequest{
        Account: "my_miner_01",
    }
client.MiningPool.ListMiningWatcherLinks(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**account:** `string` — Mining pool account name
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MiningPool.CreateMiningAccount(request) -> *gosdk.CreateMiningAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new mining account for the authenticated user. The account name must be unique within the user's accounts.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateMiningAccountRequest{
        Name: "my_miner_01",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.MiningPool.CreateMiningAccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — Mining pool account name. Must be unique. Alphanumeric characters and underscores allowed.
    
</dd>
</dl>

<dl>
<dd>

**referralCode:** `*string` — Optional referral code for account creation
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MiningPool.GetMiningAccounts(request) -> *gosdk.GetMiningAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of mining accounts for the authenticated user. Supports filtering by account name.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetMiningAccountsRequest{
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.MiningPool.GetMiningAccounts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `*string` — Optional filter to search for a specific mining account name (exact match)
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Credit Line
<details><summary><code>client.CreditLine.GetCreditLineInfo(request) -> *gosdk.CreditLine</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns an active loan.
The endpoint works on demand - contact WhiteBIT support to get access.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
The API does not cache the response.
</Note>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetCreditLineInfoRequest{
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CreditLine.GetCreditLineInfo(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Collateral Trading
<details><summary><code>client.CollateralTrading.CollateralAccountBalance(request) -> map[string]float64</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns a current [collateral balance](/glossary#balance-collateral).

<Note>
The API does not cache the response.
</Note>

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>

<Accordion title="Errors">
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "ticker": ["ticker is invalid."]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CollateralAccountBalanceRequest{
        Ticker: gosdk.String(
            "BTC",
        ),
        Request: gosdk.String(
            "{{request}}",
        ),
        Nonce: gosdk.String(
            "{{nonce}}",
        ),
    }
client.CollateralTrading.CollateralAccountBalance(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ticker:** `*string` 

[Asset](/glossary#assets) to be filtered. For example: BTC

If not specified, returns balances for all assets.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CollateralAccountBalanceSummary(request) -> []*gosdk.CollateralAccountBalanceSummaryResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves collateral account balance summary with detailed breakdown per asset.

<Note>
The API does not cache the response.
</Note>

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>

<Accordion title="Errors">
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "ticker": ["ticker is invalid."]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CollateralAccountBalanceSummaryRequest{
        Ticker: gosdk.String(
            "BTC",
        ),
        Request: gosdk.String(
            "{{request}}",
        ),
        Nonce: gosdk.String(
            "{{nonce}}",
        ),
    }
client.CollateralTrading.CollateralAccountBalanceSummary(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ticker:** `*string` 

Filter by requested asset. For example: BTC

If not specified, returns summary for all assets.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CreateCollateralLimitOrder(request) -> *gosdk.CreateCollateralLimitOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates [limit order](/glossary#limit-order) using [collateral balance](/glossary#balance-collateral).

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Note>
For open long position use **buy**, for short **sell**. To close current position, place opposite order with current position amount.
</Note>

<Note>
  - RPI orders are post-only by design and cannot be used with the IOC flag. The API returns error code `37` when both `rpi=true` and `ioc=true` are used.
</Note>


<Accordion title="Error Codes">
  - `30` - default validation error code
  - `31` - market validation failed
  - `32` - amount validation failed
  - `33` - price validation failed
  - `36` - client_order_id validation failed
  - `37` - `ioc=true` cannot be used with `postOnly=true` or `rpi=true`
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateCollateralLimitOrderRequest{
        Market: "BTC_USDT",
        Side: gosdk.CreateCollateralLimitOrderRequestSideBuy,
        Amount: "0.01",
        Price: "40000",
        ClientOrderID: gosdk.String(
            "order1987111",
        ),
        StopLoss: gosdk.String(
            "50000",
        ),
        TakeProfit: gosdk.String(
            "30000",
        ),
        PostOnly: gosdk.Bool(
            false,
        ),
        Ioc: gosdk.Bool(
            false,
        ),
        Rpi: gosdk.Bool(
            true,
        ),
        PositionSide: gosdk.CreateCollateralLimitOrderRequestPositionSideLong.Ptr(),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CollateralTrading.CreateCollateralLimitOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` — Available margin [market](/glossary#market). Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**side:** `*gosdk.CreateCollateralLimitOrderRequestSide` — Order type. Variables: 'buy' / 'sell'. For open long position use **buy**, for short **sell**.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Amount of [stock](/glossary#stock) currency to buy or sell.
    
</dd>
</dl>

<dl>
<dd>

**price:** `string` — Price in [money](/glossary#money) currency. Example: '9800'
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Identifier should be unique and contain letters, dashes, numbers, dots or underscores.
    
</dd>
</dl>

<dl>
<dd>

**stopLoss:** `*string` 

Stop loss price.

When provided, the system creates an [OTO](/glossary#one-triggers-the-other-oto) order with a stop loss condition.
    
</dd>
</dl>

<dl>
<dd>

**takeProfit:** `*string` 

Take profit price.

When provided, the system creates an [OTO](/glossary#one-triggers-the-other-oto) order with a take profit condition.
    
</dd>
</dl>

<dl>
<dd>

**postOnly:** `*bool` — Orders are guaranteed to be the [maker](/glossary#maker) order when [executed](/glossary#finished-orders).
    
</dd>
</dl>

<dl>
<dd>

**ioc:** `*bool` — An immediate or cancel order (IOC) is an order that attempts to execute all or part immediately and then cancels any unfilled portion.
    
</dd>
</dl>

<dl>
<dd>

**rpi:** `*bool` 

Enables Retail Price Improvement (RPI) mode.

RPI orders are post-only by design and cannot be used with `ioc=true`. The API returns error code `37` when both `rpi=true` and `ioc=true` are used.
    
</dd>
</dl>

<dl>
<dd>

**positionSide:** `*gosdk.CreateCollateralLimitOrderRequestPositionSide` — Defines the position direction when hedge mode is enabled. See [positionSide](/glossary#position-side)
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CreateCollateralBulkOrder(request) -> []*gosdk.CreateCollateralBulkOrderResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates multiple collateral limit orders.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateCollateralBulkOrderRequest{
        Orders: []*gosdk.CreateCollateralBulkOrderRequestOrdersItem{
            &gosdk.CreateCollateralBulkOrderRequestOrdersItem{
                Market: gosdk.String(
                    "BTC_PERP",
                ),
                Side: gosdk.CreateCollateralBulkOrderRequestOrdersItemSideBuy.Ptr(),
                Amount: gosdk.String(
                    "0.02",
                ),
                Price: gosdk.String(
                    "40000",
                ),
                ClientOrderID: gosdk.String(
                    "",
                ),
                PostOnly: gosdk.Bool(
                    false,
                ),
                Ioc: gosdk.Bool(
                    false,
                ),
                Rpi: gosdk.Bool(
                    true,
                ),
                PositionSide: gosdk.CreateCollateralBulkOrderRequestOrdersItemPositionSideLong.Ptr(),
            },
            &gosdk.CreateCollateralBulkOrderRequestOrdersItem{
                Market: gosdk.String(
                    "BTC_USDT",
                ),
                Side: gosdk.CreateCollateralBulkOrderRequestOrdersItemSideSell.Ptr(),
                Amount: gosdk.String(
                    "0.0001",
                ),
                Price: gosdk.String(
                    "41000",
                ),
                ClientOrderID: gosdk.String(
                    "",
                ),
                PostOnly: gosdk.Bool(
                    false,
                ),
                Ioc: gosdk.Bool(
                    false,
                ),
                Rpi: gosdk.Bool(
                    true,
                ),
                PositionSide: gosdk.CreateCollateralBulkOrderRequestOrdersItemPositionSideLong.Ptr(),
            },
            &gosdk.CreateCollateralBulkOrderRequestOrdersItem{
                Market: gosdk.String(
                    "ETH_BTC",
                ),
                Side: gosdk.CreateCollateralBulkOrderRequestOrdersItemSideSell.Ptr(),
                Amount: gosdk.String(
                    "0.02",
                ),
                Price: gosdk.String(
                    "0.030",
                ),
                ClientOrderID: gosdk.String(
                    "",
                ),
                PostOnly: gosdk.Bool(
                    false,
                ),
                Ioc: gosdk.Bool(
                    false,
                ),
                Rpi: gosdk.Bool(
                    true,
                ),
                PositionSide: gosdk.CreateCollateralBulkOrderRequestOrdersItemPositionSideLong.Ptr(),
            },
        },
        StopOnFail: gosdk.Bool(
            true,
        ),
        Request: gosdk.String(
            "{{request}}",
        ),
        Nonce: gosdk.String(
            "{{nonce}}",
        ),
    }
client.CollateralTrading.CreateCollateralBulkOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orders:** `[]*gosdk.CreateCollateralBulkOrderRequestOrdersItem` 
    
</dd>
</dl>

<dl>
<dd>

**stopOnFail:** `*bool` 

Controls how the bulk order processor handles failures.

When true: Processing stops at the first order that fails validation or execution. Only orders up to (but not including) the failed order are processed.

When false (default): All orders in the bulk request are processed regardless of individual failures. Each order result is returned in the response array.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CreateCollateralMarketOrder(request) -> *gosdk.CreateCollateralMarketOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a collateral market order.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateCollateralMarketOrderRequest{
        Market: "BTC_USDT",
        Side: gosdk.CreateCollateralMarketOrderRequestSideBuy,
        Amount: "0.01",
        ClientOrderID: gosdk.String(
            "order1987111",
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CollateralTrading.CreateCollateralMarketOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**side:** `*gosdk.CreateCollateralMarketOrderRequestSide` 
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**stopLoss:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**takeProfit:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**positionSide:** `*gosdk.CreateCollateralMarketOrderRequestPositionSide` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CreateCollateralStopLimitOrder(request) -> *gosdk.CreateCollateralStopLimitOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a collateral stop-limit order.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateCollateralStopLimitOrderRequest{
        Market: "BTC_USDT",
        Side: gosdk.CreateCollateralStopLimitOrderRequestSideBuy,
        Amount: "0.001",
        Price: "40000",
        ActivationPrice: "40000",
        StopLoss: gosdk.String(
            "30000",
        ),
        TakeProfit: gosdk.String(
            "50000",
        ),
        ClientOrderID: gosdk.String(
            "order1987111",
        ),
        PositionSide: gosdk.CreateCollateralStopLimitOrderRequestPositionSideLong.Ptr(),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CollateralTrading.CreateCollateralStopLimitOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**side:** `*gosdk.CreateCollateralStopLimitOrderRequestSide` 
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**price:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**activationPrice:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**stopLoss:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**takeProfit:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**positionSide:** `*gosdk.CreateCollateralStopLimitOrderRequestPositionSide` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CreateCollateralTriggerMarketOrder(request) -> *gosdk.CreateCollateralTriggerMarketOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a collateral trigger market order.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateCollateralTriggerMarketOrderRequest{
        Market: "BTC_USDT",
        Side: gosdk.CreateCollateralTriggerMarketOrderRequestSideBuy,
        Amount: "0.01",
        ActivationPrice: "40000",
        ClientOrderID: gosdk.String(
            "order1987111",
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CollateralTrading.CreateCollateralTriggerMarketOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**side:** `*gosdk.CreateCollateralTriggerMarketOrderRequestSide` 
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**activationPrice:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**stopLoss:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**takeProfit:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**positionSide:** `*gosdk.CreateCollateralTriggerMarketOrderRequestPositionSide` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CollateralAccountSummary(request) -> *gosdk.CollateralAccountSummaryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves collateral account summary.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CollateralAccountSummaryRequest{
        Request: gosdk.String(
            "{{request}}",
        ),
        Nonce: gosdk.String(
            "{{nonce}}",
        ),
    }
client.CollateralTrading.CollateralAccountSummary(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.GetOpenPositions(request) -> []*gosdk.GetOpenPositionsResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves open positions.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetOpenPositionsRequest{
        Market: gosdk.String(
            "BTC_USDT",
        ),
        Request: gosdk.String(
            "{{request}}",
        ),
        Nonce: gosdk.String(
            "{{nonce}}",
        ),
    }
client.CollateralTrading.GetOpenPositions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `*string` 

Filter by specific market. For example: BTC_USDT

If not specified, returns all open positions.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.ClosePosition(request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint closes a position.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ClosePositionRequest{
        PositionID: 123,
        PositionSide: gosdk.ClosePositionRequestPositionSideLong.Ptr(),
        Market: "BTC_USDT",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CollateralTrading.ClosePosition(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**positionID:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**positionSide:** `*gosdk.ClosePositionRequestPositionSide` 
    
</dd>
</dl>

<dl>
<dd>

**market:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.GetPositionsHistory(request) -> []*gosdk.GetPositionsHistoryResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves positions history.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetPositionsHistoryRequest{
        Market: gosdk.String(
            "BTC_USDT",
        ),
        PositionID: gosdk.Int(
            1,
        ),
        Request: gosdk.String(
            "{{request}}",
        ),
        Nonce: gosdk.String(
            "{{nonce}}",
        ),
    }
client.CollateralTrading.GetPositionsHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**positionID:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.GetFundingHistory(request) -> *gosdk.GetFundingHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves funding history.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetFundingHistoryRequest{
        Market: gosdk.String(
            "BTC_PERP",
        ),
        Limit: gosdk.Int(
            100,
        ),
        Offset: gosdk.Int(
            0,
        ),
        Request: gosdk.String(
            "{{request}}",
        ),
        Nonce: gosdk.String(
            "{{nonce}}",
        ),
    }
client.CollateralTrading.GetFundingHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `*string` 

Filter by specific market. For example: BTC_PERP

If not specified, returns funding history for all markets.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of records to return
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Number of records to skip
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.ChangeCollateralAccountLeverage(request) -> *gosdk.ChangeCollateralAccountLeverageResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint changes account leverage.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ChangeCollateralAccountLeverageRequest{
        Leverage: 5,
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CollateralTrading.ChangeCollateralAccountLeverage(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**leverage:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.GetCollateralHedgeMode(request) -> *gosdk.GetCollateralHedgeModeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves hedge mode status.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetCollateralHedgeModeRequest{
        Request: gosdk.String(
            "{{request}}",
        ),
        Nonce: gosdk.String(
            "{{nonce}}",
        ),
    }
client.CollateralTrading.GetCollateralHedgeMode(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.UpdateHedgeMode(request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint updates hedge mode.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.UpdateHedgeModeRequest{
        HedgeMode: true,
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CollateralTrading.UpdateHedgeMode(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**hedgeMode:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.GetConditionalOrders(request) -> *gosdk.GetConditionalOrdersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves active conditional orders.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetConditionalOrdersRequest{
        Market: gosdk.String(
            "BTC_USDT",
        ),
        Offset: gosdk.Int(
            0,
        ),
        Limit: gosdk.Int(
            100,
        ),
    }
client.CollateralTrading.GetConditionalOrders(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.GetOcoOrders(request) -> []*gosdk.GetOcoOrdersResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves active OCO orders.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetOcoOrdersRequest{
        Market: gosdk.String(
            "BTC_USDT",
        ),
        Offset: gosdk.Int(
            0,
        ),
        Limit: gosdk.Int(
            100,
        ),
    }
client.CollateralTrading.GetOcoOrders(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CreateCollateralOcoOrder(request) -> *gosdk.CreateCollateralOcoOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a collateral OCO order.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateCollateralOcoOrderRequest{
        Market: "BTC_USDT",
        Side: gosdk.CreateCollateralOcoOrderRequestSideBuy,
        Amount: "0.001",
        Price: "40000",
        ActivationPrice: "41000",
        StopLimitPrice: "42000",
        ClientOrderID: gosdk.String(
            "order1987111",
        ),
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CollateralTrading.CreateCollateralOcoOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**side:** `*gosdk.CreateCollateralOcoOrderRequestSide` 
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**price:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**activationPrice:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**stopLimitPrice:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CancelConditionalOrder(request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint cancels a conditional order.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CancelConditionalOrderRequest{
        Market: "BTC_USDT",
        ID: 117703764514,
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CollateralTrading.CancelConditionalOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CancelOcoOrder(request) -> *gosdk.CancelOcoOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint cancels an OCO order.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CancelOcoOrderRequest{
        Market: "BTC_USDT",
        OrderID: 117703764514,
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CollateralTrading.CancelOcoOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CancelOtoOrder(request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint cancels an OTO order.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CancelOtoOrderRequest{
        Market: "BTC_USDT",
        OtoID: 117703764514,
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.CollateralTrading.CancelOtoOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**otoID:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Market Fee
<details><summary><code>client.MarketFee.GetMarketFee() -> *gosdk.GetMarketFeeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns maker and taker fees for a specific market.

The `maker` and `taker` fields represent spot trading fees. The `futures_maker` and `futures_taker` fields represent futures trading fees.

The system calculates the effective futures fee as the lower value between the user-specific custom fee and the market-specific fee.

When the market fee is lower than the assigned custom fee, the system returns the market fee.

Example: If the custom futures taker fee equals `0.026` and the market fee equals `0.02`, the response returns `0.02`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetMarketFeeRequest{
        Market: "BTC_USDT",
    }
client.MarketFee.GetMarketFee(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` 

Market to query.

If the request includes the `market` parameter, the system returns fees for the specified market only.

When fee values are identical across markets, the response contains identical values regardless of the specified market.

Example: BTC_USDT
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Spot Trading
<details><summary><code>client.SpotTrading.TradeAccountBalance(request) -> map[string]*gosdk.TradeAccountBalanceResponseValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the [trade balance](/glossary#balance-spotbalance-trade) by currency [ticker](/glossary#ticker) or all balances.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>

<Accordion title="Errors">
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "ticker": ["Ticker field should be a string."]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "ticker": ["Currency was not found."]
  }
}
```

```json
{
  "code": 1,
  "message": "Inner validation failed",
  "errors": {
    "amount": ["Invalid argument."]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.TradeAccountBalanceRequest{}
client.SpotTrading.TradeAccountBalance(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ticker:** `*string` — Currency's [ticker](/glossary#ticker). Example: BTC
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CreateLimitOrder(request) -> *gosdk.OrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates [limit trading order](/glossary#limit-order).

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Note>
  - RPI orders do not appear in public order book feeds (`depth`, `bookTicker`). RPI orders are visible only in private active orders and in the exchange UI order book (web/mobile).
  - RPI orders are post-only by design and cannot be used with the IOC flag. The API returns error code `37` when both `rpi=true` and `ioc=true` are used.
</Note>

<Accordion title="Error Codes">
  - `30` - default validation error code
  - `31` - market validation failed
  - `32` - amount validation failed
  - `33` - price validation failed
  - `36` - client_order_id validation failed
  - `37` - `ioc=true` cannot be used with `postOnly=true` or `rpi=true`
</Accordion>

<Accordion title="Errors">
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "amount": ["Amount field is required."],
    "market": ["Market field is required."],
    "price": ["Price field is required."],
    "side": ["Side field is required."]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "side": ["Side field should contain only 'buy' or 'sell' values."]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": ["Amount field should be numeric string or number."]
  }
}
```

```json
{
  "code": 33,
  "message": "Validation failed",
  "errors": {
    "price": ["Price field should be numeric string or number."]
  }
}
```

```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market is not available."]
  }
}
```

```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market field should not be empty string."]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": [
      "Given amount is less than min amount 0.001",
      "Min amount step = 0.000001"
    ]
  }
}
```

```json
{
  "code": 36,
  "message": "Validation failed",
  "errors": {
    "client_order_id": ["ClientOrderId field should be a string."]
  }
}
```

```json
{
  "code": 36,
  "message": "Validation failed",
  "errors": {
    "client_order_id": [
      "ClientOrderId field should contain only latin letters, numbers and dashes."
    ]
  }
}
```

```json
{
  "code": 36,
  "message": "Validation failed",
  "errors": {
    "client_order_id": [
      "This client order id is already used by the current account."
    ]
  }
}
```

```json
{
  "code": 37,
  "message": "Validation failed",
  "errors": {
    "ioc": ["Either IOC or PostOnly flag in true state is allowed."]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "total": ["Total (amount * price) is less than 5.05"]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": [
      "Min amount step = 0.01"
    ]
  }
}
```

```json
{
  "code": 33,
  "message": "Validation failed",
  "errors": {
    "price": ["Price field should be at least 10", "Min price step = 0.000001"]
  }
}
```

```json
{
  "code": 33,
  "message": "Validation failed",
  "errors": {
    "price": ["Price should be greater than 0."]
  }
}
```

```json
{
  "code": 35,
  "message": "Validation failed",
  "errors": {
    "maker_fee": ["Incorrect maker fee"]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.LimitOrderRequest{
        Market: "BTC_USDT",
        Side: gosdk.LimitOrderRequestSideBuy,
        Amount: "0.001",
        Price: "9800",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.SpotTrading.CreateLimitOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` — Available [market](/glossary#market). Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**side:** `*gosdk.LimitOrderRequestSide` — Order type. Variables: 'buy' / 'sell' Example: 'buy'
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Amount of [stock](/glossary#stock) currency to buy or sell. Example: '0.001' or 0.001
    
</dd>
</dl>

<dl>
<dd>

**price:** `string` — Price in money currency. Example: '9800' or 9800
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Identifier should be unique and contain letters, dashes, numbers, dots or underscores. The identifier must be unique.
    
</dd>
</dl>

<dl>
<dd>

**postOnly:** `*bool` — [Orders](/glossary#orders) are guaranteed to be the [maker](/glossary#maker) order when [executed](/glossary#finished-orders). Variables: 'true' / 'false' Example: 'false'.
    
</dd>
</dl>

<dl>
<dd>

**ioc:** `*bool` 

Immediate-or-cancel (IOC) executes all or part of an order immediately and cancels any unfilled portion.

IOC does not support `rpi=true` because RPI uses post-only behavior by design.
The API returns error code `37` when a request sets both `ioc=true` and `rpi=true`.

Refer to [Order Parameter Rules](/guides/order-parameter-rules) for unsupported parameter combinations.
    
</dd>
</dl>

<dl>
<dd>

**bboRole:** `*int` — When the [BBO](/glossary#bbo) option is activated for Limit orders, the system selects the best market prices for execution. Variables: 1 - Queue Method / 2 - Counterparty Method. Use method 2 with ioc flag. Example: 2.
    
</dd>
</dl>

<dl>
<dd>

**stp:** `*gosdk.LimitOrderRequestStp` — Self trade prevention mode. Variables: 'no' / 'cancel_both' / 'cancel_new' / 'cancel_old'. Example: 'no'.
    
</dd>
</dl>

<dl>
<dd>

**rpi:** `*bool` 

Enables Retail Price Improvement (RPI) mode.

RPI orders use post-only behavior by design. An RPI order does not support `ioc=true`.
The API returns error code `37` when a request sets both `rpi=true` and `ioc=true`.
RPI orders do not appear in public order book feeds (`depth`, `bookTicker`). RPI orders are visible only in private active orders and in the exchange UI order book (web/mobile).
RPI executions may apply custom fees or rebates, especially when trading via sub-accounts. Use Query Market Fee / Query All Market Fees to verify effective fees.

Refer to [Order Parameter Rules](/guides/order-parameter-rules) for unsupported parameter combinations.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CreateBulkLimitOrder(request) -> gosdk.BulkLimitOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates bulk [limit trading orders](/glossary#limit-order).

<Warning>
  Limit: From 1 to 20 orders per request.
</Warning>

<Note>
  - RPI orders do not appear in public order book feeds (`depth`, `bookTicker`). RPI orders are visible only in private active orders and in the exchange UI order book (web/mobile).
  - RPI orders are post-only by design and cannot be used with the IOC flag. The API returns error code `37` when both `rpi=true` and `ioc=true` are used.
</Note>


<Accordion title="Error Codes">
  - `30` - default validation error code
  - `31` - market validation failed
  - `32` - amount validation failed
  - `33` - price validation failed
  - `36` - client_order_id validation failed
  - `37` - `ioc=true` cannot be used with `postOnly=true` or `rpi=true`
</Accordion>

<Accordion title="Errors">
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "orders": ["The orders must be an array."]
  }
}
```

Individual order errors (in multiply response):

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "amount": ["Amount field is required."],
    "market": ["Market field is required."],
    "price": ["Price field is required."],
    "side": ["Side field is required."]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "side": ["Side field should contain only 'buy' or 'sell' values."]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": ["Amount field should be numeric string or number."]
  }
}
```

```json
{
  "code": 33,
  "message": "Validation failed",
  "errors": {
    "price": ["Price field should be numeric string or number."]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateBulkLimitOrderRequest{
        Orders: []*gosdk.BulkOrderItem{
            &gosdk.BulkOrderItem{
                Side: gosdk.BulkOrderItemSideBuy.Ptr(),
                Amount: gosdk.String(
                    "0.02",
                ),
                Price: gosdk.String(
                    "40000",
                ),
                Market: gosdk.String(
                    "BTC_USDT",
                ),
                PostOnly: gosdk.Bool(
                    false,
                ),
                Ioc: gosdk.Bool(
                    false,
                ),
                ClientOrderID: gosdk.String(
                    "",
                ),
                Rpi: gosdk.Bool(
                    true,
                ),
            },
            &gosdk.BulkOrderItem{
                Side: gosdk.BulkOrderItemSideSell.Ptr(),
                Amount: gosdk.String(
                    "0.0001",
                ),
                Price: gosdk.String(
                    "41000",
                ),
                Market: gosdk.String(
                    "BTC_USDT",
                ),
                PostOnly: gosdk.Bool(
                    false,
                ),
                Ioc: gosdk.Bool(
                    false,
                ),
                ClientOrderID: gosdk.String(
                    "",
                ),
                Rpi: gosdk.Bool(
                    true,
                ),
            },
            &gosdk.BulkOrderItem{
                Side: gosdk.BulkOrderItemSideSell.Ptr(),
                Amount: gosdk.String(
                    "0.02",
                ),
                Price: gosdk.String(
                    "41000",
                ),
                Market: gosdk.String(
                    "BTC_USDT",
                ),
                PostOnly: gosdk.Bool(
                    false,
                ),
                Ioc: gosdk.Bool(
                    false,
                ),
                ClientOrderID: gosdk.String(
                    "",
                ),
                Rpi: gosdk.Bool(
                    true,
                ),
            },
        },
    }
client.SpotTrading.CreateBulkLimitOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orders:** `[]*gosdk.BulkOrderItem` — Array of limit orders
    
</dd>
</dl>

<dl>
<dd>

**stopOnFail:** `*bool` 

Controls how the bulk order processor handles failures.

When true: Processing stops at the first order that fails validation or execution. Only orders up to (but not including) the failed order are processed.

When false (default): All orders in the bulk request are processed regardless of individual failures. Each order result is returned in the response array.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CreateMarketOrder(request) -> *gosdk.OrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates [market trading order](/glossary#market-order).

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
- `30` - default validation error code
- `31` - market validation failed
- `32` - amount validation failed
- `36` - client_order_id validation failed
</Accordion>

<Accordion title="Errors">
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "amount": ["Amount field is required."],
    "market": ["Market field is required."],
    "side": ["Side field is required."]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "side": ["Side field should contain only 'buy' or 'sell' values."]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": ["Amount field should be numeric string or number."]
  }
}
```

```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market is not available."]
  }
}
```

```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market field should not be empty string."]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": ["Not enough balance."]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": [
      "Given amount is less than min amount 0.001",
      "Min amount step = 0.000001"
    ]
  }
}
```

```json
{
  "code": 36,
  "message": "Validation failed",
  "errors": {
    "client_order_id": ["ClientOrderId field should be a string."]
  }
}
```

```json
{
  "code": 36,
  "message": "Validation failed",
  "errors": {
    "client_order_id": [
      "ClientOrderId field should contain only latin letters, numbers and dashes."
    ]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.MarketOrderRequest{
        Market: "BTC_USDT",
        Side: gosdk.MarketOrderRequestSideBuy,
        Amount: "100",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.SpotTrading.CreateMarketOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*gosdk.MarketOrderRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CreateStockMarketOrder(request) -> *gosdk.OrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates buy [stock](/glossary#stock) market trading [order](/glossary#orders).

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
- `30` - default validation error code
- `31` - market validation failed
- `32` - amount validation failed
- `36` - client_order_id validation failed
</Accordion>

<Accordion title="Errors">
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "amount": ["Amount field is required."],
    "market": ["Market field is required."],
    "side": ["Side field is required."]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "side": ["Side field should contain only 'buy' or 'sell' values."]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": ["Amount field should be numeric string or number."]
  }
}
```

```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market is not available."]
  }
}
```

```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market field should not be empty string."]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": ["Not enough balance."]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": [
      "Given amount is less than min amount 0.001",
      "Min amount step = 0.000001"
    ]
  }
}
```

```json
{
  "code": 36,
  "message": "Validation failed",
  "errors": {
    "client_order_id": ["ClientOrderId field should be a string."]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.MarketOrderRequest{
        Market: "BTC_USDT",
        Side: gosdk.MarketOrderRequestSideBuy,
        Amount: "100",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.SpotTrading.CreateStockMarketOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*gosdk.MarketOrderRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CreateStopLimitOrder(request) -> *gosdk.OrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates [stop-limit trading order](/glossary#stop-limit-order).

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
- `30` - default validation error code
- `31` - market validation failed
- `32` - amount validation failed
- `33` - price validation failed
- `36` - client_order_id validation failed
</Accordion>

<Accordion title="Errors">
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "activation_price": ["Activation price field is required."],
    "amount": ["Amount field is required."],
    "market": ["Market field is required."],
    "price": ["Price field is required."],
    "side": ["Side field is required."]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "side": ["Side field should contain only 'buy' or 'sell' values."]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": ["Amount field should be numeric string or number."]
  }
}
```

```json
{
  "code": 33,
  "message": "Validation failed",
  "errors": {
    "price": ["Price field should be numeric string or number."]
  }
}
```

```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market is not available."]
  }
}
```

```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market field should not be empty string."]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": ["Not enough balance."]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": [
      "Given amount is less than min amount 0.001",
      "Min amount step = 0.000001"
    ]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "total": ["Total (amount * price) is less than 5.05"]
  }
}
```

```json
{
  "code": 36,
  "message": "Validation failed",
  "errors": {
    "client_order_id": ["ClientOrderId field should be a string."]
  }
}
```

```json
{
  "code": 36,
  "message": "Validation failed",
  "errors": {
    "client_order_id": [
      "ClientOrderId field should contain only latin letters, numbers and dashes."
    ]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.StopLimitOrderRequest{
        Market: "BTC_USDT",
        Side: gosdk.StopLimitOrderRequestSideBuy,
        Amount: "0.001",
        Price: "9800",
        ActivationPrice: "10000",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.SpotTrading.CreateStopLimitOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` — Available [market](/glossary#market). Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**side:** `*gosdk.StopLimitOrderRequestSide` — Order type. Variables: 'buy' / 'sell' Example: 'buy'
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Amount of [stock](/glossary#stock) currency to buy or sell. Example: '0.001' or 0.001
    
</dd>
</dl>

<dl>
<dd>

**price:** `string` — Price in [money](/glossary#money) currency. Example: '9800' or 9800
    
</dd>
</dl>

<dl>
<dd>

**activationPrice:** `string` — Activation price in [money](/glossary#money) currency. Example: '10000' or 10000
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Identifier should be unique and contain letters, dashes, numbers, dots or underscores. The identifier must be unique.
    
</dd>
</dl>

<dl>
<dd>

**bboRole:** `*int` — When the [BBO](/glossary#bbo) option is activated for Limit orders, the system selects the best market prices for execution. Variables: 1 - Queue Method / 2 - Counterparty Method.
    
</dd>
</dl>

<dl>
<dd>

**stp:** `*gosdk.StopLimitOrderRequestStp` — Self trade prevention mode. Variables: 'no' / 'cancel_both' / 'cancel_new' / 'cancel_old'. Example: 'no'.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CreateStopMarketOrder(request) -> *gosdk.OrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates [stop-market trading order](/glossary#stop-market-order).

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
- `30` - default validation error code
- `31` - market validation failed
- `32` - amount validation failed
- `36` - client_order_id validation failed
</Accordion>

<Accordion title="Errors">
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "activation_price": ["Activation price field is required."],
    "amount": ["Amount field is required."],
    "market": ["Market field is required."],
    "side": ["Side field is required."]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "side": ["Side field should contain only 'buy' or 'sell' values."]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": ["Amount field should be numeric string or number."]
  }
}
```

```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market is not available."]
  }
}
```

```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market field should not be empty string."]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": ["Not enough balance."]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": [
      "Given amount is less than min amount 0.001",
      "Min amount step = 0.000001"
    ]
  }
}
```

```json
{
  "code": 36,
  "message": "Validation failed",
  "errors": {
    "client_order_id": ["ClientOrderId field should be a string."]
  }
}
```

```json
{
  "code": 36,
  "message": "Validation failed",
  "errors": {
    "client_order_id": [
      "ClientOrderId field should contain only latin letters, numbers and dashes."
    ]
  }
}
```

```json
{
  "code": 36,
  "message": "Validation failed",
  "errors": {
    "client_order_id": [
      "This client order id is already used by the current account."
    ]
  }
}
```

```json
{
  "code": 32,
  "message": "Validation failed",
  "errors": {
    "amount": ["Amount should be greater than 0."]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.StopMarketOrderRequest{
        Market: "BTC_USDT",
        Side: gosdk.StopMarketOrderRequestSideBuy,
        Amount: "0.01",
        ActivationPrice: "10000",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.SpotTrading.CreateStopMarketOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` — Available [market](/glossary#market). Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**side:** `*gosdk.StopMarketOrderRequestSide` — Order type. Variables: 'buy' / 'sell' Example: 'buy'
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Amount of [money](/glossary#money) currency to buy or amount in [stock](/glossary#stock) currency to sell. Example: '0.01' or 0.01 for buy and '0.0001' for sell.
    
</dd>
</dl>

<dl>
<dd>

**activationPrice:** `string` — Activation price in [money](/glossary#money) currency. Example: '10000' or 10000
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Identifier should be unique and contain letters, dashes, numbers, dots or underscores. The identifier must be unique.
    
</dd>
</dl>

<dl>
<dd>

**stp:** `*gosdk.StopMarketOrderRequestStp` — Self trade prevention mode. Variables: 'no' / 'cancel_both' / 'cancel_new' / 'cancel_old'. Example: 'no'.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CancelOrder(request) -> *gosdk.OrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancel existing [order](/glossary#orders).

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Note>
- Modification by client_order_id takes priority over order_id.
- The request supports working only with order_id or only with client_order_id.
- Do not pass both values at the same time.
</Note>

<Accordion title="Error Codes">
- `30` - default validation error code
- `31` - market validation failed
</Accordion>

<Accordion title="Errors">
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "market": ["Market field is required."],
    "order_id": ["OrderId field is required."]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "market": ["Market is not available."]
  }
}
```

```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market is not available."]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "order_id": ["OrderId field should be an integer."]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "market": [
      "Market field should be a string.",
      "Market field format is invalid."
    ]
  }
}
```

```json
{
  "code": 2,
  "message": "Inner validation failed",
  "errors": {
    "order_id": ["Unexecuted order was not found."]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CancelOrderRequest{
        Market: "BTC_USDT",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.SpotTrading.CancelOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` — Available [market](/glossary#market). Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `*int` — Order Id. Example: 4180284841. Required if client_order_id is not set.
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Custom client order id. Example: 'customId11'. Required if order_id is not set.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CancelAllOrders(request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancels all orders that meet the conditions [order](/glossary#orders).

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
- `30` - default validation error code
- `31` - market validation failed
</Accordion>

<Accordion title="Errors">
```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market is not available."]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "type": ["The type must be an array."]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "market": [
      "Market field should be a string.",
      "Market field format is invalid."
    ]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CancelAllOrdersRequest{
        Market: gosdk.String(
            "BTC_USDT",
        ),
        Type: []gosdk.CancelAllOrdersRequestTypeItem{
            gosdk.CancelAllOrdersRequestTypeItemSpot,
            gosdk.CancelAllOrdersRequestTypeItemMargin,
            gosdk.CancelAllOrdersRequestTypeItemFutures,
        },
    }
client.SpotTrading.CancelAllOrders(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `*string` — Available [market](/glossary#market). Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**type_:** `[]*gosdk.CancelAllOrdersRequestTypeItem` — Order types value. Example: 'spot', 'margin', 'futures'
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.GetActiveOrders(request) -> []*gosdk.OrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves [active orders](/glossary#active-orders) (orders not yet executed).

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>

<Note>
Search across all markets is available only if client_order_id and order_id are not provided.
</Note>

<Accordion title="Errors">
```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market is not available"]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "limit": ["The limit may not be greater than 100."],
    "offset": ["The offset may not be greater than 10000."]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetActiveOrdersRequest{}
client.SpotTrading.GetActiveOrders(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `*string` — Available [market](/glossary#market). Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `*int` — Available order_id. Example: 3134995325
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Available client_order_id. Example: customId11
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Starting line index (OFFSET). Default: 0, Min: 0
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — LIMIT is a special clause used to limit records a particular query can return. Default: 50, Min: 1, Max: 100
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.GetExecutedOrderHistory(request) -> []*gosdk.GetExecutedOrderHistoryResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves all deals for all markets. Can be filtered by single market if needed.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>

<Note>
The endpoint can retrieve data not older than 6 months from current month. For older data, use the Report on the History page.
</Note>

<Accordion title="Errors">
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "limit": ["Limit field should be an integer."],
    "offset": ["Offset field should be an integer."]
  }
}
```

```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market field format is invalid."]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetExecutedOrderHistoryRequest{}
client.SpotTrading.GetExecutedOrderHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `*string` — Requested [market](/glossary#market). Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Filter by custom order identifier
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*int` — Start date in Unix-time format
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*int` — End date in Unix-time format
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Starting line index (OFFSET). Default: 0, Min: 0
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — LIMIT is a special clause used to limit records a particular query can return. Default: 50, Min: 1, Max: 100
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.GetOrderDeals(request) -> *gosdk.GetOrderDealsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves deals for a specific order.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetOrderDealsRequest{
        OrderID: 3134995325,
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.SpotTrading.GetOrderDeals(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderID:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.GetOrderHistory(request) -> map[string][]*gosdk.GetOrderHistoryResponseValueItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves order history.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetOrderHistoryRequest{}
client.SpotTrading.GetOrderHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.ModifyOrder(request) -> *gosdk.OrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint modifies existing [order](/glossary#orders).

Supported order types: limit, stop limit, stop market.

Request must contain one of the following parameters: amount, price, activationPrice.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Note>
- Use total parameter instead of amount for modify buy stop market order.
- Modification by client_order_id takes priority.
- The request supports working only with order_id or only with client_order_id.
- Do not pass both values at the same time.
</Note>

<Accordion title="Error Codes">
**Status 400** (client errors): 1, 2, 6, 20, 24, 101, 158

**Status 422** (business logic): 10, 11, 12, 13, 14, 15, 16, 17, 25, 27, 40, 42, 51, 103, 104, 105, 106, 111, 112, 113, 114, 115, 150, 151, 152, 153, 155, 157, 159, 160, 161, 162, 163, 250, 251, 300, 302, 330
</Accordion>

<Accordion title="Errors">
```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market is not available."]
  }
}
```

```json
{
  "code": 2,
  "message": "Inner validation failed",
  "errors": {
    "order_id": ["Unexecuted order was not found."]
  }
}
```

```json
{
  "code": 10,
  "message": "Validation failed",
  "errors": {
    "amount": ["Not enough balance."]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ModifyOrderRequest{
        Market: "BTC_USDT",
        Request: "{{request}}",
        Nonce: "{{nonce}}",
    }
client.SpotTrading.ModifyOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderID:** `*int` — Active order id. Required if client_order_id is not set.
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Identifier should be unique and contain letters, dashes, numbers, dots or underscores. Required if order_id is not set.
    
</dd>
</dl>

<dl>
<dd>

**market:** `string` — Available [market](/glossary#market). Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*string` — Amount of [stock](/glossary#stock) currency to buy or sell. Example: '0.001' or 0.001
    
</dd>
</dl>

<dl>
<dd>

**total:** `*string` — Total of [money](/glossary#money) currency to buy or sell. Example: '0.001' or 0.001
    
</dd>
</dl>

<dl>
<dd>

**price:** `*string` — Price in [money](/glossary#money) currency. Example: '9800' or 9800
    
</dd>
</dl>

<dl>
<dd>

**activationPrice:** `*string` — Activation price in [money](/glossary#money) currency. Example: '10000' or 10000
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.SetKillSwitch(request) -> *gosdk.SetKillSwitchResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates, updates, deletes [kill-switch timer](/glossary#kill-switch-timer).

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Note>
- If timeout=null - delete existing timer by market.
- If types=null - create timer by market for all order types.
</Note>

<Accordion title="Error Codes">
- `30` - default validation error code
- `31` - market validation failed
</Accordion>

<Accordion title="Errors">
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "market": ["Market field is required."],
    "timeout": ["Timeout field is required."]
  }
}
```

```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market is not available."]
  }
}
```

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "timeout": ["Timeout should be at least 5."]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.SetKillSwitchRequest{
        Market: "BTC_USDT",
        Timeout: "60",
    }
client.SpotTrading.SetKillSwitch(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `string` — Available [market](/glossary#market). Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**timeout:** `string` — Timer value. Example: '5'-'600' or null
    
</dd>
</dl>

<dl>
<dd>

**types:** `[]*gosdk.SetKillSwitchRequestTypesItem` — Order types value. Example: 'spot', 'margin', 'futures' or null
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.GetKillSwitchStatus(request) -> []*gosdk.GetKillSwitchStatusResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the status of [kill-switch timer](/glossary#kill-switch-timer).

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
- `30` - default validation error code
- `31` - market validation failed
</Accordion>

<Accordion title="Errors">
```json
{
  "code": 31,
  "message": "Validation failed",
  "errors": {
    "market": ["Market is not available."]
  }
}
```
</Accordion>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetKillSwitchStatusRequest{}
client.SpotTrading.GetKillSwitchStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**market:** `*string` — Available [market](/glossary#market). Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>
