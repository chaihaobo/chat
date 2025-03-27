package transport

import (
	"context"

	"github.com/chaihaobo/chat/application"
	"github.com/chaihaobo/chat/infrastructure"
	"github.com/chaihaobo/chat/resource"
	"github.com/chaihaobo/chat/transport/http"
	"github.com/chaihaobo/chat/transport/subscriber"
)

type (
	Transport interface {
		ServeAll() error
		ShutdownAll() error
		HTTP() http.Transport
		Subscriber() subscriber.Transport
	}

	transport struct {
		res        resource.Resource
		http       http.Transport
		subscriber subscriber.Transport
	}
)

func (t *transport) Subscriber() subscriber.Transport {
	return t.subscriber
}

func (t *transport) ShutdownAll() error {
	return t.http.Shutdown()
}

func (t *transport) ServeAll() error {
	go func() {
		if err := t.subscriber.Subscribe(); err != nil {
			t.res.Logger().Error(context.Background(), "failed to start subscriber", err)
		}
	}()

	return t.HTTP().Serve()
}

func (t *transport) HTTP() http.Transport {
	return t.http
}

func New(res resource.Resource, infra infrastructure.Infrastructure, app application.Application) Transport {
	httpTransport := http.NewTransport(res, app)
	subscriberTransport := subscriber.NewTransport(res, infra)
	return &transport{
		http:       httpTransport,
		subscriber: subscriberTransport,
	}
}
