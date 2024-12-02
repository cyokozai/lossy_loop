package main

import (
	"image"
	"image/jpeg"
	"io"
)

// EncodeJPEG: 画像をJPEG形式でエンコードします
func EncodeJPEG(w io.Writer, img image.Image, quality int) error {
	return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
}
