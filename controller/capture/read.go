package capture

import (
	"intracs_anpr_api/internal/response"
	"intracs_anpr_api/model"
	"intracs_anpr_api/types"
)

func (c *Controller) ReadById(id int) (model.Capture, error) {
	// Validation id
	result, err := c.service.ReadById(id)
	if err != nil {
		return model.Capture{}, err
	}

	return result, nil
}

func (c *Controller) Read(page types.PageFilter, date types.DateFilter, validationStatus types.ValidationStatusFilter, withPlateImage bool) ([]model.Capture, response.PaginationData, error) {

	if page.Number == 0 {
		page.Number = 1
	}

	if page.Limit == 0 {
		page.Limit = 25
	}

	result, rowCount, err := c.service.Read(page, withPlateImage, date, validationStatus)

	pagingData := response.NewPaging(int(rowCount), int(page.Number), int(page.Limit))

	return result, pagingData, err
}
