package types

type ApiResponses interface {
	GetApiResponses(status string, message string, data map[string]string) ApiResponse
}

type ApiResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

func (r ApiResponse) GetApiResponses(
	status string,
	message string,
	data map[string]string,
) ApiResponse {
	r.Status = status
	r.Message = message
	r.Data = data

	return r
}
