package ws

import (
	"context"
	"encoding/json"

	"github.com/chaihaobo/gocommon/trace"
	"github.com/gorilla/websocket"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"

	"github.com/chaihaobo/chat/model/dto/ws"
	"github.com/chaihaobo/chat/resource"
)

type (
	Connection struct {
		res            resource.Resource
		raw            *websocket.Conn
		userID         uint64
		eventCallbacks map[ws.EventType]EventCallBacks
	}
)

func (c *Connection) registerEventCallBacks(event ws.EventType, callbacks ...EventCallback) *Connection {
	c.eventCallbacks[event] = append(c.eventCallbacks[event], callbacks...)
	return c
}

func (c *Connection) close() error {
	return c.raw.Close()
}

func (c *Connection) startRead() {
	ctx, span := otel.Tracer(trace.DefaultTracerName).Start(context.Background(), "connection.startRead")
	defer span.End()
	defer c.close()
	for {
		_, rawPayload, err := c.raw.ReadMessage()
		if err != nil {
			c.res.Logger().Error(ctx, "failed to read message", err)
			return
		}
		var payload ws.Payload
		if err := json.Unmarshal(rawPayload, &payload); err != nil {
			c.res.Logger().Error(ctx, "failed to unmarshal ws payload", err)
			return
		}
		c.res.Logger().Info(ctx, "receive ws payload", zap.Any("payload", payload))
		// handle payload
		c.invokeCallbacks(ctx, &payload)
	}
}

func (c *Connection) ID() uint64 {
	return c.userID
}

func (c *Connection) invokeCallbacks(ctx context.Context, payload *ws.Payload) {
	if callbacks, ok := c.eventCallbacks[payload.Event]; ok {
		callbacks.invoke(ctx, c, payload)
	}
}

func NewConnection(res resource.Resource, userID uint64, raw *websocket.Conn) *Connection {
	return &Connection{
		res:            res,
		raw:            raw,
		userID:         userID,
		eventCallbacks: make(map[ws.EventType]EventCallBacks),
	}
}
