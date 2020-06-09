package main

import (
	"context"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {

	const greetings = "Welcome to Gin "
	var ctx = context.Background()
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPath := redisHost + ":" + redisPort

	client := redis.NewClient(&redis.Options{
		Addr: redisPath,
	})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		number, err := client.Incr(ctx, "number").Result()
		if err != nil {
			panic(err)
		}

		c.String(200, greetings + strconv.FormatInt(number, 10) + "\n")

	})

	r.Run(":80")

}
