package validationreason

import (
	"errors"
	"intracs_anpr_api/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Validation request and return code, model and error
func updateValidation(r *http.Request) (model.ValidationReason, error) {
	vars := mux.Vars(r)
	codeVar, ok := vars["code"]
	if !ok {
		return model.ValidationReason{}, errors.New("code value invalid")
	}

	result := model.ValidationReason{}
	// codeParam := r.URL.Query().Get("code")
	validStatusForm := r.FormValue("validStatus")
	reasonForm := r.FormValue("reason")

	code, err := strconv.ParseUint(codeVar, 10, 8)
	if err != nil {
		return result, errors.New("code parameter invalid and can't parse to uint8")
	}

	validStatus, err := strconv.ParseInt(validStatusForm, 10, 8)
	if err != nil {
		return result, errors.New("valid status invalid and can't parse to uint8")
	} else if validStatus < 0 && validStatus > 1 {
		return result, errors.New("valid status only accept 0 or 1 value")
	}

	result.Code = uint8(code)
	result.ValidStatus = int8(validStatus)
	result.Reason = reasonForm

	return result, nil
}

func (c *Controller) Update(r *http.Request) (model.ValidationReason, error) {

	reason, err := updateValidation(r)
	if err != nil {
		return model.ValidationReason{}, err
	}

	result, err := c.service.Update(reason)
	if err != nil {
		return model.ValidationReason{}, err
	}

	return result, nil
}
