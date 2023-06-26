package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func splitFileIntoChunks(filePath string, chunkSize int64) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	chunkNumber := 0
	buffer := make([]byte, chunkSize)

	outputDir := filepath.Join("chunks", filepath.Base(filePath))
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatal("Error creating output directory:", err)
	}

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal("Error reading file:", err)
		}
		if n == 0 {
			break
		}

		outputPath := filepath.Join(outputDir, "chunk"+strconv.Itoa(chunkNumber))
		chunkFile, err := os.Create(outputPath)
		if err != nil {
			log.Fatal("Error creating chunk file:", err)
		}
		defer chunkFile.Close()

		_, err = chunkFile.Write(buffer[:n])
		if err != nil {
			log.Fatal("Error writing chunk:", err)
		}

		chunkNumber++
	}
}

func main() {
	filePaths := []string{
		"pexels-ambientnature-atmosphere-3929990-1920x1080-30fps.mp4",
		"pexels-kelly-4208317-3840x2160-24fps.mp4",
		"pexels-cottonbro-studio-3403583-2160x4096-50fps.mp4",
		"pexels-francesco-morrone-4185375-2024x3840-24fps.mp4",
		"pexels-cottonbro-studio-2795730-3840x2160-25fps.mp4",
	}

	chunkSize := int64(1024 * 1024) // Chunk size in bytes (1MB)

	// Create the base chunks directory if it doesn't exist
	err := os.Mkdir("chunks", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal("Error creating chunks directory:", err)
	}

	for _, filePath := range filePaths {
		splitFileIntoChunks(filePath, chunkSize)
		log.Println("File", filePath, "split into chunks")
	}
}
