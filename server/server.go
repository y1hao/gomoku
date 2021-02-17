package main

type Server struct {
	Invitations map[int]*Room
	Invite      chan *Client
	Accept      chan *Client
}

func NewServer() *Server {
	return &Server{
		Invitations: make(map[int]*Room),
		Invite:      make(chan *Client),
		Accept:      make(chan *Client),
	}
}

func (s *Server) Run() {
	for {
		select {
		case c := <-s.Invite:
			s.invite(c)

		case c := <-s.Accept:
			s.accept(c)
		}
	}
}

func (s *Server) invite(c *Client) {
	room := NewRoom()
	go room.Run()
	room.Register <- c
	s.Invitations[c.code] = room
}

func (s *Server) accept(c *Client) {
	s.Invitations[c.code].Register <- c
	delete(s.Invitations, c.code)
}
