package validationreason

import (
	"intracs_anpr_api/internal/response"
	"net/http"
)

func (api *ValidationReasonApi) Update(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "PUT" {
		result, err := api.controller.Update(r)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		api.response.SuccessWithLog(w, r, "update capture validation reason successfully", result, false)
		return
	}

	api.response.MethodInvalid(w)
}
