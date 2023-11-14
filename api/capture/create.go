package capture

import (
	"fmt"
	"intracs_anpr_api/internal/response"
	"net/http"
	"time"
)

func (api *CaptureApi) Create(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	w = *response.SetDefaultHeader(w)

	if r.Method == "POST" {
		result := struct {
			Id int `json:"id"`
		}{
			Id: -1,
		}

		lastInsertId, err := api.controller.Create(r)
		if err != nil {
			api.response.ErrorWithLog(w, r, http.StatusBadRequest, err.Error())
			fmt.Printf("Running Time: %v\n", time.Since(startTime))
			return
		}

		result.Id = lastInsertId
		api.response.SuccessWithLog(w, r, "Add new capture successfully", result, true)
		fmt.Printf("Running Time: %v\n", time.Since(startTime))
		return
	}
	api.response.MethodInvalid(w)
	fmt.Printf("Running Time: %v\n", time.Since(startTime))

}
