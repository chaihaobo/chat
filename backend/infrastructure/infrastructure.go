package infrastructure

import (
	"github.com/chaihaobo/gocommon/queue"

	"github.com/chaihaobo/chat/infrastructure/cache"
	"github.com/chaihaobo/chat/infrastructure/store"
	"github.com/chaihaobo/chat/resource"
)

type (
	Infrastructure interface {
		Store() store.Store
		Cache() cache.Client
		Queue() queue.Queue
		Close() error
	}

	infrastructure struct {
		store store.Store
		cache cache.Client
		queue queue.Queue
	}
)

func (i *infrastructure) Queue() queue.Queue {
	return i.queue
}

func (i *infrastructure) Close() error {
	closeFuncs := []func() error{
		i.store.Client().Close,
		i.cache.Close,
	}
	for _, closeFun := range closeFuncs {
		if err := closeFun(); err != nil {
			return err
		}
	}
	return nil
}

func (i *infrastructure) Cache() cache.Client {
	return i.cache
}

func (i *infrastructure) Store() store.Store {
	return i.store
}

func New(res resource.Resource) (Infrastructure, error) {
	store, err := store.New(res)
	if err != nil {
		return nil, err
	}

	cacheClient, err := cache.NewClient(res)
	if err != nil {
		return nil, err
	}

	redisConfig := res.Configuration().Redis
	redisQueue, err := queue.NewRedisQueue(res.Logger(), redisConfig.Address, redisConfig.Index, redisConfig.Password)
	if err != nil {
		return nil, err
	}

	return &infrastructure{
		store: store,
		cache: cacheClient,
		queue: redisQueue,
	}, nil
}
