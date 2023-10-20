package capture

import (
	"errors"
	"intracs_anpr_api/internal/response"
	"intracs_anpr_api/model"
	"math"
	"time"
)

func (c *Controller) Read(limit int, page int, withPlateImage bool) ([]model.Capture, response.PaginationData, error) {
	currentPage := 1
	currentLimit := 25
	pageCount := 0

	if page > 0 {
		currentPage = page
	}

	if limit > 0 {
		currentLimit = limit
	}

	result, rowCount, err := c.service.Read(currentLimit, currentPage, withPlateImage)

	pageCount = int(math.Ceil((float64(rowCount) / float64(currentLimit))))

	pagingData := response.PaginationData{
		CurrentPage: currentPage,
		PageCount:   pageCount,
		LimitSize:   currentLimit,
	}

	return result, pagingData, err
}

func (c *Controller) Count(date string) (int64, error) {
	var err error
	var result int64 = 0

	_, err = time.Parse(time.DateOnly, date)
	if err != nil && date != "" {
		return result, errors.New("'today' parameter can't parse to date format")
	}

	if date != "" {
		result, err = c.service.CountByDate(date)
	} else {
		result, err = c.service.Count()
	}

	return result, err
}

func (c *Controller) CountToday() (int64, error) {
	var result int64 = 0
	var dateString = ""

	dateString = time.Now().Format(time.DateOnly)
	result, err := c.service.CountByDate(dateString)

	return result, err
}

func (c *Controller) CountBetweenDate(startDate string, endDate string) (int64, error) {
	var result int64 = 0
	var err error

	_, err = time.Parse(time.DateTime, startDate)
	if err != nil && startDate != "" {
		return result, errors.New("start date can't parse to date time format")
	}

	_, err = time.Parse(time.DateTime, endDate)
	if err != nil && endDate != "" {
		return result, errors.New("end date can't parse to date time format")
	}

	if startDate == "" && endDate != "" {
		result, err = c.service.CountBeforeDate(endDate)
	} else if startDate != "" && endDate == "" {
		result, err = c.service.CountByDate(startDate)
	} else if startDate != "" && endDate != "" {
		result, err = c.service.CountBetweenDate(startDate, endDate)
	}

	return result, err
}

func (c *Controller) ValidatedCount() (int64, error) {
	var result int64 = -1
	var err error

	result, err = c.service.ValidatedCount()

	return result, err
}

func (c *Controller) UnValidatedCount() (int64, error) {
	var result int64 = -1
	var err error

	result, err = c.service.UnValidatedCount()

	return result, err
}
