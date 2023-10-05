package api

import (
	"database/sql"
	"fmt"
	"intracs_anpr_api/handlers"
	"intracs_anpr_api/repositories"
	"intracs_anpr_api/types"
	"net/http"
	"strconv"
)

var db *sql.DB
var db_err error

func init() {
	// connect to db
	db, db_err = handlers.DBConnect()

	if db_err != nil {
		fmt.Println("Database can't reach.", db_err)
	}

	fmt.Println("Database initialized for /captured")
}

const SUCCESS_MESSAGE string = "Captured image uploaded successfully"

func insertCaptureValidation(w http.ResponseWriter, r *http.Request) (types.ImageCaptured, error) {
	var capturedImage types.ImageCapture = types.ImageCaptured{}

	data, err := capturedImage.FromRequest("-1", w, r)
	if err != nil {
		http.Error(w, "Failed to parse capture image to dict!", http.StatusBadRequest)
		fmt.Println(err)
		return data, err
	}

	return data, nil
}

func InsertCapture(w http.ResponseWriter, r *http.Request) {
	data, err := insertCaptureValidation(w, r)
	if err != nil {
		http.Error(w, "Failed to create image capture!", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	result, err := repositories.InsertCapture(w, data, db)
	if err != nil {
		http.Error(w, "Failed to insert captured image!", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed get last insert id!", http.StatusInternalServerError)
		return
	}

	response := "{\"message\", \"" + SUCCESS_MESSAGE + "\", \"data\": [{\"id\": " + strconv.Itoa(int(lastId)) + " }]}"

	fmt.Fprintln(w, response)
	fmt.Println(response)
}
