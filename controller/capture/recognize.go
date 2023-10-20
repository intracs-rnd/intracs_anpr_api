package capture

import (
	"errors"
	"time"
)

func (c *Controller) RecognizedCount(startDate string, endDate string) (int64, error) {
	var err error
	var result int64 = 0

	_, err = time.Parse(time.DateTime, startDate)
	if err != nil && startDate != "" {
		return result, errors.New("start date can't parse to date time format")
	}

	_, err = time.Parse(time.DateTime, endDate)
	if err != nil && endDate != "" {
		return result, errors.New("end date can't parse to date time format")
	}

	if startDate == "" && endDate == "" {
		result, err = c.service.RecognizedCount()
	} else if startDate == "" && endDate != "" {
		result, err = c.service.RecognizedCountBeforeDate(endDate)
	} else if startDate != "" && endDate == "" {
		result, err = c.service.RecognizedCountByDate(startDate)
	} else if startDate != "" && endDate != "" {
		result, err = c.service.RecognizedCountBetweenDate(startDate, endDate)
	}

	return result, err
}
