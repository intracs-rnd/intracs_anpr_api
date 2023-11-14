package capture

import (
	"encoding/json"
	"intracs_anpr_api/internal/response"
	"intracs_anpr_api/internal/validator"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (api *CaptureApi) UpdateValidation(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultMethodHeader(w, "PUT")

	if r.Method == "OPTIONS" {
		api.response.SuccessNoData(w, r, "")
		return
	} else if r.Method == "PUT" {
		vars := mux.Vars(r)
		captureIdVar, ok := vars["id"]
		if !ok {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, "id required")
			return
		}

		_, captureId, err := validator.CaptureId(captureIdVar, true)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		reasonCodeForm := r.FormValue("reasonCode")

		// try to check raw data on body
		if reasonCodeForm == "" {
			var requestData struct {
				ReasonCode int `json:"reasonCode"`
			}

			err = json.NewDecoder(r.Body).Decode(&requestData)
			if err != nil {
				api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			}

			reasonCodeForm = strconv.Itoa(requestData.ReasonCode)
		}

		// Send to controller
		result, err := api.controller.UpdateValidation(captureId, reasonCodeForm)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		api.response.SuccessWithLog(w, r, "Update capture validation successfully", result, false)
		return
	}

	api.response.MethodInvalid(w)
}
