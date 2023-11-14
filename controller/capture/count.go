package capture

import (
	"errors"
	"time"
)

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
