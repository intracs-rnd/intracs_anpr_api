package capture

import (
	"intracs_anpr_api/internal/response"
	"intracs_anpr_api/internal/validator"
	"intracs_anpr_api/types"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (api *CaptureApi) ReadById(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		vars := mux.Vars(r)
		captureIdVar, _ := vars["id"]
		_, captureId, err := validator.CaptureId(captureIdVar, true)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		result, err := api.controller.ReadById(captureId)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		api.response.SuccessWithLog(w, r, "read capture by id successfully", result, false)
		return
	}

	api.response.MethodInvalid(w)
}

// func (api *CaptureApi) Read(w http.ResponseWriter, r *http.Request) {
// 	w = *response.SetDefaultHeader(w)

// 	if r.Method == "GET" {
// 		params := r.URL.Query()

// 		validationStatueParam := params.Get("validation_status")
// 		withPlateImage := false

// 		pageNumber, err := strconv.Atoi(params.Get("page"))
// 		if err != nil && params.Get("page") != "" {
// 			api.response.ErrorWithLog(w, r, http.StatusBadRequest, "Page number can't parse to int")
// 			return
// 		}

// 		pageLimit, err := strconv.Atoi(params.Get("limit"))
// 		if err != nil && params.Get("limit") != "" {
// 			api.response.ErrorWithLog(w, r, http.StatusBadRequest, "Page limit can't parse to int")
// 			return
// 		}

// 		startDate, err := time.Parse(time.DateOnly, params.Get("start_date"))
// 		if err != nil && params.Get("start_date") != "" {
// 			api.response.ErrorWithLog(w, r, http.StatusBadRequest, "Start date invalid and can't parse to date (YYYY-MM-DD)")
// 			return
// 		}

// 		endDate, err := time.Parse(time.DateOnly, params.Get("end_date"))
// 		if err != nil && params.Get("end_date") != "" {
// 			api.response.ErrorWithLog(w, r, http.StatusBadRequest, "End date invalid and can't parse to date (YYYY-MM-DD)")
// 			return
// 		}

// 		withPlateImage, err = strconv.ParseBool(params.Get("with_plate_image"))
// 		if err != nil && params.Get("with_plate_image") != "" {
// 			api.response.ErrorWithLog(w, r, http.StatusBadRequest, "with plate image parameter invalid")
// 			return
// 		}

// 		result, paging, err := api.controller.Read(startDate, endDate, validationStatueParam, pageLimit, pageNumber, withPlateImage)

// 		if err != nil {
// 			api.response.ErrorWithLog(w, r, http.StatusBadRequest, "Captures can't read, cause:"+err.Error())
// 			return
// 		}

// 		api.response.SuccessWithPaging(w, "Captures read successfully", result, paging)
// 		return
// 	}

// 	api.response.MethodInvalid(w)
// }

func (api *CaptureApi) Read(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		params := r.URL.Query()

		pageNumber := params.Get("page")
		pageLimit := params.Get("limit")
		startDate := params.Get("startDate")
		endDate := params.Get("endDate")
		date := params.Get("date")
		isValidParam := params.Get("isValid")
		isValidatedParam := params.Get("isValidated")
		withPlateImageParam := params.Get("withPlateImage")

		var pageFilter types.PageFilter
		var dateFilter types.DateFilter
		var validationFilter types.ValidationStatusFilter

		pageFilter, err := pageFilter.FromString(pageLimit, pageNumber, false)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		if startDate != "" && endDate != "" {
			dateFilter, err = dateFilter.FromBetweenDateString(startDate, endDate, false)
		} else if date != "" {
			dateFilter, err = dateFilter.FromDateString(date, false)
		}
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		validationFilter, err = validationFilter.FromString(isValidParam, isValidatedParam)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		withPlateImage, err := strconv.ParseBool(withPlateImageParam)
		if err != nil && withPlateImageParam != "" {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, "with plate image parameter invalid")
			return
		}

		result, paging, err := api.controller.Read(pageFilter, dateFilter, validationFilter, withPlateImage)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, "Captures can't read, cause:"+err.Error())
			return
		}

		if len(result) == 0 {
			api.response.SuccessNoData(w, r, "Captures data not found")
			return
		}

		api.response.SuccessWithPaging(w, r, "Captures read successfully", result, paging)
		return
	}
	api.response.MethodInvalid(w)
}
