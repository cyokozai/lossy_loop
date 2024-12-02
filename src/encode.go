package main

import (
	"io"
	"log"
	"image"
	"image/jpeg"
	"github.com/chai2010/webp"
)

// Encode: 画像をJPEG・WebP形式でエンコード
func Encode(w io.Writer, img image.Image, quality int, method string) error {
	if method == "jpg" || method == "jpeg" {
		return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
	} else if method == "webp" {
		return webp.Encode(w, img, &webp.Options{Lossless: false, Quality: float32(quality)})
	} else {
		log.Fatalf("Invalid format: %v", method)

		return nil
	}
}