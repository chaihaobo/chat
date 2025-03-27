package message

import (
	"context"

	"github.com/chaihaobo/chat/infrastructure/store/client"
	"github.com/chaihaobo/chat/model/entity"
)

type (
	Repository interface {
		Save(ctx context.Context, data *entity.Message) error
	}
	repository struct {
		client client.Client
	}
)

func (r repository) Save(ctx context.Context, user *entity.Message) error {
	return r.client.DB(ctx).Save(user).Error
}

func NewRepository(client client.Client) Repository {
	return &repository{
		client: client,
	}
}
