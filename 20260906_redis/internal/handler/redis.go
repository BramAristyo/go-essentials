package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/BramAristyo/go-essentials/20260906_redis/internal/constants"
	"github.com/redis/go-redis/v9"
)

type RedisHandler struct {
	rdb *redis.Client
}

func NewRedisHandler(rdb *redis.Client) *RedisHandler {
	return &RedisHandler{rdb: rdb}
}

func (rh *RedisHandler) Push(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Job string `json:"job"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	err := rh.rdb.LPush(r.Context(), constants.JOB_KEY, body.Job).Err()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to push job", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"job":    body.Job,
	})
}

func (rh *RedisHandler) PushWithType(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Job  string `json:"job"`
		Type int    `json:"type"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	switch body.Type {
	case 1:
		err := rh.rdb.LPush(r.Context(), constants.JOB_KEY, body.Job).Err()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "failed to push job", http.StatusInternalServerError)
			return
		}
	default:
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"job":    body.Job,
	})
}

func (rh *RedisHandler) Claim(w http.ResponseWriter, r *http.Request) {
	result, err := rh.rdb.BRPop(r.Context(), 1, constants.JOB_KEY).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "empty",
				"message": "no job available",
			})
			return
		}

		http.Error(w, "failed to claim job", http.StatusInternalServerError)
		return
	}

	res := result[1] + " processed"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"result": res,
	})
}

func (rh *RedisHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	jobs, err := rh.rdb.LRange(r.Context(), constants.JOB_KEY, 0, -1).Result()
	if err != nil {
		http.Error(w, "failed to fetch jobs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"jobs": jobs,
	})
}
