package stream

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockWSServer struct {
	server     *httptest.Server
	conn       net.Conn
	connMu     sync.Mutex
	messages   [][]byte
	messagesMu sync.Mutex
}

func newMockWSServer(t *testing.T) *mockWSServer {
	m := &mockWSServer{
		messages: make([][]byte, 0),
	}

	m.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			t.Logf("upgrade error: %v", err)
			return
		}

		m.connMu.Lock()
		m.conn = conn
		m.connMu.Unlock()

		go func() {
			defer conn.Close()
			for {
				data, err := wsutil.ReadClientText(conn)
				if err != nil {
					return
				}

				m.messagesMu.Lock()
				m.messages = append(m.messages, data)
				m.messagesMu.Unlock()

				var cmd Command
				if err := json.Unmarshal(data, &cmd); err != nil {
					continue
				}

				switch cmd.Method {
				case PingRequest:
					resp := map[string]any{
						"id":     cmd.ID,
						"result": "pong",
						"error":  nil,
					}
					respData, _ := json.Marshal(resp)
					wsutil.WriteServerText(conn, respData)
				case "authorize":
					resp := map[string]any{
						"id":     cmd.ID,
						"result": map[string]any{"status": "success"},
						"error":  nil,
					}
					respData, _ := json.Marshal(resp)
					wsutil.WriteServerText(conn, respData)
				case KlineSubscribe:
					resp := map[string]any{
						"id":     cmd.ID,
						"result": map[string]any{"status": "success"},
						"error":  nil,
					}
					respData, _ := json.Marshal(resp)
					wsutil.WriteServerText(conn, respData)
				case KlineRequest:
					resp := map[string]any{
						"id":     cmd.ID,
						"result": [][]any{},
						"error":  nil,
					}
					respData, _ := json.Marshal(resp)
					wsutil.WriteServerText(conn, respData)
				case DepthSubscribe:
					resp := map[string]any{
						"id":     cmd.ID,
						"result": map[string]any{"status": "success"},
						"error":  nil,
					}
					respData, _ := json.Marshal(resp)
					wsutil.WriteServerText(conn, respData)
				case LastPriceSubscribe:
					resp := map[string]any{
						"id":     cmd.ID,
						"result": map[string]any{"status": "success"},
						"error":  nil,
					}
					respData, _ := json.Marshal(resp)
					wsutil.WriteServerText(conn, respData)
				}
			}
		}()
	}))

	return m
}

func (m *mockWSServer) getURL() string {
	return "ws" + strings.TrimPrefix(m.server.URL, "http")
}

func (m *mockWSServer) close() {
	m.connMu.Lock()
	if m.conn != nil {
		m.conn.Close()
	}
	m.connMu.Unlock()
	m.server.Close()
}

func (m *mockWSServer) sendEvent(event Event) error {
	m.connMu.Lock()
	conn := m.conn
	m.connMu.Unlock()

	if conn == nil {
		return nil
	}

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return wsutil.WriteServerText(conn, data)
}

func (m *mockWSServer) getMessages() [][]byte {
	m.messagesMu.Lock()
	defer m.messagesMu.Unlock()
	result := make([][]byte, len(m.messages))
	copy(result, m.messages)
	return result
}

func newTestStream(t *testing.T, url string, token string) (*Stream, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	stream := &Stream{
		URL:             url,
		token:           token,
		subscribes:      make(map[string]*Subscription),
		commandHandlers: make(map[int64]CommandHandler),
		errorHandler:    func(err error) {},
		randCounter:     1,
	}

	err := stream.connect(ctx)
	require.NoError(t, err)

	return stream, cancel
}

func TestNewStream_Connect(t *testing.T) {
	server := newMockWSServer(t)
	defer server.close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream := &Stream{
		URL:             server.getURL(),
		token:           "",
		subscribes:      make(map[string]*Subscription),
		commandHandlers: make(map[int64]CommandHandler),
		errorHandler:    func(err error) {},
		randCounter:     1,
	}

	err := stream.connect(ctx)
	require.NoError(t, err)
	stream.listen(ctx)
	defer stream.Close()

	assert.True(t, stream.isAlive())
}

func TestStream_Subscribe(t *testing.T) {
	server := newMockWSServer(t)
	defer server.close()

	stream, cancel := newTestStream(t, server.getURL(), "")
	defer cancel()
	defer stream.Close()

	time.Sleep(100 * time.Millisecond)

	eventReceived := make(chan struct{}, 1)
	handler := func(event KlineUpdateEvent) {
		eventReceived <- struct{}{}
	}

	sub := NewKlineSubscription(handler, "BTC_USDT", 3600)
	err := stream.Subscribe(sub)
	require.NoError(t, err)

	time.Sleep(100 * time.Millisecond)

	assert.NotNil(t, stream.subscribes[KlineUpdate])
}

func TestStream_Query(t *testing.T) {
	server := newMockWSServer(t)
	defer server.close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream := &Stream{
		URL:             server.getURL(),
		token:           "",
		subscribes:      make(map[string]*Subscription),
		commandHandlers: make(map[int64]CommandHandler),
		errorHandler:    func(err error) {},
		randCounter:     1,
	}

	err := stream.connect(ctx)
	require.NoError(t, err)
	stream.listen(ctx)
	defer stream.Close()

	time.Sleep(100 * time.Millisecond)

	responseReceived := make(chan []byte, 1)
	cmd := NewKlineCommand("BTC_USDT", time.Now().Unix()-3600, time.Now().Unix(), 3600)

	err = stream.Query(cmd, func(command Command, response []byte) {
		responseReceived <- response
	})
	require.NoError(t, err)

	select {
	case <-responseReceived:
	case <-time.After(2 * time.Second):
		t.Fatal("timeout waiting for query response")
	}
}

func TestStream_Unsubscribe(t *testing.T) {
	server := newMockWSServer(t)
	defer server.close()

	stream, cancel := newTestStream(t, server.getURL(), "")
	defer cancel()
	defer stream.Close()

	time.Sleep(100 * time.Millisecond)

	handler := func(event KlineUpdateEvent) {}
	sub := NewKlineSubscription(handler, "BTC_USDT", 3600)
	err := stream.Subscribe(sub)
	require.NoError(t, err)

	err = stream.Unsubscribe(sub)
	require.NoError(t, err)

	stream.m.Lock()
	_, exists := stream.subscribes[KlineUpdate]
	stream.m.Unlock()
	assert.False(t, exists)
}

func TestStream_Close(t *testing.T) {
	server := newMockWSServer(t)
	defer server.close()

	stream, cancel := newTestStream(t, server.getURL(), "")
	defer cancel()

	assert.True(t, stream.isAlive())

	err := stream.Close()
	require.NoError(t, err)

	assert.False(t, stream.isAlive())
}

func TestStream_Write(t *testing.T) {
	server := newMockWSServer(t)
	defer server.close()

	stream, cancel := newTestStream(t, server.getURL(), "")
	defer cancel()
	defer stream.Close()

	time.Sleep(100 * time.Millisecond)

	testMsg := []byte(`{"id":1,"method":"ping","params":[]}`)
	err := stream.write(testMsg)
	require.NoError(t, err)

	time.Sleep(100 * time.Millisecond)

	messages := server.getMessages()
	assert.Greater(t, len(messages), 0)
}

func TestStream_Authorize(t *testing.T) {
	server := newMockWSServer(t)
	defer server.close()

	stream, cancel := newTestStream(t, server.getURL(), "test-token")
	defer cancel()
	defer stream.Close()

	err := stream.authorize()
	require.NoError(t, err)
}

func TestCommand_Send(t *testing.T) {
	server := newMockWSServer(t)
	defer server.close()

	stream, cancel := newTestStream(t, server.getURL(), "")
	defer cancel()
	defer stream.Close()

	time.Sleep(100 * time.Millisecond)

	cmd := NewPingCommand()
	err := cmd.send(stream)
	require.NoError(t, err)
}

func TestNewKlineSubscription(t *testing.T) {
	called := false
	handler := func(event KlineUpdateEvent) {
		called = true
	}

	sub := NewKlineSubscription(handler, "BTC_USDT", 3600)

	assert.Equal(t, KlineSubscribe, sub.Command.Method)
	assert.Equal(t, KlineUpdate, sub.EventMethod)
	assert.NotNil(t, sub.OnEvent)
	assert.Equal(t, KlineUnsubscribe, sub.UnsubscribeMethod.Method)

	event := Event{
		Method: KlineUpdate,
		Params: []any{},
	}
	sub.OnEvent(event)
	assert.True(t, called)
}

func TestNewDepthSubscription(t *testing.T) {
	called := false
	handler := func(event MarketDepthUpdateEvent) {
		called = true
	}

	sub := NewMarketDepthSubscription(handler, "ETH_USDT", 10, "0", false)

	assert.Equal(t, DepthSubscribe, sub.Command.Method)
	assert.Equal(t, DepthUpdate, sub.EventMethod)
	assert.NotNil(t, sub.OnEvent)

	event := Event{
		Method: DepthUpdate,
		Params: []any{},
	}
	sub.OnEvent(event)
	assert.True(t, called)
}

func TestNewLastPriceSubscription(t *testing.T) {
	called := false
	handler := func(event LastPriceUpdateEvent) {
		called = true
	}

	sub := NewLastPriceSubscription(handler, []string{"BTC_USDT"})

	assert.Equal(t, LastPriceSubscribe, sub.Command.Method)
	assert.Equal(t, LastPriceUpdate, sub.EventMethod)
	assert.NotNil(t, sub.OnEvent)

	event := Event{
		Method: LastPriceUpdate,
		Params: []any{},
	}
	sub.OnEvent(event)
	assert.True(t, called)
}

func TestNewSpotBalanceSubscription(t *testing.T) {
	called := false
	handler := func(event SpotBalanceUpdateEvent) {
		called = true
	}

	sub := NewSpotBalanceSubscription(handler, "USDT", "BTC")

	assert.Equal(t, SpotBalanceSubscribe, sub.Command.Method)
	assert.Equal(t, SpotBalanceUpdate, sub.EventMethod)
	assert.NotNil(t, sub.OnEvent)

	event := Event{
		Method: SpotBalanceUpdate,
		Params: []any{},
	}
	sub.OnEvent(event)
	assert.True(t, called)
}

func TestNewMarginBalanceSubscription(t *testing.T) {
	called := false
	handler := func(event MarginBalanceUpdateEvent) {
		called = true
	}

	sub := NewMarginBalanceSubscription(handler, "USDT", "BTC")

	assert.Equal(t, MarginBalanceSubscribe, sub.Command.Method)
	assert.Equal(t, MarginBalanceUpdate, sub.EventMethod)
	assert.NotNil(t, sub.OnEvent)

	event := Event{
		Method: MarginBalanceUpdate,
		Params: []any{},
	}
	sub.OnEvent(event)
	assert.True(t, called)
}

func TestNewMarketTradesSubscription(t *testing.T) {
	called := false
	handler := func(event MarketTradesUpdateEvent) {
		called = true
	}

	sub := NewMarketTradesSubscription(handler, []string{"BTC_USDT"})

	assert.Equal(t, TradesSubscribe, sub.Command.Method)
	assert.Equal(t, TradesUpdate, sub.EventMethod)
	assert.NotNil(t, sub.OnEvent)

	event := Event{
		Method: TradesUpdate,
		Params: []any{},
	}
	sub.OnEvent(event)
	assert.True(t, called)
}

func TestNewMarketStatSubscription(t *testing.T) {
	called := false
	handler := func(event MarketStatUpdateEvent) {
		called = true
	}

	sub := NewMarketStatSubscription(handler, []string{"BTC_USDT"})

	assert.Equal(t, MarketStatSubscribe, sub.Command.Method)
	assert.Equal(t, MarketStatUpdate, sub.EventMethod)
	assert.NotNil(t, sub.OnEvent)

	event := Event{
		Method: MarketStatUpdate,
		Params: []any{},
	}
	sub.OnEvent(event)
	assert.True(t, called)
}

func TestNewDealsSubscription(t *testing.T) {
	called := false
	handler := func(event DealsUpdateEvent) {
		called = true
	}

	sub := NewDealsSubscription(handler, []string{"BTC_USDT"})

	assert.Equal(t, DealsSubscribe, sub.Command.Method)
	assert.Equal(t, DealsUpdate, sub.EventMethod)
	assert.NotNil(t, sub.OnEvent)

	event := Event{
		Method: DealsUpdate,
		Params: []any{},
	}
	sub.OnEvent(event)
	assert.True(t, called)
}

func TestNewPendingOrdersSubscription(t *testing.T) {
	called := false
	handler := func(event PendingOrdersUpdateEvent) {
		called = true
	}

	sub := NewPendingOrdersSubscription(handler, "BTC_USDT")

	assert.Equal(t, OrdersPendingSubscribe, sub.Command.Method)
	assert.Equal(t, OrdersPendingUpdate, sub.EventMethod)
	assert.NotNil(t, sub.OnEvent)

	event := Event{
		Method: OrdersPendingUpdate,
		Params: []any{},
	}
	sub.OnEvent(event)
	assert.True(t, called)
}

func TestNewOrderExecutedSubscription(t *testing.T) {
	called := false
	handler := func(event OrderExecutedUpdateEvent) {
		called = true
	}

	sub := NewOrderExecutedSubscription(handler, []string{"BTC_USDT"}, 0)

	assert.Equal(t, OrdersExecutedSubscribe, sub.Command.Method)
	assert.Equal(t, OrdersExecutedUpdate, sub.EventMethod)
	assert.NotNil(t, sub.OnEvent)

	event := Event{
		Method: OrdersExecutedUpdate,
		Params: []any{},
	}
	sub.OnEvent(event)
	assert.True(t, called)
}

func TestTransformEvent(t *testing.T) {
	event := Event{
		Method: KlineUpdate,
		Params: []any{"data1", "data2"},
	}

	result, err := TransformEvent[KlineUpdateEvent](event)
	require.NoError(t, err)
	assert.Len(t, result.Params, 2)
}

func TestStream_ConcurrentWrite(t *testing.T) {
	server := newMockWSServer(t)
	defer server.close()

	stream, cancel := newTestStream(t, server.getURL(), "")
	defer cancel()
	defer stream.Close()

	time.Sleep(100 * time.Millisecond)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			msg, _ := json.Marshal(map[string]any{
				"id":     id,
				"method": "ping",
				"params": []any{},
			})
			err := stream.write(msg)
			assert.NoError(t, err)
		}(i)
	}
	wg.Wait()

	time.Sleep(200 * time.Millisecond)
	messages := server.getMessages()
	assert.GreaterOrEqual(t, len(messages), 10)
}

func TestNewPingCommand(t *testing.T) {
	cmd := NewPingCommand()
	assert.Equal(t, PingRequest, cmd.Method)
	assert.True(t, cmd.IsQuery)
}

func TestNewTimeCommand(t *testing.T) {
	cmd := NewTimeCommand()
	assert.Equal(t, TimeRequest, cmd.Method)
	assert.True(t, cmd.IsQuery)
}

func TestNewKlineCommand(t *testing.T) {
	cmd := NewKlineCommand("BTC_USDT", 1000, 2000, 3600)
	assert.Equal(t, KlineRequest, cmd.Method)
	assert.True(t, cmd.IsQuery)
	assert.Len(t, cmd.Params, 4)
}

func TestNewMarketDepthCommand(t *testing.T) {
	cmd := NewMarketDepthCommand("BTC_USDT", 10, "0")
	assert.Equal(t, DepthRequest, cmd.Method)
	assert.True(t, cmd.IsQuery)
	assert.Len(t, cmd.Params, 3)
}

func TestNewLastPriceCommand(t *testing.T) {
	cmd := NewLastPriceCommand("BTC_USDT")
	assert.Equal(t, LastPriceRequest, cmd.Method)
	assert.True(t, cmd.IsQuery)
	assert.Len(t, cmd.Params, 1)
}

func TestNewMarketStatCommand(t *testing.T) {
	cmd := NewMarketStatCommand("BTC_USDT", 86400)
	assert.Equal(t, MarketStatRequest, cmd.Method)
	assert.True(t, cmd.IsQuery)
	assert.Len(t, cmd.Params, 2)
}

func TestNewMarketTradesCommand(t *testing.T) {
	cmd := NewMarketTradesCommand("BTC_USDT", 100, 0)
	assert.Equal(t, TradesRequest, cmd.Method)
	assert.True(t, cmd.IsQuery)
	assert.Len(t, cmd.Params, 3)
}

func TestNewSpotBalanceCommand(t *testing.T) {
	cmd := NewSpotBalanceCommand([]string{"BTC", "USDT"})
	assert.Equal(t, SpotBalanceRequest, cmd.Method)
	assert.True(t, cmd.IsQuery)
	assert.Len(t, cmd.Params, 2)
}

func TestNewMarginBalanceCommand(t *testing.T) {
	cmd := NewMarginBalanceCommand([]string{"BTC", "USDT"})
	assert.Equal(t, MarginBalanceRequest, cmd.Method)
	assert.True(t, cmd.IsQuery)
	assert.Len(t, cmd.Params, 2)
}

func TestNewDealsCommand(t *testing.T) {
	cmd := NewDealsCommand("BTC_USDT", 0, 100)
	assert.Equal(t, DealsRequest, cmd.Method)
	assert.True(t, cmd.IsQuery)
	assert.Len(t, cmd.Params, 3)
}

func TestNewPendingOrdersCommand(t *testing.T) {
	cmd := NewPendingOrdersCommand("BTC_USDT", 0, 100)
	assert.Equal(t, OrderPendingRequest, cmd.Method)
	assert.True(t, cmd.IsQuery)
	assert.Len(t, cmd.Params, 3)
}

func TestNewOrdersExecutedCommand(t *testing.T) {
	cmd := NewOrdersExecutedCommand("BTC_USDT", []int64{1, 2}, 100, 0)
	assert.Equal(t, OrdersExecutedRequest, cmd.Method)
	assert.True(t, cmd.IsQuery)
	assert.Len(t, cmd.Params, 3)
}

func TestNewAuthorizeCommand(t *testing.T) {
	cmd := NewAuthorizeCommand("test-token")
	assert.Equal(t, "authorize", cmd.Method)
	assert.True(t, cmd.IsQuery)
	assert.Len(t, cmd.Params, 2)
	assert.Equal(t, "test-token", cmd.Params[0])
	assert.Equal(t, "go-sdk", cmd.Params[1])
}
