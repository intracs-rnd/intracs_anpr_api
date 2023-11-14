package validationreason

import (
	"errors"
	"intracs_anpr_api/internal/validator"
	"intracs_anpr_api/model"
)

func (c *Controller) ReadAll() ([]model.ValidationReason, error) {
	results, err := c.service.ReadAll()
	if err != nil {
		return []model.ValidationReason{}, err
	}

	return results, nil
}

func (c *Controller) ReadByCode(codeForm string) (model.ValidationReason, error) {
	// Validation code
	isCodeNull, code, err := validator.ReasonCode(codeForm, false)
	if err != nil {
		return model.ValidationReason{}, err
	} else if isCodeNull {
		return model.ValidationReason{}, errors.New("reason code required")
	}

	result, err := c.service.Read(code)
	if err != nil {
		return model.ValidationReason{}, err
	}

	return result, nil
}
