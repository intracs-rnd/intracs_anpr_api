package validation

func (c *Controller) ValidCount() (int64, error) {
	var result int64 = -1
	var err error

	result, err = c.service.ValidCount()

	return result, err
}

func (c *Controller) InvalidCount() (int64, error) {
	var result int64 = -1
	var err error

	result, err = c.service.InvalidCount()

	return result, err
}
