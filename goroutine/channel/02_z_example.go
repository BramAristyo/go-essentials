package main

import "fmt"

func sendMessage(ch chan string) {
	ch <- "hello from channel"
}

func main() {
	ch := make(chan string)

	go sendMessage(ch)

	// queueing the data pipe
	msg := <-ch
	fmt.Println(msg)
}
