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

type CountData struct {
	Count int `json:"count"`
}

// type IdData struct {
//     Id int
// }

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

func (r *ApiResponse) SuccessWithLog(w http.ResponseWriter, req *http.Request, message string, data interface{}, showLogData bool) {
	r.Update(http.StatusText(http.StatusOK), message, data)
	fmt.Fprintln(w, r.ToJsonString())

	r.Log(req, http.StatusOK, message, data, showLogData)
}

func (r *ApiResponse) SuccessNoData(w http.ResponseWriter, req *http.Request, message string) {
	r.Update(http.StatusText(http.StatusNoContent), message, nil)
	fmt.Fprintln(w, r.ToJsonString())

	r.Log(req, http.StatusOK, message, nil, false)
}

func (r *ApiResponse) SuccessWithPaging(w http.ResponseWriter, req *http.Request, message string, data interface{}, pagination PaginationData) {
	r.UpdateWithPaging(http.StatusText(http.StatusOK), message, data, pagination)
	fmt.Fprintln(w, r.ToJsonString())

	r.Log(req, http.StatusOK, message, nil, false)
}

func (r *ApiResponse) Error(w http.ResponseWriter, status int, message string) {
	r.Update(http.StatusText(status), message, nil)
	http.Error(w, r.ToJsonString(), http.StatusBadRequest)
}

func (r *ApiResponse) ErrorWithLog(w http.ResponseWriter, req *http.Request, status int, message string) {
	r.Update(http.StatusText(status), message, nil)
	http.Error(w, r.ToJsonString(), http.StatusBadRequest)

	r.Log(req, status, message, nil, false)
}

func (r *ApiResponse) MethodInvalid(w http.ResponseWriter) *ApiResponse {
	r.Update(http.StatusText(http.StatusMethodNotAllowed), "Method invalid or not allowed", nil)
	http.Error(w, r.ToJsonString(), http.StatusMethodNotAllowed)

	return r
}

func (r *ApiResponse) Log(req *http.Request, status int, message string, response interface{}, showData bool) {
	if showData {
		fmt.Printf("[ %-8s ] %d | %-50s | %-20s |  {message: %s, response: %+v} \n",
			req.Method,
			status,
			req.URL.Path,
			http.StatusText(status),
			message,
			response,
		)
	} else {
		fmt.Printf("[ %-8s ] %d | %-50s | %-20s |  %s \n",
			req.Method,
			status,
			req.URL.Path,
			http.StatusText(status),
			message,
		)
	}
}
