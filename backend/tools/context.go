package tools

import (
	"context"

	"github.com/chaihaobo/chat/constant"
)

func ContextUserID(ctx context.Context) uint64 {
	userID, _ := ctx.Value(constant.ContextKeyUserID).(uint64)
	return userID
}

func ContextUserName(ctx context.Context) string {
	userName, _ := ctx.Value(constant.ContextKeyUserName).(string)
	return userName
}

func ContextUserAvatar(ctx context.Context) string {
	avatar, _ := ctx.Value(constant.ContextKeyUserAvatar).(string)
	return avatar
}
