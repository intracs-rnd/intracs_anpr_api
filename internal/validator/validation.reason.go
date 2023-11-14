package validator

import (
	"errors"
	"strconv"
)

// Validate code from string and required
//
// return isNil: bool, id: int, err: error
func ReasonCode(code string, required bool) (bool, uint8, error) {
	if required && len(code) <= 0 {
		return true, 0, errors.New("code required and can't null")
	} else if !required && len(code) < 0 {
		return true, 0, nil
	} else {
		result, err := strconv.ParseUint(code, 10, 8)
		if err != nil {
			return true, 0, errors.New("code can't parse to uint8")
		}

		return false, uint8(result), nil
	}
}
