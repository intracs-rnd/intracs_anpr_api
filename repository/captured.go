package repository

import (
	"database/sql"
	. "intracs_anpr_api/types"
	"net/http"
)

type CapturedRepository interface {
	Read(limitSize int, pageNumber int, w http.ResponseWriter, db *sql.DB) ([]CapturedInfo, error)
	Create(data ImageCaptured, db *sql.DB) (sql.Result, error)

	// Read image
	GetImage(id int, db *sql.DB) (CapturedImage, error)
	GetFullImage(id int, db *sql.DB) (string, error)
	GetPlateImage(id int, db *sql.DB) (string, error)

	// Captured count
	Count(db *sql.DB) (int, error)
	CountByDate(date string, db *sql.DB) (int, error)
	CountBetweenDate(startDate string, endDate string, db *sql.DB) (int, error)

	// Detected count
	DetectedCount(db *sql.DB) (int, error)
	DetectedCountByDate(date string, db *sql.DB) (int, error)
	DetectedCountBetweenDate(startDate string, endDate string, db *sql.DB) (int, error)

	// Recognized count
	RecognizedCount(db *sql.DB) (int, error)
	RecognizedCountByDate(date string, db *sql.DB) (int, error)
	RecognizedCountBetweenDate(startDate string, endDate string, db *sql.DB) (int, error)
}

func Read(limitSize int, pageNumber int, w http.ResponseWriter, db *sql.DB) ([]CapturedInfo, error) {
	offset := (pageNumber - 1) * limitSize

	q, err := db.Query("SELECT id, device_id, speed, plate_detect_conf, text_recog_conf, plate_number, width, height, x1, x2, y1, y2, captured_at FROM captured ORDER BY captured_at ASC LIMIT ? OFFSET ?", limitSize, offset)

	cap := CapturedInfo{}
	result := []CapturedInfo{}

	for q.Next() {
		// coordinates := make([]int, 4)

		cap.PlateCoordinates = make([]int, 4)

		err = q.Scan(
			&cap.Id,
			&cap.DeviceId,
			&cap.Speed,
			&cap.PlateDetectConf,
			&cap.TextRecogConf,
			&cap.PlateNumber,
			&cap.Width,
			&cap.Height,
			&cap.PlateCoordinates[0],
			&cap.PlateCoordinates[1],
			&cap.PlateCoordinates[2],
			&cap.PlateCoordinates[3],
			&cap.CapturedAt,
		)

		result = append(result, cap)
	}

	return result, err
}

func Create(data ImageCaptured, db *sql.DB) (sql.Result, error) {
	// Execute insert query
	insert, err := db.Exec(
		"INSERT INTO captured (device_id, speed, plate_detect_conf, text_recog_conf, plate_number, width, height, x1, y1, x2, y2, plate_image, full_image) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		data.DeviceId,
		data.Speed,
		data.PlateDetectConf,
		data.TextRecogConf,
		data.PlateNumber,
		data.Width,
		data.Height,
		data.PlateCoordinates[0],
		data.PlateCoordinates[1],
		data.PlateCoordinates[2],
		data.PlateCoordinates[3],
		data.PlateImage,
		data.FullImage,
	)

	return insert, err
}

func GetImage(id int, db *sql.DB) (CapturedImage, error) {
	var result CapturedImage

	err := db.QueryRow("SELECT plate_image, full_image FROM captured WHERE id = ?", id).Scan(&result.PlateImage, &result.FullImage)

	return result, err
}

func GetPlateImage(id int, db *sql.DB) (string, error) {
	var result string

	err := db.QueryRow("SELECT plate_image FROM captured WHERE id = ?", id).Scan(&result)

	return result, err
}

func GetFullImage(id int, db *sql.DB) (string, error) {
	var result string

	err := db.QueryRow("SELECT full_image FROM captured WHERE id = ?", id).Scan(&result)

	return result, err
}

func Count(db *sql.DB) (int, error) {
	var result int = 0
	err := db.QueryRow("SELECT COUNT(id) FROM captured").Scan(&result)

	return result, err
}

func CountByDate(date string, db *sql.DB) (int, error) {
	var result int = 0
	err := db.QueryRow("SELECT COUNT(id) FROM captured WHERE DATE(captured_at) = ?", date).Scan(&result)

	return result, err
}

func CountBetweenDate(startDate string, endDate string, db *sql.DB) (int, error) {
	var result int = 0
	err := db.QueryRow("SELECT COUNT(id) FROM captured WHERE DATE(captured_at) BETWEEN ? AND ?", startDate, endDate).Scan(&result)

	return result, err
}

func DetectedCount(db *sql.DB) (int, error) {
	var result int = 0
	err := db.QueryRow("SELECT COUNT(plate_detect_conf) FROM captured WHERE plate_detect_conf > 0").Scan(&result)

	return result, err
}

func DetectedCountByDate(date string, db *sql.DB) (int, error) {
	var result int = 0
	err := db.QueryRow("SELECT COUNT(plate_detect_conf) FROM captured WHERE plate_detect_conf > 0 AND DATE(captured_at) = ?", date).Scan(&result)

	return result, err
}

func DetectedCountBetweenDate(startDate string, endDate string, db *sql.DB) (int, error) {
	var result int = 0
	err := db.QueryRow("SELECT COUNT(plate_detect_conf) FROM captured WHERE plate_detect_conf > 0 AND DATE(captured_at) BETWEEN ? AND ?", startDate, endDate).Scan(&result)

	return result, err
}

func RecognizedCount(db *sql.DB) (int, error) {
	var result int = 0
	err := db.QueryRow("SELECT COUNT(text_recog_conf) FROM captured WHERE text_recog_conf > 0").Scan(&result)

	return result, err
}

func RecognizedCountByDate(date string, db *sql.DB) (int, error) {
	var result int = 0
	err := db.QueryRow("SELECT COUNT(text_recog_conf) FROM captured WHERE text_recog_conf > 0 AND DATE(captured_at) = ?", date).Scan(&result)

	return result, err
}

func RecognizedCountBetweenDate(startDate string, endDate string, db *sql.DB) (int, error) {
	var result int = 0
	err := db.QueryRow("SELECT COUNT(text_recog_conf) FROM captured WHERE text_recog_conf > 0 AND DATE(captured_at) BETWEEN ? AND ?", startDate, endDate).Scan(&result)

	return result, err
}
