package validationreason

import "gorm.io/gorm"

type ValidationReasonRepo struct {
	db *gorm.DB
}

func NewValidationReasonRepo(db *gorm.DB) *ValidationReasonRepo {
	return &ValidationReasonRepo{
		db: db,
	}
}
