package types

import (
	"errors"
	"time"
)

type DateFilter struct {
	date         time.Time
	between      DateBetweenFilter
	isSingleDate bool
}

func (d *DateFilter) FromDateString(dateStr string, required bool) (DateFilter, error) {
	if required && len(dateStr) == 0 {
		return DateFilter{}, errors.New("date is required")
	}

	date, err := time.Parse(time.DateOnly, dateStr)
	if (required || len(dateStr) > 0) && err != nil {
		return DateFilter{}, errors.New("date invalid and can't parse to date format (YYYY-MM-DD)")
	}

	return DateFilter{
		date:         date,
		isSingleDate: true,
	}, nil
}

func (d *DateFilter) FromDateTimeString(dateStr string, required bool) (DateFilter, error) {
	if required && len(dateStr) == 0 {
		return DateFilter{}, errors.New("date is required")
	}

	date, err := time.Parse(time.DateTime, dateStr)
	if (required || len(dateStr) > 0) && err != nil {
		return DateFilter{}, errors.New("date invalid and can't parse to date format (YYYY-MM-DD)")
	}

	return DateFilter{
		date:         date,
		isSingleDate: true,
	}, nil
}

func (d *DateFilter) FromBetweenDateString(start string, end string, required bool) (DateFilter, error) {
	if required && len(start) == 0 && len(end) > 0 {
		return DateFilter{}, errors.New("start date is required")
	} else if required && len(end) == 0 && len(start) > 0 {
		return DateFilter{}, errors.New("end date is required")
	} else if required && len(start) == 0 && len(end) == 0 {
		return DateFilter{}, errors.New("start date and end date is required")
	}

	startDate, err := time.Parse(time.DateOnly, start)
	if (required || len(start) > 0) && err != nil {
		return DateFilter{}, errors.New("start date invalid and can't parse to date (YYYY-MM-DD)")
	}

	endDate, err := time.Parse(time.DateOnly, end)
	if (required || len(end) > 0) && err != nil {
		return DateFilter{}, errors.New("end date invalid and can't parse to date (YYYY-MM-DD)")
	}

	return DateFilter{
		isSingleDate: false,
		between: DateBetweenFilter{
			startDate: startDate,
			endDate:   endDate,
		},
	}, nil
}

func (d *DateFilter) FromBetweenDateTimeString(start string, end string, required bool) (DateFilter, error) {
	if required && len(start) == 0 && len(end) > 0 {
		return DateFilter{}, errors.New("start date is required")
	} else if required && len(end) == 0 && len(start) > 0 {
		return DateFilter{}, errors.New("end date is required")
	} else if required && len(start) == 0 && len(end) == 0 {
		return DateFilter{}, errors.New("start date and end date is required")
	}

	startDate, err := time.Parse(time.DateTime, start)
	if (required || len(start) > 0) && err != nil {
		return DateFilter{}, errors.New("start date invalid and can't parse to date (YYYY-MM-DD)")
	}

	endDate, err := time.Parse(time.DateTime, end)
	if (required || len(end) > 0) && err != nil {
		return DateFilter{}, errors.New("end date invalid and can't parse to date (YYYY-MM-DD)")
	}

	return DateFilter{
		isSingleDate: false,
		between: DateBetweenFilter{
			startDate: startDate,
			endDate:   endDate,
		},
	}, nil
}

func (d *DateFilter) GetDate() time.Time {
	return d.date
}

func (d *DateFilter) GetStartDate() time.Time {
	return d.between.startDate
}

func (d *DateFilter) GetEndDate() time.Time {
	return d.between.endDate
}

func (d *DateFilter) IsSingleDate() bool {
	return d.isSingleDate
}

func (d *DateFilter) DateString() string {
	if d.isSingleDate {
		return d.date.Format(time.DateOnly)
	}

	return ""
}

func (d *DateFilter) DateTimeString() string {
	if d.isSingleDate {
		return d.date.Format(time.DateTime)
	}

	return ""
}

func (d *DateFilter) StartDateString() string {
	if !d.isSingleDate {
		return d.between.startDate.Format(time.DateOnly)
	}

	return ""
}

func (d *DateFilter) StartDateTimeString() string {
	if !d.isSingleDate {
		return d.between.startDate.Format(time.DateTime)
	}

	return ""
}

func (d *DateFilter) EndDateString() string {
	if !d.isSingleDate {
		return d.between.endDate.Format(time.DateOnly)
	}

	return ""
}

func (d *DateFilter) EndDateTimeString() string {
	if !d.isSingleDate {
		return d.between.endDate.Format(time.DateTime)
	}

	return ""
}

func (d *DateFilter) IsSingleDateValid() bool {
	if d.isSingleDate && !d.date.IsZero() {
		return true
	}

	return false
}

func (d *DateFilter) IsBetweenDateValid() bool {
	if !d.isSingleDate && (!d.between.startDate.IsZero() && !d.between.endDate.IsZero()) {
		return true
	}
	return false
}

func (d *DateFilter) IsStartValid() bool {
	if !d.isSingleDate && !d.between.startDate.IsZero() {
		return true
	}

	return false
}

func (d *DateFilter) IsEndValid() bool {
	if !d.isSingleDate && !d.between.endDate.IsZero() {
		return true
	}

	return false
}
