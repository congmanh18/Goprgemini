package repository

import (
	"context"

	"github.com/congmanh18/TutorTree/user-service/entity"
)

func (u *userRepoImpl) InsertUser(ctx context.Context, user *entity.User) error {
	return u.DB.Executor.WithContext(ctx).Create(&user).Error
}
