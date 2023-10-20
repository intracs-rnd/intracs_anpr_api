package model

import "time"

type User struct {
	Uuid      string    `gorm:"primaryKey" json:"uuid"`
	Fullname  string    `gorm:"not null" json:"fullname"`
	Email     string    `gorm:"unique;index;not null" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRaw struct {
	Uuid        string    `gorm:"primaryKey" json:"uuid"`
	Fullname    string    `gorm:"not null" json:"fullname"`
	Email       string    `gorm:"unique;index;not null" json:"email"`
	Password    string    `gorm:"not null" json:"password"`
	RecoveryKey string    `gorm:"not null" json:"recovery_key"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u User) TableName() string {
	return "users"
}
