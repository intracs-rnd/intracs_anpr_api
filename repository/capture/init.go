package capture

import (
	"fmt"

	"gorm.io/gorm"
)

type CaptureRepo struct {
	db *gorm.DB
}

func NewCaptureRepo(db *gorm.DB) *CaptureRepo {
	return &CaptureRepo{
		db: db,
	}
}

func getOffset(number uint, limit uint) uint {
	return (number - 1) * limit
}

func getDefaultFields() []string {
	return []string{
		"id",
		"device_id",
		"speed",
		"plate_detect_conf",
		"text_recog_conf",
		"plate_number",
		"width",
		"height",
		"x1",
		"y1",
		"x2",
		"y2",
		"IFNULL(is_valid, -1) AS is_valid",
		"IFNULL(validation_reason_code, 0) AS validation_reason_code",
		"IFNULL(validation_reasons.reason, 'Belum divalidasi') AS validation_reason_text",
		"captured_at",
	}
}

func toStartDateStr(date string) string {
	return fmt.Sprintf("%s 00:00:00", date)
}

func toEndDateStr(date string) string {
	return fmt.Sprintf("%s 23:59:59", date)
}
