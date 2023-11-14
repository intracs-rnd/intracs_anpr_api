package capture

import (
	"intracs_anpr_api/controller/capture"
	"intracs_anpr_api/internal/response"
)

type CaptureApi struct {
	controller *capture.Controller
	response   *response.ApiResponse
}

func InitCaptureApi(controller *capture.Controller) *CaptureApi {
	apiResponse := response.NewAPIResponse()
	return &CaptureApi{
		controller: controller,
		response:   apiResponse,
	}
}

var message = struct {
	startDateInvalid string
	endDateInvalid   string
	idNull           string
	idInvalid        string
}{
	startDateInvalid: "start date invalid or can't parse to date (YYYY-MM-DD)",
	endDateInvalid:   "end date invalid or can't parse to date (YYYY-MM-DD)",
	idNull:           "id must be not null",
	idInvalid:        "id can't parsing to int",
}
