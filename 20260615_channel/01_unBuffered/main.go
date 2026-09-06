package main

import "fmt"

func main() {
	// init channel
	messages := make(chan string) // unbuffered channel

	go func() {
		messages <- "ping" // this is sender (send to channel), sender receiver come from operation not a channel variables
	}()

	msg := <-messages // this is receiver (receive data from channel)
	fmt.Println(msg)
}
