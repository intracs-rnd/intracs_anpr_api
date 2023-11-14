package response

import "math"

type PaginationData struct {
	RecordCount int `json:"recordCount"`
	PageCount   int `json:"pageCount"`
	CurrentPage int `json:"currentPage"`
	LimitSize   int `json:"limitSize"`
}

func NewPaging(recordCount int, currentPage int, limit int) PaginationData {
	pageCount := int(math.Ceil((float64(recordCount) / float64(limit))))

	return PaginationData{
		RecordCount: recordCount,
		CurrentPage: currentPage,
		PageCount:   pageCount,
		LimitSize:   limit,
	}
}
