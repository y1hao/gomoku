package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/CoderYihaoWang/gomoku/server/invitationCode"
	"github.com/gorilla/websocket"
)

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second

	// Max time till next pong from peer
	pongWait = 60 * time.Second

	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

type Client struct {
	conn   *websocket.Conn
	server *WsServer
	room   *Room
	send   chan []byte
}

func newClient(conn *websocket.Conn, server *WsServer) *Client {
	return &Client{
		conn:   conn,
		server: server,
		send:   make(chan []byte),
	}
}

func ServeWs(s *WsServer, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	params := strings.Split(r.URL.Path, "/")
	client := newClient(conn, s)

	if len(params) < 2 || params[1] == "" {
		code, err := invitationCode.Get()
		if err != nil {
			log.Println(err)
			return
		}
		s.invite <- Invitation{Code: code, Client: client}
		client.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Invitation code: %d", code)))
	} else {
		code, err := strconv.Atoi(params[1])
		if err != nil || code < 0 || code > invitationCode.MaxId {
			log.Println("Invalid invitation code")
			return
		}
		s.accept <- Invitation{Code: code, Client: client}
		invitationCode.Return(code)
	}

	go client.writePump()
	go client.readPump()
}

func (c *Client) readPump() {
	defer func() {
		c.disconnect()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	// Start endless read loop, waiting for messages from client
	for {
		_, m, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}

		c.handleMessage(m)
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The WsServer closed the channel.
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
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) disconnect() {
	room := c.room
	room.unregister <- c
	for client := range room.clients {
		client.conn.WriteMessage(websocket.TextMessage, []byte("The other has left"))
	}
	if len(room.clients) == 0 {
		c.server.unregister <- room
	}
	c.conn.Close()
}

func (c *Client) handleMessage(m []byte) {
	if len(m) > 0 {
		c.room.broadcast <- m
	}
}
