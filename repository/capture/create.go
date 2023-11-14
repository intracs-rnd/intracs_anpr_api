package capture

import (
	"intracs_anpr_api/model"
)

func (repo *CaptureRepo) Create(data model.Capture) (bool, int, error) {
	isSuccess := false

	result := repo.db.Select("device_id", "speed", "plate_detect_conf", "text_recog_conf", "plate_number", "width", "height", "x1", "y1", "x2", "y2", "plate_image", "full_image").Create(&data)
	if result.Error != nil {
		return isSuccess, -1, result.Error
	}

	return isSuccess, data.Id, nil
}
