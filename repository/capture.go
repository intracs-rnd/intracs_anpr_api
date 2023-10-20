package repository

import (
	"fmt"
	"intracs_anpr_api/model"

	"gorm.io/gorm"
)

type CaptureRepo struct {
	db *gorm.DB
}

func NewCaptureRepo(db *gorm.DB) *CaptureRepo {
	return &CaptureRepo{
		db: db,
	}
}

func (repo *CaptureRepo) Read(limitSize int, pageNumber int, withPlateImage bool) ([]model.Capture, int64, error) {
	var rowCount int64 = 0
	var err error
	captures := []model.Capture{}
	capture := model.Capture{}
	offset := (pageNumber - 1) * limitSize
	fields := []string{
		"id",
		"device_id",
		"speed",
		"plate_detect_conf",
		"text_recog_conf",
		"plate_number",
		"width",
		"height",
		"x1",
		"y1",
		"x2",
		"y2",
		"captured_at",
	}

	if withPlateImage {
		fields = append(fields, "plate_image")
	}

	// repo.db.Begin()
	repo.db.Table("captures").Select("id").Count(&rowCount)
	if repo.db.Error != nil {
		// repo.db.Rollback()
		return captures, rowCount, repo.db.Error
	}

	rows, err := repo.db.Table("captures").Select(fields).Order("captured_at ASC").Limit(limitSize).Offset(offset).Rows()
	if err != nil {
		fmt.Println("Read captures data failed.", err)
		// repo.db.Rollback()
		return captures, rowCount, nil
	}
	// repo.db.Commit()

	for rows.Next() {
		capture.Info.PlateCoordinates = make([]int, 4)

		if withPlateImage {
			err = rows.Scan(
				&capture.Id,
				&capture.Info.DeviceId,
				&capture.Info.Speed,
				&capture.Info.PlateDetectConf,
				&capture.Info.TextRecogConf,
				&capture.Info.PlateNumber,
				&capture.Info.Width,
				&capture.Info.Height,
				&capture.Info.PlateCoordinates[0],
				&capture.Info.PlateCoordinates[1],
				&capture.Info.PlateCoordinates[2],
				&capture.Info.PlateCoordinates[3],
				&capture.CapturedAt,
				&capture.Image.Plate,
			)
		} else {
			err = rows.Scan(
				&capture.Id,
				&capture.Info.DeviceId,
				&capture.Info.Speed,
				&capture.Info.PlateDetectConf,
				&capture.Info.TextRecogConf,
				&capture.Info.PlateNumber,
				&capture.Info.Width,
				&capture.Info.Height,
				&capture.Info.PlateCoordinates[0],
				&capture.Info.PlateCoordinates[1],
				&capture.Info.PlateCoordinates[2],
				&capture.Info.PlateCoordinates[3],
				&capture.CapturedAt,
			)

		}

		captures = append(captures, capture)
	}

	return captures, rowCount, err
}

func (repo *CaptureRepo) Create(data model.Capture) (model.Capture, error) {
	// Execute insert query
	repo.db.Exec(
		"INSERT INTO captured (device_id, speed, plate_detect_conf, text_recog_conf, plate_number, width, height, x1, y1, x2, y2, plate_image, full_image) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		data.Info.DeviceId,
		data.Info.Speed,
		data.Info.PlateDetectConf,
		data.Info.TextRecogConf,
		data.Info.PlateNumber,
		data.Info.Width,
		data.Info.Height,
		data.Info.PlateCoordinates[0],
		data.Info.PlateCoordinates[1],
		data.Info.PlateCoordinates[2],
		data.Info.PlateCoordinates[3],
		data.Image.Plate,
		data.Image.Full,
	).Row().Scan()

	return data, repo.db.Error
}

func (repo *CaptureRepo) GetImage(id int) (model.CaptureImage, error) {
	var result model.CaptureImage

	q := repo.db.Table("captures").Select("plate_image", "full_image").Where("id = ?", id).First(&result)

	if q.Error != nil {
		fmt.Println("Failed get capture image", q.Error)
		return result, q.Error
	}

	return result, nil
}

func (repo *CaptureRepo) GetPlateImage(id int) (string, error) {
	var result string

	q := repo.db.Table("captures").Select("plate_image").Where("id = ?", id).Order("captured_at ASC")

	if q.Error != nil {
		fmt.Println("Failed get capture plate image.", q.Error)
		return result, q.Error
	}

	q.Row().Scan(&result)

	return result, nil
}

func (repo *CaptureRepo) GetFullImage(id int) (string, error) {
	var result string

	q := repo.db.Table("captures").Select("full_image").Where("id = ?", id).Order("captured_at ASC")

	if q.Error != nil {
		fmt.Println("Failed get capture full image.", q.Error)
		return result, q.Error
	}

	q.Row().Scan(&result)

	return result, nil
}

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

func (repo *CaptureRepo) ValidatedCount() (int64, error) {
	var result int64 = -1

	q := repo.db.Table("captures").Select("captures.id")
	q.Joins("RIGHT JOIN capture_validation ON captures.id = capture_validation.capture_id")
	q.Count(&result)

	return result, q.Error
}

func (repo *CaptureRepo) UnValidatedCount() (int64, error) {
	var result int64 = -1

	q := repo.db.Table("captures").Select("id")
	q.Where("id NOT IN (?)", repo.db.Table("capture_validation").Select("capture_id"))
	q.Count(&result)

	return result, q.Error
}
