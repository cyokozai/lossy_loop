package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
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

	// ファイル拡張子に応じてデコード
	var img image.Image
	switch ext {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	case ".webp":
		img, err = webp.Decode(file)
	default:
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	return img, nil
}
