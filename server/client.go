package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/CoderYihaoWang/gomoku/server/game"
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
	conn         *websocket.Conn
	disconnected bool
	Player       game.Player
	Server       *Server
	Room         *Room
	Code         int
	Send         chan []byte
}

func newClient(conn *websocket.Conn, server *Server) *Client {
	return &Client{
		conn:   conn,
		Server: server,
		Send:   make(chan []byte),
	}
}

func Serve(s *Server, w http.ResponseWriter, r *http.Request) {
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
		m, err := json.Marshal(fmt.Sprintf("Your invitation Code is: %d", code))
		if err != nil {
			fmt.Println(err)
			return
		}
		client.conn.WriteMessage(websocket.TextMessage, m)
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
	defer c.disconnect()

	log.Println(err)
	m, err := json.Marshal(Message{Error: err.Error()})
	if err != nil {
		log.Println(err)
		return
	}
	c.conn.WriteMessage(websocket.TextMessage, m)
}

func (c *Client) invite() (int, error) {
	code, err := invitationCode.Get()
	if err != nil {
		invitationCode.Return(code)
		return 0, err
	}
	c.Code = code
	c.Server.Invite <- c
	return code, nil
}

func (c *Client) accept(codeString string) error {
	code, err := strconv.Atoi(codeString)
	if err != nil || code < 0 || code > invitationCode.MaxId {
		return errors.New(fmt.Sprintf("invalid invitation Code: %s", codeString))
	}
	if _, ok := c.Server.Invitations[code]; ok {
		c.Code = code
		c.Server.Accept <- c
		invitationCode.Return(code)
		return nil
	}
	return errors.New(fmt.Sprintf("the invitation Code %d does not exist", code))
}

func (c *Client) disconnect() {
	if c.disconnected {
		return
	}
	c.disconnected = true
	c.conn.Close()
	room := c.Room
	if room == nil {
		return
	}
	room.Unregister <- c
	for client := range room.Clients {
		client.conn.WriteMessage(websocket.TextMessage, []byte("The other has left"))
	}
	if _, ok := c.Server.Invitations[c.Code]; ok {
		delete(c.Server.Invitations, c.Code)
		invitationCode.Return(c.Code)
	}
}

func (c *Client) read() {
	defer c.disconnect()

	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		c.handleMessage(data)
	}
}

func (c *Client) write() {
	defer c.disconnect()

	for {
		message, ok := <-c.Send
		if !ok {
			// The Server closed the channel.
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

func (c *Client) handleMessage(data []byte) {
	m := &Message{}
	err := json.Unmarshal(data, m)
	if err != nil {
		return
	}

	switch m.Type {
	case "chat":
		c.handleChatMessage(m)
	case "move":
		c.handleMoveMessage(m)
	}
}

func (c *Client) handleChatMessage(m *Message) {
	c.Room.Broadcast <-m
}

func (c *Client) handleMoveMessage(m *Message) {
	c.Room.Broadcast <- &Message{
		Type: "move",
		Game: c.Room.Game,
	}
}