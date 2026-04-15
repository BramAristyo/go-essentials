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

	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:9999/ws", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

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

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(msg))
	}
}
