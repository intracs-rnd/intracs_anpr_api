package capture

import (
	"errors"
	"intracs_anpr_api/model"
	"strconv"
	"time"
)

func (c *Controller) UpdateValidation(captureId int, reasonCode string) (model.Capture, error) {
	var result model.Capture
	var reason model.ValidationReason
	var err error

	// TODO: Check validation reason by ID and get validStatus
	reason, err = c.reasonController.ReadByCode(reasonCode)
	if err != nil {
		return model.Capture{}, err
	}

	if reason.Code <= 0 {
		return model.Capture{}, errors.New("reason code invalid. reason_code: " + strconv.Itoa(int(reason.Code)))
	}

	result, err = c.service.UpdateValidation(captureId, reason.Code, reason.ValidStatus)
	if err != nil {
		return model.Capture{}, err
	}

	return result, nil
}

func (c *Controller) ValidatedCount(date time.Time) (int64, error) {
	var result int64 = -1
	var err error

	if date.IsZero() {
		result, err = c.service.ValidatedCount()
	} else {
		result, err = c.service.ValidatedCountByDate(date.Format(time.DateOnly))
	}
	if err != nil {
		return -1, errors.New("fetch validated captures count failed")
	}

	return result, nil
}

func (c *Controller) UnValidatedCount(date time.Time) (int64, error) {
	var result int64 = -1
	var err error

	if date.IsZero() {
		result, err = c.service.UnValidatedCount()
	} else {
		result, err = c.service.UnValidatedCountByDate(date.Format(time.DateOnly))
	}
	if err != nil {
		return -1, errors.New("fetch unvalidated captures count failed cause:" + err.Error())
	}

	return result, err
}

func (c *Controller) ValidCount(date time.Time) (int64, error) {
	if date.IsZero() {
		result, err := c.service.ValidCount()
		if err != nil {
			return -1, errors.New("fetch valid captures count failed")
		}

		return result, nil
	} else {
		result, err := c.service.ValidCountByDate(date.Format(time.DateOnly))
		if err != nil {
			return -1, errors.New("fetch valid captures count by specific date failed")
		}

		return result, nil
	}
}

func (c *Controller) ValidCountBetweenDate(startDate time.Time, endDate time.Time) (int64, error) {
	if startDate.IsZero() || endDate.IsZero() {
		return -1, errors.New("start date or end date param must not null")
	}

	result, err := c.service.ValidCountBetweenDate(
		startDate.Format(time.DateOnly),
		endDate.Format(time.DateOnly),
	)
	if err != nil {
		return -1, err
	}

	return result, nil
}

func (c *Controller) InvalidCount(date time.Time) (int64, error) {
	if date.IsZero() {
		result, err := c.service.InvalidCount()
		if err != nil {
			return -1, errors.New("fetch invalid captures count failed")
		}

		return result, nil
	} else {
		result, err := c.service.InvalidCountByDate(date.Format(time.DateOnly))
		if err != nil {
			return -1, errors.New("fetch invalid captures count by specific date failed")
		}

		return result, nil
	}
}

func (c *Controller) InvalidCountBetweenDate(startDate time.Time, endDate time.Time) (int64, error) {
	if startDate.IsZero() || endDate.IsZero() {
		return -1, errors.New("start date or end date param must not null")
	}

	result, err := c.service.InvalidCountBetweenDate(
		startDate.Format(time.DateOnly),
		endDate.Format(time.DateOnly),
	)
	if err != nil {
		return -1, err
	}

	return result, nil
}
