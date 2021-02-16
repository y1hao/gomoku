package main

type Room struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewRoom() *Room {
	return &Room{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (r *Room) Run() {
	for {
		select {
		case c := <-r.register:
			r.registerClient(c)

		case c := <-r.unregister:
			r.unregisterClient(c)

		case m := <-r.broadcast:
			r.broadCast(m)
		}
	}
}

func (r *Room) registerClient(c *Client) {
	r.clients[c] = true
	c.room = r
}

func (r *Room) unregisterClient(c *Client) {
	if _, ok := r.clients[c]; ok {
		delete(r.clients, c)
		c.room = nil
	}
}

func (r *Room) broadCast(m []byte) {
	for c := range r.clients {
		c.send <- m
	}
}
