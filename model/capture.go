package model

type Capture struct {
	Id         int          `gorm:"primaryKey" json:"id"`
	Info       CaptureInfo  `json:"information"`
	Image      CaptureImage `json:"image"`
	CapturedAt string       `json:"captured_at"`
}

type CaptureInfo struct {
	DeviceId         string  `json:"device_id"`
	Speed            float64 `json:"speed"`
	Width            int     `json:"width"`
	Height           int     `json:"height"`
	PlateCoordinates []int   `json:"plate_coordinates"`
	PlateDetectConf  float64 `json:"plate_detect_conf"`
	TextRecogConf    float64 `json:"text_recog_conf"`
	PlateNumber      string  `json:"plate_number"`
}

type CaptureImage struct {
	Plate string `gorm:"column:plate_image" json:"plate"`
	Full  string `gorm:"column:full_image" json:"full"`
}

type Tabler interface {
	TableName() string
}

func (Capture) TableName() string {
	return "captures"
}
