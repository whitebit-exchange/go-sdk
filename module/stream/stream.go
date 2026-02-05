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

type CommandHandler struct {
	Command Command
	Handler func(command Command, response []byte)
}

type Event struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
	ID     int64         `json:"id"`
}

type Stream struct {
	Url             string
	token           string
	m               sync.Mutex
	isConnected     bool
	isAuthorized    bool
	conn            net.Conn
	rw              io.ReadWriter
	subscribes      map[string]*Subscription
	commandHandlers map[int64]CommandHandler
	errorHandler    func(err error)
	randCounter     int64
}

func NewStream(ctx context.Context, token string, errorHandler func(err error)) (*Stream, error) {
	stream := &Stream{
		Url:             "wss://api.whitebit.com/ws",
		token:           token,
		subscribes:      make(map[string]*Subscription),
		commandHandlers: make(map[int64]CommandHandler),
		errorHandler:    errorHandler,
		randCounter:     1,
	}
	err := stream.init(ctx)
	if err != nil {
		return nil, err
	}

	err = stream.authorize()
	if err != nil {
		stream.Close()
		return nil, err
	}

	return stream, err
}

func (stream *Stream) init(ctx context.Context) error {
	err := stream.connect()
	if err != nil {
		return err
	}
	stopPingChan := make(chan struct{}, 1)

	go func() {
		ticker := time.NewTicker(time.Second * 15)
		pingCommand := NewPingCommand()
		pingCommand.Id = 0
		defer ticker.Stop()
		for {
			select {
			case <-stopPingChan:
				return
			case <-ticker.C:
				if !stream.isAlive() {
					continue
				}
				commandBytes, _ := json.Marshal(pingCommand)
				err = stream.write(commandBytes)
				if err != nil {
					stream.errorHandler(err)
				}
			}
		}
	}()

	go func() {
		defer func() {
			stopPingChan <- struct{}{}
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

					reconnectErr := stream.reconnect()
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

					result, _ := json.Marshal(reply.Result)
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

	return nil
}

func (stream *Stream) isAlive() bool {
	stream.m.Lock()
	defer stream.m.Unlock()
	return stream.isConnected
}

func (stream *Stream) readMessage() ([]byte, error) {
	stream.m.Lock()
	rw := stream.rw
	stream.m.Unlock()

	if rw == nil {
		return nil, fmt.Errorf("connection is nil")
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

	if err := stream.conn.SetWriteDeadline(time.Now().Add(time.Second * 30)); err != nil {
		return err
	}

	return wsutil.WriteClientText(stream.conn, msg)
}

func (stream *Stream) authorize() error {
	if stream.token == "" {
		return nil
	}

	authorizeCommand := NewAuthorizeCommand(stream.token)
	authorizeCommand.Id = 0
	commandBytes, _ := json.Marshal(authorizeCommand)
	err := stream.write(commandBytes)
	if err != nil {
		return err
	}

	time.Sleep(time.Second * 3)

	return err
}

func (stream *Stream) Subscribe(command *Subscription) error {
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

func (stream *Stream) Unsubscribe(command Command) error {
	msg, err := json.Marshal(command)
	if err != nil {
		return fmt.Errorf("whitebitws unsubscribe command marshal error: %w", err)
	}
	return stream.write(msg)
}

func (stream *Stream) Query(command Command, callback func(command Command, response []byte)) error {
	stream.m.Lock()
	for {
		stream.randCounter++
		if stream.randCounter > 1000000 {
			stream.randCounter = 1
		}
		randId := stream.randCounter
		_, exists := stream.commandHandlers[randId]
		if !exists {
			command.Id = randId
			break
		}
	}

	msg, err := json.Marshal(command)
	if err != nil {
		stream.m.Unlock()
		return fmt.Errorf("whitebitws subscribe command marshal error: %w", err)
	}
	stream.commandHandlers[command.Id] = CommandHandler{Command: command, Handler: callback}
	stream.m.Unlock()

	return stream.write(msg)
}

func (stream *Stream) connect() error {
	conn, _, _, err := ws.Dial(context.Background(), stream.Url)
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

func (stream *Stream) reconnect() error {
	stream.Close()

	connectionErr := stream.connect()
	if connectionErr != nil {
		return connectionErr
	}

	authErr := stream.authorize()
	if authErr != nil {
		return authErr
	}

	stream.m.Lock()
	subscribes := make([]*Subscription, 0, len(stream.subscribes))
	for _, subscribe := range stream.subscribes {
		subscribes = append(subscribes, subscribe)
	}
	stream.m.Unlock()

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

func (stream *Stream) Close() error {
	stream.m.Lock()
	defer stream.m.Unlock()

	stream.isConnected = false
	if stream.conn != nil {
		return stream.conn.Close()
	}
	return nil
}
