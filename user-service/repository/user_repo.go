package repository

import (
	"context"

	"github.com/congmanh18/TutorTree/manh-core/db"
	"github.com/congmanh18/TutorTree/user-service/entity"
)

type UserRepo interface {
	InsertUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	FindUserByPhone(ctx context.Context, user *entity.User) (*entity.User, error)
}

type userRepoImpl struct {
	DB *db.Database
}

func NewUserRepo(db *db.Database) UserRepo {
	return &userRepoImpl{
		DB: db,
	}
}
