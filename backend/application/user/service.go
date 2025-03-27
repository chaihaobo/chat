package user

import (
	"context"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"

	"github.com/chaihaobo/chat/constant"
	"github.com/chaihaobo/chat/infrastructure"
	"github.com/chaihaobo/chat/infrastructure/store/repository/user"
	userdto "github.com/chaihaobo/chat/model/dto/user"
	"github.com/chaihaobo/chat/resource"
	"github.com/chaihaobo/chat/tools"
	"github.com/chaihaobo/chat/tools/crypto"
	"github.com/chaihaobo/chat/tools/jwt"
)

type (
	Service interface {
		Login(ctx context.Context, request *userdto.LoginRequest) (*userdto.LoginResponse, error)
		LoginByPassword(ctx context.Context, request *userdto.LoginByPasswordRequest) (*userdto.LoginResponse, error)
		GetUserFriends(ctx context.Context) (userdto.Users, error)
		TokenManger() jwt.TokenManager
	}

	service struct {
		res           resource.Resource
		infra         infrastructure.Infrastructure
		passwordHash  crypto.Hash
		tokenManger   jwt.TokenManager
		oauthProvider goth.Provider
	}
)

func (s *service) GetUserFriends(ctx context.Context) (userdto.Users, error) {
	userID := tools.ContextUserID(ctx)
	friends, err := s.infra.Store().Repository().User().GetFriends(ctx, userID)
	if err != nil {
		s.res.Logger().Error(ctx, "failed to get friends", err)
		return nil, constant.ErrSystemMalfunction
	}
	return userdto.NewUsers(friends), nil
}

func (s *service) TokenManger() jwt.TokenManager {
	return s.tokenManger
}

func (s *service) userRepository() user.Repository {
	return s.infra.Store().Repository().User()
}

func NewService(res resource.Resource, infra infrastructure.Infrastructure) Service {
	jwtConfig := res.Configuration().JWT
	return &service{
		res:   res,
		infra: infra,
		passwordHash: crypto.NewArgon2IDHash(&crypto.GeneratePwdParams{
			Memory:      defaultMemory,
			Iterations:  defaultIterations,
			Parallelism: defaultParallelism,
			SaltLength:  defaultSaltLength,
			KeyLength:   defaultKeyLength,
		}),

		tokenManger: jwt.NewJWTManager(
			jwtConfig.AccessTokenSecretKey,
			jwtConfig.RefreshTokenSecretKey,
			jwtConfig.AccessTokenDuration,
			jwtConfig.RefreshTokenDuration,
		),
		oauthProvider: github.New(res.Configuration().Github.ClientID, res.Configuration().Github.ClientSecret, ""),
	}
}
