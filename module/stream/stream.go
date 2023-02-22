package stream

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/spf13/cast"
	"math/rand"
	"sync"
	"time"
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
	m               *sync.Mutex
	isConnected     bool
	isAuthorized    bool
	Connection      *websocket.Conn
	subscribes      map[string]*Subscription
	commandHandlers map[int64]CommandHandler
	errorHandler    func(err error)
}

func NewStream(ctx context.Context, token string, errorHandler func(err error)) (*Stream, error) {
	stream := &Stream{
		Url:             "wss://api.whitebit.com/ws",
		token:           token,
		m:               &sync.Mutex{},
		subscribes:      map[string]*Subscription{},
		commandHandlers: map[int64]CommandHandler{},
		errorHandler:    errorHandler,
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
				_, message, err := stream.getConnection().ReadMessage()
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

				var event Event
				errUnmarshal := json.Unmarshal(message, &event)
				if err != nil {
					stream.errorHandler(errUnmarshal)
					return
				}

				//command response handling
				if event.Method == "" {
					if cast.ToInt(event.ID) == 0 {
						continue
					}
					var reply CommandReply
					errReply := json.Unmarshal(message, &reply)
					if errReply != nil {
						return
					}
					handler, exists := stream.commandHandlers[event.ID]
					if !exists {
						continue
					}

					delete(stream.commandHandlers, event.ID)
					result, _ := json.Marshal(reply.Result)
					handler.Handler(handler.Command, result)
					continue
				}

				subscribe := stream.subscribes[event.Method]
				subscribe.OnEvent(event)
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

func (stream *Stream) write(msg []byte) error {
	stream.m.Lock()
	stream.Connection.SetWriteDeadline(time.Now().Add(time.Second * 30))
	err := stream.Connection.WriteMessage(websocket.TextMessage, msg)
	stream.m.Unlock()
	return err
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
		stream.subscribes[command.EventMethod] = command
	}
	return nil
}

func (command Command) send(stream *Stream) error {
	msg, err := json.Marshal(command)
	if err != nil {
		return fmt.Errorf("whitebitws subscribe command marshal error: %w", err)
	}
	stream.getConnection().SetWriteDeadline(time.Now().Add(time.Second * 30))
	err = stream.Connection.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return fmt.Errorf("error during stream subscribe: %w", err)
	}

	return nil
}

func (stream *Stream) Unsubscribe(command Command) error {
	msg, err := json.Marshal(command)
	if err != nil {
		return fmt.Errorf("whitebitws unsubscribe command marshal error: %w", err)
	}
	stream.getConnection().SetWriteDeadline(time.Now().Add(time.Second * 30))
	err = stream.write(msg)

	if err != nil {
		return fmt.Errorf("error during stream subscribe: %w", err)
	}

	return nil
}

func (stream *Stream) Query(command Command, callback func(command Command, response []byte)) error {
	for {
		randId := int64(rand.Intn(1000000-1+1) + 1)
		_, exists := stream.commandHandlers[randId]
		if !exists {
			command.Id = randId
			break
		}
	}

	msg, err := json.Marshal(command)
	if err != nil {
		return fmt.Errorf("whitebitws subscribe command marshal error: %w", err)
	}
	stream.commandHandlers[command.Id] = CommandHandler{Command: command, Handler: callback}

	stream.getConnection().SetWriteDeadline(time.Now().Add(time.Second * 30))
	err = stream.write(msg)

	if err != nil {
		return fmt.Errorf("error during stream subscribe: %w", err)
	}

	return nil
}

func (stream *Stream) connect() error {
	c, _, err := websocket.DefaultDialer.Dial(stream.Url, nil)
	if err != nil {
		return fmt.Errorf("whitebitws connection error: %w", err)
	}

	stream.m.Lock()
	stream.Connection = c
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

	var subscribeError error
	for _, subscribe := range stream.subscribes {
		for subscribeError == nil {
			subscribeError = subscribe.send(stream)
			if subscribeError == nil {
				break
			}
		}
	}
	return nil
}

func (stream *Stream) makeDisconnected() {
	if !stream.isAlive() {
		return
	}

	for _, subscribe := range stream.subscribes {
		err := subscribe.UnsubscribeMethod.send(stream)
		if err != nil {
			break
		}
	}
	stream.Close()
}

func (stream *Stream) getConnection() *websocket.Conn {
	stream.m.Lock()
	defer stream.m.Unlock()
	return stream.Connection
}

func (stream *Stream) Close() error {
	stream.m.Lock()
	defer stream.m.Unlock()

	stream.isConnected = false
	return stream.Connection.Close()
}
