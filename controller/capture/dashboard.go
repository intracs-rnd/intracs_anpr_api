package capture

import (
	"errors"
	"intracs_anpr_api/internal/response"
	"intracs_anpr_api/model"
	"intracs_anpr_api/types"
	"time"
)

func (c *Controller) summaryBetweenDate(startDate time.Time, endDate time.Time) (model.CaptureCountSummary, error) {
	var err error
	var result model.CaptureCountSummary

	startDateString := startDate.Format(time.DateOnly)
	endDateString := endDate.Format(time.DateOnly)

	result.Valid, err = c.ValidCountBetweenDate(startDate, endDate)
	if err != nil {
		return result, errors.New("failed get valid captures count cause: " + err.Error())
	}

	result.Invalid, err = c.InvalidCountBetweenDate(startDate, endDate)
	if err != nil {
		return result, errors.New("failed get invalid captures count cause: " + err.Error())
	}

	result.Total, err = c.service.CountBetweenDate(startDateString, endDateString)
	if err != nil {
		return result, errors.New("failed get captures count, cause: " + err.Error())
	}

	result.Validated, err = c.service.ValidatedCountBetweenDate(startDateString, endDateString)
	if err != nil {
		return result, errors.New("failed get validated captures count, cause: " + err.Error())
	}

	result.Unvalidated, err = c.service.UnValidatedCountBetweenDate(startDateString, endDateString)
	if err != nil {
		return result, errors.New("failed get unvalidated captures count, cause: " + err.Error())
	}

	return result, nil
}

func (c *Controller) summary() (model.CaptureCountSummary, error) {
	var err error
	var result model.CaptureCountSummary

	result.Total, err = c.service.Count()
	if err != nil {
		return result, errors.New("failed get captures count cause: " + err.Error())
	}

	result.Valid, err = c.ValidCount(time.Time{})
	if err != nil {
		return result, errors.New("failed get valid captures count cause: " + err.Error())
	}

	result.Invalid, err = c.InvalidCount(time.Time{})
	if err != nil {
		return result, errors.New("failed get invalid captures count cause: " + err.Error())
	}

	result.Validated, err = c.service.ValidatedCount()
	if err != nil {
		return result, errors.New("failed get validated captures count cause: " + err.Error())
	}

	result.Unvalidated, err = c.service.UnValidatedCount()
	if err != nil {
		return result, errors.New("failed get unvalidated captures count, cause: " + err.Error())
	}

	return result, nil
}

func (c *Controller) Summary(dateFilter types.DateFilter) (model.CaptureCountSummary, error) {
	var result model.CaptureCountSummary
	var err error

	// Summary without date range filter
	if dateFilter.IsBetweenDateValid() {
		result, err = c.summaryBetweenDate(dateFilter.GetStartDate(), dateFilter.GetEndDate())
	} else {
		result, err = c.summary()
	}

	if err != nil {
		return model.CaptureCountSummary{}, err
	}

	return result, nil
}

func (c *Controller) EachDaySummary(pageFilter types.PageFilter, dateFilter types.DateFilter) ([]model.DateSummaryTable, response.PaginationData, error) {
	var results []model.DateSummaryTable
	var result model.DateSummaryTable
	var dateList []time.Time
	var recordCount int64 = 0
	var err error

	if pageFilter.Number == 0 {
		pageFilter.Number = 1
	}

	if pageFilter.Limit == 0 {
		pageFilter.Limit = 25
	}

	if dateFilter.IsBetweenDateValid() {
		dateList, recordCount, err = c.service.SummaryDateListBetweenDate(pageFilter, dateFilter)
	} else {
		dateList, recordCount, err = c.service.SummaryDateList(pageFilter)
	}

	if err != nil {
		return []model.DateSummaryTable{}, response.PaginationData{}, err
	}

	for _, date := range dateList {
		result.Date = date
		result.Count, err = c.Count(date.Format(time.DateOnly))
		result.Valid, err = c.ValidCount(date)
		result.Invalid, err = c.InvalidCount(date)
		result.Validated, err = c.ValidatedCount(date)
		result.Unvalidated, err = c.UnValidatedCount(date)

		if err != nil {
			return []model.DateSummaryTable{}, response.PaginationData{}, err
		}

		results = append(results, result)
	}

	pagingData := response.NewPaging(int(recordCount), int(pageFilter.Number), int(pageFilter.Limit))

	return results, pagingData, nil
}
