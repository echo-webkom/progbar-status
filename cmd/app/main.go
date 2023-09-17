package main

import (
	"github.com/gin-gonic/gin"
	redis "github.com/omfj/lol/internal"
	"github.com/omfj/lol/internal/handlers"
)

func main() {
	redis := &redis.Redis{
		Addr: "localhost:6379",
	}

	err := redis.Connect()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(gin.Logger())

	r.GET("/", handlers.HealthCheck)
	r.GET("/status", handlers.GetStatus)
	r.POST("/status", handlers.UpdateStatus)

	r.Run()
}
