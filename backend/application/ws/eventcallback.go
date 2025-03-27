package ws

import (
	"context"

	"github.com/chaihaobo/chat/model/dto/ws"
)

type (
	EventCallback interface {
		Invoke(ctx context.Context, con *Connection, payload *ws.Payload)
	}

	EventCallBacks []EventCallback

	EventCallbackFunc func(ctx context.Context, con *Connection, payload *ws.Payload)
)

func (b EventCallBacks) invoke(ctx context.Context, connection *Connection, payload *ws.Payload) {
	for _, callback := range b {
		callback.Invoke(ctx, connection, payload)
	}

}

func (e EventCallbackFunc) Invoke(ctx context.Context, con *Connection, payload *ws.Payload) {
	e(ctx, con, payload)
}
