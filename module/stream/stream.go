package stream

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

const (
	// ReadTimeout is the maximum duration for read operations on the WebSocket connection.
	ReadTimeout = 60 * time.Second
	// WriteTimeout is the maximum duration for write operations on the WebSocket connection.
	WriteTimeout = 30 * time.Second
)

// CommandHandler associates a Command with its response handler callback.
type CommandHandler struct {
	Command Command
	Handler func(command Command, response []byte)
}

// Event represents a WebSocket event received from the server.
type Event struct {
	Method string `json:"method"`
	Params []any  `json:"params"`
	ID     int64  `json:"id"`
}

// Stream manages a WebSocket connection to the WhiteBit API,
// handling subscriptions, queries, automatic reconnection, and keepalive pings.
type Stream struct {
	URL             string
	token           string
	m               sync.Mutex
	isConnected     bool
	conn            net.Conn
	rw              io.ReadWriter
	subscribes      map[string]*Subscription
	commandHandlers map[int64]CommandHandler
	errorHandler    func(err error)
	randCounter     int64
}

// NewStream creates a new WebSocket stream connection to the WhiteBit API.
// If token is non-empty, the connection will be authorized.
// The errorHandler is called for any asynchronous errors and must not be nil.
// The provided context controls the lifetime of the connection.
func NewStream(ctx context.Context, token string, errorHandler func(err error)) (*Stream, error) {
	if errorHandler == nil {
		return nil, fmt.Errorf("errorHandler must not be nil")
	}
	stream := &Stream{
		URL:             "wss://api.whitebit.com/ws",
		token:           token,
		subscribes:      make(map[string]*Subscription),
		commandHandlers: make(map[int64]CommandHandler),
		errorHandler:    errorHandler,
		randCounter:     1,
	}

	err := stream.connect(ctx)
	if err != nil {
		return nil, err
	}

	err = stream.authorize()
	if err != nil {
		stream.Close()
		return nil, err
	}

	stream.listen(ctx)

	return stream, nil
}

func (stream *Stream) listen(ctx context.Context) {
	stopPingChan := make(chan struct{}, 1)

	go func() {
		ticker := time.NewTicker(time.Second * 15)
		pingCommand := NewPingCommand()
		pingCommand.ID = 0
		defer ticker.Stop()
		for {
			select {
			case <-stopPingChan:
				return
			case <-ticker.C:
				if !stream.isAlive() {
					continue
				}
				commandBytes, err := json.Marshal(pingCommand)
				if err != nil {
					stream.errorHandler(fmt.Errorf("marshal ping command: %w", err))
					continue
				}
				if writeErr := stream.write(commandBytes); writeErr != nil {
					stream.errorHandler(writeErr)
				}
			}
		}
	}()

	go func() {
		defer func() {
			// Use non-blocking send to avoid deadlock if ping goroutine already stopped
			select {
			case stopPingChan <- struct{}{}:
			default:
			}
			stream.makeDisconnected()
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				message, err := stream.readMessage()
				if err != nil {
					stream.errorHandler(err)
					time.Sleep(time.Second)

					reconnectErr := stream.reconnect(ctx)
					if reconnectErr != nil {
						stream.errorHandler(reconnectErr)
						return
					}

					continue
				}

				if len(message) == 0 {
					continue
				}

				var event Event
				errUnmarshal := json.Unmarshal(message, &event)
				if errUnmarshal != nil {
					stream.errorHandler(errUnmarshal)
					continue
				}

				// command response handling
				if event.Method == "" {
					if event.ID == 0 {
						continue
					}
					var reply CommandReply
					errReply := json.Unmarshal(message, &reply)
					if errReply != nil {
						continue
					}

					stream.m.Lock()
					handler, exists := stream.commandHandlers[event.ID]
					if exists {
						delete(stream.commandHandlers, event.ID)
					}
					stream.m.Unlock()

					if !exists {
						continue
					}

					result, err := json.Marshal(reply.Result)
					if err != nil {
						stream.errorHandler(fmt.Errorf("marshal command result: %w", err))
						continue
					}
					handler.Handler(handler.Command, result)
					continue
				}

				stream.m.Lock()
				subscribe := stream.subscribes[event.Method]
				stream.m.Unlock()

				if subscribe != nil {
					subscribe.OnEvent(event)
				}
			}
		}
	}()
}

func (stream *Stream) isAlive() bool {
	stream.m.Lock()
	defer stream.m.Unlock()
	return stream.isConnected
}

func (stream *Stream) readMessage() ([]byte, error) {
	stream.m.Lock()
	rw := stream.rw
	conn := stream.conn
	stream.m.Unlock()

	if rw == nil || conn == nil {
		return nil, fmt.Errorf("connection is nil")
	}

	if err := conn.SetReadDeadline(time.Now().Add(ReadTimeout)); err != nil {
		return nil, fmt.Errorf("set read deadline: %w", err)
	}

	data, err := wsutil.ReadServerText(rw)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (stream *Stream) write(msg []byte) error {
	stream.m.Lock()
	defer stream.m.Unlock()

	if stream.conn == nil {
		return fmt.Errorf("connection is nil")
	}

	if err := stream.conn.SetWriteDeadline(time.Now().Add(WriteTimeout)); err != nil {
		return err
	}

	err := wsutil.WriteClientText(stream.conn, msg)

	// Clear the deadline after write
	if clearErr := stream.conn.SetWriteDeadline(time.Time{}); clearErr != nil && err == nil {
		err = clearErr
	}

	return err
}

func (stream *Stream) authorize() error {
	if stream.token == "" {
		return nil
	}

	authorizeCommand := NewAuthorizeCommand(stream.token)
	commandBytes, err := json.Marshal(authorizeCommand)
	if err != nil {
		return fmt.Errorf("authorize marshal error: %w", err)
	}

	err = stream.write(commandBytes)
	if err != nil {
		return err
	}

	stream.m.Lock()
	conn := stream.conn
	rw := stream.rw
	stream.m.Unlock()

	if conn == nil {
		return fmt.Errorf("connection is nil")
	}

	if err = conn.SetReadDeadline(time.Now().Add(10 * time.Second)); err != nil {
		return fmt.Errorf("authorize set read deadline error: %w", err)
	}
	defer conn.SetReadDeadline(time.Time{})

	data, err := wsutil.ReadServerText(rw)
	if err != nil {
		return fmt.Errorf("authorize read response error: %w", err)
	}

	var reply CommandReply
	if err = json.Unmarshal(data, &reply); err != nil {
		return fmt.Errorf("authorize unmarshal error: %w", err)
	}

	if reply.Error != nil {
		return fmt.Errorf("authorize failed: %v", reply.Error)
	}

	return nil
}

// Subscribe sends a subscription command and registers the event handler.
func (stream *Stream) Subscribe(command *Subscription) error {
	command.onError = stream.errorHandler
	err := command.send(stream)
	if err != nil {
		return fmt.Errorf("websocket send command error: %w", err)
	}
	if !command.Command.IsQuery {
		stream.m.Lock()
		stream.subscribes[command.EventMethod] = command
		stream.m.Unlock()
	}
	return nil
}

func (command Command) send(stream *Stream) error {
	msg, err := json.Marshal(command)
	if err != nil {
		return fmt.Errorf("whitebitws subscribe command marshal error: %w", err)
	}
	return stream.write(msg)
}

// Unsubscribe sends an unsubscribe command and removes the subscription from the stream.
func (stream *Stream) Unsubscribe(sub *Subscription) error {
	msg, err := json.Marshal(sub.UnsubscribeMethod)
	if err != nil {
		return fmt.Errorf("whitebitws unsubscribe command marshal error: %w", err)
	}

	stream.m.Lock()
	delete(stream.subscribes, sub.EventMethod)
	stream.m.Unlock()

	return stream.write(msg)
}

// Query sends a one-shot command and invokes the callback with the response.
func (stream *Stream) Query(command Command, callback func(command Command, response []byte)) error {
	stream.m.Lock()
	for {
		stream.randCounter++
		if stream.randCounter > 1000000 {
			stream.randCounter = 1
		}
		randID := stream.randCounter
		_, exists := stream.commandHandlers[randID]
		if !exists {
			command.ID = randID
			break
		}
	}
	stream.commandHandlers[command.ID] = CommandHandler{Command: command, Handler: callback}
	stream.m.Unlock()

	// Marshal outside mutex to improve concurrency
	msg, err := json.Marshal(command)
	if err != nil {
		// Remove the handler if marshal fails
		stream.m.Lock()
		delete(stream.commandHandlers, command.ID)
		stream.m.Unlock()
		return fmt.Errorf("whitebitws query command marshal error: %w", err)
	}

	return stream.write(msg)
}

func (stream *Stream) connect(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	conn, _, _, err := ws.Dial(ctx, stream.URL)
	if err != nil {
		return fmt.Errorf("whitebitws connection error: %w", err)
	}

	stream.m.Lock()
	stream.conn = conn
	stream.rw = conn
	stream.isConnected = true
	stream.m.Unlock()

	return nil
}

func (stream *Stream) reconnect(ctx context.Context) error {
	stream.Close()

	connectionErr := stream.connect(ctx)
	if connectionErr != nil {
		return connectionErr
	}

	authErr := stream.authorize()
	if authErr != nil {
		return authErr
	}

	stream.m.Lock()
	// Notify pending query callbacks that connection was lost
	pendingHandlers := make([]CommandHandler, 0, len(stream.commandHandlers))
	for id, handler := range stream.commandHandlers {
		pendingHandlers = append(pendingHandlers, handler)
		delete(stream.commandHandlers, id)
	}
	subscribes := make([]*Subscription, 0, len(stream.subscribes))
	for _, subscribe := range stream.subscribes {
		subscribes = append(subscribes, subscribe)
	}
	stream.m.Unlock()

	// Notify all pending query callbacks with error response (outside mutex to avoid deadlock)
	for _, handler := range pendingHandlers {
		if handler.Handler != nil {
			// Create error response
			errorResponse := []byte(`{"error":"connection lost during query"}`)
			go handler.Handler(handler.Command, errorResponse)
		}
	}

	for _, subscribe := range subscribes {
		subscribeError := subscribe.send(stream)
		if subscribeError != nil {
			return subscribeError
		}
	}
	return nil
}

func (stream *Stream) makeDisconnected() {
	if !stream.isAlive() {
		return
	}

	stream.m.Lock()
	subscribes := make([]*Subscription, 0, len(stream.subscribes))
	for _, subscribe := range stream.subscribes {
		subscribes = append(subscribes, subscribe)
	}
	stream.m.Unlock()

	for _, subscribe := range subscribes {
		err := subscribe.UnsubscribeMethod.send(stream)
		if err != nil {
			break
		}
	}
	stream.Close()
}

// Close closes the WebSocket connection.
func (stream *Stream) Close() error {
	stream.m.Lock()
	defer stream.m.Unlock()

	stream.isConnected = false
	if stream.conn != nil {
		return stream.conn.Close()
	}
	return nil
}
