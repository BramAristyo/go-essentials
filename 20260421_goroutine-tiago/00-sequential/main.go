package main

import (
	"fmt"
	"time"
)

type Order struct {
	Code     string
	PrepTime time.Duration
}

func processOrder(order Order) {
	fmt.Printf("Preparing table: %s\n", order.Code)

	time.Sleep(order.PrepTime)

	fmt.Printf("Table %s ready!\n", order.Code)
}

func main() {
	orders := []Order{
		{Code: "A", PrepTime: 5 * time.Second},
		{Code: "B", PrepTime: 7 * time.Second},
		{Code: "C", PrepTime: 3 * time.Second},
		{Code: "D", PrepTime: 6 * time.Second},
		{Code: "E", PrepTime: 5 * time.Second},
	}

	for _, order := range orders {
		processOrder(order)
	}
}
