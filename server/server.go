package main

type WsServer struct {
	invitations map[int]*Room
	invite      chan *Client
	accept      chan *Client
}

func NewWsServer() *WsServer {
	return &WsServer{
		invitations: make(map[int]*Room),
		invite:      make(chan *Client),
		accept:      make(chan *Client),
	}
}

func (s *WsServer) Run() {
	for {
		select {
		case c := <-s.invite:
			s.inviteClient(c)

		case c := <-s.accept:
			s.acceptClient(c)
		}
	}
}

func (s *WsServer) inviteClient(c *Client) {
	room := NewRoom()
	go room.Run()
	room.register <- c
	s.invitations[c.code] = room
}

func (s *WsServer) acceptClient(c *Client) {
	s.invitations[c.code].register <- c
	delete(s.invitations, c.code)
}
