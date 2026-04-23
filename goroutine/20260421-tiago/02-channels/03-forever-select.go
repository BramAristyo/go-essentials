package main

import (
	"fmt"
	"time"
)

// Use for + select to handle continuous channel events (e.g., WebSockets)
func main() {
	chPizza := make(chan string)
	chMatcha := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		chPizza <- "🍕"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		chMatcha <- "🧃"
	}()

	// Infinite loop to serve all incoming channel events
	for {
		select {
		case serve := <-chPizza:
			fmt.Println("received ", serve)
		case serve := <-chMatcha:
			fmt.Println("received ", serve)
		case <-time.After(5 * time.Second):
			panic("timeout: no more signals received")
		}
	}
}
