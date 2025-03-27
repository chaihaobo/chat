package ws

import (
	"context"
	"encoding/json"
	"sync"

	commonContext "github.com/chaihaobo/gocommon/context"
	"github.com/chaihaobo/gocommon/trace"
	"github.com/gorilla/websocket"
	"github.com/hashicorp/go-multierror"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"

	"github.com/chaihaobo/chat/model/dto/ws"
	"github.com/chaihaobo/chat/resource"
	"github.com/chaihaobo/chat/tools"
)

type (
	Connection struct {
		ctx            context.Context
		res            resource.Resource
		raw            *websocket.Conn
		eventCallbacks map[ws.EventType]EventCallBacks
	}
	Connections     []*Connection
	UserConnections struct {
		mutex       sync.Mutex
		connections map[uint64]Connections
	}
)

func (c Connections) WriteJSON(payload *ws.Payload) error {
	var multiError error
	for _, conn := range c {
		if err := conn.WriteJSON(payload); err != nil {
			multiError = multierror.Append(multiError, err)
		}
	}
	return multiError
}

func (c *Connection) WriteJSON(data any) error {
	return c.raw.WriteJSON(data)
}

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
	return tools.ContextUserID(c.ctx)
}

func (c *Connection) Context() context.Context {
	return c.ctx
}

func (c *Connection) invokeCallbacks(ctx context.Context, payload *ws.Payload) {
	if callbacks, ok := c.eventCallbacks[payload.Event]; ok {
		callbacks.invoke(ctx, c, payload)
	}
}

func NewConnection(ctx context.Context, res resource.Resource, raw *websocket.Conn) *Connection {
	return &Connection{
		ctx:            commonContext.Async(ctx),
		res:            res,
		raw:            raw,
		eventCallbacks: make(map[ws.EventType]EventCallBacks),
	}
}

func (uc *UserConnections) Put(userID uint64, conn *Connection) {
	uc.mutex.Lock()
	defer uc.mutex.Unlock()
	uc.connections[userID] = append(uc.connections[userID], conn)
}

func (uc *UserConnections) Get(userID uint64) Connections {
	uc.mutex.Lock()
	defer uc.mutex.Unlock()
	return uc.connections[userID]
}

func NewUserConnections() *UserConnections {
	return &UserConnections{
		connections: make(map[uint64]Connections),
	}
}
