package capture

import (
	"intracs_anpr_api/model"
	"intracs_anpr_api/types"

	"gorm.io/gorm"
)

func getValidationStatusFilter(db *gorm.DB, validation types.ValidationStatusFilter) *gorm.DB {
	if validation.IsValidActive {
		if validation.IsValid() {
			db.Where("IFNULL(is_valid, -1) = 1")
		} else {
			db.Where("IFNULL(is_valid, -1) = 0")
		}
	} else if validation.IsValidatedActive {
		if validation.IsValidated() {
			db.Where("IFNULL(is_valid, -1) != -1")
		} else {
			db.Where("IFNULL(is_valid, -1) = -1")
		}
	}

	return db
}

func (repo *CaptureRepo) ReadById(captureId int) (model.Capture, error) {
	var result model.Capture
	fields := getDefaultFields()

	q := repo.db.Select(fields).First(&result, captureId).Joins("JOIN validation_reasons ON captures.validation_reason_code = validation_reasons.code")
	if q.Error != nil {
		return model.Capture{}, q.Error
	}

	return result, nil
}

func (repo *CaptureRepo) Read(page types.PageFilter, withPlateImage bool, date types.DateFilter, validationStatus types.ValidationStatusFilter) ([]model.Capture, int64, error) {
	var rowCount int64 = 0
	captures := []model.Capture{}
	fields := getDefaultFields()

	// date filter
	startDate := ""
	endDate := ""

	if date.IsSingleDateValid() {
		startDate = toStartDateStr(date.DateString())
		endDate = toEndDateStr(date.DateString())
	} else if date.IsBetweenDateValid() {
		startDate = toStartDateStr(date.StartDateString())
		endDate = toEndDateStr(date.EndDateString())
	}

	if withPlateImage {
		fields = append(fields, "plate_image")
	}

	q := repo.db.Table("captures").Select([]string{"id", "is_valid"})
	if date.IsSingleDateValid() || date.IsBetweenDateValid() {
		q.Where("captured_at BETWEEN ? AND ?", startDate, endDate)
	}
	q = getValidationStatusFilter(q, validationStatus)

	q.Count(&rowCount)
	if repo.db.Error != nil {
		return captures, rowCount, repo.db.Error
	}

	q = repo.db.Table("captures").Select(fields)
	q.Joins("JOIN validation_reasons ON captures.validation_reason_code = validation_reasons.code")
	if date.IsSingleDateValid() || date.IsBetweenDateValid() {
		q.Where("captured_at BETWEEN ? AND ?", startDate, endDate)
	}
	q = getValidationStatusFilter(q, validationStatus)
	q.Limit(int(page.Limit)).
		Offset(int(page.GetOffset())).
		Order("captured_at ASC")

	rows, err := q.Rows()
	if err != nil {
		return captures, rowCount, err
	}

	for rows.Next() {
		capture := model.Capture{}

		if withPlateImage {
			err = rows.Scan(
				&capture.Id,
				&capture.Info.DeviceId,
				&capture.Info.Speed,
				&capture.Info.PlateDetectConf,
				&capture.Info.TextRecogConf,
				&capture.Info.PlateNumber,
				&capture.Info.Width,
				&capture.Info.Height,
				&capture.Info.PlateCoordinates.X1,
				&capture.Info.PlateCoordinates.Y1,
				&capture.Info.PlateCoordinates.X2,
				&capture.Info.PlateCoordinates.Y2,
				&capture.IsValid,
				&capture.ValidationReasonCode,
				&capture.ValidationReason,
				&capture.CapturedAt,
				&capture.Image.Plate,
			)
		} else {
			err = rows.Scan(
				&capture.Id,
				&capture.Info.DeviceId,
				&capture.Info.Speed,
				&capture.Info.PlateDetectConf,
				&capture.Info.TextRecogConf,
				&capture.Info.PlateNumber,
				&capture.Info.Width,
				&capture.Info.Height,
				&capture.Info.PlateCoordinates.X1,
				&capture.Info.PlateCoordinates.Y1,
				&capture.Info.PlateCoordinates.X2,
				&capture.Info.PlateCoordinates.Y2,
				&capture.IsValid,
				&capture.ValidationReasonCode,
				&capture.ValidationReason,
				&capture.CapturedAt,
			)
		}
		if err != nil {
			return captures, rowCount, err
		}

		captures = append(captures, capture)
	}
	return captures, rowCount, nil
}
