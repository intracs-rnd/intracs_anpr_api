package capture

import (
	"intracs_anpr_api/controller/validationreason"
	"intracs_anpr_api/model"
	"intracs_anpr_api/repository/capture"
	"intracs_anpr_api/types"
	"time"
)

type repositories interface {
	ReadById(captureId int) (model.Capture, error)
	Read(page types.PageFilter, withPlateImage bool, date types.DateFilter, validationStatus types.ValidationStatusFilter) ([]model.Capture, int64, error)
	// Read(limitSize int, pageNumber int, withPlateImage bool) ([]model.Capture, int64, error)
	Create(data model.Capture) (bool, int, error)

	// Summary
	SummaryDateList(pageFilter types.PageFilter) ([]time.Time, int64, error)
	SummaryDateListBetweenDate(pageFilter types.PageFilter, dateFilter types.DateFilter) ([]time.Time, int64, error)

	// Read image
	GetImage(id int) (model.CaptureImage, error)
	GetPlateImage(id int) (string, error)
	GetFullImage(id int) (string, error)

	// Captures Count
	Count() (int64, error)
	CountByDate(date string) (int64, error)
	CountBeforeDate(date string) (int64, error)
	CountBetweenDate(startDate string, endDate string) (int64, error)

	// Detected Count
	DetectedCount() (int64, error)
	DetectedCountByDate(date string) (int64, error)
	DetectedCountBeforeDate(date string) (int64, error)
	DetectedCountBetweenDate(startDate string, endDate string) (int64, error)

	// Recognized count
	RecognizedCount() (int64, error)
	RecognizedCountByDate(date string) (int64, error)
	RecognizedCountBeforeDate(date string) (int64, error)
	RecognizedCountBetweenDate(startDate string, endDate string) (int64, error)

	// Validation
	ValidCount() (int64, error)
	ValidCountByDate(date string) (int64, error)
	ValidCountBetweenDate(startDate string, endDate string) (int64, error)

	InvalidCount() (int64, error)
	InvalidCountByDate(date string) (int64, error)
	InvalidCountBetweenDate(startDate string, endDate string) (int64, error)

	// Validated count
	ValidatedCount() (int64, error)
	ValidatedCountByDate(date string) (int64, error)
	ValidatedCountBetweenDate(startDate string, endDate string) (int64, error)

	UnValidatedCount() (int64, error)
	UnValidatedCountByDate(date string) (int64, error)
	UnValidatedCountBetweenDate(startDate string, endDate string) (int64, error)

	// Validation
	UpdateValidation(captureId int, reasonCode uint8, validStatus int8) (model.Capture, error)
}

type Controller struct {
	service          repositories
	reasonController *validationreason.Controller
}

func InitController(captureRepo *capture.CaptureRepo, reasonController *validationreason.Controller) *Controller {
	return &Controller{
		service:          captureRepo,
		reasonController: reasonController,
	}
}
