package handler

import "github.com/redis/go-redis/v9"

type StatusHandler struct {
	rdb *redis.Client
}

func NewStatusHandler(rdb *redis.Client) *StatusHandler {
	return &StatusHandler{rdb: rdb}
}
