package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	Code     string
	PrepTime time.Duration
}

func processOrder(wg *sync.WaitGroup, order Order) {
	defer wg.Done()

	fmt.Printf("Preparing table: %s\n", order.Code)

	time.Sleep(order.PrepTime)

	fmt.Printf("Table %s ready!\n", order.Code)
}

func main() {
	var wg sync.WaitGroup
	orders := []Order{
		{Code: "A", PrepTime: 5 * time.Second},
		{Code: "B", PrepTime: 7 * time.Second},
		{Code: "C", PrepTime: 3 * time.Second},
		{Code: "D", PrepTime: 6 * time.Second},
		{Code: "E", PrepTime: 5 * time.Second},
	}

	for _, order := range orders {
		wg.Add(1)
		go processOrder(&wg, order)
	}

	wg.Wait()
}
