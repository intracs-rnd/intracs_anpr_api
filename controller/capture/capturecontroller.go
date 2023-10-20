package capture

import (
	"intracs_anpr_api/model"
	"intracs_anpr_api/repository"
)

type repositories interface {
	Read(limitSize int, pageNumber int, withPlateImage bool) ([]model.Capture, int64, error)
	Create(data model.Capture) (model.Capture, error)

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

	// Validated count
	ValidatedCount() (int64, error)
	UnValidatedCount() (int64, error)
}

type Controller struct {
	service repositories
}

func InitController(captureRepo *repository.CaptureRepo) *Controller {
	return &Controller{
		service: captureRepo,
	}
}
