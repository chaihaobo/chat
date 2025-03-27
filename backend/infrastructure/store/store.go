package store

import (
	"github.com/chaihaobo/chat/infrastructure/store/client"
	"github.com/chaihaobo/chat/infrastructure/store/repository"
	"github.com/chaihaobo/chat/resource"
)

type (
	Store interface {
		Client() client.Client
		Repository() repository.Repository
	}
	store struct {
		client     client.Client
		repository repository.Repository
	}
)

func (s *store) Repository() repository.Repository {
	return s.repository
}

func (s *store) Client() client.Client {
	return s.client
}

func New(res resource.Resource) (Store, error) {
	client, err := client.New(res)
	if err != nil {
		return nil, err
	}
	return &store{
		client:     client,
		repository: repository.New(client),
	}, nil

}
