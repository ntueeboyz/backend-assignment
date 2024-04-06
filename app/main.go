package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var (
    rdb            *redis.Client
    ctx            = context.Background()
)

func initRedis() {
    rdb = redis.NewClient(&redis.Options{
        Addr:     "redis:6379", // Redis server address
        Password: "",               // no password set
        DB:       0,                // use default DB
    })
}

func main() {
    initRedis()
    router := gin.Default()

    // Define routes and associate them with their handlers
    router.POST("/api/v1/ad", createAd)
    router.GET("/api/v1/ad", listAds)

    router.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
