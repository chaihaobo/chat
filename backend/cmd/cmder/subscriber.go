package cmder

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/chaihaobo/chat/cmd/core"
	"github.com/chaihaobo/chat/tools"
	"github.com/chaihaobo/chat/transport/subscriber"
)

// subscriberCmd will run the subscriber to handle the queue message
var subscriberCmd = &cobra.Command{
	Use:   "serveSubscriber",
	Short: "will start the subscriber process",
}

func NewSubscriber() core.Cmder {
	return core.CmderFunc(func(ctx *core.Context) *cobra.Command {
		subscriberCmd.Run = func(cmd *cobra.Command, args []string) {
			listenSubscriber(ctx, ctx.Transport.Subscriber())
		}
		return httpCmd
	})
}

func listenSubscriber(ctx *core.Context, subscriber subscriber.Transport) {
	go func() {
		if err := subscriber.Subscribe(); err != nil {
			ctx.Resource.Logger().Error(context.Background(), "listen subscriber failed", err)
		}
	}()
	tools.GracefulShutdown(subscriber.Shutdown, ctx.Infrastructure.Close, ctx.Resource.Close)
}
