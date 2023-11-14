package validationreason

import (
	"intracs_anpr_api/model"
	"time"

	"gorm.io/gorm"
)

func (repo *ValidationReasonRepo) Update(data model.ValidationReason) (model.ValidationReason, error) {
	result := model.ValidationReason{}
	result.UpdatedAt = time.Now()

	err := repo.db.Transaction(func(tx *gorm.DB) error {
		tx.Model(&data).Select("valid_status", "reason", "updated_at").Updates(data)
		if tx.Error != nil {
			return tx.Error
		}

		tx.First(&result, data.Code)
		if tx.Error != nil {
			return tx.Error
		}

		return nil
	})

	// repo.db.Table("").Model(&data).Select("valid_status", "reason", "updated_at").Updates(data)

	// repo.db.Last(&result)

	if err != nil {
		return model.ValidationReason{}, err
	}

	return result, nil
}
