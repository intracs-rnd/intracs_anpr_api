package capture

import (
	"errors"
	"intracs_anpr_api/model"
	"net/http"
	"strconv"
	"strings"
)

func base64Validation(image string) bool {
	prefix := "data:image/jpeg;base64"

	if strings.HasPrefix(image, prefix) {
		return true
	}
	return false
}

func plateCoordinatesValidation(coordinate string) (model.Coordinate, error) {
	var err error
	coordinates := strings.Split(coordinate, ",")
	result := model.Coordinate{}
	coorInt := make([]int64, 4)

	if len(coordinates) != 4 {
		return model.Coordinate{}, errors.New("plate coordinates not have 4 data with format (x1, y1, x2, y2)")
	} else {
		for i := 0; i < len(coordinates); i++ {
			coorInt[i], err = strconv.ParseInt(strings.TrimSpace(coordinates[i]), 10, 16)
			if err != nil {
				return model.Coordinate{}, errors.New("plate coordinates index " + strconv.Itoa(i) + " (" + coordinates[i] + ") invalid and can't parse to int")
			}
		}

		result = model.Coordinate{
			X1: int16(coorInt[0]),
			Y1: int16(coorInt[1]),
			X2: int16(coorInt[2]),
			Y2: int16(coorInt[3]),
		}

		return result, nil
	}
}

func createCaptureValidation(r *http.Request) (model.Capture, error) {
	var err error
	result := model.Capture{}

	result.Info.DeviceId = r.FormValue("device_id")

	result.Info.Speed, err = strconv.ParseFloat(r.FormValue("speed"), 64)
	if err != nil {
		return model.Capture{}, errors.New("form speed invalid and can't parse to float")
	}

	result.Info.Width, err = strconv.Atoi(r.FormValue("width"))
	if err != nil {
		return model.Capture{}, errors.New("form width invalid and can't parse to int")
	}

	result.Info.Height, err = strconv.Atoi(r.FormValue("height"))
	if err != nil {
		return model.Capture{}, errors.New("form height invalid and can't parse to int")
	}

	result.Info.PlateDetectConf, err = strconv.ParseFloat(r.FormValue("plate_detect_conf"), 64)
	if err != nil {
		return model.Capture{}, errors.New("form plate detect confidence invalid and can't parse to float")
	}

	result.Info.TextRecogConf, err = strconv.ParseFloat(r.FormValue("text_recog_conf"), 64)
	if err != nil {
		return model.Capture{}, errors.New("form text recognized confidence invalid and can't parse to float")
	}

	result.Info.PlateCoordinates, err = plateCoordinatesValidation(r.FormValue("plate_coordinate"))
	if err != nil {
		return model.Capture{}, err
	}

	if !base64Validation(r.FormValue("plate_image")) {
		return model.Capture{}, errors.New("form plate image invalid " + r.FormValue("plate_image"))
	} else {
		result.Image.Plate = r.FormValue("plate_image")
	}

	if !base64Validation(r.FormValue("full_image")) {
		return model.Capture{}, errors.New("form full image invalid")
	} else {
		result.Image.Full = r.FormValue("full_image")
	}

	// Parse all request form to Capture model
	result.Info.PlateNumber = r.FormValue("plate_number")

	return result, nil
}

func (c *Controller) Create(r *http.Request) (int, error) {
	capture, err := createCaptureValidation(r)
	if err != nil {
		return -1, err
	}

	_, lastInsertId, err := c.service.Create(capture)
	if err != nil {
		return -1, err
	}

	return lastInsertId, nil
}
