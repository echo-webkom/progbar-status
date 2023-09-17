package models

import (
	redis "github.com/omfj/lol/internal"
)

func SetStatus(status string) error {
	return redis.RDB.Set("status", status, 0).Err()
}

func GetStatus() (string, error) {
	return redis.RDB.Get("status").Result()
}
