package route

import (
	"intracs_anpr_api/api"

	"github.com/gorilla/mux"
)

func GetAll(api *api.Api) *mux.Router {
	route := mux.NewRouter().PathPrefix("/api").Subrouter()

	// Read capture record
	// route.HandleFunc("/1.0/capture", api.Capture.Read)
	route.HandleFunc("/1.0/capture/{id:[0-9]+}", api.Capture.ReadById).Methods("GET")
	route.HandleFunc("/1.0/capture", api.Capture.Read).Methods("GET")
	route.HandleFunc("/1.0/capture", api.Capture.Create).Methods("POST")

	// Read capture summary
	route.HandleFunc("/1.0/capture/summary", api.Capture.Summary).Methods("GET")
	route.HandleFunc("/1.0/capture/summary/eachday", api.Capture.SummaryEachDay).Methods("GET")

	// Create new capture record
	// route.HandleFunc("/1.0/capture/add", api.Capture.Create)

	// Read capture image
	route.HandleFunc("/1.0/capture/image/{id:[0-9]+}", api.Capture.GetImage)
	route.HandleFunc("/1.0/capture/image/{id:[0-9]+}/plate", api.Capture.GetPlateImage)
	route.HandleFunc("/1.0/capture/image/{id:[0-9]+}/full", api.Capture.GetFullImage)

	// Read capture count
	route.HandleFunc("/1.0/capture/count", api.Capture.Count)
	route.HandleFunc("/1.0/capture/count/today", api.Capture.Count)
	route.HandleFunc("/1.0/capture/detection/count", api.Capture.DetectedCount)
	route.HandleFunc("/1.0/capture/recognition/count", api.Capture.RecognizedCount)
	route.HandleFunc("/1.0/capture/validated/count", api.Capture.ValidatedCount)
	route.HandleFunc("/1.0/capture/unvalidated/count", api.Capture.UnValidatedCount)

	// Read capture validation count
	route.HandleFunc("/1.0/capture/validation/valid/count", api.Capture.ValidCount)
	route.HandleFunc("/1.0/capture/validation/invalid/count", api.Capture.InvalidCount)

	// Update capture validation
	route.HandleFunc("/1.0/capture/validation/{id:[0-9]+}", api.Capture.UpdateValidation)

	route.HandleFunc("/1.0/capture/validation/reason", api.ValidationReason.ReadAll).Methods("GET")
	route.HandleFunc("/1.0/capture/validation/reason", api.ValidationReason.Create).Methods("POST")
	route.HandleFunc("/1.0/capture/validation/reason/{code:[0-9]+}", api.ValidationReason.Update).Methods("PUT")

	return route
}
