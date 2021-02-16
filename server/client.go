package main

import (
	"errors"
	"fmt"
	"github.com/CoderYihaoWang/gomoku/server/invitationCode"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"strings"
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
		code, err := client.invite()
		if err != nil {
			client.error(err)
			return
		}
		client.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Invitation code: %d", code)))
	} else {
		err := client.accept(params[1])
		if err != nil {
			client.error(err)
			return
		}
	}

	go client.write()
	go client.read()
}

func (c *Client) error(err error) {
	log.Println(err)
	c.conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
	c.disconnect()
}

func (c *Client) invite() (int, error) {
	code, err := invitationCode.Get()
	if err != nil {
		invitationCode.Return(code)
		return 0, err
	}
	c.server.invite <- Invitation{Code: code, Client: c}
	return code, nil
}

func (c *Client) accept(codeString string) error {
	code, err := strconv.Atoi(codeString)
	if err != nil || code < 0 || code > invitationCode.MaxId {
		return errors.New(fmt.Sprintf("invalid invitation code: %s", codeString))
	}
	if _, ok := c.server.invitations[code]; ok {
		c.server.accept <- Invitation{Code: code, Client: c}
		invitationCode.Return(code)
		return nil
	}
	return errors.New(fmt.Sprintf("the invitation code %d does not exist", code))
}

func (c *Client) disconnect() {
	c.conn.Close()
	room := c.room
	if room == nil {
		return
	}
	room.unregister <- c
	for client := range room.clients {
		client.conn.WriteMessage(websocket.TextMessage, []byte("The other has left"))
	}
	if len(room.clients) == 0 {
		c.server.unregister <- room
	}
}

func (c *Client) read() {
	defer func() {
		c.disconnect()
	}()

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

func (c *Client) write() {
	for {
		message, ok := <-c.send
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
	}
}

func (c *Client) handleMessage(m []byte) {
	if len(m) > 0 {
		c.room.broadcast <- m
	}
}
