package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	inputDir := "./input/"
	outputDir := "./output/"

	// 入力ディレクトリ内の全てのファイルを取得
	files, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatalf("Failed to read input directory: %v", err)
	}

	// コマンドライン引数から品質を取得（デフォルトは80）
	quality := 80
	if len(os.Args) > 1 {
		if os.Args[1] == "random" {
			quality = 1 + int(time.Now().UnixNano()) % 100
		} else {
			q, err := strconv.Atoi(os.Args[1])
			if err == nil && q >= 1 && q <= 100 {
				quality = q
			}
		}
	}

	for _, file := range files {
		if !file.IsDir() {
			inputPath := filepath.Join(inputDir, file.Name())
			outputPath := filepath.Join(outputDir, "image_"+time.Now().Format("20060102_150405")+".jpg")

			// 画像をデコード
			img, err := DecodeImage(inputPath, filepath.Ext(inputPath))
			if err != nil {
				log.Printf("Failed to decode image %v: %v", inputPath, err)
				continue
			}

			// 画像をエンコードして保存
			err = SaveImage(outputPath, img, quality)
			if err != nil {
				log.Printf("Failed to save image %v: %v", outputPath, err)
				continue
			}

			log.Printf("Image successfully compressed and saved as %v", outputPath)
		}
	}
}
