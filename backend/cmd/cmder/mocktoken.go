package cmder

import (
	"context"

	"github.com/samber/lo"
	"github.com/spf13/cobra"

	"github.com/chaihaobo/chat/cmd/core"
	"github.com/chaihaobo/chat/tools/jwt"
)

// httpCmd represents the base command when called without any subcommands
var mockTokenCmd = &cobra.Command{
	Use:   "mockToken",
	Short: "will generate mock tokens",
}

func NewMockToken() core.Cmder {
	return core.CmderFunc(func(ctx *core.Context) *cobra.Command {
		mockTokenCmd.Run = func(cmd *cobra.Command, args []string) {
			println(lo.Must(ctx.Application.User().TokenManger().GenerateAccessToken(context.Background(), &jwt.UserForToken{ID: 1})))
			println(lo.Must(ctx.Application.User().TokenManger().GenerateAccessToken(context.Background(), &jwt.UserForToken{ID: 2})))
		}
		return mockTokenCmd
	})
}
