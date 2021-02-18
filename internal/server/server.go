package server

import (
	"sync"

	"github.com/CoderYihaoWang/gomoku/internal/game"
)

type Server struct {
	Invitations   map[int]*Room
	InvitationsMu sync.Mutex
	Invite        chan *Client
	Accept        chan *Client
	Rematch chan *Client
	rematchRequests map[*Client]bool
	rematchRequestsMu sync.Mutex
}

func New() *Server {
	return &Server{
		Invitations: make(map[int]*Room),
		Invite:      make(chan *Client),
		Accept:      make(chan *Client),
		Rematch: make(chan *Client),
		rematchRequests: make(map[*Client]bool),
	}
}

func (s *Server) Run() {
	for {
		select {
		case c := <-s.Invite:
			s.invite(c)

		case c := <-s.Accept:
			s.accept(c)

			case c := <-s.Rematch:
				s.rematch(c)
		}
	}
}

func (s *Server) invite(c *Client) {
	room := NewRoom()
	go room.Run()
	room.Register <- c

	s.InvitationsMu.Lock()
	defer s.InvitationsMu.Unlock()

	s.Invitations[c.Code] = room
}

func (s *Server) accept(c *Client) {
	s.InvitationsMu.Lock()
	defer s.InvitationsMu.Unlock()

	room := s.Invitations[c.Code]
	room.Register <- c
	delete(s.Invitations, c.Code)

	g := game.New()
	room.StartGame <- g
}

func (s *Server) rematch(c *Client) {
	s.rematchRequestsMu.Lock()
	defer s.rematchRequestsMu.Unlock()

	s.rematchRequests[c] = true
	if len(s.rematchRequests) == 2 {
		g := game.New()
		c.Room.StartGame <- g
	}
}