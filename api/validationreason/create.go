package validationreason

import (
	"intracs_anpr_api/internal/response"
	"net/http"
)

func (api *ValidationReasonApi) Create(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "POST" {
		result, err := api.controller.Create(r)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		api.response.SuccessWithLog(w, r, "create new capture validation reason successfully", result, false)
		return
	}

	api.response.MethodInvalid(w)
}
