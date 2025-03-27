package user

import (
	"github.com/samber/lo"

	"github.com/chaihaobo/chat/model/entity"
)

type (
	Users []*User
	User  struct {
		ID       uint64 `json:"id"`
		UserName string `json:"username"`
		Avatar   string `json:"avatar"`
	}
)

func NewUsers(users []*entity.User) Users {
	return lo.Map(users, func(item *entity.User, index int) *User {
		return &User{
			ID:       item.ID,
			UserName: item.Username,
			Avatar:   item.Avatar,
		}
	})
}
