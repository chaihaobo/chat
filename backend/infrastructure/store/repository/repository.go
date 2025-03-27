package repository

import (
	"github.com/chaihaobo/chat/infrastructure/store/client"
	"github.com/chaihaobo/chat/infrastructure/store/repository/message"
	"github.com/chaihaobo/chat/infrastructure/store/repository/user"
)

type (
	Repository interface {
		User() user.Repository
		Message() message.Repository
	}
	repository struct {
		userRepository    user.Repository
		messageRepository message.Repository
	}
)

func (r *repository) Message() message.Repository {
	return r.messageRepository
}

func (r *repository) User() user.Repository {
	return r.userRepository
}

func New(client client.Client) Repository {
	return &repository{
		userRepository:    user.NewRepository(client),
		messageRepository: message.NewRepository(client),
	}
}
