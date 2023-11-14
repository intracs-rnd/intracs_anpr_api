package validationreason

import (
	"intracs_anpr_api/controller/validationreason"
	"intracs_anpr_api/internal/response"
)

type ValidationReasonApi struct {
	controller *validationreason.Controller
	response   *response.ApiResponse
}

func InitValidationReasonApi(controller *validationreason.Controller) *ValidationReasonApi {
	apiResponse := response.NewAPIResponse()
	return &ValidationReasonApi{
		controller: controller,
		response:   apiResponse,
	}
}
