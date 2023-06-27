package config

import (
	"github.com/go-redis/redis"
	"implementation-retry-uploadfiles-golang/usecase"
)

func NewRedisFileUploader(addr, password string) usecase.FileUploader {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return &usecase.RedisFileUploader{
		Client: client,
	}
}
