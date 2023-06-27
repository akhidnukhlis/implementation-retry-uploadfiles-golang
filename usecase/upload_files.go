package usecase

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// FileUploader interface defines the methods for file upload operations
type FileUploader interface {
	UploadFile(fileName string) error
	Close() error
}

// RedisFileUploader implements FileUploader using Redis client library
type RedisFileUploader struct {
	Client *redis.Client
}

func (r *RedisFileUploader) UploadFile(fileName string) error {
	// Simulate file upload failure
	// return errors.New("Failed to upload file")

	// Upload file to Redis
	key := fmt.Sprintf("files:%s", fileName)
	err := r.Client.Set(key, "file content", 0).Err()
	if err != nil {
		return fmt.Errorf("Failed to upload file: %s", err)
	}

	return nil
}

func (r *RedisFileUploader) Close() error {
	return r.Client.Close()
}

// RetryFileUploader implements FileUploader with retry mechanism
type RetryFileUploader struct {
	uploader      FileUploader
	maxAttempts   int
	retryInterval time.Duration
}

func NewRetryFileUploader(uploader FileUploader, maxAttempts int, retryInterval time.Duration) FileUploader {
	return &RetryFileUploader{
		uploader:      uploader,
		maxAttempts:   maxAttempts,
		retryInterval: retryInterval,
	}
}

func (r *RetryFileUploader) UploadFile(fileName string) error {
	attempts := 0

	for attempts < r.maxAttempts {
		attempts++
		fmt.Printf("Attempting to upload file, attempt #%d\n", attempts)

		err := r.uploader.UploadFile(fileName)
		if err == nil {
			return nil // File upload successful, exit the function
		}

		fmt.Printf("Failed to upload file: %s\n", err)

		if attempts < r.maxAttempts {
			fmt.Printf("Retrying upload after %v\n", r.retryInterval)
			time.Sleep(r.retryInterval)
		}
	}

	return errors.New("Failed to upload file after multiple attempts")
}

func (r *RetryFileUploader) Close() error {
	return r.uploader.Close()
}
