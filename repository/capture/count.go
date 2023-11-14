package capture

import "fmt"

func (repo *CaptureRepo) Count() (int64, error) {
	var result int64

	q := repo.db.Table("captures").Select("id").Count(&result)

	if q.Error != nil {
		fmt.Println("Failed get capture count.", q.Error)
		return result, q.Error
	}
	return result, nil
}

func (repo *CaptureRepo) CountByDate(date string) (int64, error) {
	var result int64

	q := repo.db.Table("captures").Select("id").Where("DATE(captured_at) = ?", date).Count(&result)

	if q.Error != nil {
		fmt.Println("Failed get capture count by date.", q.Error)
		return result, q.Error
	}
	return result, nil
}

func (repo *CaptureRepo) CountBeforeDate(date string) (int64, error) {
	var result int64

	q := repo.db.Table("captures").Select("id").Where("DATE(captured_at) <= ?", date).Count(&result)

	if q.Error != nil {
		fmt.Println("Failed get capture count before date.", q.Error)
		return result, q.Error
	}
	return result, nil
}

func (repo *CaptureRepo) CountBetweenDate(startDate string, endDate string) (int64, error) {
	var result int64

	q := repo.db.Table("captures").Select("id").Where("DATE(captured_at) BETWEEN ? AND ?", startDate, endDate).Count(&result)

	if q.Error != nil {
		fmt.Println("Failed get capture count between date.", q.Error)
		return result, q.Error
	}
	return result, nil
}

func (repo *CaptureRepo) DetectedCount() (int64, error) {
	var result int64

	q := repo.db.Table("captures").Select("plate_detect_conf").Where("plate_detect_conf > 0").Count(&result)

	if q.Error != nil {
		fmt.Println("Failed get capture detected count.", q.Error)
		return result, q.Error
	}
	return result, nil
}

func (repo *CaptureRepo) DetectedCountByDate(date string) (int64, error) {
	var result int64

	q := repo.db.Table("captures").Select("plate_detect_conf").Where("plate_detect_conf > 0 AND DATE(captured_at) = ?", date).Count(&result)

	if q.Error != nil {
		fmt.Println("Failed get capture detected count by date.", q.Error)
		return result, q.Error
	}
	return result, nil
}

func (repo *CaptureRepo) DetectedCountBeforeDate(date string) (int64, error) {
	var result int64

	q := repo.db.Table("captures").Select("plate_detect_conf").Where("plate_detect_conf > 0 AND DATE(captured_at) <= ?", date).Count(&result)

	if q.Error != nil {
		fmt.Println("Failed get capture detected count by date.", q.Error)
		return result, q.Error
	}
	return result, nil
}

func (repo *CaptureRepo) DetectedCountBetweenDate(startDate string, endDate string) (int64, error) {
	var result int64

	q := repo.db.Table("captures").Select("plate_detect_conf").Where("plate_detect_conf > 0 AND DATE(captured_at) BETWEEN ? AND ?", startDate, endDate).Count(&result)

	if q.Error != nil {
		fmt.Println("Failed get capture detected count between date.", q.Error)
		return result, q.Error
	}
	return result, nil
}

func (repo *CaptureRepo) RecognizedCount() (int64, error) {
	var result int64

	q := repo.db.Table("captures").Select("text_recog_conf").Where("text_recog_conf > 0").Count(&result)

	if q.Error != nil {
		fmt.Println("Failed get capture recognized count.", q.Error)
		return result, q.Error
	}
	return result, nil
}

func (repo *CaptureRepo) RecognizedCountByDate(date string) (int64, error) {
	var result int64

	q := repo.db.Table("captures").Select("text_recog_conf").Where("text_recog_conf > 0 AND DATE(captured_at) = ?", date).Count(&result)

	if q.Error != nil {
		fmt.Println("Failed get capture recognized count by date.", q.Error)
		return result, q.Error
	}
	return result, nil
}

func (repo *CaptureRepo) RecognizedCountBeforeDate(date string) (int64, error) {
	var result int64

	q := repo.db.Table("captures").Select("text_recog_conf").Where("text_recog_conf > 0 AND DATE(captured_at) <= ?", date).Count(&result)

	if q.Error != nil {
		fmt.Println("Failed get capture recognized count by date.", q.Error)
		return result, q.Error
	}
	return result, nil
}

func (repo *CaptureRepo) RecognizedCountBetweenDate(startDate string, endDate string) (int64, error) {
	var result int64

	q := repo.db.Table("captures").Select("text_recog_conf").Where("text_recog_conf > 0 AND DATE(captured_at) BETWEEN ? AND ?", startDate, endDate).Count(&result)

	if q.Error != nil {
		fmt.Println("Failed get capture recognized count between date.", q.Error)
		return result, q.Error
	}
	return result, nil
}

func (repo *CaptureRepo) ValidCount() (int64, error) {
	var result int64 = -1

	q := repo.db.Table("captures").Select("id").Where("is_valid = 1").Count(&result)

	return result, q.Error
}

func (repo *CaptureRepo) ValidCountByDate(date string) (int64, error) {
	var result int64 = -1

	q := repo.db.Table("captures").Select("id").Where("is_valid = 1 AND DATE(captured_at) = DATE(?)", date).Count(&result)

	return result, q.Error
}

func (repo *CaptureRepo) ValidCountBetweenDate(startDate string, endDate string) (int64, error) {
	var result int64 = -1

	q := repo.db.Table("captures").Select("id").
		Where("is_valid = 1 AND DATE(captures.captured_at) BETWEEN DATE(?) AND DATE(?)", startDate, endDate).
		Count(&result)

	return result, q.Error
}

func (repo *CaptureRepo) InvalidCount() (int64, error) {
	var result int64 = -1

	q := repo.db.Table("captures").Select("id").Where("is_valid = 0").Count(&result)

	return result, q.Error
}

func (repo *CaptureRepo) InvalidCountByDate(date string) (int64, error) {
	var result int64 = -1

	q := repo.db.Table("captures").Select("id").Where("is_valid = 0 AND DATE(captured_at) = DATE(?)", date).Count(&result)

	return result, q.Error
}

func (repo *CaptureRepo) InvalidCountBetweenDate(startDate string, endDate string) (int64, error) {
	var result int64 = -1

	q := repo.db.Table("captures").Select("id").
		Where("is_valid = 0 AND DATE(captures.captured_at) BETWEEN ? AND ?", startDate, endDate).
		Count(&result)

	return result, q.Error
}

func (repo *CaptureRepo) ValidatedCount() (int64, error) {
	var result int64 = -1

	q := repo.db.Table("captures").Select("id").
		Where("is_valid != -1").
		Count(&result)

	return result, q.Error
}

func (repo *CaptureRepo) ValidatedCountByDate(date string) (int64, error) {
	var result int64 = -1

	q := repo.db.Table("captures").Select("id").
		Where("is_valid != -1 AND DATE(captured_at) = DATE(?)", date).
		Count(&result)

	return result, q.Error
}

func (repo *CaptureRepo) ValidatedCountBetweenDate(startDate string, endDate string) (int64, error) {
	var result int64 = -1

	q := repo.db.Table("captures").Select("id").
		Where("is_valid != -1 AND Date(captured_at) BETWEEN ? AND ?", startDate, endDate).
		Count(&result)

	return result, q.Error
}

func (repo *CaptureRepo) UnValidatedCount() (int64, error) {
	var result int64 = -1

	q := repo.db.Table("captures").Select("id").
		Where("is_valid = -1").
		Count(&result)

	return result, q.Error
}

func (repo *CaptureRepo) UnValidatedCountByDate(date string) (int64, error) {
	var result int64 = -1

	q := repo.db.Table("captures").Select("id").
		Where("is_valid = -1 AND DATE(captured_at) = DATE(?)", date).
		Count(&result)

	return result, q.Error
}

func (repo *CaptureRepo) UnValidatedCountBetweenDate(startDate string, endDate string) (int64, error) {
	var result int64 = -1

	q := repo.db.Table("captures").Select("id").
		Where("is_valid = -1 AND Date(captured_at) BETWEEN Date(?) AND Date(?)", startDate, endDate).
		Count(&result)

	return result, q.Error
}
