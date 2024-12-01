package main

import (
	"image"
	"image/jpeg"
	"os"
)

// EncodeJPEG: 画像をJPEG形式でエンコードします
func EncodeJPEG(file *os.File, img image.Image, quality int) error {
	return jpeg.Encode(file, img, &jpeg.Options{Quality: quality})
}
