package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/CoderYihaoWang/gomoku/internal/game"
	"github.com/CoderYihaoWang/gomoku/internal/invitationCode"
	"github.com/CoderYihaoWang/gomoku/internal/message"
	"github.com/gorilla/websocket"
)

const (
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 2048
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
}

type Client struct {
	Conn         *websocket.Conn
	disconnected bool
	Player       game.Player
	Server       *Server
	Room         *Room
	Code         int
	Send         chan []byte
}

func newClient(conn *websocket.Conn, server *Server) *Client {
	return &Client{
		Conn:   conn,
		Server: server,
		Send:   make(chan []byte),
	}
}

func ServeWs(s *Server, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	path := r.URL.Path[len("/ws"):]
	params := strings.Split(path, "/")
	client := newClient(conn, s)

	if len(params) < 2 || params[1] == "" {
		code, m := client.invite()
		if m != nil {
			client.error(m)
			return
		}
		data, _ := json.Marshal(message.NewInvitationCode(strconv.Itoa(code)))
		client.Conn.WriteMessage(websocket.TextMessage, data)
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
	c.Conn.WriteMessage(websocket.TextMessage, data)
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

	c.Conn.Close()

	room := c.Room
	if room == nil {
		return
	}

	room.Unregister <- c

	room.ClientsMu.Lock()
	defer room.ClientsMu.Unlock()

	for client := range room.Clients {
		m, _ := json.Marshal(message.NewOpponentLeft())
		client.Conn.WriteMessage(websocket.TextMessage, m)
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

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		c.handleMessage(data)
	}
}

func (c *Client) write() {
	defer c.disconnect()

	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case m, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(m)

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) handleMessage(data []byte) {
	if len(data) == 0 {
		return
	}
	m := &message.Message{}
	err := json.Unmarshal(data, m)
	if err != nil {
		m, _ := json.Marshal(message.NewInvalidMessageFormat())
		c.Conn.WriteMessage(websocket.TextMessage, m)
		return
	}

	switch m.Type {
	case message.Chat:
		c.handleChatMessage(m)

	case message.Move:
		c.handleMoveMessage(m)

	case message.NextGame:
		c.handleNextGame()
	}
}

func (c *Client) handleChatMessage(m *message.Message) {
	c.Room.Broadcast <- m
}

func (c *Client) handleMoveMessage(m *message.Message) {
	move := m.Move
	status := c.Room.Game
	err := status.Move(move)
	if err != nil {
		data, _ := json.Marshal(message.NewInvalidMove())
		c.Conn.WriteMessage(websocket.TextMessage, data)
		return
	}
	c.Room.Broadcast <- message.NewStatus(status)
}

func (c *Client) handleNextGame() {
	if c.Room.Game.Winner == game.None {
		data, _ := json.Marshal(message.NewInvalidOperation())
		c.Conn.WriteMessage(websocket.TextMessage, data)
		return
	}
	c.Room.Rematch <- c
}
