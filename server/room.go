package main

type Room struct {
	clients map[*Client]bool
	register chan *Client
	unregister chan *Client
	broadcast chan []byte
}

func NewRoom() *Room {
	return &Room{
		clients: make(map[*Client]bool),
		register: make(chan *Client),
		unregister: make(chan *Client),
		broadcast: make(chan []byte),
	}
}

func (r *Room) Run() {
	for {
		select {
		case c := <-r.register:
			r.RegisterClient(c)

		case c := <-r.unregister:
			r.UnregisterClient(c)

		case m := <-r.broadcast:
			r.BroadCast(m)
		}
	}
}

func (r *Room) RegisterClient(c *Client) {
	r.clients[c] = true
}

func (r *Room) UnregisterClient(c *Client) {
	if _, ok := r.clients[c]; ok {
		delete(r.clients, c)
	}
}

func (r *Room) BroadCast(m []byte) {
	for c := range r.clients {
		c.send <- m
	}
}

