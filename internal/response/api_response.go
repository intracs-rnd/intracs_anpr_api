package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiResponses interface {
	New() *ApiResponse
	Update(status string, message string, data interface{}) *ApiResponse
	ToString() string
}

type ApiResponse struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination,omitempty"`
}

type PaginationData struct {
	PageCount   int `json:"page_count"`
	CurrentPage int `json:"current_page"`
	LimitSize   int `json:"limit_size"`
}

type CountData struct {
	Count int `json:"count"`
}

func NewAPIResponse() *ApiResponse {
	return &ApiResponse{}
}

func (r *ApiResponse) UpdateWithPaging(status string, message string, data interface{}, paging PaginationData) {
	r.Status = status
	r.Message = message

	if data != nil {
		r.Data = data
		r.Pagination = paging
	} else {
		r.Data = make([]string, 0)
		r.Pagination = nil
	}
	// return &r
}

func (r *ApiResponse) Update(status string, message string, data interface{}) *ApiResponse {
	r.Status = status
	r.Message = message
	r.Pagination = nil

	if data != nil {
		r.Data = data
	} else {
		r.Data = make([]string, 0)
	}
	return r
}

func (r *ApiResponse) ToJsonString() string {
	jsonData, _ := json.MarshalIndent(&r, "", "  ")

	return string(jsonData)
}

func (r *ApiResponse) Success(w http.ResponseWriter, message string, data interface{}) {
	r.Update(http.StatusText(http.StatusOK), message, data)
	fmt.Fprintln(w, r.ToJsonString())
}

func (r *ApiResponse) SuccessWithPaging(w http.ResponseWriter, message string, data interface{}, pagination PaginationData) {
	r.UpdateWithPaging(http.StatusText(http.StatusOK), message, data, pagination)
	fmt.Fprintln(w, r.ToJsonString())
}

func (r *ApiResponse) Error(w http.ResponseWriter, status int, message string) {
	r.Update(http.StatusText(status), message, nil)
	http.Error(w, r.ToJsonString(), http.StatusBadRequest)
}

func (r *ApiResponse) MethodInvalid(w http.ResponseWriter) {
	r.Update(http.StatusText(http.StatusMethodNotAllowed), "Method invalid or not allowed", nil)
	http.Error(w, r.ToJsonString(), http.StatusMethodNotAllowed)
}
