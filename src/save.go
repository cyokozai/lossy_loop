package main

import (
	"fmt"
	"image"
	"os"
)

// SaveImage: 指定されたパスに画像を保存します
func SaveImage(path string, img image.Image, quality int) error {
	// 出力ファイルを作成
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	// JPEG形式で保存
	err = EncodeJPEG(file, img, quality)
	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	return nil
}
