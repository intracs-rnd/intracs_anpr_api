package validationreason

import (
	"errors"
	"intracs_anpr_api/model"
)

func (repo *ValidationReasonRepo) Create(data model.ValidationReason) (model.ValidationReason, error) {
	result := model.ValidationReason{}

	q := repo.db.Create(&data)
	if q.Error != nil {
		return result, q.Error
	}

	if q.RowsAffected == 0 {
		return result, errors.New("failed create new validation reason cause database have a problem")
	}

	q.Last(&result)

	return result, nil
}
