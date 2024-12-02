package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"github.com/schollz/progressbar/v3"
)

func main() {
	inputDir, outputDir := "./input/", "./output/"

	// 入力ディレクトリ内の全てのファイルを取得
	files, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatalf("Failed to read input directory: %v", err)
	}

	// outputDirが存在しない場合、ディレクトリを作成
	_, err = os.Stat(outputDir)
	if os.IsNotExist(err) {
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

	// 非可逆圧縮回数
	maxIter := 10
	if len(os.Args) > 2 {
		if os.Args[2] == "random" {
			maxIter = 1 + int(time.Now().UnixNano()) % 1000
		} else {
			q, err := strconv.Atoi(os.Args[2])
			if err == nil && q >= 1 && q <= 1000 {
				maxIter = q
			} else {
				log.Fatalf("Max iteration must be between 1 and 1000")
			}
		}
	}

	method := os.Args[3]

	for _, file := range files {
		inputPath  := filepath.Join(inputDir, file.Name())
		outputPath := filepath.Join(outputDir, "image_"+time.Now().Format("20060102_150405")+".jpg")

		log.Printf("Compressing %v with quality %v and %v iterations\n", file.Name(), quality, maxIter)

		// プログレスバーの初期化
		example := "[" + file.Name() + " -> " + outputPath + "]"
		bar := progressbar.NewOptions(maxIter,
			progressbar.OptionSetWidth(50),              // プログレスバーの幅
			progressbar.OptionSetPredictTime(true),      // 残り時間の予測
			progressbar.OptionSetDescription(example), 	 // バーの前に表示する説明
			progressbar.OptionSetRenderBlankState(true), // 空の状態でも表示
		)

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
			err = Encode(&buf, img, quality, method)
			if err != nil {
				log.Printf("Failed to encode image during iteration %v: %v", i, err)

				break
			}
			
			// 画像をデコード
			img, err = DecodeImageFromReader(&buf, "." + method)
			if err != nil {
				log.Printf("Failed to decode image during iteration %v: %v", i, err)
				
				break
			}
			
			if method == "webp" {
				method = "jpg"
			} else if method == "jpg" || method == "jpeg" {
				method = "webp"
			} else {
				log.Fatalf("Invalid format: %v", method)

				break
			}

			// プログレスバーを更新
			err = bar.Add(1)
			if err != nil {
				fmt.Println("Failed to update progress bar", err)
				
				break
			}
		}

		// 結果を保存
		SaveImage(outputPath, img, 100, method)

		log.Printf("Image successfully compressed and saved as %v", outputPath)
	}
}