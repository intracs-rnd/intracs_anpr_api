package capture

import (
	"intracs_anpr_api/internal/response"
	"net/http"
)

func (api *CaptureApi) Create(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "POST" {
		result := struct {
			Id int `json:"id"`
		}{
			Id: -1,
		}

		lastInsertId, err := api.controller.Create(r)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		result.Id = lastInsertId
		api.response.SuccessWithLog(w, r, "Add new capture successfully", result, true)
		return
	}
	api.response.MethodInvalid(w)

}
