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
	var inputPath, outputPath string

	// ユーザーから入力画像のパスを取得
	inputPath 
	
	// 出力画像のパス
	outputPath := "./out/" + "image_" + time.Now().Format("20060102_150405") + ".jpg"

	// コマンドライン引数から品質を取得（デフォルトは80）
	quality := 80
	if len(os.Args) == 1 {
		if os.Args[1] == "random" {
			quality = 1 + int(time.Now().UnixNano()) % 100
		} else {
			q, err := strconv.Atoi(os.Args[1])
			if err == nil && q >= 1 && q <= 100 {
				quality = q
			}
		}
	}

	// 画像をデコード
	img, err := DecodeImage(inputPath, filepath.Ext(inputPath))
	if err != nil {
		log.Fatalf("Failed to decode image: %v", err)
	}

	// 画像をエンコードして保存
	err = SaveImage(outputPath, img, quality)
	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}

	log.Printf("Image successfully compressed and saved as %v", outputPath)

	return 0
}
