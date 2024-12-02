package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	inputDir, outputDir := "./input/", "./output/"

	fmt.Printf("Input directory: %v\n", inputDir) // 入力ディレクトリ内の全てのファイルを取得
	files, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatalf("Failed to read input directory: %v", err)
	}

	fmt.Printf("Output directory: %v\n", outputDir) // outputDirが存在しない場合、ディレクトリを作成
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.Mkdir(outputDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create output directory: %v", err)
		}
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
	} else {
		fmt.Println("Usage: go run main.go [quality]")
		fmt.Println("     : ./lussyloop [quality]")
		fmt.Println("  quality: 1 ~ 100 (default: 80)")
		fmt.Println("         : random (random quality)")
	}

	// 非可逆変換回数
	maxIter := 1000
	if len(os.Args) > 2 {
		if os.Args[2] == "random" {
			maxIter = 1 + int(time.Now().UnixNano()) % 1000
		} else {
			q, err := strconv.Atoi(os.Args[2])
			if err == nil && q >= 1 && q <= 1000 {
				log.Fatalf("Max iteration must be between 1 and 1000")
				maxIter = q
			}
		}
	}

	fmt.Printf("Quality: %v\n", quality)

	for _, file := range files {
		inputPath := filepath.Join(inputDir, file.Name())
		outputPath := filepath.Join(outputDir, "image_"+time.Now().Format("20060102_150405")+".jpg")

		fmt.Printf("Compressing %v with quality %v...\n", inputPath, quality)

		// 画像をデコード
		img, err := DecodeImage(inputPath, filepath.Ext(inputPath))
		if err != nil {
			log.Printf("Failed to decode image %v: %v", inputPath, err)

			continue
		}

		// 1000回非可逆変換
		for i := 0; i < maxIter; i++ {
			var buf bytes.Buffer
			
			// 画像をエンコード
			err = EncodeJPEG(&buf, img, quality)
			if err != nil {
				log.Printf("Failed to encode image during iteration %v: %v", i, err)

				break
			}
			
			// 画像をデコード
			img, err = DecodeImageFromReader(&buf, ".jpg")
			if err != nil {
				log.Printf("Failed to decode image during iteration %v: %v", i, err)
				
				break
			}
		}

		// 最終結果を保存
		err = SaveImage(outputPath, img, 80)
		if err != nil {
			log.Printf("Failed to save image %v: %v", outputPath, err)

			continue
		} else {
			fmt.Printf("Image successfully compressed and saved as %v\n", outputPath)
			log.Printf("Image successfully compressed and saved as %v", outputPath)
		}
	}
}