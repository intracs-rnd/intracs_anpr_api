package main

import (
	"fmt"
	"net/http"

	// Third party
	_ "github.com/go-sql-driver/mysql"

	// Project
	"intracs_anpr_api/api"
	"intracs_anpr_api/controller"
	"intracs_anpr_api/database"
	"intracs_anpr_api/repository"
)

func main() {
	db := database.GetDatabase()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	api := api.InitApi(controllers)

	// Users API
	http.HandleFunc("/api/v1/login", api.User.Login)

	// Captures API
	captureApi := api.Capture
	http.HandleFunc("/api/v1/capture", captureApi.Read)
	http.HandleFunc("/api/v1/capture/add", captureApi.Create)
	http.HandleFunc("/api/v1/capture/image", captureApi.GetImage)
	http.HandleFunc("/api/v1/capture/image/plate", captureApi.GetPlateImage)
	http.HandleFunc("/api/v1/capture/image/full", captureApi.GetFullImage)
	http.HandleFunc("/api/v1/capture/count", captureApi.Count)
	http.HandleFunc("/api/v1/capture/count/today", captureApi.Count)
	http.HandleFunc("/api/v1/capture/detection/count", captureApi.DetectedCount)
	http.HandleFunc("/api/v1/capture/recognition/count", captureApi.RecognizedCount)
	http.HandleFunc("/api/v1/capture/validated/count", captureApi.ValidatedCount)
	http.HandleFunc("/api/v1/capture/unvalidated/count", captureApi.UnValidatedCount)

	// Captures Validation API
	capValidationApi := api.CaptureValidation
	http.HandleFunc("/api/v1/capture/validation/valid/count", capValidationApi.ValidCount)
	http.HandleFunc("/api/v1/capture/validation/invalid/count", capValidationApi.InvalidCount)

	fmt.Println("Starting API server at http://localhost:8881/")
	http.ListenAndServe(":8881", nil)
}
