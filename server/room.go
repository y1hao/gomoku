package main

import (
	"encoding/json"
	"github.com/CoderYihaoWang/gomoku/server/game"
)

type Room struct {
	Clients    map[*Client]bool
	Game *game.Game
	Register   chan *Client
	Unregister chan *Client
	StartGame chan *game.Game
	Broadcast  chan *Message

}

func NewRoom() *Room {
	return &Room{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		StartGame: make(chan *game.Game),
		Broadcast:  make(chan *Message),
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

		case g := <-r.StartGame:
			r.startGame(g)
		}
	}
}

func (r *Room) register(c *Client) {
	r.Clients[c] = true
	c.Room = r
}

func (r *Room) unregister(c *Client) {
	if _, ok := r.Clients[c]; ok {
		delete(r.Clients, c)
		c.Room = nil
	}
}

func (r *Room) startGame(g *game.Game) {
	r.Game = g
}

func (r *Room) broadcast(m *Message) {
	data, err := json.Marshal(m)
	if err != nil {
		return
	}
	for c := range r.Clients {
		c.Send <- data
	}
}
