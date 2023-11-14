package controller

import (
	"intracs_anpr_api/controller/capture"
	"intracs_anpr_api/controller/user"
	"intracs_anpr_api/controller/validationreason"
	"intracs_anpr_api/repository"
)

type Controllers struct {
	User             *user.Controller
	Capture          *capture.Controller
	ValidationReason *validationreason.Controller
}

func InitControllers(repositories *repository.Repositories) *Controllers {
	validationReason := validationreason.InitController(repositories.ValidationReason)

	return &Controllers{
		User:             user.InitController(repositories.User),
		Capture:          capture.InitController(repositories.Capture, validationReason),
		ValidationReason: validationReason,
	}
}
