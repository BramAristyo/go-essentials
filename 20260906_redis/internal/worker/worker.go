package worker

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Worker struct {
	rdb    *redis.Client
	jobKey string
}

func NewWorker(rdb *redis.Client, jobKey string) *Worker {
	return &Worker{
		rdb:    rdb,
		jobKey: jobKey,
	}
}

func (w *Worker) Run(ctx context.Context) {
	fmt.Println("worker is ready ...")

	for {
		// BRPop = Blocking Right Pop
		res, err := w.rdb.BRPop(context.Background(), 0, w.jobKey).Result()
		if err != nil {
			fmt.Printf("Error :%s/n", err)
			continue
		}

		fmt.Println("Job received :", res)
	}
}
