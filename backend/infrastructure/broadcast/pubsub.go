package broadcast

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"

	"github.com/chaihaobo/chat/infrastructure/cache"
	"github.com/chaihaobo/chat/resource"
)

type (
	Pubsub[T any] interface {
		Publish(ctx context.Context, data *T) error
		Subscribe(consumeFunc func(ctx context.Context, data *T))
	}

	pubsub[T any] struct {
		channel     string
		client      *redis.Client
		redisPubsub *redis.PubSub
		res         resource.Resource
	}
)

func (p pubsub[T]) Publish(ctx context.Context, data *T) error {
	rawMessage, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.client.Publish(ctx, p.channel, string(rawMessage)).Err()
}

func (p pubsub[T]) Subscribe(consumeFunc func(ctx context.Context, data *T)) {
	channel := p.redisPubsub.Channel()
	for message := range channel {
		ctx := context.Background()
		var payload T
		if err := json.Unmarshal([]byte(message.Payload), &payload); err != nil {
			p.res.Logger().Error(ctx, "failed to unmarshal broadcast payload", err)
			continue
		}
		consumeFunc(ctx, &payload)
	}
}

func NewPubsub[T any](cacheClient cache.Client, channel string) Pubsub[T] {
	redisPubsub := cacheClient.Raw().Subscribe(context.Background(), channel)
	return &pubsub[T]{
		channel:     channel,
		client:      cacheClient.Raw(),
		redisPubsub: redisPubsub,
	}
}
