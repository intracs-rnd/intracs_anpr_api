package capture

import (
	"intracs_anpr_api/model"

	"gorm.io/gorm"
)

func (repo *CaptureRepo) UpdateValidation(captureId int, reasonId uint8, validStatus int8) (model.Capture, error) {
	var result model.Capture

	// Get validation reason valid status
	// Update capture valid status from validation reason table
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		tx.Table("captures").Where("id = ?", captureId).Updates(map[string]interface{}{
			"is_valid":               validStatus,
			"validation_reason_code": reasonId,
		})

		if tx.Error != nil {
			return tx.Error
		}

		tx.First(&result, captureId)
		if tx.Error != nil {
			return tx.Error
		}

		return nil
	})

	if err != nil {
		return model.Capture{}, nil
	}

	return result, nil
}
