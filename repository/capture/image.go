package capture

import (
	"fmt"
	"intracs_anpr_api/model"
)

func (repo *CaptureRepo) GetImage(id int) (model.CaptureImage, error) {
	var result model.CaptureImage

	q := repo.db.Table("captures").Select("plate_image", "full_image").Where("id = ?", id).First(&result)

	if q.Error != nil {
		fmt.Println("Failed get capture image", q.Error)
		return result, q.Error
	}

	return result, nil
}

func (repo *CaptureRepo) GetPlateImage(id int) (string, error) {
	var result string

	q := repo.db.Table("captures").Select("plate_image").Where("id = ?", id).Order("captured_at ASC")

	if q.Error != nil {
		fmt.Println("Failed get capture plate image.", q.Error)
		return result, q.Error
	}

	q.Row().Scan(&result)

	return result, nil
}

func (repo *CaptureRepo) GetFullImage(id int) (string, error) {
	var result string

	q := repo.db.Table("captures").Select("full_image").Where("id = ?", id).Order("captured_at ASC")

	if q.Error != nil {
		fmt.Println("Failed get capture full image.", q.Error)
		return result, q.Error
	}

	q.Row().Scan(&result)

	return result, nil
}
