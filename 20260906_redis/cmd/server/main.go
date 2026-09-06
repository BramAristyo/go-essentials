package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/BramAristyo/go-essentials/20260906_redis/internal/constants"
	"github.com/BramAristyo/go-essentials/20260906_redis/internal/handler"
	"github.com/BramAristyo/go-essentials/20260906_redis/internal/worker"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w := worker.NewWorker(rdb, constants.JOB_KEY)
	go w.Run(context.Background())

	redisHandler := handler.NewRedisHandler(rdb)

	http.HandleFunc("POST /jobs", redisHandler.Push)
	http.HandleFunc("GET /jobs", redisHandler.GetAll)
	http.HandleFunc("POST /jobs/claim", redisHandler.Claim)

	http.HandleFunc("POST /jobs/with-type", redisHandler.PushWithType)

	defer rdb.Close()
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
