package types

import (
	"strconv"
	"strings"
)

type ImageCapture interface {
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
	PlateImage       string
	FullImage        string
	CapturedAt       string
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
