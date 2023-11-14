package validationreason

import (
	"intracs_anpr_api/model"
	"intracs_anpr_api/repository/validationreason"
)

type repositories interface {
	ReadAll() ([]model.ValidationReason, error)
	Read(code uint8) (model.ValidationReason, error)
	Create(data model.ValidationReason) (model.ValidationReason, error)
	Update(data model.ValidationReason) (model.ValidationReason, error)
}

type Controller struct {
	service repositories
}

func InitController(validationReasonRepo *validationreason.ValidationReasonRepo) *Controller {
	return &Controller{
		service: validationReasonRepo,
	}
}
