package tools

import (
	"context"

	"github.com/chaihaobo/chat/constant"
)

func ContextUserID(ctx context.Context) uint64 {
	userID, _ := ctx.Value(constant.ContextKeyUserID).(uint64)
	return userID
}
