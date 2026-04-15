package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {

	var clientInitial string
	fmt.Print("Input client initial : ")
	fmt.Scanln(&clientInitial)

	// Initiate a WebSocket handshake with the server to establish a persistent full-duplex connection.
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:9999/ws", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	// Launch a dedicated goroutine to simulate periodic message ingress, maintaining active interaction.
	go func() {
		for {
			err := conn.WriteMessage(websocket.TextMessage, []byte("Hello world from "+clientInitial))

			if err != nil {
				log.Fatal(err)
				return
			}

			time.Sleep(2 * time.Second)
		}
	}()

	// The main execution thread acts as an event listener, processing incoming frames from the server.
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(msg))
	}
}
