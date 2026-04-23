package main

import (
	"fmt"
	"time"
)

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

	// select picks the first available channel communication
	select {
	case serve := <-chPizza:
		fmt.Println("received ", serve)
	case serve := <-chMatcha:
		fmt.Println("received ", serve)
	case <-time.After(5 * time.Second):
		panic("timeout: service too slow!")
	}
}
