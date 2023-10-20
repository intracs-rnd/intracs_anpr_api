package controller

import (
	"intracs_anpr_api/controller/capture"
	capture_validation "intracs_anpr_api/controller/capture/validation"
	"intracs_anpr_api/controller/user"
	"intracs_anpr_api/repository"
)

type Controllers struct {
	User              *user.Controller
	Capture           *capture.Controller
	CaptureValidation *capture_validation.Controller
}

func InitControllers(repositories *repository.Repositories) *Controllers {
	return &Controllers{
		User:              user.InitController(repositories.User),
		Capture:           capture.InitController(repositories.Capture),
		CaptureValidation: capture_validation.InitController(repositories.CapValidation),
	}
}
