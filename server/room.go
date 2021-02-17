package main

type Room struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
}

func NewRoom() *Room {
	return &Room{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
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
		}
	}
}

func (r *Room) register(c *Client) {
	r.Clients[c] = true
	c.room = r
}

func (r *Room) unregister(c *Client) {
	if _, ok := r.Clients[c]; ok {
		delete(r.Clients, c)
		c.room = nil
	}
}

func (r *Room) broadcast(m []byte) {
	for c := range r.Clients {
		c.send <- m
	}
}
