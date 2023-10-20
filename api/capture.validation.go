package api

import (
	"intracs_anpr_api/controller/capture/validation"
	"intracs_anpr_api/internal/response"
	"net/http"
)

type CaptureValidationApi struct {
	controller *validation.Controller
	response   *response.ApiResponse
}

func InitCaptureValidationApi(controller *validation.Controller) *CaptureValidationApi {
	apiResponse := response.NewAPIResponse()

	return &CaptureValidationApi{
		controller: controller,
		response:   apiResponse,
	}
}

func (api *CaptureValidationApi) ValidCount(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		result, err := api.controller.ValidCount()
		if err != nil {
			api.response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		api.response.Success(w, "get valid captures count successfully", result)
		return
	}
	api.response.MethodInvalid(w)
}

func (api *CaptureValidationApi) InvalidCount(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		result, err := api.controller.InvalidCount()
		if err != nil {
			api.response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		api.response.Success(w, "get invalid captures count successfully", result)
		return
	}
	api.response.MethodInvalid(w)
}
