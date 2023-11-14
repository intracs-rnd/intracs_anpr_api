package capture

import (
	"intracs_anpr_api/internal/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (api *CaptureApi) GetImage(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		vars := mux.Vars(r)
		idVar := vars["id"]

		if idVar == "" {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, message.idNull)
			return
		}

		id, err := strconv.Atoi(idVar)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, message.idInvalid)
			return
		}

		result, err := api.controller.GetImage(id)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, "failed fetch images. "+err.Error())
			return
		}

		api.response.SuccessWithLog(w, r, "Get captures image successfully", result, false)
		return
	}

	api.response.MethodInvalid(w)
}

func (api *CaptureApi) GetPlateImage(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		vars := mux.Vars(r)
		idVar := vars["id"]

		if idVar == "" {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, message.idNull)
			return
		}

		id, err := strconv.Atoi(idVar)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, message.idInvalid)
			return
		}

		result, err := api.controller.GetPlateImage(id)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, "failed fetch plate images. "+err.Error())
			return
		}

		api.response.SuccessWithLog(w, r, "Get captured plate image successfully", result, false)
		return
	}
}

func (api *CaptureApi) GetFullImage(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		vars := mux.Vars(r)
		idVar := vars["id"]

		if idVar == "" {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, message.idNull)
			return
		}

		id, err := strconv.Atoi(idVar)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, message.idInvalid)
			return
		}

		result, err := api.controller.GetFullImage(id)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, "failed fetch full images. "+err.Error())
			return
		}

		api.response.SuccessWithLog(w, r, "Get captured full image successfully", result, false)
		return
	}
}
