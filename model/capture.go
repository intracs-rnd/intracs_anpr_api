package model

import "time"

type Capture struct {
	Id                   int          `gorm:"column:id;primaryKey" json:"id"`
	Info                 CaptureInfo  `gorm:"embedded" json:"information"`
	Image                CaptureImage `gorm:"embedded" json:"image"`
	IsValid              int8         `gorm:"column:is_valid" json:"isValid"`
	ValidationReasonCode uint8        `gorm:"column:validation_reason_code" json:"validationReasonCode"`
	ValidationReason     string       `gorm:"column:validation_reason_text" json:"validationReason"`
	CapturedAt           string       `gorm:"column:captured_at" json:"capturedAt"`
}

type CaptureInfo struct {
	DeviceId         string     `gorm:"column:device_id" json:"deviceId"`
	Speed            float64    `gorm:"column:speed" json:"speed"`
	Width            int        `gorm:"column:width" json:"width"`
	Height           int        `gorm:"column:height" json:"height"`
	PlateCoordinates Coordinate `gorm:"embedded" json:"plateCoordinates"`
	PlateDetectConf  float64    `gorm:"column:plate_detect_conf" json:"plateDetectConf"`
	TextRecogConf    float64    `gorm:"column:text_recog_conf" json:"textRecogConf"`
	PlateNumber      string     `gorm:"column:plate_number" json:"plateNumber"`
}

type CaptureImage struct {
	Plate string `gorm:"column:plate_image" json:"plate"`
	Full  string `gorm:"column:full_image" json:"full"`
}

type CaptureCountSummary struct {
	Total       int64 `json:"total"`
	Valid       int64 `json:"valid"`
	Invalid     int64 `json:"invalid"`
	Validated   int64 `json:"validated"`
	Unvalidated int64 `json:"unvalidated"`
}

type DateSummaryTable struct {
	Date        time.Time `gorm:"column:captured_at" json:"date"`
	Count       int64     `gorm:"column:captured_count" json:"count"`
	Valid       int64     `gorm:"column:valid_count" json:"valid"`
	Invalid     int64     `gorm:"column:invalid_count" json:"invalid"`
	Validated   int64     `gorm:"column:validated_count" json:"validated"`
	Unvalidated int64     `gorm:"column:unvalidated_count" json:"unvalidated"`
}

type Coordinate struct {
	X1 int16 `gorm:"column:x1" json:"x1"`
	Y1 int16 `gorm:"column:y1" json:"y1"`
	X2 int16 `gorm:"column:x2" json:"x2"`
	Y2 int16 `gorm:"column:y2" json:"y2"`
}

type Tabler interface {
	TableName() string
}

func (Capture) TableName() string {
	return "captures"
}
