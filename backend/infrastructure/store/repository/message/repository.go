package message

import (
	"context"

	"github.com/chaihaobo/chat/infrastructure/store/client"
	"github.com/chaihaobo/chat/infrastructure/store/repository/querytypes"
	"github.com/chaihaobo/chat/model/entity"
)

type (
	Repository interface {
		Save(ctx context.Context, data *entity.Message) error
		GetRecentlyMessages(ctx context.Context, query *querytypes.RecentlyMessageQuery) (entity.Messages, error)
	}
	repository struct {
		client client.Client
	}
)

func (r repository) GetRecentlyMessages(ctx context.Context, query *querytypes.RecentlyMessageQuery) (entity.Messages, error) {
	result := make(entity.Messages, 0)
	if err := r.client.DB(ctx).Clauses(query.ToClauses()...).Order("id desc").Offset(query.Offset).Limit(query.Limit).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r repository) Save(ctx context.Context, user *entity.Message) error {
	return r.client.DB(ctx).Save(user).Error
}

func NewRepository(client client.Client) Repository {
	return &repository{
		client: client,
	}
}
