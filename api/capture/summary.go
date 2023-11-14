package capture

import (
	"intracs_anpr_api/internal/response"
	"intracs_anpr_api/types"
	"net/http"
)

func (api *CaptureApi) Summary(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		params := r.URL.Query()

		startDate := params.Get("startDate")
		endDate := params.Get("endDate")

		var dateFilter types.DateFilter
		dateFilter, err := dateFilter.FromBetweenDateString(startDate, endDate, false)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		result, err := api.controller.Summary(dateFilter)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		api.response.SuccessWithLog(w, r, "fetch captures summary successfully", result, false)
		return
	}

	api.response.MethodInvalid(w)
}

func (api *CaptureApi) SummaryEachDay(w http.ResponseWriter, r *http.Request) {
	w = *response.SetDefaultHeader(w)

	if r.Method == "GET" {
		var pageFilter types.PageFilter
		var dateFilter types.DateFilter

		params := r.URL.Query()

		page := params.Get("page")
		limit := params.Get("limit")
		startDate := params.Get("startDate")
		endDate := params.Get("endDate")

		pageFilter, err := pageFilter.FromString(limit, page, false)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		dateFilter, err = dateFilter.FromBetweenDateString(startDate, endDate, false)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		result, paging, err := api.controller.EachDaySummary(pageFilter, dateFilter)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			return
		}

		api.response.SuccessWithPaging(w, r, "fetch captures summary each day successfully", result, paging)
		return
	}
	api.response.MethodInvalid(w)
}
