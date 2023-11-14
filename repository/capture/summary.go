package capture

import (
	"errors"
	"intracs_anpr_api/types"
	"time"
)

func (repo *CaptureRepo) SummaryDateList(pageFilter types.PageFilter) ([]time.Time, int64, error) {
	var result []time.Time
	var dateString string
	var count int64

	q := repo.db.Table("captures")
	q.Select("DATE(captured_at)")
	q.Group("DATE(captured_at)")
	q.Count(&count)

	q.Limit(int(pageFilter.Limit))
	q.Offset(int(pageFilter.GetOffset()))
	q.Order("captured_at")

	rows, err := q.Rows()
	if err != nil {
		return []time.Time{}, 0, err
	}

	// Get row count

	for rows.Next() {
		rows.Scan(&dateString)

		date, err := time.Parse(time.RFC3339, dateString)
		if err != nil {
			return []time.Time{}, 0, errors.New("failed parse captured at to date cause: " + err.Error())
		}

		result = append(result, date)
	}

	return result, count, nil
}

func (repo *CaptureRepo) SummaryDateListBetweenDate(pageFilter types.PageFilter, dateFilter types.DateFilter) ([]time.Time, int64, error) {
	var result []time.Time
	var dateString string
	var count int64

	q := repo.db.Table("captures")
	q.Select("DATE(captured_at)")
	q.Where("DATE(captured_at) BETWEEN ? AND ?", dateFilter.StartDateString(), dateFilter.EndDateString())
	q.Group("DATE(captured_at)")
	q.Limit(int(pageFilter.Limit))
	q.Offset(int(pageFilter.GetOffset()))
	q.Order("captured_at")

	rows, err := q.Rows()
	if err != nil {
		return []time.Time{}, 0, err
	}

	q.Count(&count)

	for rows.Next() {
		rows.Scan(&dateString)

		date, err := time.Parse(time.RFC3339, dateString)
		if err != nil {
			return []time.Time{}, 0, errors.New("failed parse captured at to date cause: " + err.Error())
		}

		result = append(result, date)
	}

	return result, count, nil
}
