package capture

import (
	"intracs_anpr_api/internal/response"
	"net/http"
	"strconv"
	"time"
)

func (api *CaptureApi) Count(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		var result int64 = -1
		params := r.URL.Query()

		date := params.Get("date")
		today := params.Get("today")
		startDate := params.Get("startDate")
		endDate := params.Get("endDate")
		isToday := false

		isToday, err := strconv.ParseBool(today)
		if err != nil && today != "" {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, "today parameter invalid and can't parse to boolean")
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
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		api.response.SuccessWithLog(w, r, "get captures count successfully", result, false)
		return
	}

	api.response.MethodInvalid(w)
}

func (api *CaptureApi) DetectedCount(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		params := r.URL.Query()

		startDate := params.Get("startDate")
		endDate := params.Get("endDate")

		result, err := api.controller.DetectedCount(startDate, endDate)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		api.response.SuccessWithLog(w, r, "get captures detected count successfully", result, false)
		return
	}

	api.response.MethodInvalid(w)
}

func (api *CaptureApi) RecognizedCount(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		params := r.URL.Query()

		startDate := params.Get("startDate")
		endDate := params.Get("endDate")

		result, err := api.controller.RecognizedCount(startDate, endDate)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		api.response.SuccessWithLog(w, r, "get captures recognized count successfully", result, false)
		return
	}

	api.response.MethodInvalid(w)
}

func (api *CaptureApi) ValidatedCount(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		result, err := api.controller.ValidatedCount(time.Time{})

		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		api.response.SuccessWithLog(w, r, "get validated count successfully", result, false)
		return
	}
	api.response.MethodInvalid(w)
}

func (api *CaptureApi) UnValidatedCount(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		result, err := api.controller.UnValidatedCount(time.Time{})

		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		api.response.SuccessWithLog(w, r, "get unvalidated count successfully", result, false)
		return
	}
	api.response.MethodInvalid(w)
}

func (api *CaptureApi) ValidCount(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		var result int64
		params := r.URL.Query()

		startDateParam := params.Get("startDate")
		endDateParam := params.Get("endDate")

		startDate, err := time.Parse(time.DateOnly, startDateParam)
		if err != nil && startDateParam != "" {
			api.response.Error(w, http.StatusBadRequest, "start date invalid and can't parse to date (YYYY-MM-DD)")
			return
		}

		endDate, err := time.Parse(time.DateOnly, endDateParam)
		if err != nil && endDateParam != "" {
			api.response.Error(w, http.StatusBadRequest, "end date invalid and can't parse to date (YYYY-MM-DD)")
			return
		}

		if startDateParam == "" && endDateParam == "" {
			result, err = api.controller.ValidCount(time.Time{})
		} else {
			result, err = api.controller.ValidCountBetweenDate(startDate, endDate)
		}
		if err != nil {
			api.response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		api.response.SuccessWithLog(w, r, "get valid captures count successfully", result, false)
		return
	}
	api.response.MethodInvalid(w)
}

func (api *CaptureApi) InvalidCount(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		var result int64
		params := r.URL.Query()

		startDateParam := params.Get("startDate")
		endDateParam := params.Get("endDate")

		startDate, err := time.Parse(time.DateOnly, startDateParam)
		if err != nil && startDateParam != "" {
			api.response.Error(w, http.StatusBadRequest, "start date invalid and can't parse to date (YYYY-MM-DD)")
			return
		}

		endDate, err := time.Parse(time.DateOnly, endDateParam)
		if err != nil && endDateParam != "" {
			api.response.Error(w, http.StatusBadRequest, "end date invalid and can't parse to date (YYYY-MM-DD)")
			return
		}

		if startDateParam == "" && endDateParam == "" {
			result, err = api.controller.InvalidCount(time.Time{})
		} else {
			result, err = api.controller.InvalidCountBetweenDate(startDate, endDate)
		}
		if err != nil {
			api.response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		api.response.Success(w, "get invalid captures count successfully", result)
		return
	}
	api.response.MethodInvalid(w)
}
