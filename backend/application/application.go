package application

import (
	"github.com/chaihaobo/chat/application/health"
	"github.com/chaihaobo/chat/application/message"
	"github.com/chaihaobo/chat/application/user"
	"github.com/chaihaobo/chat/application/ws"
	"github.com/chaihaobo/chat/infrastructure"
	"github.com/chaihaobo/chat/resource"
)

type (
	Application interface {
		Health() health.Service
		User() user.Service
		Ws() ws.Service
		Message() message.Service
	}

	application struct {
		health  health.Service
		user    user.Service
		ws      ws.Service
		message message.Service
	}
)

func (a *application) Message() message.Service {
	return a.message
}

func (a *application) Ws() ws.Service {
	return a.ws
}

func (a *application) User() user.Service {
	return a.user
}

func (a *application) Health() health.Service {
	return a.health
}

func New(res resource.Resource, infra infrastructure.Infrastructure) Application {
	return &application{
		health:  health.NewService(res, infra),
		user:    user.NewService(res, infra),
		ws:      ws.NewService(res, infra),
		message: message.NewService(res, infra),
	}
}
