package types

import (
	"errors"
	"strconv"
)

type PageFilter struct {
	Required bool
	Limit    uint
	Number   uint
}

func (p *PageFilter) FromString(limit string, number string, required bool) (PageFilter, error) {
	var defaultLimit uint = 25
	var defaultNumber uint = 1

	if required && limit == "" {
		return *p, errors.New("page limit can't null")
	}

	if required && number == "" {
		return *p, errors.New("page number can't null")
	}

	pageLimit, err := strconv.ParseUint(limit, 10, 32)
	if required && err != nil {
		return *p, errors.New("page limit invalid or can't parse to uint")
	} else if pageLimit > 0 {
		defaultLimit = uint(pageLimit)
	}

	pageNumber, err := strconv.ParseUint(number, 10, 32)
	if required && err != nil {
		return *p, errors.New("page number invalid or can't parse to uint")
	} else if pageNumber > 0 {
		defaultNumber = uint(pageNumber)
	}

	return PageFilter{
		Required: required,
		Limit:    defaultLimit,
		Number:   defaultNumber,
	}, nil
}

func (p *PageFilter) IsValid() bool {
	if p.Limit > 0 && p.Number > 0 {
		return true
	}
	return false
}

func (p *PageFilter) GetOffset() uint {
	return (p.Number - 1) * p.Limit
}
