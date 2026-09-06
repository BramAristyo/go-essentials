package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/BramAristyo/go-essentials/20260906_redis/internal/handler"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(pong)

	redisHandler := handler.NewRedisHandler(rdb)

	http.HandleFunc("POST /jobs", redisHandler.Push)
	http.HandleFunc("GET /jobs", redisHandler.GetAll)
	http.HandleFunc("POST /jobs/claim", redisHandler.Claim)

	defer rdb.Close()
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
