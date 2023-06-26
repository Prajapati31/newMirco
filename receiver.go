package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func reassembleVideos(chunksDir, outputDir string) {
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatal("Error creating output directory:", err)
	}

	fileInfos, err := os.ReadDir(chunksDir)
	if err != nil {
		log.Fatal("Error reading chunks directory:", err)
	}

	for _, fileInfo := range fileInfos {
		chunkDir := filepath.Join(chunksDir, fileInfo.Name())
		chunkFiles, err := os.ReadDir(chunkDir)
		if err != nil {
			log.Fatal("Error reading chunk directory:", err)
		}

		outputPath := filepath.Join(outputDir, fileInfo.Name()+".mp4")
		outputFile, err := os.Create(outputPath)
		if err != nil {
			log.Fatal("Error creating output file:", err)
		}
		defer outputFile.Close()

		for _, chunkFile := range chunkFiles {
			chunkPath := filepath.Join(chunkDir, chunkFile.Name())
			chunk, err := os.Open(chunkPath)
			if err != nil {
				log.Fatal("Error opening chunk file:", err)
			}
			defer chunk.Close()

			_, err = io.Copy(outputFile, chunk)
			if err != nil {
				log.Fatal("Error writing chunk to output file:", err)
			}
		}
		log.Println("File reassembled:", outputPath)
	}
}

func main() {
	chunksDir := "chunks"
	outputDir := "output"

	reassembleVideos(chunksDir, outputDir)
}
