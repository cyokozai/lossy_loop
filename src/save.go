package main

import (
	"log"
	"os"
	"image"
)

// SaveImage: 指定されたパスに画像を保存します
func SaveImage(path string, img image.Image, quality int, method string) {
	// 出力ファイルを作成
	file, err := os.Create(path)
	if err != nil {
		log.Println("Failed to create output file:", err)
		return
	}
	defer file.Close()

	// 画像をエンコードしてファイルに書き込む
	err = Encode(file, img, quality, method)
	if err != nil {
		log.Println("Failed to encode image:", err)
		return
	}

	log.Println("Image successfully saved:", path)
}
