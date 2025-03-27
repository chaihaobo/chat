package controller

import (
	"github.com/chaihaobo/chat/application"
	"github.com/chaihaobo/chat/resource"
	"github.com/chaihaobo/chat/transport/http/controller/health"
	"github.com/chaihaobo/chat/transport/http/controller/user"
	"github.com/chaihaobo/chat/transport/http/controller/ws"
)

type (
	Controller interface {
		Health() health.Controller
		User() user.Controller
		Ws() ws.Controller
	}

	controllers struct {
		healthController health.Controller
		userController   user.Controller
		wsController     ws.Controller
	}
)

func (c *controllers) Ws() ws.Controller {
	return c.wsController
}

func (c *controllers) User() user.Controller {
	return c.userController
}

func (c *controllers) Health() health.Controller {
	return c.healthController
}

func New(res resource.Resource, app application.Application) Controller {
	return &controllers{
		healthController: health.NewController(res, app),
		userController:   user.NewController(res, app),
		wsController:     ws.NewController(res, app),
	}
}
