package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("cooking start")
		time.Sleep(2 * time.Second)
		ch <- "🍕🍕🍕"
		fmt.Println("done!")
	}()

	fmt.Println("wait pizza .... ")
	// Blocks execution until the channel receives a value
	pizza := <-ch
	fmt.Println("nom nom ... ", pizza)
}
