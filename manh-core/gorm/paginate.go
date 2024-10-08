package gorm

import (
	"github.com/congmanh18/TutorTree/manh-core/record"
	"gorm.io/gorm"
)

func Paginate(pagination *record.Pagination, txCountTotalRows *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return nil
}
