package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	maxMessageSize = 512
	pongWait       = 5 * time.Second
	pingPeriod     = (pongWait * 9) / 10
)

// Client represents a distinct WebSocket participant, encapsulating the stateful connection
// and providing a decoupled interface for asynchronous bi-directional communication.
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

// read facilitates the ingress of data by continuously monitoring the WebSocket socket.
// It implements the message-pumping pattern, forwarding raw data to the Hub's broadcast channel.
func (c *Client) read() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(appData string) error {
		log.Print("Yes, iam!")
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}

			break
		}
		c.hub.broadcast <- message
	}
}

// write implements the egress logic, flushing queued messages from the send channel to the network.
// It manages heartbeat synchronization via a Ticker to ensure connection persistence.
func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			log.Print("hey client are u still alive ?")
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
