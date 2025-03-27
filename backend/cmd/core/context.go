package core

import (
	"github.com/chaihaobo/chat/application"
	"github.com/chaihaobo/chat/infrastructure"
	"github.com/chaihaobo/chat/resource"
	"github.com/chaihaobo/chat/transport"
)

type Context struct {
	Resource       resource.Resource
	Infrastructure infrastructure.Infrastructure
	Application    application.Application
	Transport      transport.Transport
}

func NewContext(res resource.Resource, infra infrastructure.Infrastructure, app application.Application, tsp transport.Transport) *Context {
	return &Context{
		Resource:       res,
		Infrastructure: infra,
		Application:    app,
		Transport:      tsp,
	}
}
