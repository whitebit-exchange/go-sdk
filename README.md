<h1 align="center">WhiteBit Go SDK</h1>

<p align="center">
  <strong>Official Go SDK for the WhiteBit API — trade, query, and manage your crypto portfolio programmatically.</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-≥ 1.21-00ADD8?style=flat-square&logo=go" alt="Go ≥ 1.21" />
  <img src="https://img.shields.io/badge/license-Apache_2.0-green?style=flat-square" alt="Apache 2.0 license" />
</p>

---

## Prerequisites

| Requirement | Details |
|-------------|---------|
| **Go ≥ 1.21** | Required runtime |
| **WhiteBit account** | Sign up at [whitebit.com](https://whitebit.com) |
| **WhiteBit API key** | Profile → API keys → Create key (Read and/or Trade permissions) |

---

## Installation

```bash
go get github.com/whitebit-exchange/go-sdk
```

---

## Quick Start

### 1. Get your API credentials

1. Log in to [whitebit.com](https://whitebit.com) → **Profile → API keys**
2. Create a new key — choose **Read** and/or **Trade** permissions as needed
3. Copy your **API Key** and **Token**

> Public endpoints (market data, tickers, order book) work without credentials. Private endpoints (account, trading) require both.

### 2. Initialize the client

**Public endpoints only** (market data, tickers, order book):

```go
import (
    "context"
    whitebitclient "github.com/whitebit-exchange/go-sdk/client"
)

client := whitebitclient.NewClient()
```

**Private endpoints** (account, trading — requires HMAC signing):

```go
import (
    "context"
    whitebitclient "github.com/whitebit-exchange/go-sdk/client"
    "github.com/whitebit-exchange/go-sdk/option"
    wbauth "github.com/whitebit-exchange/go-sdk/auth"
)

client := whitebitclient.NewClient(
    option.WithTxcApikey("YOUR_API_KEY"),
    option.WithHTTPClient(wbauth.NewHmacClient("YOUR_API_SECRET")),
)
```

> **Note:** WhiteBit private endpoints use HMAC-SHA512 signing (`X-TXC-PAYLOAD` + `X-TXC-SIGNATURE`).
> `wbauth.NewHmacClient` handles this automatically — no manual signing needed.

---

## Usage Examples

```go
ctx := context.Background()

// Market data (no credentials required)
tickers, _ := client.PublicAPIV4.GetMarketActivity(ctx)
depth, _   := client.PublicAPIV4.GetOrderbook(ctx, &publicapiv4.GetOrderbookRequest{Market: "BTC_USDT"})

// Account
balance, _ := client.AccountEndpoints.GetTradingBalance(ctx)

// Spot trading
order, _   := client.SpotTrading.CreateLimitOrder(ctx, &spottrading.CreateLimitOrderRequest{
    Market: "BTC_USDT",
    Side:   "buy",
    Amount: "0.01",
    Price:  "95000",
})
client.SpotTrading.CancelOrder(ctx, &spottrading.CancelOrderRequest{
    Market:  "BTC_USDT",
    OrderId: order.OrderId,
})

// Main account — transfer & withdraw
client.Transfer.Transfer(ctx, &transfer.TransferRequest{From: "main", To: "spot", Ticker: "USDT", Amount: "100"})
client.Withdraw.CreateWithdraw(ctx, &withdraw.CreateWithdrawRequest{Ticker: "USDT", Amount: "500", Address: "0x..."})
```

---

## Available Modules

| Module | Description |
|--------|-------------|
| `PublicAPIV4` | Tickers, order book, trade history, klines, assets |
| `SpotTrading` | Limit, market, stop-limit, stop-market, bulk orders |
| `CollateralTrading` | Collateral orders, OCO, positions |
| `AccountEndpoints` | Trading balance, open orders, order history |
| `MainAccount` | Main balances, deposit addresses, fee info |
| `Transfer` | Transfer between main and trade accounts |
| `Withdraw` | Withdrawal requests |
| `Codes` | WhiteBit codes — create, apply, history |
| `Fees` | Trading fees |
| `SubAccount` | Sub-account management |

---

## Resources

| | |
|---|---|
| [WhiteBIT API Documentation](https://docs.whitebit.com) | Official API reference |
| [API Platform Overview](https://docs.whitebit.com/private/http-trade-v4/) | REST, WebSocket, authentication, rate limits |
| [Use with AI](https://github.com/whitebit-exchange/whitebit-mcp) | Use API docs with Claude, Cursor, VS Code via MCP |
| [GitHub Repository](https://github.com/whitebit-exchange/go-sdk) | Source code |
| [Releases](https://github.com/whitebit-exchange/go-sdk/releases) | Binaries and changelog |
| [Contributing](CONTRIBUTING.md) | Development setup and contribution guide |
| [Report an Issue](https://github.com/whitebit-exchange/go-sdk/issues) | Bug reports and feature requests |
| [WhiteBIT Exchange](https://whitebit.com) | The exchange |

---

## License

[Apache 2.0](LICENSE.md)
