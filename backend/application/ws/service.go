package ws

import (
	"context"
	"sync"

	"github.com/gorilla/websocket"

	"github.com/chaihaobo/chat/infrastructure"
	"github.com/chaihaobo/chat/model/dto/ws"
	"github.com/chaihaobo/chat/resource"
	"github.com/chaihaobo/chat/tools"
	"github.com/chaihaobo/chat/tools/gorecovery"
)

type (
	// Service 管理所有的客户连接
	Service interface {
		AcceptWsConnection(ctx context.Context, con *websocket.Conn)
	}
	service struct {
		res         resource.Resource
		infra       infrastructure.Infrastructure
		connections sync.Map
	}
)

func (s *service) AcceptWsConnection(ctx context.Context, con *websocket.Conn) {
	connection := NewConnection(s.res, tools.ContextUserID(ctx), con)
	s.connections.Store(connection.ID(), connection)
	connection.registerEventCallBacks(ws.EventSendMessage, EventCallbackFunc(s.handleSendMessageEvent))
	gorecovery.Go(connection.startRead)
}

func (s *service) handleSendMessageEvent(ctx context.Context, con *Connection, payload *ws.Payload) {
	var request ws.MessageSend
	payload.ScanDataTo(&request)
	toUser := request.To
	s.sendPayload(ctx, toUser, ws.NewPayload(ws.EventReceiveMessage, &ws.MessageReceive{
		From:    con.ID(),
		Content: request.Content,
	}))
}

func (s *service) sendPayload(ctx context.Context, userID uint64, payload *ws.Payload) {
	if con, ok := s.connections.Load(userID); ok {
		if err := con.(*Connection).raw.WriteJSON(payload); err != nil {
			s.res.Logger().Error(ctx, "failed to send ws payload", err)
		}
	}

}

func NewService(res resource.Resource, infra infrastructure.Infrastructure) Service {
	return &service{
		res:   res,
		infra: infra,
	}
}
