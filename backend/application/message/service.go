package message

import (
	"context"

	"github.com/chaihaobo/chat/infrastructure"
	"github.com/chaihaobo/chat/model/dto/message"
	"github.com/chaihaobo/chat/resource"
)

type (
	Service interface {
		// GetFriendRecentlyMessages 获取当前用户某个好友最近的信息
		GetFriendRecentlyMessages(ctx context.Context, request *message.GetRecentlyMessagesRequest) (*message.GetRecentlyMessagesResponse, error)
	}

	service struct {
		res   resource.Resource
		infra infrastructure.Infrastructure
	}
)

func (s service) GetFriendRecentlyMessages(ctx context.Context, request *message.GetRecentlyMessagesRequest) (*message.GetRecentlyMessagesResponse, error) {
	if err := s.res.Validator().Struct(request); err != nil {
		return nil, err
	}
	s.infra.Store().Repository().Message().GetRecentlyMessages()

	//TODO implement me
	panic("implement me")
}

func NewService(res resource.Resource, infra infrastructure.Infrastructure) Service {
	return &service{
		res:   res,
		infra: infra,
	}
}
