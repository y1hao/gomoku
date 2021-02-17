package main

type Server struct {
	invitations map[int]*Room
	invite      chan *Client
	accept      chan *Client
}

func NewServer() *Server {
	return &Server{
		invitations: make(map[int]*Room),
		invite:      make(chan *Client),
		accept:      make(chan *Client),
	}
}

func (s *Server) Run() {
	for {
		select {
		case c := <-s.invite:
			s.sendInvitation(c)

		case c := <-s.accept:
			s.acceptInvitation(c)
		}
	}
}

func (s *Server) sendInvitation(c *Client) {
	room := NewRoom()
	go room.Run()
	room.register <- c
	s.invitations[c.code] = room
}

func (s *Server) acceptInvitation(c *Client) {
	s.invitations[c.code].register <- c
	delete(s.invitations, c.code)
}
