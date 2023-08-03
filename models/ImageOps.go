package models

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
)

func Resize(filepath string, width, height uint) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	newImage := resize.Resize(width, height, img, resize.Lanczos3)
	buffer := new(bytes.Buffer)

	if err = jpeg.Encode(buffer, newImage, nil); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
