package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Client struct {
	hub       *Hub
	conn      *websocket.Conn
	userID    string
	send      chan SocketEvent
	writeKill chan bool
}

func (c *Client) read() {
	var socketEvent SocketEvent

	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	setSocketPayloadReadConfig(c)

	for {
		fmt.Println("reading message.....")
		_, payload, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure, websocket.CloseNoStatusReceived) {
				log.Println("read error: ", err)
			}
			c.writeKill <- true
			break
		}

		decoder := json.NewDecoder(bytes.NewReader(payload))
		decodeErr := decoder.Decode(&socketEvent)

		if decodeErr != nil {
			log.Println("decode error: ", err)
			c.writeKill <- true
			break
		}

		// handle socket event
		handleSocketEvent(c, socketEvent)
	}
}

func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case payload, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				return
			}

			// encode data
			buffer := new(bytes.Buffer)
			json.NewEncoder(buffer).Encode(payload)
			finallPayload := buffer.Bytes()

			writer, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println("create writer error: ", err)
				return
			}
			writer.Write(finallPayload)

			n := len(c.send)
			for i := 0; i < n; i++ {
				json.NewEncoder(buffer).Encode(<-c.send)
				writer.Write(buffer.Bytes())
			}

			if err := writer.Close(); err != nil {
				log.Println("close writer err: ", err)
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}

		case <-c.writeKill:
			return
		}
	}

}

func setSocketPayloadReadConfig(c *Client) {
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
}
