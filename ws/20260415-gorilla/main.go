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
	hub := newHub()
	go hub.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		client := &Client{
			hub: hub, conn: conn, send: make(chan []byte, 256),
		}

		client.hub.register <- client

		go client.read()
		go client.write()
	})

	fmt.Println("Running on port :9999")

	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal("Server err: ", err)
	}

}
