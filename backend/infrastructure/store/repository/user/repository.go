package user

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/chaihaobo/chat/infrastructure/store/client"
	"github.com/chaihaobo/chat/model/entity"
)

type (
	Repository interface {
		GetByUsername(ctx context.Context, username string) (*entity.User, error)
		GetByGithubUserID(ctx context.Context, githubUserID string) (*entity.User, error)
		Save(ctx context.Context, user *entity.User) error
		GetFriends(ctx context.Context, userID uint64) ([]*entity.User, error)
	}
	repository struct {
		client client.Client
	}
)

func (r repository) GetFriends(ctx context.Context, userID uint64) ([]*entity.User, error) {
	sql := `
select t2.* from friend t1
left join user t2 on t1.to_user_id = t2.id
where t1.user_id = ?;
`
	result := make([]*entity.User, 0)
	if err := r.client.DB(ctx).Raw(sql, userID).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r repository) Save(ctx context.Context, user *entity.User) error {
	return r.client.DB(ctx).Save(user).Error
}

func (r repository) GetByGithubUserID(ctx context.Context, githubUserID string) (*entity.User, error) {
	var data entity.User
	err := r.client.DB(ctx).Where("github_user_id = ?", githubUserID).Last(&data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &data, nil
}

func (r repository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	var data entity.User
	if result := r.client.DB(ctx).Where("username = ?", username).Find(&data); result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}
	return &data, nil
}

func NewRepository(client client.Client) Repository {
	return &repository{
		client: client,
	}
}
