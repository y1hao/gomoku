package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"sync"

	"github.com/CoderYihaoWang/gomoku/internal/game"
	"github.com/CoderYihaoWang/gomoku/internal/message"
)

type Room struct {
	Clients           map[*Client]bool
	ClientsMu         sync.Mutex
	Game              *game.Game
	Register          chan *Client
	Unregister        chan *Client
	StartGame         chan struct{}
	Broadcast         chan *message.Message
	Rematch           chan *Client
	rematchRequests   map[*Client]bool
	rematchRequestsMu sync.Mutex
}

func NewRoom() *Room {
	return &Room{
		Clients:         make(map[*Client]bool),
		Register:        make(chan *Client),
		Unregister:      make(chan *Client),
		StartGame:       make(chan struct{}),
		Broadcast:       make(chan *message.Message),
		Rematch:         make(chan *Client),
		rematchRequests: make(map[*Client]bool),
	}
}

func (r *Room) Run() {
	for {
		select {
		case c := <-r.Register:
			r.register(c)

		case c := <-r.Unregister:
			r.unregister(c)

		case m := <-r.Broadcast:
			r.broadcast(m)

		case <-r.StartGame:
			r.startGame()

		case c := <-r.Rematch:
			r.registerRematch(c)
		}
	}
}

func (r *Room) register(c *Client) {
	r.ClientsMu.Lock()
	defer r.ClientsMu.Unlock()

	r.Clients[c] = true
	c.Room = r
}

func (r *Room) unregister(c *Client) {
	r.ClientsMu.Lock()
	defer r.ClientsMu.Unlock()

	if _, ok := r.Clients[c]; ok {
		delete(r.Clients, c)
		c.Room = nil
	}
}

func (r *Room) startGame() {
	r.ClientsMu.Lock()
	defer r.ClientsMu.Unlock()

	r.Game = game.New()

	clients := make([]*Client, 0, 2)
	for c := range r.Clients {
		clients = append(clients, c)
	}

	clients[0].Player = game.Black
	m, _ := json.Marshal(message.NewAssignPlayer(game.Black))
	clients[0].Conn.WriteMessage(websocket.TextMessage, m)

	clients[1].Player = game.White
	m, _ = json.Marshal(message.NewAssignPlayer(game.White))
	clients[1].Conn.WriteMessage(websocket.TextMessage, m)

	m, _ = json.Marshal(message.NewStatus(r.Game))
	clients[0].Conn.WriteMessage(websocket.TextMessage, m)
	clients[1].Conn.WriteMessage(websocket.TextMessage, m)
}

func (r *Room) registerRematch(c *Client) {
	r.rematchRequestsMu.Lock()
	defer r.rematchRequestsMu.Unlock()

	r.rematchRequests[c] = true
	if len(r.rematchRequests) == 2 {
		r.startGame()
	}
}

func (r *Room) broadcast(m *message.Message) {
	r.ClientsMu.Lock()
	defer r.ClientsMu.Unlock()

	data, err := json.Marshal(m)
	if err != nil {
		return
	}
	for c := range r.Clients {
		c.Send <- data
	}
}
