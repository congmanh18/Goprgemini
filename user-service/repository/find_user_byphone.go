package repository

import (
	"context"

	"github.com/congmanh18/TutorTree/user-service/entity"
)

func (u *userRepoImpl) FindUserByPhone(ctx context.Context, user *entity.User) (*entity.User, error) {
	var userEntity entity.User

	var findQuery = u.DB.Executor.WithContext(ctx).
		Where("phone = ?", user.Phone)

	if err := findQuery.Find(&userEntity).Error; err != nil {
		return nil, err
	}

	return &userEntity, nil
}
