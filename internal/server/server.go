package server

import (
	"sync"
)

type Server struct {
	Invitations   map[int]*Room
	InvitationsMu sync.Mutex
	Invite        chan *Client
	Accept        chan *Client
}

func New() *Server {
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
	r := NewRoom()
	go r.Run()
	r.Register <- c

	s.InvitationsMu.Lock()
	defer s.InvitationsMu.Unlock()

	s.Invitations[c.Code] = r
}

func (s *Server) accept(c *Client) {
	s.InvitationsMu.Lock()
	defer s.InvitationsMu.Unlock()

	r := s.Invitations[c.Code]
	r.Register <- c
	delete(s.Invitations, c.Code)

	r.StartGame <- struct{}{}
}
