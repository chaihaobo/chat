package user

import (
	"github.com/samber/lo"

	"github.com/chaihaobo/chat/model/entity"
)

type (
	Users []*User
	User  struct {
		ID       uint64 `json:"id" mapstructure:"id"`
		UserName string `json:"username" mapstructure:"username"`
		Avatar   string `json:"avatar" mapstructure:"avatar"`
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

func NewUser(user *entity.User) *User {
	return &User{
		ID:       user.ID,
		UserName: user.Username,
		Avatar:   user.Avatar,
	}
}
