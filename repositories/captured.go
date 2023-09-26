package repositories

import (
	"database/sql"
	"intracs_anpr_api/handlers"
	. "intracs_anpr_api/types"
	"net/http"
	"strconv"
)

func InsertCapture(w http.ResponseWriter, data ImageCaptured) (*sql.Rows, error) {
	// Connect to database
	db, _ := handlers.DBConnect()

	// Execute insert query
	insert, err := db.Query(
		"INSERT INTO captured (device_id, speed, width, height, x1, y1, x2, y2, plate_image, full_image) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		data.DeviceId,
		strconv.FormatFloat(data.Speed, 'E', 3, 64),
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
	defer insert.Close()

	return insert, nil
}
