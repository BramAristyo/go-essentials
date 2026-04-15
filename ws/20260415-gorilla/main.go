package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	// Initialize the Hub as a long-lived orchestrator to handle asynchronous event multiplexing.
	hub := newHub()
	go hub.run()

	// Handle WebSocket upgrade requests, transitioning HTTP connections to full-duplex TCP communication.
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		client := &Client{
			hub: hub, conn: conn, send: make(chan []byte, 256),
		}

		// Enqueue the new client into the Hub's registration system to synchronize session state.
		client.hub.register <- client

		// Spawn concurrent goroutines for read/write pumps, enabling non-blocking bi-directional throughput.
		go client.read()
		go client.write()
	})

	fmt.Println("Running on port :9999")

	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal("Server err: ", err)
	}

}
