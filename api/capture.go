package api

import (
	"fmt"
	"intracs_anpr_api/controller/capture"
	"intracs_anpr_api/internal/response"
	"net/http"
	"strconv"
)

type CaptureApi struct {
	controller *capture.Controller
	response   *response.ApiResponse
}

func InitCaptureApi(controller *capture.Controller) *CaptureApi {
	apiResponse := response.NewAPIResponse()
	return &CaptureApi{
		controller: controller,
		response:   apiResponse,
	}
}

func (api *CaptureApi) Create(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "POST" {
		fmt.Println(r.Form)
	}
}

func (api *CaptureApi) Read(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		params := r.URL.Query()

		pageNumberParam := params.Get("page")
		limitParam := params.Get("limit")
		withPlateForm := params.Get("with_plate_image")
		withPlateImage := false

		pageNumber, err := strconv.Atoi(pageNumberParam)
		if err != nil && pageNumberParam != "" {
			api.response.Error(w, http.StatusBadRequest, "Page number can't parse to int")
			return
		}

		pageLimit, err := strconv.Atoi(limitParam)
		if err != nil && limitParam != "" {
			api.response.Error(w, http.StatusBadRequest, "Page limit can't parse to int")
			return
		}

		withPlateImage, err = strconv.ParseBool(withPlateForm)
		if err != nil && withPlateForm != "" {
			api.response.Error(w, http.StatusBadRequest, "with plate image parameter invalid")
			return
		}

		result, paging, err := api.controller.Read(pageLimit, pageNumber, withPlateImage)
		if err != nil {
			api.response.Error(w, http.StatusBadRequest, "Captures can't read, cause:"+err.Error())
			return
		}

		api.response.SuccessWithPaging(w, "Captures read successfully", result, paging)
		return
	}

	api.response.MethodInvalid(w)
}

func (api *CaptureApi) GetImage(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {

		params := r.URL.Query()
		idQuery := params.Get("id")

		if idQuery == "" {
			api.response.Error(w, http.StatusBadRequest, "id must be not null")
			return
		}

		id, err := strconv.Atoi(idQuery)
		if err != nil {
			api.response.Error(w, http.StatusBadRequest, "id can't parsing to int")
			return
		}

		fmt.Println(id)
		result, err := api.controller.GetImage(id)
		if err != nil {
			api.response.Error(w, http.StatusBadRequest, "failed fetch images. "+err.Error())
			return
		}

		api.response.Success(w, "Get captures image successfully", result)
		return
	}

	api.response.MethodInvalid(w)
}

func (api *CaptureApi) GetPlateImage(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {

		params := r.URL.Query()
		idQuery := params.Get("id")

		if idQuery == "" {
			api.response.Error(w, http.StatusBadRequest, "id must be not null")
			return
		}

		id, err := strconv.Atoi(idQuery)
		if err != nil {
			api.response.Error(w, http.StatusBadRequest, "id can't parsing to int")
			return
		}

		fmt.Println(id)
		result, err := api.controller.GetPlateImage(id)
		if err != nil {
			api.response.Error(w, http.StatusBadRequest, "failed fetch plate images. "+err.Error())
			return
		}

		api.response.Success(w, "Get captured plate image successfully", result)
		return
	}
}

func (api *CaptureApi) GetFullImage(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {

		params := r.URL.Query()
		idQuery := params.Get("id")

		if idQuery == "" {
			api.response.Error(w, http.StatusBadRequest, "id must be not null")
			return
		}

		id, err := strconv.Atoi(idQuery)
		if err != nil {
			api.response.Error(w, http.StatusBadRequest, "id can't parsing to int")
			return
		}

		result, err := api.controller.GetFullImage(id)
		if err != nil {
			api.response.Error(w, http.StatusBadRequest, "failed fetch full images. "+err.Error())
			return
		}

		api.response.Success(w, "Get captured full image successfully", result)
		return
	}
}

func (api *CaptureApi) Count(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		var result int64 = -1
		params := r.URL.Query()

		date := params.Get("date")
		today := params.Get("today")
		startDate := params.Get("start_date")
		endDate := params.Get("end_date")
		isToday := false

		isToday, err := strconv.ParseBool(today)
		if err != nil && today != "" {
			api.response.Error(w, http.StatusBadRequest, "today parameter invalid and can't parse to boolean")
			return
		}

		if isToday {
			result, err = api.controller.CountToday()
		} else if startDate != "" || endDate != "" {
			result, err = api.controller.CountBetweenDate(startDate, endDate)
		} else {
			result, err = api.controller.Count(date)
		}

		if err != nil {
			api.response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		api.response.Success(w, "get captures count successfully", result)
		return
	}

	api.response.MethodInvalid(w)
}

func (api *CaptureApi) DetectedCount(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		params := r.URL.Query()

		startDate := params.Get("start_date")
		endDate := params.Get("end_date")

		result, err := api.controller.DetectedCount(startDate, endDate)
		if err != nil {
			api.response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		api.response.Success(w, "get captures detected count successfully", result)
		return
	}

	api.response.MethodInvalid(w)
}

func (api *CaptureApi) RecognizedCount(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		params := r.URL.Query()

		startDate := params.Get("start_date")
		endDate := params.Get("end_date")

		result, err := api.controller.RecognizedCount(startDate, endDate)
		if err != nil {
			api.response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		api.response.Success(w, "get captures recognized count successfully", result)
		return
	}

	api.response.MethodInvalid(w)
}

func (api *CaptureApi) ValidatedCount(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		result, err := api.controller.ValidatedCount()

		if err != nil {
			api.response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		api.response.Success(w, "get validated count successfully", result)
		return
	}
	api.response.MethodInvalid(w)
}

func (api *CaptureApi) UnValidatedCount(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		result, err := api.controller.UnValidatedCount()

		if err != nil {
			api.response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		api.response.Success(w, "get unvalidated count successfully", result)
		return
	}
	api.response.MethodInvalid(w)
}
