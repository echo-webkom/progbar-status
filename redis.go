package main

import (
	"github.com/go-redis/redis"
)

var (
	RDB *redis.Client
)

type Redis struct {
	Addr string
}

func (d *Redis) Connect() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     d.Addr,
		Password: "",
		DB:       0,
	})

	err := RDB.Ping().Err()

	return err
}
