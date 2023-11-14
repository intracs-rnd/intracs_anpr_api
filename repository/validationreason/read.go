package validationreason

import "intracs_anpr_api/model"

func (repo *ValidationReasonRepo) Read(code uint8) (model.ValidationReason, error) {
	var result model.ValidationReason

	repo.db.First(&result, code)

	if repo.db.Error != nil {
		return model.ValidationReason{}, repo.db.Error
	}

	return result, nil
}

func (repo *ValidationReasonRepo) ReadAll() ([]model.ValidationReason, error) {
	var results []model.ValidationReason

	repo.db.Find(&results)

	if repo.db.Error != nil {
		return []model.ValidationReason{}, repo.db.Error
	}

	return results, nil
}
