package api

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"intracs_anpr_api/repositories"
	"intracs_anpr_api/types"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func imageToBase64(image multipart.File) (string, error) {
	imageIO, err := io.ReadAll(image)
	if err != nil {
		return "Error Reading the File", err
	}

	contentType := http.DetectContentType(imageIO)

	switch contentType {
	case "image/jpg":
		imageArray, err := jpeg.Decode(bytes.NewReader(imageIO))
		if err != nil {
			fmt.Println("Error Decoding the Image")
		}

		jpegOption := jpeg.Options{Quality: 100}
		var imageBuffer bytes.Buffer
		if err := jpeg.Encode(&imageBuffer, imageArray, &jpegOption); err != nil {
			fmt.Println("Error Encoding the PNG")
		}
		imageIO = imageBuffer.Bytes()
	}

	imageBase64 := "data:" + contentType + ";base64," + base64.StdEncoding.EncodeToString(imageIO)

	return imageBase64, nil
}

func handleBase64Image(field string, r *http.Request) (string, multipart.File, *multipart.FileHeader, error) {
	imageFile, header, err := r.FormFile(field)
	if err != nil {
		fmt.Println("Error Retrieving the File")
		return "", nil, nil, err
	}
	defer imageFile.Close()

	imageString, err := imageToBase64(imageFile)
	if err != nil {
		fmt.Println("Error Converting the Image to Base64")
		return "", nil, nil, err
	}

	return imageString, imageFile, header, nil
}

func InsertCapture(w http.ResponseWriter, r *http.Request) {
	// Maximum file size 10 MB
	r.ParseMultipartForm(1 << 2)

	// get form values
	speed := r.FormValue("speed")
	deviceId := r.FormValue("device_id")
	width := r.FormValue("width")
	height := r.FormValue("height")
	plateCoordinate := strings.Split(r.FormValue("plate_coordinate"), ",")

	// get plate image
	plateImage, _, _, err := handleBase64Image("plate_image", r)
	if err != nil {
		http.Error(w, "Failed convert plate image to Base64!", http.StatusBadRequest)
		return
	}

	// get full image
	fullImage, _, _, err := handleBase64Image("full_image", r)
	if err != nil {
		http.Error(w, "Failed convert full image to Base64!", http.StatusBadRequest)
		return
	}

	var capturedImage types.ImageCapture
	capturedImage = types.ImageCaptured{}

	data, err := capturedImage.FromString(
		"-1",
		deviceId,
		speed,
		width,
		height,
		plateCoordinate,
		plateImage,
		fullImage,
	)
	if err != nil {
		http.Error(w, "Failed to create image capture!", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// insert captured image
	rows, err := repositories.InsertCapture(w, data)
	if err != nil {
		http.Error(w, "Failed to insert captured image!", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer rows.Close()

	// resultData := map[string]string{
	// 	"device_id":         deviceId,
	// 	"speed":             speed,
	// 	"width":             width,
	// 	"height":            height,
	// 	"plate_coordinates": "[" + strings.Join(plateCoordinate, ",") + "]",
	// }

	// var apiResponse types.ApiResponses
	// apiResponse = types.ApiResponse{}
	// apiResponse = apiResponse.GetApiResponses("success", "image captured successfully", resultData)

	// response, err := json.Marshal(apiResponse)
	// if err != nil {
	// 	http.Error(w, "Failed to marshal response!", http.StatusInternalServerError)
	// }

	// fmt.Println(response)
	// fmt.Fprintln(w, string(response))
	fmt.Fprintln(w, "Captured image uploaded successfully")
}
