package message

import (
	"context"

	"github.com/chaihaobo/chat/constant"
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
	request.FillDefault()
	recentlyMessages, total, err := s.infra.Store().Repository().Message().GetRecentlyMessages(ctx, request.ToQuery(ctx))
	if err != nil {
		s.res.Logger().Error(ctx, "failed to get recently message", err)
		return nil, constant.ErrSystemMalfunction
	}
	return message.NewGetRecentlyMessagesResponse(request, recentlyMessages, total), nil
}

func NewService(res resource.Resource, infra infrastructure.Infrastructure) Service {
	return &service{
		res:   res,
		infra: infra,
	}
}
