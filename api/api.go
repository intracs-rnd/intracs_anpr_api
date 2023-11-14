package api

import (
	"intracs_anpr_api/api/capture"
	"intracs_anpr_api/api/validationreason"
	"intracs_anpr_api/controller"
)

type Api struct {
	User             *UserApi
	Capture          *capture.CaptureApi
	ValidationReason *validationreason.ValidationReasonApi
}

func InitApi(controller *controller.Controllers) *Api {
	return &Api{
		User:             InitUserApi(controller.User),
		Capture:          capture.InitCaptureApi(controller.Capture),
		ValidationReason: validationreason.InitValidationReasonApi(controller.ValidationReason),
	}
}
