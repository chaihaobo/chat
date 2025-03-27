package subscriber

import (
	"github.com/chaihaobo/chat/infrastructure"
	"github.com/chaihaobo/chat/resource"
)

type (
	Transport interface {
		Subscribe() error
		Shutdown() error
	}

	transport struct {
		res   resource.Resource
		infra infrastructure.Infrastructure
	}
)

func (t transport) Subscribe() error {
	return t.infra.Queue().RunSubscriber()
}

func (t transport) Shutdown() error {
	t.infra.Queue().Shutdown()
	return nil
}

func NewTransport(res resource.Resource, infra infrastructure.Infrastructure) Transport {
	return &transport{
		res:   res,
		infra: infra,
	}
}
