package entity

import (
	"time"

	"github.com/congmanh18/TutorTree/manh-core/record"
)

type User struct {
	record.BaseEntity
	Phone    *string
	Password *string
	Avatar   *string
	Birthday time.Time
	Email    *string
}

func (u *User) TableName() string {
	return "users"
}
