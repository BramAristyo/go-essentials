package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	defer rdb.Close()
	ctx := context.Background()

	// Set & Get
	rdb.Set(ctx, "username", "camelia", 0)
	name, _ := rdb.Get(ctx, "username").Result()

	fmt.Println("name :", name)

	// Set & Get with TTL
	rdb.Set(ctx, "otp", "1234567910", 10*time.Second)
	otp, _ := rdb.Get(ctx, "otp").Result()

	fmt.Println("OTP:", otp)

	// check TTL left
	time.Sleep(3 * time.Second)
	ttl, _ := rdb.TTL(ctx, "otp").Result()
	fmt.Println("TTL OTP: ", ttl)

	// INCR
	// rdb.Set(ctx, "visit_count", 0, 0)
	rdb.Incr(ctx, "visit_count")
	rdb.Incr(ctx, "visit_count")
	rdb.Incr(ctx, "visit_count")

	visitCount, _ := rdb.Get(ctx, "visit_count").Result()
	fmt.Println("Visit Count: ", visitCount)

	// exists
	exists, _ := rdb.Exists(ctx, "username").Result()
	fmt.Println("username is exists? ", exists)

	// delete
	rdb.Del(ctx, "username")
	exists, _ = rdb.Exists(ctx, "username").Result()
}
