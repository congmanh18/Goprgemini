package record

import "time"

import (
	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        *string         `json:"id,omitempty"`
	CreatedAt *time.Time      `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt *time.Time      `gorm:"autoCreateTime" json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}
