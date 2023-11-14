package types

import (
	"errors"
	"strconv"
)

type ValidationStatusFilter struct {
	isValid           bool
	IsValidActive     bool
	isValidated       bool
	IsValidatedActive bool
}

func (v *ValidationStatusFilter) FromString(isValid string, isValidated string) (ValidationStatusFilter, error) {
	if isValid != "" && isValidated != "" {
		return *v, errors.New("inconsistency is_valid and is_validated parameter, just can use one of them")
	}

	if isValid != "" {
		valid, err := strconv.ParseBool(isValid)
		if err != nil {
			return *v, errors.New("parameter is_valid invalid or can't parse to bool")
		} else {
			v.isValid = valid
			v.IsValidActive = true

			v.isValidated = false
			v.IsValidatedActive = false
		}
	} else if isValidated != "" {
		validated, err := strconv.ParseBool(isValidated)
		if err != nil {
			return *v, errors.New("parameter is_validated invalid or can't parse to bool")
		} else {
			v.isValid = false
			v.IsValidActive = false

			v.isValidated = validated
			v.IsValidatedActive = true
		}
	}

	return *v, nil
}

func (v *ValidationStatusFilter) IsValid() bool {
	if v.isValid && v.IsValidActive {
		return true
	}
	return false
}

func (v *ValidationStatusFilter) IsValidated() bool {
	if v.isValidated && v.IsValidatedActive {
		return true
	}
	return false
}
