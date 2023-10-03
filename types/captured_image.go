package types

import (
	"net/http"
	"strconv"
	"strings"
)

type ImageCapture interface {
	FromRequest(id string, w http.ResponseWriter, r *http.Request) (ImageCaptured, error)
	FromString(id string,
		deviceId string,
		speed string,
		width string,
		height string,
		plateCoordinates []string,
		plateImage string,
		fullImage string) (ImageCaptured, error)
}

type ImageCaptured struct {
	Id               string
	DeviceId         string
	Speed            float64
	Width            int
	Height           int
	PlateCoordinates []int
	PlateDetectConf  float64
	TextRecogConf    float64
	PlateNumber      string
	PlateImage       string
	FullImage        string
	CapturedAt       string
}

func (data ImageCaptured) FromRequest(id string, w http.ResponseWriter, r *http.Request) (ImageCaptured, error) {
	speed := r.FormValue("speed")
	width := r.FormValue("width")
	height := r.FormValue("height")
	deviceId := r.FormValue("device_id")
	fullImage := r.FormValue("full_image")
	plateImage := r.FormValue("plate_image")
	plateNumber := r.FormValue("plate_number")
	textRecogConf := r.FormValue("text_recog_conf")
	plateDetectConf := r.FormValue("plate_detect_conf")
	plateCoordinates := strings.Split(r.FormValue("plate_coordinate"), ",")

	// Speed
	speedData, err := strconv.ParseFloat(speed, 64)
	if err != nil {
		http.Error(w, "Failed parse speed data to float!", http.StatusBadRequest)
		return ImageCaptured{}, err
	}

	widthData, err := strconv.Atoi(width)
	if err != nil {
		http.Error(w, "Failed parse width data to int!", http.StatusBadRequest)
		return ImageCaptured{}, err
	}

	heightData, err := strconv.Atoi(height)
	if err != nil {
		http.Error(w, "Failed parse height data to int!", http.StatusBadRequest)
		return ImageCaptured{}, err
	}

	plateDetectConfData, err := strconv.ParseFloat(plateDetectConf, 64)
	if err != nil {
		http.Error(w, "Failed parse plate detect conf data to float!", http.StatusBadRequest)
		return ImageCaptured{}, err
	}

	textRecogConfData, err := strconv.ParseFloat(textRecogConf, 64)
	if err != nil {
		http.Error(w, "Failed parse text recog conf data to float!", http.StatusBadRequest)
		return ImageCaptured{}, err
	}

	x1, err := strconv.Atoi(strings.TrimSpace(plateCoordinates[0]))
	if err != nil {
		http.Error(w, "Failed parse x1 data to int!", http.StatusBadRequest)
		return ImageCaptured{}, err
	}

	y1, err := strconv.Atoi(strings.TrimSpace(plateCoordinates[1]))
	if err != nil {
		http.Error(w, "Failed parse y1 data to int!", http.StatusBadRequest)
		return ImageCaptured{}, err
	}

	x2, err := strconv.Atoi(strings.TrimSpace(plateCoordinates[2]))
	if err != nil {
		http.Error(w, "Failed parse x2 data to int!", http.StatusBadRequest)
		return ImageCaptured{}, err
	}

	y2, err := strconv.Atoi(strings.TrimSpace(plateCoordinates[3]))
	if err != nil {
		http.Error(w, "Failed parse y1 data to int!", http.StatusBadRequest)
		return ImageCaptured{}, err
	}

	data.Id = id
	data.DeviceId = deviceId
	data.Speed = speedData
	data.Width = widthData
	data.Height = heightData
	data.PlateNumber = plateNumber
	data.PlateImage = plateImage
	data.FullImage = fullImage
	data.PlateDetectConf = plateDetectConfData
	data.TextRecogConf = textRecogConfData

	// Convert string coordinates to float
	data.PlateCoordinates = make([]int, 4)
	data.PlateCoordinates[0] = x1
	data.PlateCoordinates[1] = y1
	data.PlateCoordinates[2] = x2
	data.PlateCoordinates[3] = y2

	return data, nil
}

func (data ImageCaptured) FromString(
	id string,
	deviceId string,
	speed string,
	width string,
	height string,
	plateCoordinates []string,
	plateImage string,
	fullImage string,
) (ImageCaptured, error) {
	speedData, err := strconv.ParseFloat(speed, 64)
	if err != nil {
		return ImageCaptured{}, err
	}

	widthData, err := strconv.Atoi(width)
	if err != nil {
		return ImageCaptured{}, err
	}

	heightData, err := strconv.Atoi(height)
	if err != nil {
		return ImageCaptured{}, err
	}

	x1, err := strconv.Atoi(strings.TrimSpace(plateCoordinates[0]))
	if err != nil {
		return ImageCaptured{}, err
	}

	y1, err := strconv.Atoi(strings.TrimSpace(plateCoordinates[1]))
	if err != nil {
		return ImageCaptured{}, err
	}

	x2, err := strconv.Atoi(strings.TrimSpace(plateCoordinates[2]))
	if err != nil {
		return ImageCaptured{}, err
	}

	y2, err := strconv.Atoi(strings.TrimSpace(plateCoordinates[3]))
	if err != nil {
		return ImageCaptured{}, err
	}

	data.Id = id
	data.DeviceId = deviceId
	data.Speed = speedData
	data.Width = widthData
	data.Height = heightData
	data.PlateImage = plateImage
	data.FullImage = fullImage

	// Convert string coordinates to float
	data.PlateCoordinates = make([]int, 4)
	data.PlateCoordinates[0] = x1
	data.PlateCoordinates[1] = y1
	data.PlateCoordinates[2] = x2
	data.PlateCoordinates[3] = y2

	return data, nil
}
