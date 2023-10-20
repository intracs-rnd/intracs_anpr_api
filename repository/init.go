package repository

import (
	"gorm.io/gorm"
)

type Repositories struct {
	User          *UserRepo
	Capture       *CaptureRepo
	CapValidation *CaptureValidationRepo
}

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := NewUserRepo(db)
	captureRepo := NewCaptureRepo(db)
	capValidationRepo := NewCaptureValidationRepo(db)

	return &Repositories{
		User:          userRepo,
		Capture:       captureRepo,
		CapValidation: capValidationRepo,
	}
}
