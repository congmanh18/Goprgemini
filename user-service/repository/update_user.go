package repository

import (
	"context"
	"reflect"

	"github.com/congmanh18/TutorTree/manh-core/gorm"
	"github.com/congmanh18/TutorTree/user-service/entity"
)

func (u *userRepoImpl) UpdateUser(ctx context.Context, user *entity.User) error {
	var omitFields = gorm.OmitFields(user, func(fieldValue reflect.Value) bool {
		return fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil()
	})

	return u.DB.Executor.WithContext(ctx).
		Omit(omitFields...).
		Updates(user).Error
}
