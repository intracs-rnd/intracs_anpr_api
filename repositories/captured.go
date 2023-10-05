package repositories

import (
	"database/sql"
	. "intracs_anpr_api/types"
	"net/http"
)

func InsertCapture(w http.ResponseWriter, data ImageCaptured, db *sql.DB) (sql.Result, error) {
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

	if err != nil {
		return nil, err
	}
	// defer insert.Close()

	return insert, nil
}
