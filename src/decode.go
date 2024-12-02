package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"github.com/chai2010/webp"
)

// DecodeImage: 指定されたパスの画像をデコードします
func DecodeImage(path, ext string) (image.Image, error) {
	// 入力ファイルを開く
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	return DecodeImageFromReader(file, ext)
}

// DecodeImageFromReader: 指定されたリーダーから画像をデコードします
func DecodeImageFromReader(r io.Reader, ext string) (image.Image, error) {
	// ファイル拡張子に応じてデコード
	var img image.Image
	var err error
	switch ext {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(r)
	case ".png":
		img, err = png.Decode(r)
	case ".webp":
		img, err = webp.Decode(r)
	default:
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	return img, nil
}
