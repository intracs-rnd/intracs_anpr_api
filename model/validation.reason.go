package model

import (
	"time"

	"gorm.io/gorm"
)

type ValidationReason struct {
	Code        uint8          `gorm:"column:code;primaryKey;-:create" json:"code"`
	ValidStatus int8           `gorm:"column:valid_status" json:"validStatus"`
	Reason      string         `gorm:"column:reason" json:"reason"`
	CreateAt    time.Time      `gorm:"column:created_at;autoCreateTime:milli" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoUpdateTime:milli" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;" json:"deletedAt"`
}
