package main

import (
	"log"
)

// Hub serves as a central message dispatcher and state manager for all active WebSocket connections.
// It acts as a fan-out controller, orchestrating the distribution of message payloads 
// while managing client registration and unregistration via synchronized channel operations.
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

// newHub implements the factory pattern, initializing the Hub with thread-safe data structures
// and buffered channels for optimized asynchronous message passing.
func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte, 256),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			log.Printf("Broadcast from client: %s", string(message))
			for client := range h.clients {
				select {
				case client.send <- message:
				// Gracefully terminate connections that fail to keep pace with the broadcast frequency.
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
