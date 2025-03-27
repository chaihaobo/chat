package ws

import (
	"context"

	"github.com/chaihaobo/gocommon/queue"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"github.com/chaihaobo/chat/constant"
	"github.com/chaihaobo/chat/infrastructure"
	"github.com/chaihaobo/chat/infrastructure/broadcast"
	userdto "github.com/chaihaobo/chat/model/dto/user"
	"github.com/chaihaobo/chat/model/dto/ws"
	"github.com/chaihaobo/chat/model/entity"
	"github.com/chaihaobo/chat/resource"
	"github.com/chaihaobo/chat/tools/gorecovery"
)

type (
	// Service 管理所有的客户连接
	Service interface {
		AcceptWsConnection(ctx context.Context, con *websocket.Conn)
	}
	service struct {
		res              resource.Resource
		infra            infrastructure.Infrastructure
		userConnections  *UserConnections
		messageBroadcast broadcast.Pubsub[message]
	}
)

func (s *service) AcceptWsConnection(ctx context.Context, con *websocket.Conn) {
	connection := NewConnection(ctx, s.res, con)
	s.userConnections.Put(connection.ID(), connection)
	connection.registerEventCallBacks(ws.EventSendMessage, EventCallbackFunc(s.handleSendMessageEvent))
	gorecovery.Go(connection.startRead)
}

func (s *service) handleSendMessageEvent(ctx context.Context, con *Connection, payload *ws.Payload) {
	var request ws.MessageSend
	payload.ScanDataTo(&request)
	// 改成用mq实现,支持分布式
	message := newMessage(con, &request)
	if err := s.messageBroadcast.Publish(ctx, message); err != nil {
		s.res.Logger().Error(ctx, "failed to publish message", err)
		return
	}
	s.res.Logger().Info(ctx, "send message", zap.Any("message", message))
}

func (s *service) handleBroadcastMessage(ctx context.Context, message *message) {
	s.sendPayload(ctx, message.To.ID, ws.NewPayload(ws.EventReceiveMessage, &ws.MessageReceive{
		From: &userdto.User{
			ID:       message.From.ID,
			UserName: message.From.UserName,
			Avatar:   message.From.Avatar,
		},
		Content: message.Content,
	}))
	// 消息持久化
	msg := message.toEntity()
	if err := s.infra.Queue().Publish(ctx, constant.TopicSentMessage, msg); err != nil {
		s.res.Logger().Error(ctx, "failed to publish message sent event", err)
	}
}

func (s *service) sendPayload(ctx context.Context, to uint64, payload *ws.Payload) {
	if connections := s.userConnections.Get(to); len(connections) > 0 {
		if err := connections.WriteJSON(payload); err != nil {
			s.res.Logger().Error(ctx, "failed to send ws payload", err)
		}
	}
}

func (s *service) handleMessageSentEvent(ctx context.Context, topic string, message *entity.Message) error {
	s.res.Logger().Info(ctx, "receive message sent event", zap.Any("message", message))
	// 持久化
	if err := s.infra.Store().Repository().Message().Save(ctx, message); err != nil {
		s.res.Logger().Error(ctx, "failed to save message", err)
		return constant.ErrSystemMalfunction
	}
	return nil
}

func NewService(res resource.Resource, infra infrastructure.Infrastructure) Service {
	pubsub := broadcast.NewPubsub[message](infra.Cache(), "message")
	svc := &service{
		res:              res,
		infra:            infra,
		messageBroadcast: pubsub,
		userConnections:  NewUserConnections(),
	}
	gorecovery.Go(func() {
		res.Logger().Info(context.Background(), "message broadcast started")
		pubsub.Subscribe(svc.handleBroadcastMessage)
	})
	infra.Queue().SubscribeTo(constant.TopicSentMessage, queue.CreateSubscriber[*entity.Message](svc.handleMessageSentEvent))

	return svc
}
