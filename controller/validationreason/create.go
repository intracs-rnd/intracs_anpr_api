package validationreason

import (
	"errors"
	"intracs_anpr_api/model"
	"net/http"
	"strconv"
)

func requestValidation(r *http.Request) (model.ValidationReason, error) {
	validStatusForm := r.FormValue("valid_status")
	reasonForm := r.FormValue("reason")
	result := model.ValidationReason{}

	validStatus, err := strconv.ParseInt(validStatusForm, 10, 8)
	if err != nil {
		return result, errors.New("valid status invalid and can't parse to int8")
	}

	if validStatus != 0 && validStatus != 1 {
		return result, errors.New("valid status only accept 0 or 1 value")
	}

	if len(reasonForm) <= 0 {
		return result, errors.New("validation reason must have value and can't null")
	}

	result.ValidStatus = int8(validStatus)
	result.Reason = reasonForm

	return result, nil
}

func (c *Controller) Create(r *http.Request) (model.ValidationReason, error) {
	reason, err := requestValidation(r)
	if err != nil {
		return model.ValidationReason{}, err
	}

	result, err := c.service.Create(reason)
	if err != nil {
		return model.ValidationReason{}, err
	}

	return result, nil
}
