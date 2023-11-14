package validationreason

import (
	"intracs_anpr_api/internal/response"
	"net/http"
)

func (api *ValidationReasonApi) ReadAll(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		result, err := api.controller.ReadAll()
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		api.response.SuccessWithLog(w, r, "fetch all capture validation reason", result, false)
		return
	}

	api.response.MethodInvalid(w)
}
