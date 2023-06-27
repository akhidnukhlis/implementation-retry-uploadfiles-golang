package main

import (
	"fmt"
	"implementation-retry-uploadfiles-golang/config"
	"implementation-retry-uploadfiles-golang/usecase"
	"log"
	"time"
)

func main() {
	var filePath = "files/example.txt"

	// Initialize Redis client and file uploader
	fileUploader := config.NewRedisFileUploader("localhost:6379", "password")

	// Wrap file uploader with retry mechanism
	retryFileUploader := usecase.NewRetryFileUploader(fileUploader, 3, time.Second*5)

	// Upload file
	err := retryFileUploader.UploadFile(filePath)
	if err != nil {
		log.Println("Failed to upload file:", err)
	}

	fmt.Println("File uploaded successfully")

	// Close Redis client connection
	err = fileUploader.Close()
	if err != nil {
		log.Println("Error closing Redis client:", err)
	}
}
