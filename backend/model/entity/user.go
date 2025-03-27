package entity

import (
	"github.com/chaihaobo/chat/tools/jwt"
)

type User struct {
	BaseEntity
	Username     string `gorm:"column:username"`
	Password     string `gorm:"column:password"`
	Avatar       string `gorm:"column:avatar"`
	GithubUserID string `gorm:"column:github_user_id"`
}

func (User) TableName() string {
	return "user"
}

func (u User) ToJWTClaims() *jwt.UserForToken {
	return &jwt.UserForToken{
		ID:       u.ID,
		UserName: u.Username,
		Avatar:   u.Avatar,
	}
}
