package api

import (
	"intracs_anpr_api/controller"
)

type Api struct {
	User              *UserApi
	Capture           *CaptureApi
	CaptureValidation *CaptureValidationApi
}

func InitApi(controller *controller.Controllers) *Api {
	return &Api{
		User:              InitUserApi(controller.User),
		Capture:           InitCaptureApi(controller.Capture),
		CaptureValidation: InitCaptureValidationApi(controller.CaptureValidation),
	}
}
