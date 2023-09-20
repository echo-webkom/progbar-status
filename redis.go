package main

import (
	"os"

	"github.com/go-redis/redis"
)

var (
	RDB *redis.Client
)

func ConnectToRedis() error {
	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		redisUrl = "localhost:6379"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")

	RDB = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: redisPassword,
		DB:       0,
	})

	err := RDB.Ping().Err()

	return err
}
