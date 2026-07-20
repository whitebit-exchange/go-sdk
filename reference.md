# Reference
<details><summary><code>client.ConvertEstimate(request) -> *sdk.ConvertEstimateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a quote for converting one currency to another. Quote lifetime is 10 seconds, then quote will be expired.

The minimum convert size is derived from the resulting proceeds, not a fixed per-pair floor: a request is rejected as too small when `rate × amount` rounds to zero in the target currency. Account balance is pre-checked at estimate time and re-checked at confirm time, because balance can change within the 10-second quote window. There is an absolute server-side maximum on the conversion amount.

<Note>
The endpoint can be used to obtain a pre-execution price estimate for a market order. Call the endpoint with the desired amount before placing a market order to see the approximate execution price.
</Note>

<Note>
Error `message` values may be returned as translation keys (for example `validation.required`) rather than finalized English strings. Treat the `code` and the field name under `errors` as the stable contract.
</Note>

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
request := &sdk.ConvertEstimateRequest{
        From: "BTC",
        To: "USDT",
        Direction: sdk.ConvertEstimateRequestDirectionTo,
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

**direction:** `*sdk.ConvertEstimateRequestDirection` — Convert amount direction, defines in which currency corresponding "amount" field is populated. Use "to" in case amount is in "to" currency, use "from" if amount is in "from" currency
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Amount to convert or receive. The value is silently truncated to 8 decimal places before evaluation; excess decimals do not raise an error.
    
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

<details><summary><code>client.ConvertConfirm(request) -> *sdk.ConvertConfirmResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint confirms an estimated quote.

An expired quote returns code `20` (`api.converter.quoteExpired`), which is distinct from a quote that could not be found (code `0`). Balance is re-checked at confirm time and can fail with code `37` even when the estimate succeeded, because balance may change within the 10-second quote window.

<Note>
A re-confirmed (already used) quote and a quote that never existed both return code `0` with the `quoteId` field key `frontendServerSide.converter.quoteInvalid`. The response alone does not distinguish "already used" from "never existed".
</Note>

<Note>
Error `message` values may be returned as translation keys (for example `api.converter.quoteExpired`) rather than finalized English strings. Treat the `code` and the field name under `errors` as the stable contract.
</Note>

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
request := &sdk.ConvertConfirmRequest{
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

<details><summary><code>client.ConvertHistory(request) -> *sdk.ConvertHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns convert history, sorted by `id` descending (newest first).

The `from`–`to` window is capped at 30 days per request, even though data is retained for 6 months. A wider range is rejected with code `30` (`api.validation.dateTime.maxRange`).

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Note>
Error `message` values may be returned as translation keys (for example `api.validation.dateTime.maxRange`) rather than finalized English strings. Treat the `code` and the field name under `errors` as the stable contract.
</Note>

**Note:** The endpoint can retrieve data not older than 6 months from the current month. For older data, use the Report on the History page.
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
request := &sdk.ConvertHistoryRequest{
        FromTicker: sdk.String(
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

**from:** `*string` — From time filter (Unix seconds). Must be no more than 30 days before `to` and no older than 6 months. Example: 1699260637. Default: now()
    
</dd>
</dl>

<dl>
<dd>

**to:** `*string` — To time filter (Unix seconds). Must be no more than 30 days after `from`. Example: 1699260637. Default: now()
    
</dd>
</dl>

<dl>
<dd>

**quoteID:** `*string` — Quote Id. Example: 4050
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — How many records to receive. Allowed range: 1–100. Default: 100
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Number of records to skip for pagination. Minimum: 0. Default: 0
    
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

## PublicAPIV4
<details><summary><code>client.PublicAPIV4.MaintenanceStatus() -> *sdk.GetAPIV4PublicPlatformStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the current maintenance status of the WhiteBIT platform. Use the response to detect scheduled downtime and pause trading automation during maintenance windows. The `status` field returns `"system operational"` when all platform services are available, or `"system maintenance"` when the platform is undergoing planned maintenance.
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

<details><summary><code>client.PublicAPIV4.MarketInfo() -> []*sdk.GetAPIV4PublicMarketsResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves configuration and trading rules for all available spot, futures, and TradFi futures markets. Use the response to discover tradeable pairs, check minimum order sizes, and read fee schedules. Each entry includes precision settings, fee ratios, and order-size constraints for the market.

<Note>
Market configuration is reference data, re-synced from the database approximately every 10 seconds. Polling more frequently returns identical data. The cache is shared across all callers.
</Note>

<Note>
TradFi futures markets are region-gated. Markets not available in a given region are omitted from the response entirely and do not appear under any other market type.
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

<details><summary><code>client.PublicAPIV4.MarketActivity() -> map[string]*sdk.GetAPIV4PublicTickerResponseValue</code></summary>
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

<details><summary><code>client.PublicAPIV4.AssetStatusList() -> map[string]*sdk.Asset</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the deposit and withdrawal status for every supported asset. Use the response to check whether deposits and withdrawals are enabled, read per-network fee and limit details, and determine required blockchain confirmation counts. The response includes crypto assets, fiat currencies, and fiat payment methods.

<Note>
Asset status is reference data, re-synced approximately once per minute. Polling more frequently returns identical data. The cache is shared across all callers.
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

<details><summary><code>client.PublicAPIV4.Orderbook(Market) -> *sdk.OrderbookResponse</code></summary>
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
request := &sdk.GetAPIV4PublicOrderbookMarketRequest{
        Market: "BTC_USDT",
        Limit: sdk.Int(
            100,
        ),
        Level: sdk.Int(
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

**level:** `*int` — Aggregation level for price grouping. Level 0 applies no aggregation. Levels 1–5 provide increasing aggregation of the order book.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PublicAPIV4.Depth(Market) -> *sdk.OrderbookResponse</code></summary>
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
request := &sdk.GetAPIV4PublicOrderbookDepthMarketRequest{
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

<details><summary><code>client.PublicAPIV4.RecentTrades(Market) -> []*sdk.GetAPIV4PublicTradesMarketResponseItem</code></summary>
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
request := &sdk.GetAPIV4PublicTradesMarketRequest{
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

**type_:** `*sdk.GetAPIV4PublicTradesMarketRequestType` — Filter by trade side. Omit to return both buy and sell trades.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PublicAPIV4.Fee() -> map[string]*sdk.FeeInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the [fee](/glossary#fee) schedule and deposit/withdrawal limits for every supported asset. Use the response to display fee estimates before a user initiates a deposit or withdrawal. The response is keyed by currency ticker; each entry contains deposit and withdrawal fee amounts and min/max transfer limits.

<Note>
The fee schedule is reference data, re-synced approximately once per minute. Polling more frequently returns identical data. The cache is shared across all callers.
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

<details><summary><code>client.PublicAPIV4.ServerTime() -> *sdk.GetAPIV4PublicTimeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the current server time as a Unix timestamp. Use the response to synchronize local clocks before generating HMAC signatures for authenticated requests. The endpoint takes no parameters and has no request-validation errors; it returns HTTP 200 on success and fails only at the infrastructure level (see the API description).

<Note>
The server time is computed per request and is not cached.
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

The endpoint checks API availability by returning a simple health-check response. Use the endpoint to verify network connectivity and confirm the API server is reachable. A successful response contains the string `"pong"`. The endpoint takes no parameters and has no request-validation errors; it returns HTTP 200 on success and fails only at the infrastructure level (see the API description).

<Note>
The health-check response is generated per request and is not cached.
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

<details><summary><code>client.PublicAPIV4.CollateralMarketsList() -> *sdk.GetAPIV4PublicCollateralMarketsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns the list of [market](/glossary#market) pair names available for [collateral](/glossary#collateral) trading. Use the response to determine which markets support margin positions. Each item in the result array is a market pair name in `BASE_QUOTE` format (e.g., `BTC_USDT`).

<Note>
The collateral market list is reference data, re-synced approximately every 10 seconds. Polling more frequently returns identical data. The cache is shared across all callers.
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

<details><summary><code>client.PublicAPIV4.AvailableFuturesMarketsList() -> *sdk.GetAPIV4PublicFuturesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns detailed information for all available futures markets. Use the response to read current pricing, open interest, funding rates, and leverage bracket configuration. Each entry includes the predicted next funding rate, settlement timestamps, and maximum allowed position sizes per leverage level.

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

<details><summary><code>client.PublicAPIV4.FundingHistory(Market) -> []*sdk.GetAPIV4PublicFundingHistoryMarketResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns the funding rate history for a specified futures market. Use the response to analyze historical funding rate trends and settlement prices. Results are sorted by funding time in descending order and support offset-based pagination via `limit` and `offset` parameters.

<Warning>
Rate limit 2000 requests/10 sec.
</Warning>

<Note>
This endpoint supports pagination. Use `limit` (default: 100, max: 100) and `offset` (default: 0, max: 1000000) to page through results.
</Note>

<Note>
The response is a plain array with no `total`, `has_more`, or cursor — a returned count below `limit` marks the last page (an empty array means no further records).
</Note>

<Note>
Funding history is served per request at the API layer, with no application-level cache.
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
request := &sdk.GetAPIV4PublicFundingHistoryMarketRequest{
        Market: "BTC_PERP",
        StartDate: sdk.Int(
            1752480000,
        ),
        EndDate: sdk.Int(
            1752537600,
        ),
        Limit: sdk.Int(
            100,
        ),
        Offset: sdk.Int(
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

**limit:** `*int` — Number of records to return. Default: 100, Maximum: 100
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Number of records to skip. Maximum: 1000000
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Main Account
<details><summary><code>client.MainAccount.GetMainBalance(request) -> map[string]*sdk.GetMainBalanceResponseValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the [main balance](/glossary#balance-main) by currency [ticker](/glossary#ticker) or all balances.

An unknown or non-existent `ticker` is not an error: the endpoint returns a zero balance (`{"main_balance": "0"}`) for it. Omitting `ticker` returns balances for every currency.

<Accordion title="Errors">
```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "ticker": ["validation.string"]
  }
}
```
</Accordion>

Beyond the validation error above, this endpoint can return only the [common authentication errors](/api-reference/authentication).

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
request := &sdk.GetMainBalanceRequest{
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**nonce:** `int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MainAccount.GetDepositWithdrawHistory(request) -> *sdk.GetDepositWithdrawHistoryResponse</code></summary>
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
- `AML frozen` - 21
- `Uncredited` - 22
- `Pending` - 15

`Successful` (3, 7) is a final state. A credited deposit does not later transition to `AML frozen` (21) or `Uncredited` (22); those checks happen before crediting.

**Travel Rule Deposit check status codes:**
- `Awaiting verification` - 27: The transaction has been frozen due to the lack of data required under the Travel Rule. The user is required to provide this data manually through the exchange interface.
- `Confirmation in progress` - 28: The Travel Rule data provided by the user is currently being verified by WhiteBIT.

⚠️ Due to regulatory requirements in Turkey and [EU](/glossary#european-economic-area-eea), the system places every inbound crypto deposit on hold (frozen) until confirming the transaction's origin. The sender must provide certain details if the transaction is from another Virtual Asset Service Provider (VASP) or verify the address if from a self-hosted wallet. The system credits deposited funds to the account only after successful verification.

**Withdraw status codes:**
- `Pending` - 1, 2, 6, 10–17 (withdrawal in progress).
- `Successful` - 3, 7
- `Canceled` - 4
- `Unconfirmed by user` - 5
- `AML frozen` - 21
- `Partially successful` - 18

<Warning>
Rate limit: 200 requests/10 sec.
</Warning>

<Warning>
Requests with `limit` values above 100 return large payloads. Use high limits only when necessary and ensure the client application can handle large response sizes.
</Warning>

<Note>
The API does not cache the response.
</Note>

<Note>
Results are sorted newest first (descending by timestamp).
</Note>

<Note>
**No date filtering:** the endpoint does not accept `startDate` / `endDate` parameters, and pagination is capped at `offset + limit ≤ 10000` (requests beyond the cap return `Offset is too big. Please use offset + limit less than 10000.`). To read more than 10,000 records, narrow the result set with the available filters (`transactionMethod`, `ticker`, `status`) and paginate within each subset; for a complete history export beyond the cap, use the Report on the History page.
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
request := &sdk.GetDepositWithdrawHistoryRequest{
        TransactionMethod: sdk.Int(
            1,
        ),
        Ticker: sdk.String(
            "BTC",
        ),
        Limit: sdk.Int(
            100,
        ),
        Offset: sdk.Int(
            0,
        ),
        Status: []int{
            3,
            7,
        },
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**limit:** `*int` — LIMIT is a special clause used to limit records a particular query can return. Default: 50, Min: 1, Max: 500
    
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

**nonce:** `int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Deposit
<details><summary><code>client.Deposit.GetDepositAddress(request) -> *sdk.GetDepositAddressResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves a deposit address of the cryptocurrency.

<Note>
Sub-accounts use this endpoint with their own API key once deposits are enabled for the
account. Crypto deposits are disabled by default — to enable them, contact your assigned
Account Manager or email institutional@whitebit.com.
</Note>

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
request := &sdk.GetDepositAddressRequest{
        Ticker: "BTC",
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**nonce:** `int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Deposit.GetFiatDepositURL(request) -> *sdk.GetFiatDepositURLResponse</code></summary>
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
    "successLink": ["Your domain scheme incorrect. Use https only"],
    "failureLink": ["Your domain scheme incorrect. Use https only"]
  }
}
```

```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "ticker": ["Currency is not depositable via API"]
  }
}
```

```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "user": ["User not verified"]
  }
}
```

```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "amount": ["Amount is too big for deposit"]
  }
}
```

```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "amount": ["Daily limit reached"]
  }
}
```

```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "amount": ["Expiration date cannot be used for this provider"]
  }
}
```

```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "customer.birthDate": ["You must be at least 18 years old"]
  }
}
```

```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "address": ["Invalid credit card number"]
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
request := &sdk.GetFiatDepositURLRequest{
        Ticker: "UAH",
        Provider: "VISAMASTER",
        Amount: "100",
        UniqueID: "{{generateID}}",
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**uniqueID:** `string` — Unique transaction identifier on client's side. Any string up to 255 characters; not validated as a UUID.
    
</dd>
</dl>

<dl>
<dd>

**customer:** `*sdk.GetFiatDepositURLRequestCustomer` — Customer information (required for USD/EUR with VISAMASTER [provider](/glossary#provider))
    
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

**nonce:** `int` — Unique request identifier
    
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
request := &sdk.IssueCardTokenRequest{
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

<details><summary><code>client.Deposit.RefundDeposit(request) -> *sdk.RefundDepositResponse</code></summary>
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
request := &sdk.RefundDepositRequest{
        TransactionID: "54bffeb7-7a8f-43f8-bcd8-f14ec10fee85",
        Address: "0x1234567890abcdef1234567890abcdef12345678",
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**transactionID:** `string` — Transaction UUID of the deposit. Obtain from the [deposit.canceled](/platform/webhook) webhook (`uniqueId` field) or from the deposit/withdraw history in the WhiteBIT interface.
    
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

**nonce:** `int` — A unique identifier for the request. Use a monotonically increasing value such as a Unix timestamp in milliseconds.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Deposit.CreateNewAddress(request) -> *sdk.CreateNewAddressResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a new address even when the last created address is not used. The endpoint is not available by default, contact support@whitebit.com to get permissions to use the endpoint.

<Note>
For sub-accounts, crypto deposits must also be enabled for the account (disabled by
default). To enable them, contact your assigned Account Manager or email institutional@whitebit.com.
</Note>

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
request := &sdk.CreateNewAddressRequest{
        Ticker: "XLM",
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**type_:** `*sdk.CreateNewAddressRequestType` — Address type, available for specific currencies list (see address types table in endpoint description)
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## JWT
<details><summary><code>client.Jwt.IssueJwtToken(request) -> *sdk.IssueJwtTokenResponse</code></summary>
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
request := &sdk.IssueJwtTokenRequest{
        Request: "{{request}}",
        NonceWindow: sdk.Bool(
            false,
        ),
        Nonce: 1594297865000,
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

**nonce:** `int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jwt.GetWebSocketToken(request) -> *sdk.GetWebSocketTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The V4 endpoint can be used to retrieve the WebSocket token for user.
The token is required to authorize WebSocket connections for private API access.

<Accordion title="Errors">
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "user": ["user not found"]
  }
}
```
</Accordion>

Beyond the error above, this endpoint can return only the [common authentication errors](/api-reference/authentication).

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
request := &sdk.GetWebSocketTokenRequest{
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**nonce:** `int` — Unique request identifier
    
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
request := &sdk.CreateWithdrawRequest{
        Ticker: "ETH",
        Amount: "0.9",
        Address: "0x0964A6B8F794A4B8d61b62652dB27ddC9844FB4c",
        UniqueID: "24529041",
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**uniqueID:** `string` 

Unique transaction identifier. Any string up to 255 characters; not validated as a UUID.

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

**customerIP:** `*string` 

End-customer IP address forwarded to the [fiat](/glossary#fiat) [provider](/glossary#provider) for antifraud checks before the withdrawal is processed.

⚠️ Required if currency [ticker](/glossary#ticker) is USD or EUR with VISAMASTER [provider](/glossary#provider).
    
</dd>
</dl>

<dl>
<dd>

**beneficiary:** `*sdk.CreateWithdrawRequestBeneficiary` 

Beneficiary information.

⚠️ Required if currency [ticker](/glossary#ticker) is one of: UAH_IBAN, USD_VISAMASTER, EUR_VISAMASTER, USD, EUR.

Per-field requirements vary by currency and provider. Card-related fields (`cardToken`, `card.*`, `cardTokenSave`, `fingerprintSession`) apply only to card-acquiring rails; bank-related fields (`bank.*`) apply to bank-rail withdrawals; `tin` is required for UAH_IBAN; `phone`, `email`, and `birthDate` are required for VISAMASTER/Mercuryo rails. See `/asset-status-list` for the active provider per currency.
    
</dd>
</dl>

<dl>
<dd>

**travelRule:** `*sdk.CreateWithdrawRequestTravelRule` 

Travel Rule information for regulatory compliance.

⚠️ Required if currency is crypto and the account is from [EEA](/glossary#european-economic-area-eea)

See [Travel Rule Overview](/api-reference/travel-rule/overview) for complete documentation.

**Legacy format:** The API still accepts the old flat format (`type`, `vasp`, `name`, `address` fields), but this format will not pass Travel Rule verification. To complete Travel Rule compliance, use the new structured format with `walletType`, `beneficiary`, and `vasp` objects.
    
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

**nonce:** `int` — Unique request identifier
    
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
request := &sdk.WithdrawRequest{
        Ticker: "ETH",
        Amount: "0.9",
        Address: "0x0964A6B8F794A4B8d61b62652dB27ddC9844FB4c",
        UniqueID: "24529041",
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**uniqueID:** `string` — Unique transaction identifier. Any string up to 255 characters; not validated as a UUID. ⚠️ Generate a new unique ID for each withdrawal request.
    
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

**customerIP:** `*string` — End-customer IP address forwarded to the [fiat](/glossary#fiat) [provider](/glossary#provider) for antifraud checks before the withdrawal is processed. ⚠️ Required if currency [ticker](/glossary#ticker) is USD or EUR with VISAMASTER [provider](/glossary#provider).
    
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

**nonce:** `int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Withdraw.CreateExpressWithdrawToken(request) -> *sdk.CreateExpressWithdrawTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a signed, single-use Express Withdraw payment token that charges a specific amount from a WhiteBIT user's balance to the partner's [Main balance](/glossary#balance-main) in an instant, off-chain, zero-[fee](/glossary#fee) internal transfer. The response returns a URL that embeds the token; the paying user confirms the exact [ticker](/glossary#ticker) and amount on the WhiteBIT-hosted confirmation surface.

Token and payment constraints:
- Each token is single-use: WhiteBIT marks the token used at confirmation and rejects any replay.
- Each token expires 90 seconds after creation; the `expireAt` response field carries the authoritative expiry timestamp. Generate the token as close as possible to the moment of presenting the URL to the user.
- The [ticker](/glossary#ticker) must be a withdrawal-enabled cryptocurrency; the endpoint rejects [fiat](/glossary#fiat) tickers.
- Each payment is capped at the equivalent of 10,000 USDT; WhiteBIT enforces the cap at token creation and re-enforces the cap at confirmation.
- WhiteBIT rejects self-payments: the paying user and the token creator must be different WhiteBIT accounts.
- The endpoint is idempotent per `externalId`: re-submitting the same `externalId` with an identical `ticker` and `amount` while the token is still valid returns the same token instead of creating a duplicate charge. After the token expires, the same `externalId` receives a fresh token.

<Note>
Standard private-API rate limits apply — see [Rate limits](/api-reference/rate-limits). The endpoint carries no endpoint-specific limit.
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
request := &sdk.CreateExpressWithdrawTokenRequest{
        Ticker: "USDT",
        Amount: "25.50",
        ExternalID: "order-100294",
        Request: "{{request}}",
        Nonce: 1594297865000,
    }
client.Withdraw.CreateExpressWithdrawToken(
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

Currency [ticker](/glossary#ticker) to charge. Example: USDT

⚠️ The ticker must be a withdrawal-enabled cryptocurrency; the endpoint rejects [fiat](/glossary#fiat) tickers. Use [Asset Status endpoint](/public/http-v4/asset-status-list) to check the withdrawal status of a currency.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` 

Amount to charge in the specified [ticker](/glossary#ticker). Numeric string.

⚠️ The amount converted to USDT-equivalent must not exceed 10,000; the endpoint rejects larger amounts with error code `191`.
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `string` — Partner-side reference for the payment (order or invoice identifier), unique per partner account. The identifier powers idempotency and replay protection: a pending `externalId` with an identical `ticker` and `amount` returns the same token; the endpoint rejects an already-paid `externalId` with error code `19`.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` — Unique request identifier
    
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
request := &sdk.TransferBetweenBalancesRequest{
        Method: sdk.TransferBetweenBalancesRequestMethodDeposit.Ptr(),
        Ticker: "XLM",
        Amount: "0.9",
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**method:** `*sdk.TransferBetweenBalancesRequestMethod` 

Transfer method.

⚠️ We highly recommend to use **from** and **to** fields, which provides more flexibility. This way will be deprecated in future.

Example: **deposit** to transfer from [main](/glossary#balance-main) to [trade](/glossary#balance-spotbalance-trade) / **withdraw** to transfer from [trade](/glossary#balance-spotbalance-trade) balance to [main](/glossary#balance-main). For [collateral balances](/glossary#balance-collateral) use **collateral-deposit** to transfer from main to collateral and **collateral-withdraw** to transfer from collateral to main

**Not required** if **from** and **to** are set.
    
</dd>
</dl>

<dl>
<dd>

**from:** `*sdk.TransferBetweenBalancesRequestFrom` 

Balance FROM which funds will move to. Acceptable values: [**main**](/glossary#balance-main), [**spot**](/glossary#balance-spotbalance-trade), [**collateral**](/glossary#balance-collateral)

**Not required** if **method** is set.
    
</dd>
</dl>

<dl>
<dd>

**to:** `*sdk.TransferBetweenBalancesRequestTo` 

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

**nonce:** `int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Codes
<details><summary><code>client.Codes.CreateCode(request) -> *sdk.CreateCodeResponse</code></summary>
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
request := &sdk.CreateCodeRequest{
        Ticker: "ETH",
        Amount: "0.002",
        Passphrase: sdk.String(
            "some passphrase",
        ),
        Description: sdk.String(
            "some description",
        ),
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**amount:** `string` — Amount to transfer. Up to 18 decimal places, value greater than zero and capped at 1e17 (10^17), and not exceeding the [main balance](/glossary#balance-main).
    
</dd>
</dl>

<dl>
<dd>

**passphrase:** `*string` — Passphrase for applying [WhiteBIT codes](/glossary#whitebit-codes). Passphrase must contain only latin letters, numbers and symbols (like !@#$%^, no whitespaces). Max: 25 symbols.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — Additional text description for [code](/glossary#whitebit-codes). Visible only for creator. Max: 280 symbols.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Codes.ApplyCode(request) -> *sdk.ApplyCodeResponse</code></summary>
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

<Note>
To avoid leaking whether a code exists, most failure modes — invalid format, expired,
non-existent, or wrong passphrase — collapse to one generic rejection on field `code`.
Only two cases are distinguishable at the API surface: a code that has already been
applied, and a code created by the same account.
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
request := &sdk.ApplyCodeRequest{
        Code: "WBe11f4fce-2a53-4edc-b195-66b693bd77e3ETH",
        Passphrase: sdk.String(
            "some passphrase",
        ),
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**nonce:** `int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Codes.GetMyCodes(request) -> *sdk.GetMyCodesResponse</code></summary>
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

<Note>
Results are sorted by creation date, newest first. Pagination is capped at `offset + limit ≤ 10000`; for a complete history export beyond the cap, use the Report on the History page.
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
request := &sdk.GetMyCodesRequest{
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**nonce:** `int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Codes.GetCodesHistory(request) -> *sdk.GetCodesHistoryResponse</code></summary>
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

<Note>
Results are sorted by date, newest first.
</Note>

<Note>
**No date filtering:** the endpoint does not accept `startDate` / `endDate` parameters, and pagination is capped at `offset + limit ≤ 10000` (`limit` ≤ 100). For a complete history export beyond the cap, use the Report on the History page.
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
request := &sdk.GetCodesHistoryRequest{
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**nonce:** `int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Fees
<details><summary><code>client.Fees.GetFees(request) -> []*sdk.MainAccountFeeInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an array of objects containing deposit/withdrawal [fees](/glossary#fee) for the corresponding currencies.
Zero value in amount fields means that the setting is disabled.

The endpoint takes no input beyond the signed request envelope and returns the full per-currency fee schedule on success. It can return only the [common authentication errors](/api-reference/authentication).

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
request := &sdk.GetFeesRequest{
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**nonce:** `int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sub-Account
<details><summary><code>client.SubAccount.CreateSubAccount(request) -> *sdk.SubAccount</code></summary>
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

<Note>
Crypto deposits are disabled by default. Once deposits are enabled for the account, the
capability applies to the account and its sub-accounts; enablement is not available via
API. To request it, contact your assigned Account Manager or email institutional@whitebit.com.
Once enabled, a sub-account generates deposit addresses through the standard
[deposit-address endpoint](/api-reference/account-wallet/get-cryptocurrency-deposit-address)
using its own API key with deposit permission.
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
request := &sdk.CreateSubAccountRequest{
        Alias: "trading_bot",
        Permissions: &sdk.CreateSubAccountRequestPermissions{
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

**permissions:** `*sdk.CreateSubAccountRequestPermissions` 
    
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
request := &sdk.DeleteSubAccountRequest{
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
request := &sdk.EditSubAccountRequest{
        ID: "8e667b4a-0b71-4988-8af5-9474dbfaeb51",
        Alias: "training",
        Permissions: &sdk.EditSubAccountRequestPermissions{
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

**permissions:** `*sdk.EditSubAccountRequestPermissions` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccount.ListSubAccounts(request) -> *sdk.ListSubAccountsResponse</code></summary>
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
request := &sdk.ListSubAccountsRequest{}
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

<details><summary><code>client.SubAccount.Transfer(request) -> *sdk.SubAccountTransferResponse</code></summary>
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
request := &sdk.SubAccountTransferRequest{
        ID: "8e667b4a-0b71-4988-8af5-9474dbfaeb51",
        Direction: sdk.SubAccountTransferRequestDirectionMainToSub,
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

**direction:** `*sdk.SubAccountTransferRequestDirection` — Transfer direction
    
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
request := &sdk.BlockSubAccountRequest{
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
request := &sdk.UnblockSubAccountRequest{
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

<details><summary><code>client.SubAccount.GetSubAccountBalances(request) -> map[string][]*sdk.GetSubAccountBalancesResponseValueItem</code></summary>
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
request := &sdk.GetSubAccountBalancesRequest{
        ID: "8e667b4a-0b71-4988-8af5-9474dbfaeb51",
        Ticker: sdk.String(
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

<details><summary><code>client.SubAccount.GetSubAccountTransferHistory(request) -> *sdk.GetSubAccountTransferHistoryResponse</code></summary>
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

<Note>
Results are sorted by transaction id descending (newest transfer first). The response is a plain array with no `total`, `has_more`, or cursor — a returned count below `limit` marks the last page (an empty array means no further records).
</Note>

<Note>
**No date filtering:** the endpoint does not accept `startDate` / `endDate` parameters, and pagination is capped at `offset + limit ≤ 10000` (`limit` ≤ 100). The required `id` parameter already scopes results to a single sub-account; for a complete history export beyond the cap, use the Report on the History page.
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
request := &sdk.GetSubAccountTransferHistoryRequest{
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

**direction:** `*sdk.GetSubAccountTransferHistoryRequestDirection` — Transfer direction (optional)
    
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

<details><summary><code>client.SubAccount.ListUnconfirmedSubAccountWithdrawals(request) -> *sdk.ListUnconfirmedSubAccountWithdrawalsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns a paginated list of withdrawal transactions in `unconfirmed_by_main_account` status,
created by [sub-accounts](/glossary#sub-account) of the authenticated main account and awaiting main account confirmation.
Results are ordered by creation time, newest first.

<Note>
The sub-account withdrawal endpoints are not available by default. To request access, contact institutional@whitebit.com.
A `404 Not Found` response indicates the endpoint is not enabled for the account.
</Note>

<Note>
The sub-account feature must be enabled for the region.
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
request := &sdk.ListUnconfirmedSubAccountWithdrawalsRequest{
        Limit: sdk.Int(
            100,
        ),
        Offset: sdk.Int(
            0,
        ),
    }
client.SubAccount.ListUnconfirmedSubAccountWithdrawals(
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

**limit:** `*int` — Number of records to return.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Number of records to skip.
    
</dd>
</dl>

<dl>
<dd>

**subAccountID:** `*string` — Filter by specific sub-account external ID. If omitted, returns withdrawals from all sub-accounts.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccount.ConfirmSubAccountWithdrawal(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint confirms a single withdrawal transaction created by a [sub-account](/glossary#sub-account)
of the authenticated main account. Confirmation is the main account action that approves a withdrawal
held in `unconfirmed_by_main_account` status and releases it for processing.

Identify the target transaction by its external id, obtained from
[List Unconfirmed Sub-Account Withdrawals](/api-reference/sub-accounts/list-unconfirmed-sub-account-withdrawals).
Confirmation is the only main account action on an unconfirmed withdrawal; a withdrawal left unconfirmed
expires after a retention period. A successful call returns an empty object.

<Note>
The sub-account withdrawal endpoints are not available by default. To request access, contact institutional@whitebit.com.
A `404 Not Found` response indicates the endpoint is not enabled for the account.
</Note>

<Note>
The sub-account feature must be enabled for the region.
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
request := &sdk.ConfirmSubAccountWithdrawalRequest{
        ID: "f47ac10b-58cc-4372-a567-0e02b2c3d479",
    }
client.SubAccount.ConfirmSubAccountWithdrawal(
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

**id:** `string` — External id of the withdrawal transaction to confirm.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubAccount.GetSubAccountKycURL(request) -> *sdk.GetSubAccountKycURLResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint generates a temporary KYC verification link for a [sub-account](/glossary#sub-account).

<Note>
The sub-account must meet all of the following conditions before a KYC URL can be generated:
- The sub-account must be activated (have an associated user).
- The sub-account must be active (not locked or blocked).
- The sub-account must not have shared KYC enabled.
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
request := &sdk.GetSubAccountKycURLRequest{
        ID: "8e667b4a-0b71-4988-8af5-9474dbfaeb51",
    }
client.SubAccount.GetSubAccountKycURL(
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

**id:** `string` — Sub-account external ID. Must belong to the authenticated main account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sub-Account API Keys
<details><summary><code>client.SubAccountAPIKeys.CreateSubAccountAPIKey(request) -> *sdk.SubAccountAPIKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a new API key for a [sub-account](/glossary#sub-account). Each sub-account supports up to 50 API keys, independent from the main account and from other sub-accounts.

<Note>
A `type: 2` key carries deposit and withdrawal permissions, but crypto deposits must also
be enabled for the account. Deposits are disabled by default — to enable them, contact your
assigned Account Manager or email institutional@whitebit.com.
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
request := &sdk.CreateSubAccountAPIKeyRequest{
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
request := &sdk.EditSubAccountAPIKeyRequest{
        APIKeyID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
        Title: "Trading Bot Key",
        URLs: []*sdk.EditSubAccountAPIKeyRequestURLsItem{
            &sdk.EditSubAccountAPIKeyRequestURLsItem{
                URL: sdk.String(
                    "/api/v4/main-account/withdraw",
                ),
                Enable: sdk.Bool(
                    false,
                ),
            },
            &sdk.EditSubAccountAPIKeyRequestURLsItem{
                URL: sdk.String(
                    "/api/v4/main-account/balance",
                ),
                Enable: sdk.Bool(
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

**urls:** `[]*sdk.EditSubAccountAPIKeyRequestURLsItem` — Array of URL objects for API key restrictions
    
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
request := &sdk.DeleteSubAccountAPIKeyRequest{
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

<details><summary><code>client.SubAccountAPIKeys.ListSubAccountAPIKeys(request) -> *sdk.ListSubAccountAPIKeysResponse</code></summary>
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

<Note>
Results are sorted by api-key id descending (newest key first). The response is a plain array with no `total`, `has_more`, or cursor — a returned count below `limit` marks the last page (an empty array means no further records).
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
request := &sdk.ListSubAccountAPIKeysRequest{}
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
request := &sdk.ResetSubAccountAPIKeyRequest{
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

<details><summary><code>client.SubAccountAPIKeys.ListSubAccountAPIKeyIPAddresses(request) -> *sdk.ListSubAccountAPIKeyIPAddressesResponse</code></summary>
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
request := &sdk.ListSubAccountAPIKeyIPAddressesRequest{
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

<details><summary><code>client.SubAccountAPIKeys.CreateSubAccountAPIKeyIPAddress(request) -> *sdk.CreateSubAccountAPIKeyIPAddressResponse</code></summary>
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
request := &sdk.CreateSubAccountAPIKeyIPAddressRequest{
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

<details><summary><code>client.SubAccountAPIKeys.DeleteSubAccountAPIKeyIPAddress(request) -> *sdk.DeleteSubAccountAPIKeyIPAddressResponse</code></summary>
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
request := &sdk.DeleteSubAccountAPIKeyIPAddressRequest{
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

## Credit Line
<details><summary><code>client.CreditLine.GetCreditLineInfo(request) -> *sdk.CreditLine</code></summary>
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
request := &sdk.GetCreditLineInfoRequest{
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**nonce:** `int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Travel Rule
<details><summary><code>client.TravelRule.GetTravelRuleVasps(request) -> *sdk.GetTravelRuleVaspsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the list of Virtual Asset Service Providers (VASPs) that can be used when submitting travel rule data for deposits and withdrawals.

Use the returned `vaspId` when submitting travel rule data if the destination/originating VASP is in this list. If the VASP is not in the list, use `vaspName` as a fallback.

<Accordion title="Errors">
**API disabled for region (422):**
```json
{
  "code": 0,
  "message": "Something went wrong",
  "errors": {
    "error": ["This endpoint is disabled for your region"]
  }
}
```
</Accordion>

Beyond the errors above, this endpoint can return only the [common authentication errors](/api-reference/authentication).

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
request := &sdk.GetTravelRuleVaspsRequest{
        Request: "{{request}}",
        Nonce: 1594297865000,
    }
client.TravelRule.GetTravelRuleVasps(
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

**nonce:** `int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TravelRule.SubmitTravelRuleDepositVerification(request) -> *sdk.SubmitTravelRuleDepositVerificationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submits originator information for a deposit that is ready for the travel rule verification.

**Wallet types:**
- `hosted` - VASP-hosted wallet (exchange, custodian). Requires `vaspData` object.
- `unhosted` - Self-custody wallet (hardware wallet, software wallet). No VASP required.

**Party types:**
- `individual` - Natural person. Requires `firstName` and `lastName`.
- `entity` - Legal entity. Requires `fullName`.

<Accordion title="Errors">
**Validation failed (400):**
```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "originator.firstName": ["The first name field is required for individual."],
    "originator.address": ["The originator address field is required."]
  }
}
```

**Transaction not found (422):**
```json
{
  "code": 0,
  "message": "Something went wrong",
  "errors": {
    "uniqueId": ["Transaction not found"]
  }
}
```

**Transaction must be deposit (422):**
```json
{
  "code": 0,
  "message": "Something went wrong",
  "errors": {
    "uniqueId": ["Transaction must be a deposit"]
  }
}
```

**Transaction not ready for verification (422):**
```json
{
  "code": 0,
  "message": "Something went wrong",
  "errors": {
    "uniqueId": ["Transaction is not ready for verification"]
  }
}
```

**API disabled for region (422):**
```json
{
  "code": 0,
  "message": "Something went wrong",
  "errors": {
    "error": ["This endpoint is disabled for your region"]
  }
}
```

**Missing VASP for hosted wallet (400):**
```json
{
  "code": 0,
  "message": "Validation failed",
  "errors": {
    "vaspData": ["Either vaspId or vaspName is required for hosted wallets"]
  }
}
```
</Accordion>

Beyond the errors above, this endpoint can return only the [common authentication errors](/api-reference/authentication).

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
request := &sdk.TravelRuleDepositVerificationRequest{
        UniqueID: "550e8400-e29b-41d4-a716-446655440000",
        WalletType: sdk.TravelRuleDepositVerificationRequestWalletTypeHosted,
        Originator: &sdk.TravelRuleOriginator{
            Type: sdk.TravelRuleOriginatorTypeIndividual,
            FirstName: sdk.String(
                "Alice",
            ),
            LastName: sdk.String(
                "Johnson",
            ),
            ResidenceCountry: "NLD",
            WalletAddress: "0x9876543210fedcba9876543210fedcba98765432",
            Address: &sdk.TravelRuleAddress{
                Country: "NLD",
                City: "Amsterdam",
                AddressLine1: "Damrak 1",
            },
        },
        VaspData: &sdk.TravelRuleVasp{
            VaspID: sdk.String(
                "vasp-002",
            ),
        },
        Request: "{{request}}",
        Nonce: 1594297865000,
    }
client.TravelRule.SubmitTravelRuleDepositVerification(
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

**uniqueID:** `string` — Transaction external ID (from deposit/withdraw history)
    
</dd>
</dl>

<dl>
<dd>

**walletType:** `*sdk.TravelRuleDepositVerificationRequestWalletType` 

Wallet type:
- `hosted` - VASP-hosted wallet (exchange, custodian). Requires `vaspData` object.
- `unhosted` - Self-custody wallet (hardware wallet, software wallet).
    
</dd>
</dl>

<dl>
<dd>

**originator:** `*sdk.TravelRuleOriginator` 
    
</dd>
</dl>

<dl>
<dd>

**vaspData:** `*sdk.TravelRuleVasp` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` — Request signature
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` — Unique request identifier
    
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

The endpoint returns the current [collateral balance](/glossary#balance-collateral) for one or all assets. The response maps each asset ticker to its collateral balance amount. Use the optional `ticker` parameter to filter results to a single asset.

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
request := &sdk.CollateralAccountBalanceRequest{
        Ticker: sdk.String(
            "BTC",
        ),
        Request: sdk.String(
            "{{request}}",
        ),
        Nonce: sdk.Int(
            1594297865000,
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

**nonce:** `*int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CollateralAccountBalanceSummary(request) -> []*sdk.CollateralAccountBalanceSummaryResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns a detailed [collateral balance](/glossary#balance-collateral) summary with a per-asset breakdown. Each record includes the current balance, borrowed amount, and available balance with and without borrowing capacity. Use the optional `ticker` parameter to filter results to a single asset.

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
request := &sdk.CollateralAccountBalanceSummaryRequest{
        Ticker: sdk.String(
            "BTC",
        ),
        Request: sdk.String(
            "{{request}}",
        ),
        Nonce: sdk.Int(
            1594297865000,
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

**nonce:** `*int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CreateCollateralLimitOrder(request) -> *sdk.CreateCollateralLimitOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a [limit order](/glossary#limit-order) using [collateral balance](/glossary#balance-collateral). The order executes at the specified price or better. Use `buy` to open or increase a long position and `sell` to open or increase a short position. To close a position, place an opposite-side order matching the position amount.

**Order validation rules** (per-market, from `GET /api/v4/public/markets`):
- `amount` must have at most `stockPrec` decimal places
- `price` must have at most `moneyPrec` decimal places
- `amount` must be ≥ `minAmount`
- `amount × price` must be ≥ `minTotal`
- `amount × price` must be ≤ `maxTotal` (when `maxTotal` is not `"0"`)

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Note>
For open long position use **buy**, for short **sell**. To close current position, place opposite order with current position amount.
</Note>

<Note>
  - RPI orders are post-only by design and cannot be used with the IOC flag. The API returns error code `40` when both `rpi=true` and `ioc=true` are used.
</Note>


<Accordion title="Error Codes">
  - `30` - default validation error code. Also returned when `reduceOnly=true` is combined with `stopLoss` or `takeProfit`
  - `31` - market validation failed
  - `32` - amount validation failed
  - `33` - price validation failed
  - `36` - clientOrderId validation failed
  - `37` - `ioc=true` cannot be combined with `postOnly=true`
  - `40` - `ioc=true` cannot be combined with `rpi=true`
  - `43` - `rpi=true` is not allowed for the account
  - `10` - insufficient balance to place the order
  - `111` - resulting position would exceed the market maximum
  - `112` - pending orders value would exceed the allowed maximum
  - `113` - position side cannot be changed while open positions or orders exist
  - `114` - hedge mode position side does not match (sent `BOTH` or omitted `positionSide` in hedge mode, or sent `LONG`/`SHORT` in one-way mode)
  - `115` - order would open a position in the opposite direction (one-way mode)
  - `116` - reduce-only validation failed (no position exists or order side matches position direction)
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
request := &sdk.CreateCollateralLimitOrderRequest{
        Market: "BTC_USDT",
        Side: sdk.CreateCollateralLimitOrderRequestSideBuy,
        Amount: "0.01",
        Price: "40000",
        ClientOrderID: sdk.String(
            "order1987111",
        ),
        StopLoss: sdk.String(
            "50000",
        ),
        TakeProfit: sdk.String(
            "30000",
        ),
        PostOnly: sdk.Bool(
            false,
        ),
        Ioc: sdk.Bool(
            false,
        ),
        Rpi: sdk.Bool(
            true,
        ),
        PositionSide: sdk.CreateCollateralLimitOrderRequestPositionSideLong.Ptr(),
        ReduceOnly: sdk.Bool(
            false,
        ),
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**side:** `*sdk.CreateCollateralLimitOrderRequestSide` — Order type. Variables: 'buy' / 'sell'. For open long position use **buy**, for short **sell**.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Amount of [stock](/glossary#stock) currency to buy or sell. Minimum and step values are market-dependent — query the [market info](/api-reference/market-data/market-info) endpoint for constraints.
    
</dd>
</dl>

<dl>
<dd>

**price:** `string` — Limit order price in [money](/glossary#money) currency. Minimum price step is market-dependent — query the [market info](/api-reference/market-data/market-info) endpoint for constraints.
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Custom client order identifier. Uniqueness is enforced only among the account's open (pending) orders on the same market — once a previous order is filled or canceled, the same identifier can be reused, including on the same market. Contains only letters, numbers, dashes, dots, or underscores.
    
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

**postOnly:** `*bool` — When `true`, guarantees the order executes as a [maker](/glossary#maker) order. The system rejects the order if it would immediately match as taker. Default: `false`.
    
</dd>
</dl>

<dl>
<dd>

**ioc:** `*bool` — When `true`, the order executes all or part immediately and cancels any unfilled portion. Cannot be combined with `postOnly=true` or `rpi=true`.
    
</dd>
</dl>

<dl>
<dd>

**rpi:** `*bool` 

Enables Retail Price Improvement (RPI) mode.

RPI orders are post-only by design and cannot be used with `ioc=true`. The API returns error code `40` when both `rpi=true` and `ioc=true` are used.
    
</dd>
</dl>

<dl>
<dd>

**positionSide:** `*sdk.CreateCollateralLimitOrderRequestPositionSide` 

Position direction. Optional at the request layer but functionally required when hedge mode is enabled. See [positionSide](/glossary#position-side).

- **One-way mode** (default account mode): the field is ignored. Orders always use `BOTH`, and the response returns `positionSide: "BOTH"` whether the field is sent or omitted.
- **Hedge mode**: the field MUST be `LONG` or `SHORT`. Sending `BOTH`, omitting the field, or sending a value that does not match the account's mode causes the trade service to reject the order with error code `114` (`Hedge mode position side does not match`).
    
</dd>
</dl>

<dl>
<dd>

**reduceOnly:** `*bool` — When `true`, the order can only reduce or close an existing position — the order cannot increase the position or open a new one. If the order amount exceeds the current position size, the system reduces the order to match — the response returns the adjusted amount. Cannot be combined with `stopLoss` or `takeProfit`. The API returns error code `116` if no open position exists or the order side matches the position direction. See [reduce-only](/glossary#reduce-only).
    
</dd>
</dl>

<dl>
<dd>

**stp:** `*sdk.CreateCollateralLimitOrderRequestStp` 

Self-trade prevention mode. Allowed values: `no` (self-trades allowed), `cb` (cancel both the new and the existing order), `cn` (cancel the new order, keep the existing), `co` (cancel the existing order, place the new one). Default: `no`.

Legacy values `cancel_both`, `cancel_new`, `cancel_old` are deprecated: the API accepts the legacy values with identical behavior until a deprecation deadline is announced, then rejects the legacy values. Responses always return the abbreviated form, regardless of which variant the request used.

See [Self-Trade Prevention](/platform/self-trade-prevention).
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CreateCollateralBulkOrder(request) -> []*sdk.CreateCollateralBulkOrderResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates multiple collateral [limit orders](/glossary#limit-order) in a single request. Each order in the `orders` array is validated and processed individually. The `stopOnFail` parameter controls whether processing stops at the first failure or continues through all orders. The response array contains a result or error object for each submitted order, in the same order as the request.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
  - `30` - default validation error code (per-order). Also returned when `reduceOnly=true` is combined with `stopLoss` or `takeProfit`
  - `31` - market validation failed
  - `32` - amount validation failed
  - `33` - price validation failed
  - `36` - clientOrderId validation failed
  - `37` - `ioc=true` cannot be used with `postOnly=true` or `rpi=true`
  - `10` - insufficient balance to place the order
  - `111` - resulting position would exceed the market maximum
  - `112` - pending orders value would exceed the allowed maximum
  - `113` - position side cannot be changed while open positions or orders exist
  - `114` - hedge mode position side does not match (per-order; sent `BOTH` or omitted `positionSide` in hedge mode, or sent `LONG`/`SHORT` in one-way mode)
  - `115` - order would open a position in the opposite direction (one-way mode)
  - `116` - reduce-only validation failed (no position exists or order side matches position direction). For bulk orders, this error appears per-order inside the response array.
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
request := &sdk.CreateCollateralBulkOrderRequest{
        Orders: []*sdk.CreateCollateralBulkOrderRequestOrdersItem{
            &sdk.CreateCollateralBulkOrderRequestOrdersItem{
                Market: sdk.String(
                    "BTC_PERP",
                ),
                Side: sdk.CreateCollateralBulkOrderRequestOrdersItemSideBuy.Ptr(),
                Amount: sdk.String(
                    "0.02",
                ),
                Price: sdk.String(
                    "40000",
                ),
                ClientOrderID: sdk.String(
                    "",
                ),
                PostOnly: sdk.Bool(
                    false,
                ),
                Ioc: sdk.Bool(
                    false,
                ),
                Rpi: sdk.Bool(
                    true,
                ),
                PositionSide: sdk.CreateCollateralBulkOrderRequestOrdersItemPositionSideLong.Ptr(),
                ReduceOnly: sdk.Bool(
                    false,
                ),
            },
            &sdk.CreateCollateralBulkOrderRequestOrdersItem{
                Market: sdk.String(
                    "BTC_USDT",
                ),
                Side: sdk.CreateCollateralBulkOrderRequestOrdersItemSideSell.Ptr(),
                Amount: sdk.String(
                    "0.0001",
                ),
                Price: sdk.String(
                    "41000",
                ),
                ClientOrderID: sdk.String(
                    "",
                ),
                PostOnly: sdk.Bool(
                    false,
                ),
                Ioc: sdk.Bool(
                    false,
                ),
                Rpi: sdk.Bool(
                    true,
                ),
                PositionSide: sdk.CreateCollateralBulkOrderRequestOrdersItemPositionSideLong.Ptr(),
                ReduceOnly: sdk.Bool(
                    true,
                ),
            },
            &sdk.CreateCollateralBulkOrderRequestOrdersItem{
                Market: sdk.String(
                    "ETH_BTC",
                ),
                Side: sdk.CreateCollateralBulkOrderRequestOrdersItemSideSell.Ptr(),
                Amount: sdk.String(
                    "0.02",
                ),
                Price: sdk.String(
                    "0.030",
                ),
                ClientOrderID: sdk.String(
                    "",
                ),
                PostOnly: sdk.Bool(
                    false,
                ),
                Ioc: sdk.Bool(
                    false,
                ),
                Rpi: sdk.Bool(
                    true,
                ),
                PositionSide: sdk.CreateCollateralBulkOrderRequestOrdersItemPositionSideLong.Ptr(),
                ReduceOnly: sdk.Bool(
                    false,
                ),
            },
        },
        StopOnFail: sdk.Bool(
            true,
        ),
        Request: sdk.String(
            "{{request}}",
        ),
        Nonce: sdk.Int(
            1594297865000,
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

**orders:** `[]*sdk.CreateCollateralBulkOrderRequestOrdersItem` 
    
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

**nonce:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CreateCollateralMarketOrder(request) -> *sdk.CreateCollateralMarketOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a [market order](/glossary#market-order) using [collateral balance](/glossary#balance-collateral). The order executes immediately at the best available market price. Optionally attach `stopLoss` and `takeProfit` prices to create an [OTO](/glossary#one-triggers-the-other-oto) order that activates after the market order fills.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
  - `30` - default validation error code. Also returned when `reduceOnly=true` is combined with `stopLoss` or `takeProfit`
  - `31` - market validation failed
  - `32` - amount validation failed
  - `36` - clientOrderId validation failed
  - `10` - insufficient balance to place the order
  - `111` - resulting position would exceed the market maximum
  - `112` - pending orders value would exceed the allowed maximum
  - `113` - position side cannot be changed while open positions or orders exist
  - `114` - hedge mode position side does not match (sent `BOTH` or omitted `positionSide` in hedge mode, or sent `LONG`/`SHORT` in one-way mode)
  - `115` - order would open a position in the opposite direction (one-way mode)
  - `116` - reduce-only validation failed (no position exists or order side matches position direction)
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
request := &sdk.CreateCollateralMarketOrderRequest{
        Market: "BTC_USDT",
        Side: sdk.CreateCollateralMarketOrderRequestSideBuy,
        Amount: "0.01",
        ClientOrderID: sdk.String(
            "order1987111",
        ),
        ReduceOnly: sdk.Bool(
            false,
        ),
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**market:** `string` — Available margin [market](/glossary#market). Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**side:** `*sdk.CreateCollateralMarketOrderRequestSide` — Order direction. Use `buy` to open or increase a long position and `sell` to open or increase a short position.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Amount of [stock](/glossary#stock) currency to buy or sell. Minimum and step values are market-dependent — query the [market info](/api-reference/market-data/market-info) endpoint for constraints.
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Custom client order identifier. Uniqueness is enforced only among the account's open (pending) orders on the same market — once a previous order is filled or canceled, the same identifier can be reused, including on the same market. Contains only letters, numbers, dashes, dots, or underscores.
    
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

**positionSide:** `*sdk.CreateCollateralMarketOrderRequestPositionSide` 

Position direction. Optional at the request layer but functionally required when hedge mode is enabled. See [positionSide](/glossary#position-side).

- **One-way mode** (default account mode): the field is ignored. Orders always use `BOTH`, and the response returns `positionSide: "BOTH"` whether the field is sent or omitted.
- **Hedge mode**: the field MUST be `LONG` or `SHORT`. Sending `BOTH`, omitting the field, or sending a value that does not match the account's mode causes the trade service to reject the order with error code `114` (`Hedge mode position side does not match`).
    
</dd>
</dl>

<dl>
<dd>

**reduceOnly:** `*bool` — When `true`, the order can only reduce or close an existing position — the order cannot increase the position or open a new one. If the order amount exceeds the current position size, the system reduces the order to match — the response returns the adjusted amount. Cannot be combined with `stopLoss` or `takeProfit`. The API returns error code `116` if no open position exists or the order side matches the position direction. See [reduce-only](/glossary#reduce-only).
    
</dd>
</dl>

<dl>
<dd>

**stp:** `*sdk.CreateCollateralMarketOrderRequestStp` 

Self-trade prevention mode. Allowed values: `no` (self-trades allowed), `cb` (cancel both the new and the existing order), `cn` (cancel the new order, keep the existing), `co` (cancel the existing order, place the new one). Default: `no`.

Legacy values `cancel_both`, `cancel_new`, `cancel_old` are deprecated: the API accepts the legacy values with identical behavior until a deprecation deadline is announced, then rejects the legacy values. Responses always return the abbreviated form, regardless of which variant the request used.

See [Self-Trade Prevention](/platform/self-trade-prevention).
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CreateCollateralStopLimitOrder(request) -> *sdk.CreateCollateralStopLimitOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a collateral [stop-limit order](/glossary#stop-limit-order) using [collateral balance](/glossary#balance-collateral). The order remains inactive until the market price reaches `activation_price`, then places a limit order at `price`. Optionally attach `stopLoss` and `takeProfit` prices to create an [OTO](/glossary#one-triggers-the-other-oto) order that activates after the stop-limit order fills.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
  - `30` - default validation error code. Also returned when `reduceOnly=true` is combined with `stopLoss` or `takeProfit`
  - `31` - market validation failed
  - `32` - amount validation failed
  - `33` - price validation failed
  - `36` - clientOrderId validation failed
  - `10` - insufficient balance to place the order
  - `111` - resulting position would exceed the market maximum
  - `112` - pending orders value would exceed the allowed maximum
  - `113` - position side cannot be changed while open positions or orders exist
  - `114` - hedge mode position side does not match (sent `BOTH` or omitted `positionSide` in hedge mode, or sent `LONG`/`SHORT` in one-way mode)
  - `115` - order would open a position in the opposite direction (one-way mode)
  - `116` - reduce-only validation failed (no position exists or order side matches position direction)
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
request := &sdk.CreateCollateralStopLimitOrderRequest{
        Market: "BTC_USDT",
        Side: sdk.CreateCollateralStopLimitOrderRequestSideBuy,
        Amount: "0.001",
        Price: "40000",
        ActivationPrice: "40000",
        StopLoss: sdk.String(
            "30000",
        ),
        TakeProfit: sdk.String(
            "50000",
        ),
        ClientOrderID: sdk.String(
            "order1987111",
        ),
        PositionSide: sdk.CreateCollateralStopLimitOrderRequestPositionSideLong.Ptr(),
        ReduceOnly: sdk.Bool(
            false,
        ),
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**market:** `string` — Available margin [market](/glossary#market). Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**side:** `*sdk.CreateCollateralStopLimitOrderRequestSide` — Order direction. Use `buy` to open or increase a long position and `sell` to open or increase a short position.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Amount of [stock](/glossary#stock) currency to buy or sell. Minimum and step values are market-dependent — query the [market info](/api-reference/market-data/market-info) endpoint for constraints.
    
</dd>
</dl>

<dl>
<dd>

**price:** `string` — Limit order price in [money](/glossary#money) currency. The order executes at the specified price or better after activation.
    
</dd>
</dl>

<dl>
<dd>

**activationPrice:** `string` — Trigger price in [money](/glossary#money) currency. The stop-limit order activates when the market price reaches the specified value.
    
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

**clientOrderID:** `*string` — Custom client order identifier. Uniqueness is enforced only among the account's open (pending) orders on the same market — once a previous order is filled or canceled, the same identifier can be reused, including on the same market. Contains only letters, numbers, dashes, dots, or underscores.
    
</dd>
</dl>

<dl>
<dd>

**positionSide:** `*sdk.CreateCollateralStopLimitOrderRequestPositionSide` 

Position direction. Optional at the request layer but functionally required when hedge mode is enabled. See [positionSide](/glossary#position-side).

- **One-way mode** (default account mode): the field is ignored. Orders always use `BOTH`, and the response returns `positionSide: "BOTH"` whether the field is sent or omitted.
- **Hedge mode**: the field MUST be `LONG` or `SHORT`. Sending `BOTH`, omitting the field, or sending a value that does not match the account's mode causes the trade service to reject the order with error code `114` (`Hedge mode position side does not match`).
    
</dd>
</dl>

<dl>
<dd>

**reduceOnly:** `*bool` — When `true`, the order can only reduce or close an existing position — the order cannot increase the position or open a new one. If the order amount exceeds the current position size, the system reduces the order to match — the response returns the adjusted amount. Cannot be combined with `stopLoss` or `takeProfit`. The API returns error code `116` if no open position exists or the order side matches the position direction. See [reduce-only](/glossary#reduce-only).
    
</dd>
</dl>

<dl>
<dd>

**stp:** `*sdk.CreateCollateralStopLimitOrderRequestStp` 

Self-trade prevention mode. Allowed values: `no` (self-trades allowed), `cb` (cancel both the new and the existing order), `cn` (cancel the new order, keep the existing), `co` (cancel the existing order, place the new one). Default: `no`.

Legacy values `cancel_both`, `cancel_new`, `cancel_old` are deprecated: the API accepts the legacy values with identical behavior until a deprecation deadline is announced, then rejects the legacy values. Responses always return the abbreviated form, regardless of which variant the request used.

See [Self-Trade Prevention](/platform/self-trade-prevention).
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CreateCollateralTriggerMarketOrder(request) -> *sdk.CreateCollateralTriggerMarketOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a collateral trigger [market order](/glossary#market-order) using [collateral balance](/glossary#balance-collateral). The order remains inactive until the market price reaches `activation_price`, then executes immediately at the best available market price. Optionally attach `stopLoss` and `takeProfit` prices to create an [OTO](/glossary#one-triggers-the-other-oto) order that activates after the trigger market order fills.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
  - `30` - default validation error code. Also returned when `reduceOnly=true` is combined with `stopLoss` or `takeProfit`
  - `31` - market validation failed
  - `32` - amount validation failed
  - `36` - clientOrderId validation failed
  - `10` - insufficient balance to place the order
  - `111` - resulting position would exceed the market maximum
  - `112` - pending orders value would exceed the allowed maximum
  - `113` - position side cannot be changed while open positions or orders exist
  - `114` - hedge mode position side does not match (sent `BOTH` or omitted `positionSide` in hedge mode, or sent `LONG`/`SHORT` in one-way mode)
  - `115` - order would open a position in the opposite direction (one-way mode)
  - `116` - reduce-only validation failed (no position exists or order side matches position direction)
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
request := &sdk.CreateCollateralTriggerMarketOrderRequest{
        Market: "BTC_USDT",
        Side: sdk.CreateCollateralTriggerMarketOrderRequestSideBuy,
        Amount: "0.01",
        ActivationPrice: "40000",
        ClientOrderID: sdk.String(
            "order1987111",
        ),
        ReduceOnly: sdk.Bool(
            false,
        ),
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**market:** `string` — Available margin [market](/glossary#market). Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**side:** `*sdk.CreateCollateralTriggerMarketOrderRequestSide` — Order direction. Use `buy` to open or increase a long position and `sell` to open or increase a short position.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Amount of [stock](/glossary#stock) currency to buy or sell. Minimum and step values are market-dependent — query the [market info](/api-reference/market-data/market-info) endpoint for constraints.
    
</dd>
</dl>

<dl>
<dd>

**activationPrice:** `string` — Trigger price in [money](/glossary#money) currency. The trigger market order activates when the market price reaches the specified value.
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Custom client order identifier. Uniqueness is enforced only among the account's open (pending) orders on the same market — once a previous order is filled or canceled, the same identifier can be reused, including on the same market. Contains only letters, numbers, dashes, dots, or underscores.
    
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

**positionSide:** `*sdk.CreateCollateralTriggerMarketOrderRequestPositionSide` 

Position direction. Optional at the request layer but functionally required when hedge mode is enabled. See [positionSide](/glossary#position-side).

- **One-way mode** (default account mode): the field is ignored. Orders always use `BOTH`, and the response returns `positionSide: "BOTH"` whether the field is sent or omitted.
- **Hedge mode**: the field MUST be `LONG` or `SHORT`. Sending `BOTH`, omitting the field, or sending a value that does not match the account's mode causes the trade service to reject the order with error code `114` (`Hedge mode position side does not match`).
    
</dd>
</dl>

<dl>
<dd>

**reduceOnly:** `*bool` — When `true`, the order can only reduce or close an existing position — the order cannot increase the position or open a new one. If the order amount exceeds the current position size, the system reduces the order to match — the response returns the adjusted amount. Cannot be combined with `stopLoss` or `takeProfit`. The API returns error code `116` if no open position exists or the order side matches the position direction. See [reduce-only](/glossary#reduce-only).
    
</dd>
</dl>

<dl>
<dd>

**stp:** `*sdk.CreateCollateralTriggerMarketOrderRequestStp` 

Self-trade prevention mode. Allowed values: `no` (self-trades allowed), `cb` (cancel both the new and the existing order), `cn` (cancel the new order, keep the existing), `co` (cancel the existing order, place the new one). Default: `no`.

Legacy values `cancel_both`, `cancel_new`, `cancel_old` are deprecated: the API accepts the legacy values with identical behavior until a deprecation deadline is announced, then rejects the legacy values. Responses always return the abbreviated form, regardless of which variant the request used.

See [Self-Trade Prevention](/platform/self-trade-prevention).
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CollateralAccountSummary(request) -> *sdk.CollateralAccountSummaryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns a collateral account summary including total equity, used margin, free margin, unrealized profit and loss, and the current leverage level. The `marginFraction` field indicates the ratio of used margin to total equity.

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
request := &sdk.CollateralAccountSummaryRequest{
        Request: sdk.String(
            "{{request}}",
        ),
        Nonce: sdk.Int(
            1594297865000,
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

**nonce:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.GetOpenPositions(request) -> []*sdk.GetOpenPositionsResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns all open [collateral](/glossary#balance-collateral) positions for the authenticated account. Each position includes entry price, unrealized PnL, margin allocation, liquidation price, and take-profit/stop-loss configuration. Use the optional `market` parameter to filter results to a single trading pair.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
  - `30` - default validation error code (returned when the optional `market` filter is malformed)
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
request := &sdk.GetOpenPositionsRequest{
        Market: sdk.String(
            "BTC_USDT",
        ),
        Request: sdk.String(
            "{{request}}",
        ),
        Nonce: sdk.Int(
            1594297865000,
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

**nonce:** `*int` — Unique request identifier
    
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

The endpoint closes an open [collateral](/glossary#balance-collateral) position at the current market price. The system places a market order in the opposite direction to fully close the specified position. Any attached take-profit or stop-loss orders are cancelled automatically.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
  - `30` - default validation error code (for example, a missing or malformed `positionId` or `market`)
  - `104` - position not found. Returned whether the `positionId` does not exist, the position is already closed, or it is not owned by the account — these cases are not distinguished
  - `10` - insufficient balance to fund the closing market order
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
request := &sdk.ClosePositionRequest{
        PositionID: 123,
        PositionSide: sdk.ClosePositionRequestPositionSideLong.Ptr(),
        Market: "BTC_USDT",
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**positionID:** `int` — Unique identifier of the position to close. Obtain from the [open positions](/api-reference/collateral-trading/open-positions) endpoint.
    
</dd>
</dl>

<dl>
<dd>

**positionSide:** `*sdk.ClosePositionRequestPositionSide` — Defines the position direction when hedge mode is enabled. See [positionSide](/glossary#position-side)
    
</dd>
</dl>

<dl>
<dd>

**market:** `string` — Market of the position to close. Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.GetPositionsHistory(request) -> []*sdk.GetPositionsHistoryResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns the history of [collateral](/glossary#balance-collateral) position state changes for the authenticated account. Each record represents a position event (open, partial close, full close, or liquidation) and includes the order details that triggered the change. Use the optional `market` and `positionId` parameters to filter results.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
  - `30` - default validation error code (invalid pagination — `limit` outside 1–100 or negative `offset` — or a date filter that violates `startDate` ≤ `endDate` ≤ `now + 1s`)
</Accordion>

<Note>
**Date filter window:** `startDate` and `endDate` are optional and have no defaults. The endpoint enforces no maximum window and no lower-bound floor. The only ordering constraint is `startDate` ≤ `endDate` ≤ `now + 1s` — requests that violate the ordering are rejected with a validation error.
</Note>

<Warning>
**Breaking change — April 29, 2026.** The `positionSide` field is no longer returned in the Position History response. Use `side` (same enum: `LONG`, `SHORT`, `BOTH`) plus `isHedge` (boolean) instead. Integrations reading `positionSide` from `/api/v4/collateral-account/positions/history` must migrate before consuming the new response.
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
request := &sdk.GetPositionsHistoryRequest{
        Market: sdk.String(
            "BTC_USDT",
        ),
        PositionID: sdk.Int(
            1,
        ),
        Request: sdk.String(
            "{{request}}",
        ),
        Nonce: sdk.Int(
            1594297865000,
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

Filter by specific market. Example: BTC_USDT

If not specified, returns position history for all markets.
    
</dd>
</dl>

<dl>
<dd>

**positionID:** `*int` — Filter by specific position identifier. If not specified, returns history for all positions.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*int` — Start of the query window as a Unix timestamp in seconds. Optional, no default. Must be ≤ `endDate`.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*int` — End of the query window as a Unix timestamp in seconds. Optional, no default. Must be ≥ `startDate` and ≤ `now + 1s`; violating values are rejected with a validation error.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.GetFundingHistory(request) -> *sdk.GetFundingHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns the funding rate payment history for [collateral](/glossary#balance-collateral) positions. Each record includes the funding rate, settlement price, position amount, and the resulting funding payment. Use the optional `market` parameter to filter results to a single trading pair. The response supports pagination via `limit` and `offset` parameters. Results are ordered by funding time (`fundingTime`), newest first.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>

<Note>
This endpoint supports pagination. Use `limit` (default: 100) and `offset` (default: 0) to page through results. The response does not include a `total` field — detect the last page when `records.length < limit`. An empty `records` array means you have paged past the end; receiving exactly `limit` records does not guarantee that another page exists.
</Note>

<Accordion title="Error Codes">
  - `30` - default validation error code (invalid pagination — `limit` outside 1–100 or negative `offset`)
  - `31` - market validation failed (the `market` filter is unknown or not available for collateral trading)
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
request := &sdk.GetFundingHistoryRequest{
        Market: sdk.String(
            "BTC_PERP",
        ),
        Limit: sdk.Int(
            100,
        ),
        Offset: sdk.Int(
            0,
        ),
        Request: sdk.String(
            "{{request}}",
        ),
        Nonce: sdk.Int(
            1594297865000,
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

**nonce:** `*int` — Unique request identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.ChangeCollateralAccountLeverage(request) -> *sdk.ChangeCollateralAccountLeverageResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint changes the leverage level for the [collateral](/glossary#balance-collateral) trading account. Leverage determines the ratio of borrowed funds to collateral and directly affects margin requirements and liquidation thresholds. Accepted values: `1`, `2`, `3`, `5`, `10`, `20`, `50`, `100`.

Each leverage level has a corresponding bracket defining the maximum position size for the tier. When a position exceeds the bracket limit, the system applies higher tiers with progressively lower leverage. Query market-specific brackets via `GET /api/v4/public/futures`.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Note>
A market's `max_leverage` field (from `GET /api/v4/public/futures`) may be lower than `100`. Setting leverage above a market's maximum results in an error.
</Note>

<Warning>
Changing leverage affects **all open positions** across margin and futures trading. Decreasing leverage increases margin requirements — if available funds are insufficient to support the new level, the request returns an error.
</Warning>

<Accordion title="Error Codes">
  - `30` - invalid `leverage` value (out of range, non-integer, or wrong type). Setting leverage above a market's `max_leverage` also surfaces here as an out-of-range value
  - `17` - the requested leverage is valid but available balance is insufficient to support it
  - `113` - leverage cannot be changed while open positions or orders exist
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
request := &sdk.ChangeCollateralAccountLeverageRequest{
        Leverage: 5,
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**leverage:** `int` — Target leverage level. Accepted values: `1`, `2`, `3`, `5`, `10`, `20`, `50`, `100`. The effective maximum depends on the market's `max_leverage`.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.GetCollateralHedgeMode(request) -> *sdk.GetCollateralHedgeModeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns the current [hedge mode](/glossary#hedge-mode) status for the collateral trading account. When hedge mode is enabled (`true`), the account supports simultaneous long and short positions on the same market. When disabled (`false`), the account operates in one-way mode.

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
request := &sdk.GetCollateralHedgeModeRequest{
        Request: sdk.String(
            "{{request}}",
        ),
        Nonce: sdk.Int(
            1594297865000,
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

**nonce:** `*int` 
    
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

The endpoint enables or disables [hedge mode](/glossary#hedge-mode) for the collateral trading account. When hedge mode is enabled (`true`), the account supports simultaneous long and short positions on the same market. When disabled (`false`), the account operates in one-way mode.

<Warning>
Rate limit: 1000 requests/10 sec.
</Warning>

<Warning>
Switching between one-way mode and hedge mode requires **no open positions**. Close all futures positions before toggling the mode. If the switch does not take effect immediately after closing positions, wait approximately 15 seconds and retry.
</Warning>

<Accordion title="Error Codes">
  - `30` - default validation error code (for example, a missing or non-boolean `hedgeMode` value)
  - `113` - hedge mode cannot be changed while open positions or orders exist
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
request := &sdk.UpdateHedgeModeRequest{
        HedgeMode: true,
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**hedgeMode:** `bool` — Set to `true` to enable hedge mode (simultaneous long and short positions) or `false` to use one-way mode.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.GetCollateralAccountAdlQuantile(request) -> []*sdk.GetCollateralAccountAdlQuantileResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns the [Auto-Deleveraging (ADL)](/glossary#auto-deleveraging-adl) quantile for each perpetual market in which the authenticated account holds an open position. Each entry exposes the deleveraging-priority value for the long and short sides of the position, where `0` indicates the lowest deleveraging priority and `4` indicates the highest. The endpoint returns an empty array when the account has no perpetual positions.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>

<Note>
Only perpetual markets (markets with the `_PERP` suffix) are returned. Spot and margin markets are not included.
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
request := &sdk.GetCollateralAccountAdlQuantileRequest{
        Request: "{{request}}",
        Nonce: 1594297865000,
    }
client.CollateralTrading.GetCollateralAccountAdlQuantile(
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

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.GetConditionalOrders(request) -> *sdk.GetConditionalOrdersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns active (unexecuted) conditional orders for the authenticated account. Conditional orders include [OCO](/glossary#one-cancels-the-other-oco) and [OTO](/glossary#one-triggers-the-other-oto) types. The response uses polymorphic structure — each record contains a `type` field (`oco` or `oto`) that determines the record shape. Use the optional `market` parameter to filter results.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
  - `30` - default validation error code (invalid pagination — `limit` outside 1–100 or negative `offset`)
  - `31` - market validation failed (the `market` filter is unknown or not available for collateral trading)
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
request := &sdk.GetConditionalOrdersRequest{
        Market: sdk.String(
            "BTC_USDT",
        ),
        Offset: sdk.Int(
            0,
        ),
        Limit: sdk.Int(
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

Filter by specific market. Example: BTC_USDT

If not specified, returns conditional orders for all markets.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Number of records to skip for pagination.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of records to return per page.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.GetOcoOrders(request) -> []*sdk.GetOcoOrdersResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns active (unexecuted) [OCO](/glossary#one-cancels-the-other-oco) orders for the authenticated account. Each OCO order contains a `stop_loss` and `take_profit` leg. When one leg executes, the system cancels the other automatically. Use the optional `market` parameter to filter results.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>

<Note>
This endpoint supports pagination. Use `limit` (default: 50) and `offset` (default: 0) to page through results. The response does not include a `total` field — detect the last page when fewer than `limit` OCO orders are returned. An empty array means you have paged past the end; receiving exactly `limit` orders does not guarantee that another page exists.
</Note>

<Accordion title="Error Codes">
  - `30` - default validation error code (invalid pagination — `limit` outside 1–100 or negative `offset`)
  - `31` - market validation failed (the `market` filter is unknown or not available for collateral trading)
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
request := &sdk.GetOcoOrdersRequest{
        Market: sdk.String(
            "BTC_USDT",
        ),
        Offset: sdk.Int(
            0,
        ),
        Limit: sdk.Int(
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

Filter by specific market. Example: BTC_USDT

If not specified, returns OCO orders for all markets.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Number of records to skip for pagination.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of records to return per page.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CreateCollateralOcoOrder(request) -> *sdk.CreateCollateralOcoOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a collateral [OCO](/glossary#one-cancels-the-other-oco) (one-cancels-the-other) order using [collateral balance](/glossary#balance-collateral). An OCO order combines a limit order (take-profit leg) and a stop-limit order (stop-loss leg) into a single conditional group. When one leg executes, the system cancels the other automatically.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
  - `30` - default validation error code. Also returned when `reduceOnly=true` is combined with `stopLoss` or `takeProfit`
  - `31` - market validation failed
  - `32` - amount validation failed
  - `33` - price validation failed
  - `36` - clientOrderId validation failed
  - `10` - insufficient balance to place the order
  - `111` - resulting position would exceed the market maximum
  - `112` - pending orders value would exceed the allowed maximum
  - `113` - position side cannot be changed while open positions or orders exist
  - `114` - hedge mode position side does not match (sent `BOTH` or omitted `positionSide` in hedge mode, or sent `LONG`/`SHORT` in one-way mode)
  - `115` - order would open a position in the opposite direction (one-way mode)
  - `116` - reduce-only validation failed (no position exists or order side matches position direction)
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
request := &sdk.CreateCollateralOcoOrderRequest{
        Market: "BTC_USDT",
        Side: sdk.CreateCollateralOcoOrderRequestSideBuy,
        Amount: "0.001",
        Price: "40000",
        ActivationPrice: "41000",
        StopLimitPrice: "42000",
        ClientOrderID: sdk.String(
            "order1987111",
        ),
        ReduceOnly: sdk.Bool(
            false,
        ),
        PositionSide: sdk.CreateCollateralOcoOrderRequestPositionSideLong.Ptr(),
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**market:** `string` — Available margin [market](/glossary#market). Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**side:** `*sdk.CreateCollateralOcoOrderRequestSide` — Order direction. Use `buy` to open or increase a long position and `sell` to open or increase a short position.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Amount of [stock](/glossary#stock) currency for both legs of the OCO order. Minimum and step values are market-dependent — query the [market info](/api-reference/market-data/market-info) endpoint for constraints.
    
</dd>
</dl>

<dl>
<dd>

**price:** `string` — Limit order price in [money](/glossary#money) currency for the take-profit leg.
    
</dd>
</dl>

<dl>
<dd>

**activationPrice:** `string` — Trigger price in [money](/glossary#money) currency for the stop-loss leg. The stop-limit order activates when the market price reaches the specified value.
    
</dd>
</dl>

<dl>
<dd>

**stopLimitPrice:** `string` — Execution price in [money](/glossary#money) currency for the stop-loss leg. After activation, the stop-loss leg places a limit order at the specified price.
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Custom client order identifier. Uniqueness is enforced only among the account's open (pending) orders on the same market — once a previous order is filled or canceled, the same identifier can be reused, including on the same market. Contains only letters, numbers, dashes, dots, or underscores.
    
</dd>
</dl>

<dl>
<dd>

**reduceOnly:** `*bool` — When `true`, both legs of the OCO order can only reduce or close an existing position — neither leg can increase the position or open a new one. If the order amount exceeds the current position size, the system reduces the order to match — the response returns the adjusted amount. The API returns error code `116` if no open position exists or the order side matches the position direction. See [reduce-only](/glossary#reduce-only).
    
</dd>
</dl>

<dl>
<dd>

**positionSide:** `*sdk.CreateCollateralOcoOrderRequestPositionSide` 

Position direction. Optional at the request layer but functionally required when hedge mode is enabled. See [positionSide](/glossary#position-side). Both legs of the OCO inherit the value.

- **One-way mode** (default account mode): the field is ignored. Orders always use `BOTH`, and the response returns `positionSide: "BOTH"` on each leg whether the field is sent or omitted.
- **Hedge mode**: the field MUST be `LONG` or `SHORT`. Sending `BOTH`, omitting the field, or sending a value that does not match the account's mode causes the trade service to reject the order with error code `114` (`Hedge mode position side does not match`).
    
</dd>
</dl>

<dl>
<dd>

**stp:** `*sdk.CreateCollateralOcoOrderRequestStp` 

Self-trade prevention mode. The value applies to both legs of the OCO order. Allowed values: `no` (self-trades allowed), `cb` (cancel both the new and the existing order), `cn` (cancel the new order, keep the existing), `co` (cancel the existing order, place the new one). Default: `no`.

Legacy values `cancel_both`, `cancel_new`, `cancel_old` are deprecated: the API accepts the legacy values with identical behavior until a deprecation deadline is announced, then rejects the legacy values. Responses always return the abbreviated form, regardless of which variant the request used.

See [Self-Trade Prevention](/platform/self-trade-prevention).
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
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

The endpoint cancels an active conditional order ([OCO](/glossary#one-cancels-the-other-oco) or [OTO](/glossary#one-triggers-the-other-oto)) on the specified market. Both legs of the conditional order are cancelled. Use the [query unexecuted conditional orders](/api-reference/collateral-trading/query-unexecuted-conditional-orders) endpoint to obtain the conditional order `id` before cancellation.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
  - `30` - default validation error code
  - `31` - market validation failed
  - `2` - conditional order not found. Returned whether the `id` does not exist, or the order was already filled or already cancelled — these cases are not distinguished
</Accordion>

<Accordion title="Errors">
```json
{
  "code": 2,
  "message": "Inner validation failed",
  "errors": {
    "id": ["Unexecuted order was not found."]
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
request := &sdk.CancelConditionalOrderRequest{
        Market: "BTC_USDT",
        ID: 117703764514,
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**market:** `string` — Market of the conditional order to cancel. Example: BTC_USDT
    
</dd>
</dl>

<dl>
<dd>

**id:** `int` — Conditional order identifier. Obtain from the [query unexecuted conditional orders](/api-reference/collateral-trading/query-unexecuted-conditional-orders) endpoint.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CollateralTrading.CancelOcoOrder(request) -> *sdk.CancelOcoOrderResponse</code></summary>
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

<Accordion title="Error Codes">
  - `30` - default validation error code
  - `31` - market validation failed
  - `2` - OCO order not found. Returned whether the `orderId` does not exist, or the order was already filled or already cancelled — these cases are not distinguished
</Accordion>

<Accordion title="Errors">
```json
{
  "code": 2,
  "message": "Inner validation failed",
  "errors": {
    "orderId": ["Unexecuted order was not found."]
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
request := &sdk.CancelOcoOrderRequest{
        Market: "BTC_USDT",
        OrderID: 117703764514,
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**nonce:** `int` 
    
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

<Accordion title="Error Codes">
  - `30` - default validation error code
  - `31` - market validation failed
  - `2` - OTO order not found. Returned whether the `otoId` does not exist, or the order was already filled or already cancelled — these cases are not distinguished
</Accordion>

<Accordion title="Errors">
```json
{
  "code": 2,
  "message": "Inner validation failed",
  "errors": {
    "otoId": ["Unexecuted order was not found."]
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
request := &sdk.CancelOtoOrderRequest{
        Market: "BTC_USDT",
        OtoID: 117703764514,
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Market Fee
<details><summary><code>client.MarketFee.GetMarketFee() -> *sdk.GetMarketFeeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the account's default spot and futures maker and taker fees, plus any custom per-market overrides.

The `maker` and `taker` fields represent default spot trading fees. The `futures_maker` and `futures_taker` fields represent default futures trading fees. The `custom_fee` object lists per-market overrides, keyed by market name.

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
request := &sdk.GetMarketFeeRequest{
        Market: sdk.String(
            "BTC_USDT",
        ),
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

**market:** `*string` — Optional. Currently ignored by the API — all market fees are returned regardless of the value provided. Retained for backward compatibility. Example: BTC_USDT
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Spot Trading
<details><summary><code>client.SpotTrading.TradeAccountBalance(request) -> map[string]*sdk.TradeAccountBalanceResponseValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the [trade balance](/glossary#balance-spotbalance-trade) by currency [ticker](/glossary#ticker) or all balances. When the `ticker` parameter is provided, the response contains a single currency entry. When omitted, the response contains all currencies with non-zero balances. Each entry includes the `available` balance (funds ready to trade) and the `freeze` balance (funds locked in open orders).

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
request := &sdk.TradeAccountBalanceRequest{}
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

**nonce:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CreateLimitOrder(request) -> *sdk.OrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a [limit trading order](/glossary#limit-order). The order remains on the order book until filled, cancelled, or expired. Minimum and maximum values for `amount` and `price` are market-dependent — query `GET /api/v4/public/markets` for per-market constraints.

**Order validation rules** (per-market, from `GET /api/v4/public/markets`):
- `amount` must have at most `stockPrec` decimal places
- `price` must have at most `moneyPrec` decimal places
- `amount` must be ≥ `minAmount`
- `amount × price` must be ≥ `minTotal`
- `amount × price` must be ≤ `maxTotal` (when `maxTotal` is not `"0"`)

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Note>
  - RPI orders do not appear in public order book feeds (`depth`, `bookTicker`). RPI orders are visible only in private active orders and in the exchange UI order book (web/mobile).
  - RPI orders are post-only by design and cannot be used with the IOC flag. The API returns error code `40` when both `rpi=true` and `ioc=true` are used.
  - `retail=true` marks the order as a retail-source taker eligible to match RPI-maker liquidity. The Retail flag must be enabled on the account; contact the account manager to enable it.
  - `retail=true` and `rpi=true` cannot be combined. The API returns error code `41` when both flags are set.
  - `retail=true` has no effect on a `postOnly=true` order. Post-only orders are makers and cannot be retail takers.
</Note>

<Accordion title="Error Codes">
  - `30` - default validation error code
  - `31` - market validation failed
  - `32` - amount validation failed
  - `33` - price validation failed
  - `36` - clientOrderId validation failed
  - `37` - `ioc=true` cannot be combined with `postOnly=true`
  - `40` - `ioc=true` cannot be combined with `rpi=true`
  - `41` - `retail=true` cannot be combined with `rpi=true`
  - `42` - `retail=true` is not allowed for the account
  - `43` - `rpi=true` is not allowed for the account
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
    "clientOrderId": ["ClientOrderId field should be a string."]
  }
}
```

```json
{
  "code": 36,
  "message": "Validation failed",
  "errors": {
    "clientOrderId": [
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
    "clientOrderId": [
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

```json
{
  "code": 41,
  "message": "Validation failed",
  "errors": {
    "retail": ["api.tradeErrors.flagsCantBeCombined.rpiRetail"]
  }
}
```

```json
{
  "code": 42,
  "message": "Validation failed",
  "errors": {
    "retail": ["api.validation.retail.not_allowed"]
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
request := &sdk.LimitOrderRequest{
        Market: "BTC_USDT",
        Side: sdk.LimitOrderRequestSideBuy,
        Amount: "0.001",
        Price: "9800",
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**market:** `string` — Trading pair. Format: `BASE_QUOTE` (e.g., `BTC_USDT`). Query `GET /api/v4/public/markets` for available markets.
    
</dd>
</dl>

<dl>
<dd>

**side:** `*sdk.LimitOrderRequestSide` — Order side. Allowed values: `buy`, `sell`.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Order quantity in base (stock) currency. Minimum and maximum values are market-dependent. Query `GET /api/v4/public/markets` for `minAmount`, `minTotal`, `maxTotal`. Precision: `stockPrec`.
    
</dd>
</dl>

<dl>
<dd>

**price:** `string` — Limit price per unit in quote (money) currency. Minimum and maximum values are market-dependent. Precision: `moneyPrec`.
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Custom client order identifier. Uniqueness is enforced only among the account's open (pending) orders on the same market — once a previous order is filled or canceled, the same identifier can be reused, including on the same market. Contains only letters, numbers, dashes, dots, or underscores.
    
</dd>
</dl>

<dl>
<dd>

**postOnly:** `*bool` — Post-only flag. When `true`, the order executes only as a [maker](/glossary#maker) order and the system rejects the order if it would match immediately. Default: `false`.
    
</dd>
</dl>

<dl>
<dd>

**ioc:** `*bool` 

Immediate-or-cancel (IOC) flag. When `true`, the matching engine executes all or part of the order immediately and cancels any unfilled portion. Default: `false`.

IOC does not support `rpi=true` because RPI uses post-only behavior by design.
The API returns error code `40` when a request sets both `ioc=true` and `rpi=true`.

Refer to [Order Parameter Rules](/guides/order-parameter-rules) for unsupported parameter combinations.
    
</dd>
</dl>

<dl>
<dd>

**bboRole:** `*int` — Best Bid/Offer ([BBO](/glossary#bbo)) execution method. The system selects the best market price for execution. `1` = Queue method, `2` = Counterparty method. Use method `2` with the `ioc` flag.
    
</dd>
</dl>

<dl>
<dd>

**stp:** `*sdk.LimitOrderRequestStp` 

Self-trade prevention mode. Allowed values: `no` (self-trades allowed), `cb` (cancel both the new and the existing order), `cn` (cancel the new order, keep the existing), `co` (cancel the existing order, place the new one). Default: `no`.

Legacy values `cancel_both`, `cancel_new`, `cancel_old` are deprecated: the API accepts the legacy values with identical behavior until a deprecation deadline is announced, then rejects the legacy values. Responses always return the abbreviated form, regardless of which variant the request used.

See [Self-Trade Prevention](/platform/self-trade-prevention).
    
</dd>
</dl>

<dl>
<dd>

**rpi:** `*bool` 

Enables Retail Price Improvement (RPI) mode. Default: `false`.

RPI orders use post-only behavior by design. An RPI order does not support `ioc=true`.
The API returns error code `40` when a request sets both `rpi=true` and `ioc=true`.
RPI orders do not appear in public order book feeds (`depth`, `bookTicker`). RPI orders are visible only in private active orders and in the exchange UI order book (web/mobile).
RPI executions may apply custom fees or rebates, especially when trading via sub-accounts. Use Query Market Fees to verify effective fees.

Refer to [Order Parameter Rules](/guides/order-parameter-rules) for unsupported parameter combinations.
    
</dd>
</dl>

<dl>
<dd>

**retail:** `*bool` 

Retail-source taker flag. When `true`, the order is eligible to match against orders submitted by RPI makers and may receive price improvement at execution. Default: `false`.

The Retail flag must be enabled on the account before a private-API request can set `retail=true`. Contact the account manager to enable the Retail flag.

The Retail flag cannot be combined with `rpi`. The API returns error code `41` when a request sets both `retail=true` and `rpi=true`.

The flag has no effect on a `postOnly=true` order. Post-only orders are [makers](/glossary#maker); only takers carry the retail designation.

Refer to [Retail flag](/glossary#retail-flag) and [Order Parameter Rules](/guides/order-parameter-rules) for unsupported parameter combinations.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CreateBulkLimitOrder(request) -> sdk.BulkLimitOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates bulk [limit trading orders](/glossary#limit-order). Each order in the batch follows the same validation rules as a single limit order. The `stopOnFail` parameter controls whether processing stops at the first failure or continues through all orders. The response contains a result-or-error pair for each submitted order.

<Warning>
  Limit: From 1 to 20 orders per request.
</Warning>

<Note>
  - RPI orders do not appear in public order book feeds (`depth`, `bookTicker`). RPI orders are visible only in private active orders and in the exchange UI order book (web/mobile).
  - RPI orders are post-only by design and cannot be used with the IOC flag. The API returns error code `40` when both `rpi=true` and `ioc=true` are used.
  - `retail=true` marks the order as a retail-source taker eligible to match RPI-maker liquidity. The Retail flag must be enabled on the account; contact the account manager to enable it.
  - `retail=true` and `rpi=true` cannot be combined. The API returns error code `41` when both flags are set on an item.
  - `retail=true` has no effect on a `postOnly=true` item. Post-only orders are makers and cannot be retail takers.
</Note>


<Accordion title="Error Codes">
  - `30` - default validation error code
  - `31` - market validation failed
  - `32` - amount validation failed
  - `33` - price validation failed
  - `36` - clientOrderId validation failed
  - `37` - `ioc=true` cannot be combined with `postOnly=true`
  - `40` - `ioc=true` cannot be combined with `rpi=true`
  - `41` - `retail=true` cannot be combined with `rpi=true`
  - `42` - `retail=true` is not allowed for the account
  - `43` - `rpi=true` is not allowed for the account
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

```json
{
  "code": 41,
  "message": "Validation failed",
  "errors": {
    "retail": ["api.tradeErrors.flagsCantBeCombined.rpiRetail"]
  }
}
```

```json
{
  "code": 42,
  "message": "Validation failed",
  "errors": {
    "retail": ["api.validation.retail.not_allowed"]
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
request := &sdk.CreateBulkLimitOrderRequest{
        Orders: []*sdk.BulkOrderItem{
            &sdk.BulkOrderItem{
                Side: sdk.BulkOrderItemSideBuy.Ptr(),
                Amount: sdk.String(
                    "0.02",
                ),
                Price: sdk.String(
                    "40000",
                ),
                Market: sdk.String(
                    "BTC_USDT",
                ),
                PostOnly: sdk.Bool(
                    false,
                ),
                Ioc: sdk.Bool(
                    false,
                ),
                ClientOrderID: sdk.String(
                    "",
                ),
                Rpi: sdk.Bool(
                    true,
                ),
                Retail: sdk.Bool(
                    false,
                ),
            },
            &sdk.BulkOrderItem{
                Side: sdk.BulkOrderItemSideSell.Ptr(),
                Amount: sdk.String(
                    "0.0001",
                ),
                Price: sdk.String(
                    "41000",
                ),
                Market: sdk.String(
                    "BTC_USDT",
                ),
                PostOnly: sdk.Bool(
                    false,
                ),
                Ioc: sdk.Bool(
                    false,
                ),
                ClientOrderID: sdk.String(
                    "",
                ),
                Rpi: sdk.Bool(
                    false,
                ),
                Retail: sdk.Bool(
                    true,
                ),
            },
            &sdk.BulkOrderItem{
                Side: sdk.BulkOrderItemSideSell.Ptr(),
                Amount: sdk.String(
                    "0.02",
                ),
                Price: sdk.String(
                    "41000",
                ),
                Market: sdk.String(
                    "BTC_USDT",
                ),
                PostOnly: sdk.Bool(
                    false,
                ),
                Ioc: sdk.Bool(
                    false,
                ),
                ClientOrderID: sdk.String(
                    "",
                ),
                Rpi: sdk.Bool(
                    false,
                ),
                Retail: sdk.Bool(
                    false,
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

**orders:** `[]*sdk.BulkOrderItem` — Array of limit orders
    
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

**nonce:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CreateMarketOrder(request) -> *sdk.OrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a [market trading order](/glossary#market-order). The matching engine executes the order immediately at the best available price. For buy orders, `amount` represents the total in quote (money) currency to spend. For sell orders, `amount` represents the quantity in base (stock) currency to sell. Minimum and maximum values are market-dependent. Query `GET /api/v4/public/markets` for `minAmount`, `minTotal`, and `maxTotal`.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
- `30` - default validation error code
- `31` - market validation failed
- `32` - amount validation failed
- `36` - clientOrderId validation failed
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
    "clientOrderId": ["ClientOrderId field should be a string."]
  }
}
```

```json
{
  "code": 36,
  "message": "Validation failed",
  "errors": {
    "clientOrderId": [
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
request := &sdk.MarketOrderRequest{
        Market: "BTC_USDT",
        Side: sdk.MarketOrderRequestSideBuy,
        Amount: "100",
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**market:** `string` — Trading pair. Format: `BASE_QUOTE` (e.g., `BTC_USDT`). Query `GET /api/v4/public/markets` for available markets.
    
</dd>
</dl>

<dl>
<dd>

**side:** `*sdk.MarketOrderRequestSide` — Order side. Allowed values: `buy`, `sell`.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — For buy orders: total in quote (money) currency to spend. For sell orders: quantity in base (stock) currency to sell. Minimum and maximum values are market-dependent. Query `GET /api/v4/public/markets` for `minAmount`, `minTotal`, `maxTotal`.
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Custom client order identifier. Uniqueness is enforced only among the account's open (pending) orders on the same market — once a previous order is filled or canceled, the same identifier can be reused, including on the same market. Contains only letters, numbers, dashes, dots, or underscores.
    
</dd>
</dl>

<dl>
<dd>

**stp:** `*sdk.MarketOrderRequestStp` 

Self-trade prevention mode. Allowed values: `no` (self-trades allowed), `cb` (cancel both the new and the existing order), `cn` (cancel the new order, keep the existing), `co` (cancel the existing order, place the new one). Default: `no`.

Legacy values `cancel_both`, `cancel_new`, `cancel_old` are deprecated: the API accepts the legacy values with identical behavior until a deprecation deadline is announced, then rejects the legacy values. Responses always return the abbreviated form, regardless of which variant the request used.

See [Self-Trade Prevention](/platform/self-trade-prevention).
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CreateStockMarketOrder(request) -> *sdk.OrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a [stock](/glossary#stock) market trading [order](/glossary#orders). Unlike `POST /api/v4/order/market`, the `amount` parameter always represents the quantity in the base (stock) currency for both buy and sell sides. The matching engine executes the order immediately at the best available price. Minimum and maximum values are market-dependent. Query `GET /api/v4/public/markets` for `minAmount`, `minTotal`, and `maxTotal`.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
- `30` - default validation error code
- `31` - market validation failed
- `32` - amount validation failed
- `36` - clientOrderId validation failed
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
    "clientOrderId": ["ClientOrderId field should be a string."]
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
request := &sdk.StockMarketOrderRequest{
        Market: "BTC_USDT",
        Side: sdk.StockMarketOrderRequestSideBuy,
        Amount: "0.001",
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**market:** `string` — Trading pair. Format: `BASE_QUOTE` (e.g., `BTC_USDT`). Query `GET /api/v4/public/markets` for available markets.
    
</dd>
</dl>

<dl>
<dd>

**side:** `*sdk.StockMarketOrderRequestSide` — Order side. Allowed values: `buy`, `sell`.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Order quantity in base (stock) currency for both buy and sell sides. To place a market order specifying the quote (money) currency amount instead, use `POST /api/v4/order/market`. Minimum and maximum values are market-dependent. Query `GET /api/v4/public/markets` for `minAmount`, `minTotal`, `maxTotal`.
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Custom client order identifier. Uniqueness is enforced only among the account's open (pending) orders on the same market — once a previous order is filled or canceled, the same identifier can be reused, including on the same market. Contains only letters, numbers, dashes, dots, or underscores.
    
</dd>
</dl>

<dl>
<dd>

**stp:** `*sdk.StockMarketOrderRequestStp` 

Self-trade prevention mode. Allowed values: `no` (self-trades allowed), `cb` (cancel both the new and the existing order), `cn` (cancel the new order, keep the existing), `co` (cancel the existing order, place the new one). Default: `no`.

Legacy values `cancel_both`, `cancel_new`, `cancel_old` are deprecated: the API accepts the legacy values with identical behavior until a deprecation deadline is announced, then rejects the legacy values. Responses always return the abbreviated form, regardless of which variant the request used.

See [Self-Trade Prevention](/platform/self-trade-prevention).
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CreateStopLimitOrder(request) -> *sdk.OrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a [stop-limit trading order](/glossary#stop-limit-order). The order remains inactive until the market price reaches the `activation_price`, at which point the system places a limit order at the specified `price`. For buy orders, activation triggers when the market price rises to or above `activation_price`. For sell orders, activation triggers when the market price falls to or below `activation_price`. Minimum and maximum values for `amount`, `price`, and `activation_price` are market-dependent. Query `GET /api/v4/public/markets` for `minAmount`, `minTotal`, `maxTotal`, `stockPrec` (amount precision), and `moneyPrec` (price precision).

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
- `30` - default validation error code
- `31` - market validation failed
- `32` - amount validation failed
- `33` - price validation failed
- `36` - clientOrderId validation failed
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
    "clientOrderId": ["ClientOrderId field should be a string."]
  }
}
```

```json
{
  "code": 36,
  "message": "Validation failed",
  "errors": {
    "clientOrderId": [
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
request := &sdk.StopLimitOrderRequest{
        Market: "BTC_USDT",
        Side: sdk.StopLimitOrderRequestSideBuy,
        Amount: "0.001",
        Price: "9800",
        ActivationPrice: "10000",
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**market:** `string` — Trading pair. Format: `BASE_QUOTE` (e.g., `BTC_USDT`). Query `GET /api/v4/public/markets` for available markets.
    
</dd>
</dl>

<dl>
<dd>

**side:** `*sdk.StopLimitOrderRequestSide` — Order side. Allowed values: `buy`, `sell`.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — Order quantity in base (stock) currency. Minimum and maximum values are market-dependent. Query `GET /api/v4/public/markets` for `minAmount`, `minTotal`, `maxTotal`. Precision: `stockPrec`.
    
</dd>
</dl>

<dl>
<dd>

**price:** `string` — Limit price per unit in quote (money) currency applied after the stop triggers. Minimum and maximum values are market-dependent. Precision: `moneyPrec`.
    
</dd>
</dl>

<dl>
<dd>

**activationPrice:** `string` — Trigger price in quote (money) currency. For buy orders, the stop triggers when the market price rises to or above the specified price. For sell orders, the stop triggers when the market price falls to or below the specified price. Precision: `moneyPrec`.
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Custom client order identifier. Uniqueness is enforced only among the account's open (pending) orders on the same market — once a previous order is filled or canceled, the same identifier can be reused, including on the same market. Contains only letters, numbers, dashes, dots, or underscores.
    
</dd>
</dl>

<dl>
<dd>

**bboRole:** `*int` — Best Bid/Offer ([BBO](/glossary#bbo)) execution method. The system selects the best market price for execution after the stop triggers. `1` = Queue method, `2` = Counterparty method.
    
</dd>
</dl>

<dl>
<dd>

**stp:** `*sdk.StopLimitOrderRequestStp` 

Self-trade prevention mode. Allowed values: `no` (self-trades allowed), `cb` (cancel both the new and the existing order), `cn` (cancel the new order, keep the existing), `co` (cancel the existing order, place the new one). Default: `no`.

Legacy values `cancel_both`, `cancel_new`, `cancel_old` are deprecated: the API accepts the legacy values with identical behavior until a deprecation deadline is announced, then rejects the legacy values. Responses always return the abbreviated form, regardless of which variant the request used.

See [Self-Trade Prevention](/platform/self-trade-prevention).
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CreateStopMarketOrder(request) -> *sdk.OrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates a [stop-market trading order](/glossary#stop-market-order). The order remains inactive until the market price reaches the `activation_price`, at which point the system executes a market order immediately at the best available price. For buy orders, `amount` represents the total in quote currency and activation triggers when the market price rises to or above `activation_price`. For sell orders, `amount` represents the quantity in base currency and activation triggers when the market price falls to or below `activation_price`. Minimum and maximum values are market-dependent. Query `GET /api/v4/public/markets` for `minAmount`, `minTotal`, and `maxTotal`.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Accordion title="Error Codes">
- `30` - default validation error code
- `31` - market validation failed
- `32` - amount validation failed
- `36` - clientOrderId validation failed
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
    "clientOrderId": ["ClientOrderId field should be a string."]
  }
}
```

```json
{
  "code": 36,
  "message": "Validation failed",
  "errors": {
    "clientOrderId": [
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
    "clientOrderId": [
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
request := &sdk.StopMarketOrderRequest{
        Market: "BTC_USDT",
        Side: sdk.StopMarketOrderRequestSideBuy,
        Amount: "0.01",
        ActivationPrice: "10000",
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**market:** `string` — Trading pair. Format: `BASE_QUOTE` (e.g., `BTC_USDT`). Query `GET /api/v4/public/markets` for available markets.
    
</dd>
</dl>

<dl>
<dd>

**side:** `*sdk.StopMarketOrderRequestSide` — Order side. Allowed values: `buy`, `sell`.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `string` — For buy orders: total in quote (money) currency to spend. For sell orders: quantity in base (stock) currency to sell. Minimum and maximum values are market-dependent. Query `GET /api/v4/public/markets` for `minAmount`, `minTotal`, `maxTotal`.
    
</dd>
</dl>

<dl>
<dd>

**activationPrice:** `string` — Trigger price in quote (money) currency. For buy orders, the stop triggers when the market price rises to or above the specified price. For sell orders, the stop triggers when the market price falls to or below the specified price. Precision: `moneyPrec`.
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Custom client order identifier. Uniqueness is enforced only among the account's open (pending) orders on the same market — once a previous order is filled or canceled, the same identifier can be reused, including on the same market. Contains only letters, numbers, dashes, dots, or underscores.
    
</dd>
</dl>

<dl>
<dd>

**stp:** `*sdk.StopMarketOrderRequestStp` 

Self-trade prevention mode. Allowed values: `no` (self-trades allowed), `cb` (cancel both the new and the existing order), `cn` (cancel the new order, keep the existing), `co` (cancel the existing order, place the new one). Default: `no`.

Legacy values `cancel_both`, `cancel_new`, `cancel_old` are deprecated: the API accepts the legacy values with identical behavior until a deprecation deadline is announced, then rejects the legacy values. Responses always return the abbreviated form, regardless of which variant the request used.

See [Self-Trade Prevention](/platform/self-trade-prevention).
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CancelOrder(request) -> *sdk.OrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint cancels an existing [order](/glossary#orders). Provide either `orderId` or `clientOrderId` to identify the target order. The response returns the final state of the cancelled order.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Note>
- Cancellation by clientOrderId takes priority over orderId.
- The request supports working only with orderId or only with clientOrderId.
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
    "orderId": ["OrderId field is required."]
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
    "orderId": ["OrderId field should be an integer."]
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
    "orderId": ["Unexecuted order was not found."]
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
request := &sdk.CancelOrderRequest{
        Market: "BTC_USDT",
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**orderID:** `*int` — Order Id. Example: 4180284841. Required if clientOrderId is not set.
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Custom client order id. Example: 'customId11'. Required if orderId is not set.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.CancelBulkOrders(request) -> sdk.BulkCancelOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint cancels up to 100 [orders](/glossary#orders) in a single request. Each item identifies a target order by `market` plus exactly one of `orderId` or `clientOrderId`. The response is an array whose items match the input order one-to-one — `response[i]` corresponds to `request.orders[i]`.

<Warning>
Rate limit: 10000 requests/10 sec.
</Warning>

<Warning>
Limit: From 1 to 100 orders per request.
</Warning>

<Note>
- Provide exactly one of `orderId` or `clientOrderId` per item. Sending both, or neither, returns a per-item validation error.
- The endpoint always processes every item independently. There is no `stopOnFail`-style switch.
- When the caller is not authenticated, the API returns the standard authorization error and skips per-item validation.
</Note>

<Accordion title="Error Codes">
- `30` — validation failure (per-item)
- `404` — order not found (per-item)
- `500` — trade service unavailable (per-item)
</Accordion>

<Accordion title="Errors">
**Per-item errors (returned inside the response array):**

Element is not a valid object:
```json
{
  "result": null,
  "error": {
    "code": 30,
    "message": "Validation failed",
    "errors": { "request": ["Invalid order format"] }
  }
}
```

`market` field is missing:
```json
{
  "result": null,
  "error": {
    "code": 30,
    "message": "Validation failed",
    "errors": { "market": ["validation.required"] }
  }
}
```

Specified market does not exist:
```json
{
  "result": null,
  "error": {
    "code": 30,
    "message": "Validation failed",
    "errors": { "market": ["validation.market_not_exist"] }
  }
}
```

Neither `orderId` nor `clientOrderId` provided:
```json
{
  "result": null,
  "error": {
    "code": 30,
    "message": "Validation failed",
    "errors": { "request": ["validation.required"] }
  }
}
```

Both `orderId` and `clientOrderId` provided:
```json
{
  "result": null,
  "error": {
    "code": 30,
    "message": "Validation failed",
    "errors": { "request": ["api.validation.order.chooseOneId"] }
  }
}
```

Order not found (returned for both `orderId` and `clientOrderId` lookups):
```json
{
  "result": null,
  "error": {
    "code": 404,
    "message": "Order not found",
    "errors": {
      "orderId": ["Order does not exist or already cancelled"]
    }
  }
}
```

Trade service unavailable (returned when the trade service is unreachable or returns an unparseable response):
```json
{
  "result": null,
  "error": {
    "code": 500,
    "message": "Service temporary unavailable",
    "errors": { "error": ["Service temporary unavailable"] }
  }
}
```

**Request-level errors (HTTP 422, returned as a standard error envelope, not as an array):**

`orders` field is missing or is not an array:
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": { "orders": ["validation.required"] }
}
```

`orders` contains more than 100 elements:
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": { "orders": ["validation.between"] }
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
request := &sdk.CancelBulkOrdersRequest{
        Orders: []*sdk.BulkCancelOrderItem{
            &sdk.BulkCancelOrderItem{
                Market: "BTC_USDT",
                OrderID: sdk.Int(
                    4326248250,
                ),
            },
            &sdk.BulkCancelOrderItem{
                Market: "ETH_USDT",
                ClientOrderID: sdk.String(
                    "my-client-id",
                ),
            },
        },
        Request: "{{request}}",
        Nonce: 1594297865000,
    }
client.SpotTrading.CancelBulkOrders(
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

**orders:** `[]*sdk.BulkCancelOrderItem` — Array of orders to cancel. From 1 to 100 items per request.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
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

The endpoint cancels all open [orders](/glossary#orders) that match the specified filters. Use the `market` parameter to target a single trading pair, or omit the parameter to cancel across all markets. The `type` parameter filters by order type (`spot`, `margin`, `futures`). When omitted, the endpoint targets all order types.

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
request := &sdk.CancelAllOrdersRequest{
        Market: sdk.String(
            "BTC_USDT",
        ),
        Type: []sdk.CancelAllOrdersRequestTypeItem{
            sdk.CancelAllOrdersRequestTypeItemSpot,
            sdk.CancelAllOrdersRequestTypeItemMargin,
            sdk.CancelAllOrdersRequestTypeItemFutures,
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

**type_:** `[]*sdk.CancelAllOrdersRequestTypeItem` — Order types to target. Valid values: "spot" — standard spot orders. "margin" — marginal orders placed on spot markets. Note: the "margin" value is not the same as the collateral account balance; "collateral" in other endpoints refers to the funding account, whereas "margin" here refers specifically to the order type. "futures" — marginal orders placed on futures markets (e.g., BTC_PERP). If omitted, the API targets all order types.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.GetActiveOrders(request) -> []*sdk.OrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves [active orders](/glossary#active-orders) (orders not yet executed). The response includes limit, stop-limit, and stop-market orders that remain open on the order book. Use the `market` parameter to filter by trading pair, or omit the parameter to retrieve orders across all markets. The endpoint supports pagination with `limit` and `offset` parameters.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>

<Note>
Search across all markets is available only if clientOrderId and orderId are not provided.
</Note>

<Note>
This endpoint supports pagination. Use `limit` (default: 50, max: 100) and `offset` (default: 0, max: 4294967295) to page through results. The response does not include a `total` field — detect the last page when fewer than `limit` orders are returned. An empty array means you have paged past the end; receiving exactly `limit` orders does not guarantee that another page exists.
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
    "offset": ["The offset may not be greater than 4294967295."]
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
request := &sdk.GetActiveOrdersRequest{}
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

**market:** `*string` — Trading pair to filter by. Format: `BASE_QUOTE` (e.g., `BTC_USDT`). Omit to retrieve orders across all markets.
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `*int` — Filter by a specific order identifier. Returns only the matching active order.
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Filter by custom client order identifier. Returns only the matching active order.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Number of records to skip. Default: `0`. Maximum: `4294967295`.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of records to return. Default: `50`. Minimum: `1`. Maximum: `100`.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.GetExecutedOrderHistory(request) -> *sdk.GetExecutedOrderHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves executed order history for all trading types — spot, margin, and futures — across all markets. Can be filtered by a single market if needed. Results are ordered by trade time, newest first.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>

<Warning>
Requests with `limit` values above 100 return large payloads. Use high limits only when necessary and ensure the client application can handle large response sizes.
</Warning>

<Note>
The endpoint can retrieve data not older than 6 months from the current month. For older data, use the Report on the History page.
</Note>

<Note>
This endpoint supports pagination. Use `limit` (default: 50, max: 500) and `offset` (default: 0) to page through results. The response does not include a `total` field — detect the last page when fewer than `limit` records are returned (sum the records across all markets when no `market` filter is set). An empty response means you have paged past the end; receiving exactly `limit` records does not guarantee that another page exists.
</Note>

<Note>
For B2B accounts, canceled orders are not recorded in the history. To obtain canceled-order data, contact support or the assigned account manager.
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

```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "orderHistory": ["OrderHistory was not found."]
  }
}
```
Returned when `clientOrderId` is supplied but no order matches it on the calling account.
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
request := &sdk.GetExecutedOrderHistoryRequest{}
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

**clientOrderID:** `*string` — Look up by custom client order identifier. When supplied, the endpoint switches to single-order lookup mode and returns the matching order's deal history. Returns `422` with `"OrderHistory was not found."` if no order matches on the calling account.
    
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

**limit:** `*int` — LIMIT is a special clause used to limit records a particular query can return. Default: 50, Min: 1, Max: 500
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.GetOrderDeals(request) -> *sdk.GetOrderDealsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves individual trade fills (deals) for a specific order. Each deal represents a partial or full execution of the order against a counterparty. The response includes pagination and returns deal details such as price, amount, fee, and execution role (maker or taker).

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>

<Note>
This endpoint supports pagination. Use `limit` (default: 50) and `offset` (default: 0) to page through results. The response does not include a `total` field — detect the last page when `records.length < limit`. An empty `records` array means you have paged past the end; receiving exactly `limit` records does not guarantee that another page exists.
</Note>

<Note>
An unknown or not-owned `orderId` is **not** an error. The endpoint always returns HTTP 200 with the paged-list envelope; a non-matching `orderId` simply filters down to an empty `records` array.
</Note>

<Note>
The endpoint can retrieve data not older than 6 months from the current month. For older data, use the Report on the History page. An order older than this window returns an empty `records` array even when the order was filled.
</Note>

<Accordion title="Error Codes">
  - `30` - default validation error code (for example, a missing or malformed `orderId`, or invalid pagination)
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
request := &sdk.GetOrderDealsRequest{
        OrderID: 3134995325,
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**orderID:** `int` — Identifier of the order to retrieve deals for.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Number of records to skip. Default: `0`.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of records to return. Default: `50`.
    
</dd>
</dl>

<dl>
<dd>

**request:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.GetOrderHistory(request) -> map[string][]*sdk.GetOrderHistoryResponseValueItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the history of executed and cancelled orders. The response groups orders by market name. Use the `market` parameter to filter by a single trading pair, or omit the parameter to retrieve orders across all markets. The endpoint supports pagination with `limit` (default 50, max 500) and `offset` parameters. Results are ordered by time, newest first.

<Warning>
Rate limit: 12000 requests/10 sec.
</Warning>

<Warning>
Requests with `limit` values above 100 return large payloads. Use high limits only when necessary and ensure the client application can handle large response sizes.
</Warning>

<Note>
This endpoint supports pagination. Use `limit` (default: 50, max: 500) and `offset` (default: 0) to page through results. The response does not include a `total` field — detect the last page when fewer than `limit` records are returned (sum the records across all markets when no `market` filter is set). An empty response means you have paged past the end; receiving exactly `limit` records does not guarantee that another page exists.
</Note>

<Note>
**Date filter window:** the maximum span between `startDate` and `endDate` is **31 days**, and the earliest reachable date is **6 months ago (00:00 UTC)**. Requests that exceed the 31-day window or fall below the 6-month floor are rejected with a validation error. `endDate` values greater than the current time are silently clamped to `now`.
</Note>

<Note>
For B2B accounts, canceled orders are not recorded in the history. To obtain canceled-order data, contact support or the assigned account manager.
</Note>

<Accordion title="Errors">
```json
{
  "code": 30,
  "message": "Validation failed",
  "errors": {
    "orderHistory": ["OrderHistory was not found."]
  }
}
```
Returned when `clientOrderId` is supplied but no order matches it on the calling account.
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
request := &sdk.GetOrderHistoryRequest{}
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

**market:** `*string` — Trading pair to filter by. Format: `BASE_QUOTE` (e.g., `BTC_USDT`). Omit to retrieve orders across all markets.
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Look up a specific order by the custom client identifier. When supplied, the endpoint switches to single-order lookup mode and the `startDate`, `endDate`, and `status` filters are ignored. Returns `422` with `"OrderHistory was not found."` if no order matches. Ignored when `orderId` is also provided.
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `*int` — Look up a specific order by the exchange-assigned identifier. When supplied, the endpoint switches to single-order lookup mode and the `startDate`, `endDate`, and `status` filters are ignored. Returns an empty result on no match (no `422`). Takes precedence over `clientOrderId` when both are supplied.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*sdk.GetOrderHistoryRequestStatus` — Filter list-mode results by order status. Ignored when `orderId` or `clientOrderId` is supplied.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*int` — Start of the query window as a Unix timestamp in seconds. Default: `now - 1 month`. The earliest reachable date is 6 months ago (00:00 UTC) — requests with an older `startDate` are rejected with a validation error.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*int` — End of the query window as a Unix timestamp in seconds. Default: `now`. Values greater than the current time are silently clamped to `now`. The maximum span between `startDate` and `endDate` is 31 days.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Number of records to skip. Default: `0`.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of records to return. Default: `50`. Minimum: `1`. Maximum: `500`.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.GetDelistingOrderHistory(request) -> []*sdk.GetDelistingOrderHistoryResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint returns the authenticated account's delisting-related order history from the order-history index. The response combines two categories of orders, each identified by the `delistingKind` field: reverse close orders that the platform generates when a market is delisted (`delistingKind` = `reverse`, `clientOrderId` prefixed with `delisting-`), and active orders canceled at the moment of delisting (`delistingKind` = `canceled`). The endpoint returns a flat array sorted by `finishAt` descending, then `id` descending.

Set the `status` parameter to narrow the response: `filled` returns only reverse close orders that reached the `filled` outcome, and `delisting` returns only orders canceled during delisting. Omit `status` to return both categories merged in a single array.

<Note>
The endpoint supports pagination via `limit` (default: 500, max: 500) and `offset` (default: 0); the sum of `offset` and `limit` must not exceed 10000. A response that returns fewer than `limit` records indicates the last page.
</Note>

<Note>
**Date filter window:** the date range filters orders by the `finishAt` timestamp. The maximum span between `startDate` and `endDate` is **31 days**. `endDate` values greater than the current time are clamped to `now`.
</Note>

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
request := &sdk.GetDelistingOrderHistoryRequest{}
client.SpotTrading.GetDelistingOrderHistory(
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

**market:** `*string` — Trading pair to filter by. Format: `BASE_QUOTE` (e.g., `BTC_USDT`), matching the pattern `^[A-Z0-9]+_[A-Z0-9]+$`. Omit to retrieve delisting orders across all markets.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*sdk.GetDelistingOrderHistoryRequestStatus` — Category filter — distinct from the response `status` field. `filled` returns only reverse close orders (`clientOrderId` prefixed with `delisting-`), which carry `delistingKind` = `reverse` and response `status` = `filled`. `delisting` returns only orders canceled at delisting, which carry `delistingKind` = `canceled` and response `status` = `canceled`. Omit to return both categories merged.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*int` — Start of the query window as a Unix timestamp in seconds. Default: `now - 30 days`. Must not be later than `endDate`.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*int` — End of the query window as a Unix timestamp in seconds. Default: `now`. Values greater than the current time are clamped to `now`. The maximum span between `startDate` and `endDate` is 31 days.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Number of records to skip. Default: `0`. The sum of `offset` and `limit` must not exceed 10000.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of records to return. Default: `500`. Minimum: `1`. Maximum: `500`.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.ModifyOrder(request) -> *sdk.OrderResponse</code></summary>
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
- Modification by clientOrderId takes priority.
- The request supports working only with orderId or only with clientOrderId.
- Do not pass both values at the same time.
</Note>

<Note>
**WebSocket impact:** Each call to the endpoint cancels the original order and
creates a replacement with a **new `orderId`**. Clients subscribed to the
`ordersPending_update` WebSocket channel will receive:
- `event_id=3` (cancel) for the old order
- `event_id=1` (new) for the replacement

Update any `orderId` references after a successful modify response.
Use `clientOrderId` for stable order tracking across modifications.
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
    "orderId": ["Unexecuted order was not found."]
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
request := &sdk.ModifyOrderRequest{
        Market: "BTC_USDT",
        Request: "{{request}}",
        Nonce: 1594297865000,
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

**orderID:** `*int` — Active order id. Required if clientOrderId is not set.
    
</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Identifier should be unique and contain letters, dashes, numbers, dots or underscores. Required if orderId is not set.
    
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

**nonce:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.SetKillSwitch(request) -> *sdk.SetKillSwitchResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint creates, updates, or deletes a [kill-switch timer](/glossary#kill-switch-timer). The kill-switch acts as a safety mechanism for automated trading systems — the timer automatically cancels all open orders for the specified market if the client fails to reset the timer before expiration. Set `timeout` to a value between `5` and `600` (seconds) to create or update a timer. Set `timeout` to `null` to delete an existing timer.

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
request := &sdk.SetKillSwitchRequest{
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

**types:** `[]*sdk.SetKillSwitchRequestTypesItem` — Order types to target. Valid values: "spot" — standard spot orders. "margin" — marginal orders placed on spot markets. Note: the "margin" value is not the same as the collateral account balance; "collateral" in other endpoints refers to the funding account, whereas "margin" here refers specifically to the order type. "futures" — marginal orders placed on futures markets (e.g., BTC_PERP). If omitted, the API targets all order types.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**nonce:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SpotTrading.GetKillSwitchStatus(request) -> []*sdk.GetKillSwitchStatusResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The endpoint retrieves the status of active [kill-switch timers](/glossary#kill-switch-timer). The response returns an array of timer objects for the specified market, or for all markets if the `market` parameter is omitted. Each timer object includes the start time, scheduled cancellation time, and targeted order types.

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
request := &sdk.GetKillSwitchStatusRequest{}
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

**nonce:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>
