package message

import (
	"github.com/gin-gonic/gin"

	"github.com/chaihaobo/chat/application"
	"github.com/chaihaobo/chat/model/dto/message"
	"github.com/chaihaobo/chat/resource"
)

type (
	Controller interface {
		GetFriendRecentlyMessages(ctx *gin.Context) (*message.GetRecentlyMessagesResponse, error)
	}

	controller struct {
		app application.Application
		res resource.Resource
	}
)

func (c controller) GetFriendRecentlyMessages(ctx *gin.Context) (*message.GetRecentlyMessagesResponse, error) {
	var request message.GetRecentlyMessagesRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		return nil, err
	}
	return c.app.Message().GetFriendRecentlyMessages(ctx, &request)
}

func NewController(res resource.Resource, app application.Application) Controller {
	return &controller{
		app: app,
		res: res,
	}
}
