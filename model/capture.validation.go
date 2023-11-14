package model

import "time"

type CaptureValidation struct {
	CaptureId int       `gorm:"column:capture_id" json:"capture_id"`
	UserUuid  string    `json:"user_uuid"`
	IsValid   bool      `json:"is_valid"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// type ValidationReason struct {
// 	Code      int       `gorm:"primaryKey" json:"code"`
// 	Reason    string    `json:"reason"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// 	DeletedAt time.Time `json:"deleted_at"`
// }
