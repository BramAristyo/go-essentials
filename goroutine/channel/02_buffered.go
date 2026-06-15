package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// UnBuffered only process one channel (blocking) or: make(chan string, 0)
// Sender         Channel        Receiver
//   |               |              |
//   |---- ping ---->|              |
//   |  BLOCK        |              |
//   |               |<---- ready --|
//   |               |              |
//   |---- transfer --------------->|

// Buffered channel can process concurently many channel tunneling
// +--------+
// | ping1  |
// | ping2  |
// | ping3  |
// +--------+
