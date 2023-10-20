package api

// import (
// 	"database/sql"
// 	"fmt"
// 	"intracs_anpr_api/handlers"
// 	"intracs_anpr_api/repositories"
// 	. "intracs_anpr_api/types"
// 	"net/http"
// 	"strconv"
// )

// type CapturedApi interface {
// 	init()
// 	insertValidation(w http.ResponseWriter, r *http.Request) (ImageCaptured, error)
// 	Create(w http.ResponseWriter, r *http.Request)
// 	Read(w http.ResponseWriter, r *http.Request)
// 	GetImage(w http.ResponseWriter, r *http.Request)
// 	Count(w http.ResponseWriter, r *http.Request)
// 	DetectionCount(w http.ResponseWriter, r *http.Request)
// 	RecognitionCount(w http.ResponseWriter, r *http.Request)
// }

// var db *sql.DB
// var db_err error
// var repo repositories.CapturedRepository

// func init() {
// 	// connect to db
// 	db, db_err = handlers.DBConnect()

// 	if db_err != nil {
// 		fmt.Println("Database can't reach.", db_err)
// 	}

// 	fmt.Println("Database initialized for /captured")
// }

// const SUCCESS_MESSAGE string = "Captured image uploaded successfully"

// func insertValidation(w http.ResponseWriter, r *http.Request) (ImageCaptured, error) {
// 	var capturedImage ImageCapture = ImageCaptured{}

// 	data, err := capturedImage.FromRequest("-1", w, r)
// 	if err != nil {
// 		http.Error(w, "Failed to parse capture image to dict!", http.StatusBadRequest)
// 		fmt.Println(err)
// 		return data, err
// 	}

// 	return data, nil
// }

// func Create(w http.ResponseWriter, r *http.Request) {
// 	data, err := insertValidation(w, r)
// 	if err != nil {
// 		http.Error(w, "Failed to create image capture!", http.StatusBadRequest)
// 		fmt.Println(err)
// 		return
// 	}

// 	result, err := repo.Create(data, db)
// 	if err != nil {
// 		http.Error(w, "Failed to insert captured image!", http.StatusInternalServerError)
// 		fmt.Println(err)
// 		return
// 	}

// 	lastId, err := result.LastInsertId()
// 	if err != nil {
// 		http.Error(w, "Failed get last insert id!", http.StatusInternalServerError)
// 		return
// 	}

// 	response := "{\"message\", \"" + SUCCESS_MESSAGE + "\", \"data\": [{\"id\": " + strconv.Itoa(int(lastId)) + " }]}"

// 	fmt.Fprintln(w, response)
// 	fmt.Println(response)
// }

// func GetImage(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		// var response ApiResponse = ApiResponse{}
// 		// var id int = -1

// 		idForm := r.FormValue("id")

// 		if idForm != "" {
// 			idValue, err := strconv.Atoi(idForm)

// 			if err != nil {
// 				fmt.Fprintln(w, idValue)
// 			}

// 			fmt.Fprintln(w, "Success")
// 			fmt.Fprintln(w, idValue)
// 		}
// 		fmt.Fprintln(w, "Failed")

// 	}
// }

// func Read(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		var response ApiResponse = ApiResponse{}
// 		var page int = 1
// 		var limit int = 25
//
// 		// Get page number and limit from request
// 		pageNumberForm := r.FormValue("page")
// 		pageLimitForm := r.FormValue("limit")
//
// 		// Convert page number and limit to int
// 		pageNumber, _ := strconv.Atoi(pageNumberForm)
// 		pageLimit, _ := strconv.Atoi(pageLimitForm)
//
// 		if pageNumber > 0 {
// 			page = pageNumber
// 		}
//
// 		if pageLimit > 0 {
// 			limit = pageLimit
// 		}

// 		// Read data from database
// 		data, err := repo.Read(limit, page, w, db)
// 		if err != nil {
// 			http.Error(w, "Failed fetching data from database", http.StatusBadRequest)
// 			fmt.Println(err)
// 			return
// 		}

// 		// Create response
// 		response.Status = "success"
// 		response.Message = "fetch captured data successfully"
// 		response.Data = data

// 		fmt.Fprintln(w, response.ToJsonString())
// 		fmt.Println("Reading captured data")
// 	}
// }

// func Count(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		var response ApiResponse = ApiResponse{}
// 		var data int = 0
// 		var err error

// 		startDateForm := r.FormValue("start_date")
// 		endDateForm := r.FormValue("end_date")

// 		if startDateForm != "" && endDateForm != "" {
// 			data, err = repo.CountBetweenDate(startDateForm, endDateForm, db)
// 		} else if startDateForm != "" && endDateForm == "" {
// 			data, err = repo.CountByDate(startDateForm, db)
// 		} else if startDateForm == "" && endDateForm == "" {
// 			data, err = repo.Count(db)
// 		}

// 		if err != nil {
// 			response.Status = "failed"
// 			response.Message = "failed fetch captured count"
// 			response.Data = make([]int, 0)

// 			http.Error(w, response.ToJsonString(), http.StatusBadRequest)
// 			fmt.Println(err)
// 			return
// 		}

// 		response.Status = "success"
// 		response.Message = "fetching captures count successfully"
// 		response.Data = CountData{Count: data}

// 		fmt.Fprintln(w, response.ToJsonString())
// 	}
// }

// func CapturesDetectionCount(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		var response ApiResponse = ApiResponse{}
// 		var data int
// 		var err error

// 		startDateForm := r.FormValue("start_date")
// 		endDateForm := r.FormValue("end_date")

// 		if startDateForm != "" && endDateForm != "" {
// 			data, err = repo.DetectedCountBetweenDate(startDateForm, endDateForm, db)
// 		} else if startDateForm != "" && endDateForm == "" {
// 			data, err = repo.DetectedCountByDate(startDateForm, db)
// 		} else if startDateForm == "" && endDateForm == "" {
// 			data, err = repo.DetectedCount(db)
// 		}

// 		if err != nil {
// 			response.Status = "failed"
// 			response.Message = "failed fetch captured detection count"
// 			response.Data = make([]int, 0)

// 			http.Error(w, "Failed fetch captured detection count", http.StatusBadRequest)
// 			fmt.Println(err)
// 			return
// 		}

// 		response.Status = "success"
// 		response.Message = "fetching captures detection count successfully"
// 		response.Data = CountData{Count: data}

// 		fmt.Fprintln(w, response.ToJsonString())
// 	}
// }

// func CapturesRecognitionCount(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		var response ApiResponse = ApiResponse{}
// 		var data int
// 		var err error

// 		startDateForm := r.FormValue("start_date")
// 		endDateForm := r.FormValue("end_date")

// 		if startDateForm != "" && endDateForm != "" {
// 			data, err = repo.RecognizedCountBetweenDate(startDateForm, endDateForm, db)
// 		} else if startDateForm != "" && endDateForm == "" {
// 			data, err = repo.RecognizedCountByDate(startDateForm, db)
// 		} else if startDateForm == "" && endDateForm == "" {
// 			data, err = repo.RecognizedCount(db)
// 		}

// 		if err != nil {
// 			response.Status = "failed"
// 			response.Message = "failed fetch captured recognition count"
// 			response.Data = make([]int, 0)

// 			http.Error(w, "Failed fetch captured recognition count", http.StatusBadRequest)
// 			fmt.Println(err)
// 			return
// 		}

// 		response.Status = "success"
// 		response.Message = "fetching captures recognition count successfully"
// 		response.Data = CountData{Count: data}

// 		fmt.Fprintln(w, response.ToJsonString())
// 	}
// }
