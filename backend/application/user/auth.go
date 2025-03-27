package user

import (
	"context"
	"net/url"
	"time"

	"github.com/markbates/goth"
	"github.com/samber/lo"
	"go.uber.org/zap"

	"github.com/chaihaobo/chat/constant"
	"github.com/chaihaobo/chat/model/dto/user"
	"github.com/chaihaobo/chat/model/entity"
)

const (
	defaultMemory      = 16 * 1024
	defaultIterations  = 2
	defaultParallelism = 2
	defaultSaltLength  = 16
	defaultKeyLength   = 32
)

func (s *service) Login(ctx context.Context, request *user.LoginRequest) (*user.LoginResponse, error) {
	if err := s.res.Validator().Struct(request); err != nil {
		return nil, err
	}
	// github的code
	code := request.Code
	session := lo.Must(s.oauthProvider.BeginAuth(""))
	_, err := session.Authorize(s.oauthProvider, url.Values{
		"code": []string{code},
	})
	if err != nil {
		s.res.Logger().Error(ctx, "failed to authorize oauth provider", err)
		return nil, constant.ErrAuthCodeInvalid
	}
	oauthUser, err := s.oauthProvider.FetchUser(session)
	if err != nil {
		s.res.Logger().Error(ctx, "failed to fetch oauth user", err)
		return nil, constant.ErrAuthCodeInvalid
	}
	s.res.Logger().Info(ctx, "oauth user", zap.Any("oauthUser", oauthUser))

	user, err := s.infra.Store().Repository().User().GetByGithubUserID(ctx, oauthUser.UserID)
	if err != nil {
		s.res.Logger().Error(ctx, "failed to get user by github user id", err)
		return nil, constant.ErrSystemMalfunction
	}
	if user == nil {
		user, err = s.createOauthUser(ctx, oauthUser)
	}

	return s.generateLoginResponse(ctx, user), nil
}

func (s *service) createOauthUser(ctx context.Context, oauthUser goth.User) (*entity.User, error) {
	user := &entity.User{
		BaseEntity: entity.BaseEntity{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username:     oauthUser.Name,
		Avatar:       oauthUser.AvatarURL,
		GithubUserID: oauthUser.UserID,
	}
	if err := s.infra.Store().Repository().User().Save(ctx, user); err != nil {
		s.res.Logger().Error(ctx, "failed to save user", err)
		return nil, constant.ErrSystemMalfunction
	}
	return user, nil
}
func (s *service) generateLoginResponse(ctx context.Context, userEntity *entity.User) *user.LoginResponse {
	jwtClaims := userEntity.ToJWTClaims()
	return &user.LoginResponse{
		ID:           userEntity.ID,
		Avatar:       userEntity.Avatar,
		AccessToken:  lo.Must(s.tokenManger.GenerateAccessToken(ctx, jwtClaims)),
		RefreshToken: lo.Must(s.tokenManger.GenerateRefreshToken(ctx, jwtClaims)),
	}
}

func (s *service) LoginByPassword(ctx context.Context, request *user.LoginByPasswordRequest) (*user.LoginResponse, error) {
	if err := s.res.Validator().Struct(request); err != nil {
		return nil, err
	}
	user, err := s.userRepository().GetByUsername(ctx, request.Username)
	if err != nil {
		s.res.Logger().Error(ctx, "failed to get user by username", err)
		return nil, constant.ErrSystemMalfunction
	}
	if user == nil {
		return nil, constant.ErrUserNotFound
	}
	if request.Password != user.Password {
		return nil, constant.ErrUserPasswordWrong
	}

	return s.generateLoginResponse(ctx, user), nil
}
