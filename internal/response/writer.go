package response

import "net/http"

func SetDefaultHeader(w http.ResponseWriter) *http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	return &w
}
