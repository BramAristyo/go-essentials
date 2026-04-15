package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:9999/ws", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	go func() {
		for {
			err := conn.WriteMessage(websocket.TextMessage, []byte("Hello world"))

			if err != nil {
				log.Fatal(err)
				return
			}

			time.Sleep(5 * time.Second)
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(msg))
	}
}
