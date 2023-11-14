package repository

import (
	"intracs_anpr_api/repository/capture"
	"intracs_anpr_api/repository/validationreason"

	"gorm.io/gorm"
)

type Repositories struct {
	User             *UserRepo
	Capture          *capture.CaptureRepo
	ValidationReason *validationreason.ValidationReasonRepo
}

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := NewUserRepo(db)
	captureRepo := capture.NewCaptureRepo(db)
	validationReasonRepo := validationreason.NewValidationReasonRepo(db)

	return &Repositories{
		User:             userRepo,
		Capture:          captureRepo,
		ValidationReason: validationReasonRepo,
	}
}
