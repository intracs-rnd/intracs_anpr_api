package validation

import (
	"intracs_anpr_api/model"
	"intracs_anpr_api/repository"
)

type repositories interface {
	Read(limitSize int, page int) ([]model.CaptureValidation, error)

	ValidCount() (int64, error)
	InvalidCount() (int64, error)
}

type Controller struct {
	service repositories
}

func InitController(captureValidationRepo *repository.CaptureValidationRepo) *Controller {
	return &Controller{
		service: captureValidationRepo,
	}
}
