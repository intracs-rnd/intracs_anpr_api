package types

import (
	"errors"
	"time"
)



type DateBetweenFilter struct {
	startDate time.Time
	endDate   time.Time
}

func (d *DateBetweenFilter) FromDateString(start string, end string) (DateBetweenFilter, error) {
	startDate, err := time.Parse(time.DateOnly, start)
	if err != nil {
		return *d, errors.New("start date invalid and can't parse to date (YYYY-MM-DD)")
	}

	endDate, err := time.Parse(time.DateOnly, end)
	if err != nil {
		return *d, errors.New("end date invalid and can't parse to date (YYYY-MM-DD)")
	}

	return DateBetweenFilter{
		startDate: startDate,
		endDate:   endDate,
	}, nil
}

func (d DateBetweenFilter) FromDateTimeString(start string, end string) (*DateBetweenFilter, error) {
	startDate, err := time.Parse(time.DateTime, start)
	if err != nil {
		return &d, errors.New("start date invalid and can't parse to date time (YYYY-MM-DD hh:mm:ss)")
	}

	endDate, err := time.Parse(time.DateTime, end)
	if err != nil {
		return &d, errors.New("end date invalid and can't parse to date time (YYYY-MM-DD hh:mm:ss)")
	}

	return &DateBetweenFilter{
		startDate: startDate,
		endDate:   endDate,
	}, nil
}

func (d *DateBetweenFilter) StartDateString() string {
	return d.startDate.Format(time.DateOnly)
}

func (d *DateBetweenFilter) StartDateTimeString() string {
	return d.startDate.Format(time.DateTime)
}

func (d *DateBetweenFilter) EndDateString() string {
	return d.endDate.Format(time.DateOnly)
}

func (d *DateBetweenFilter) EndDateTimeString() string {
	return d.endDate.Format(time.DateTime)
}

func (d *DateBetweenFilter) IsAllValid() bool {
	if d.startDate.IsZero() || d.endDate.IsZero() {
		return false
	}
	return true
}

func (d *DateBetweenFilter) IsStartValid() bool {
	if d.startDate.IsZero() {
		return false
	}
	return true
}

func (d *DateBetweenFilter) IsEndValid() bool {
	if d.endDate.IsZero() {
		return false
	}
	return true
}
