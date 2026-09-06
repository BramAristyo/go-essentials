package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctxTO, cancelTO := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)

	defer cancelTO()

	go slowQuery(ctxTO)

	time.Sleep(5 * time.Second)

	// context cancel

	ctxCan, cancel := context.WithCancel(context.Background())

	go worker(ctxCan)

	time.Sleep(4 * time.Second)

	fmt.Println("cancel worker...")
	cancel()

	time.Sleep(2 * time.Second)

}

func slowQuery(ctx context.Context) {
	fmt.Println("slow query started ... ")

	select {

	case <-time.After(5 * time.Second):
		fmt.Println("Success .. ")

	case <-ctx.Done():
		fmt.Println("timeout :", ctx.Err())

	}
}

func worker(ctx context.Context) {
	for {
		select {

		case <-ctx.Done():
			fmt.Println("worker stopped")
			return

		default:
			fmt.Println("worker running ...")
			time.Sleep(1 * time.Second)
		}
	}
}
