package helpers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func ValidateImage(data string) (image.Image, string, error) {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return nil, "", errors.New("Imagen invalida")
	}

	imageType := data[11:idx]

	unbase, err := base64.StdEncoding.DecodeString(data[idx+8:])

	if err != nil {
		return nil, "", errors.New("Error unbase")
	}

	r := bytes.NewReader(unbase)
	var img image.Image

	switch imageType {
	case "png":
		img, err = png.Decode(r)
		if err != nil {
			return nil, "", errors.New("Bad png")
		}
	case "jpeg":
		img, err = jpeg.Decode(r)
		if err != nil {
			return nil, "", errors.New("Bad jpeg")
		}

	default:
		return nil, "", errors.New("Formato de imagen no válido, solo es permitido PNG, JPEG")
	}
	return img, imageType, err
}

func SaveImageToDisk(img image.Image, name string, imageType string, folder string) (string, error) {
	var filenameFull string
	pathFile := folder + name
	switch imageType {
	case "png":
		filenameFull = pathFile + ".png"

		f, err := os.OpenFile("./public"+filenameFull, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return "", errors.New("Cannot open file")
		}

		png.Encode(f, img)
	case "jpeg":
		filenameFull = pathFile + ".jpeg"

		f, err := os.OpenFile("./public"+filenameFull, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return "", errors.New("Cannot open file")
		}

		jpeg.Encode(f, img, nil)
	}
	return filenameFull, nil
}
