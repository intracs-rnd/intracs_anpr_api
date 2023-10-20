package types

import "encoding/json"

type ApiResponses interface {
	// GetApiResponses(status string, message string, data []byte) *ApiResponses

	ToString() string
}

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CountData struct {
	Count int `json:"count"`
}

func (r ApiResponse) Update(status string, message string, data interface{}) *ApiResponse {
	return &ApiResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func (r ApiResponse) ToJsonString() string {
	jsonData, _ := json.MarshalIndent(r, "", "  ")

	return string(jsonData)
}

// func (r ApiResponse) GetApiResponses(
// 	status string,
// 	message string,
// 	data []byte,
// ) *ApiResponses {
// 	r.Status = status
// 	r.Message = message
// 	r.Data = data

// 	return r
// }
