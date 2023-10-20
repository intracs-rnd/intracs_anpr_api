package capture

import "intracs_anpr_api/model"

func (c *Controller) GetImage(id int) (model.CaptureImage, error) {
	result, err := c.service.GetImage(id)

	return result, err
}

func (c *Controller) GetPlateImage(id int) (string, error) {
	result, err := c.service.GetPlateImage(id)

	return result, err
}

func (c *Controller) GetFullImage(id int) (string, error) {
	result, err := c.service.GetFullImage(id)

	return result, err
}
