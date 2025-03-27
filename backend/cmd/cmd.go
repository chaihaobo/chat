package cmd

import (
	"github.com/chaihaobo/chat/application"
	"github.com/chaihaobo/chat/cmd/cmder"
	"github.com/chaihaobo/chat/cmd/core"
	"github.com/chaihaobo/chat/infrastructure"
	"github.com/chaihaobo/chat/resource"
	"github.com/chaihaobo/chat/transport"
)

func Execute() error {
	ctx, err := initialContext()
	if err != nil {
		return err
	}
	return cmder.NewRoot().Command(ctx).Execute()
}

func initialContext() (*core.Context, error) {
	res, err := resource.New("./configuration.yaml")
	if err != nil {
		return nil, err
	}

	infra, err := infrastructure.New(res)
	if err != nil {
		return nil, err
	}
	app := application.New(res, infra)
	tsp := transport.New(res, infra, app)
	ctx := core.NewContext(res, infra, app, tsp)
	return ctx, nil
}
