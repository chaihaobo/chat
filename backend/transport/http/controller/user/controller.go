package user

import (
	"github.com/gin-gonic/gin"

	"github.com/chaihaobo/chat/application"
	"github.com/chaihaobo/chat/model/dto/user"
	"github.com/chaihaobo/chat/resource"
)

type (
	Controller interface {
		Login(ctx *gin.Context) (*user.LoginResponse, error)
		LoginByPassword(ctx *gin.Context) (*user.LoginResponse, error)
	}

	controller struct {
		app application.Application
		res resource.Resource
	}
)

func NewController(res resource.Resource, app application.Application) Controller {
	return &controller{
		app: app,
		res: res,
	}
}
