package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/CoderYihaoWang/gomoku/internal/game"
	"github.com/CoderYihaoWang/gomoku/internal/invitationCode"
	"github.com/CoderYihaoWang/gomoku/internal/message"
	"github.com/gorilla/websocket"
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

	prefix := "/ws"
	path := r.URL.Path[len(prefix):]
	params := strings.Split(path, "/")
	client := newClient(conn, s)

	if len(params) < 2 || params[1] == "" {
		code, m := client.invite()
		if m != nil {
			client.error(m)
			return
		}
		data, _ := json.Marshal(message.NewInvitationCode(strconv.Itoa(code)))
		client.conn.WriteMessage(websocket.TextMessage, data)
	} else {
		m := client.accept(params[1])
		if m != nil {
			client.error(m)
			return
		}
	}

	go client.write()
	go client.read()
}

func (c *Client) error(m *message.Message) {
	data, _ := json.Marshal(m)
	c.conn.WriteMessage(websocket.TextMessage, data)
}

func (c *Client) invite() (int, *message.Message) {
	code, err := invitationCode.Get()
	if err != nil {
		return 0, message.NewInsufficientInvitationCode()
	}
	c.Code = code
	c.Server.Invite <- c
	return code, nil
}

func (c *Client) accept(codeString string) *message.Message {
	code, err := strconv.Atoi(codeString)
	if err != nil || code < 0 || code > invitationCode.MaxId {
		return message.NewInvalidInvitationCode(codeString)
	}

	c.Server.InvitationsMu.Lock()
	defer c.Server.InvitationsMu.Unlock()

	_, ok := c.Server.Invitations[code]
	if !ok {
		return message.NewInvalidInvitationCode(codeString)
	}

	c.Code = code
	c.Server.Accept <- c
	invitationCode.Return(code)
	return nil
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

	room.ClientsMu.Lock()
	defer room.ClientsMu.Unlock()

	for client := range room.Clients {
		m, _ := json.Marshal(message.NewOpponentLeft())
		client.conn.WriteMessage(websocket.TextMessage, m)
	}

	c.Server.InvitationsMu.Lock()
	defer c.Server.InvitationsMu.Unlock()

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
	m := &message.Message{}
	err := json.Unmarshal(data, m)
	if err != nil {
		return
	}

	switch m.Type {
	case message.Chat:
		c.handleChatMessage(m)
	case message.Move:
		c.handleMoveMessage(m)
	case message.Close:
		c.handleCloseMessage(m)
	}
}

func (c *Client) handleChatMessage(m *message.Message) {
	c.Room.Broadcast <- m
}

func (c *Client) handleMoveMessage(m *message.Message) {
	c.Room.Broadcast <- message.NewStatus(game.New())
}

func (c *Client) handleCloseMessage(_ *message.Message) {
	c.disconnect()
}
