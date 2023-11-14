package validator

import (
	"errors"
	"strconv"
)

// Validate capture ID from string and required
//
// return isNil: bool, id: int, err: error
func CaptureId(id string, required bool) (bool, int, error) {
	if required && len(id) <= 0 {
		return true, -1, errors.New("id required and can't null")
	} else if !required && len(id) < 0 {
		return true, -1, nil
	} else {
		result, err := strconv.Atoi(id)
		if err != nil {
			return true, -1, errors.New("id can't parse to int")
		}

		return false, result, nil
	}
}
