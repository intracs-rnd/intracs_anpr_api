package repository

import (
	"intracs_anpr_api/model"

	"gorm.io/gorm"
)

type CaptureValidationRepo struct {
	db *gorm.DB
}

func NewCaptureValidationRepo(db *gorm.DB) *CaptureValidationRepo {
	return &CaptureValidationRepo{
		db: db,
	}
}

func (repo *CaptureValidationRepo) Read(limitSize int, page int) ([]model.CaptureValidation, error) {
	result := []model.CaptureValidation{}
	var err error

	return result, err
}

func (repo *CaptureValidationRepo) ValidCount() (int64, error) {
	var result int64 = -1

	q := repo.db.Table("capture_validation").Select("capture_id").Where("is_valid = 1").Count(&result)

	return result, q.Error
}

func (repo *CaptureValidationRepo) InvalidCount() (int64, error) {
	var result int64 = -1

	q := repo.db.Table("capture_validation").Select("capture_id").Where("is_valid = 0").Count(&result)

	return result, q.Error
}
