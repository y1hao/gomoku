package main

type WsServer struct {
	rooms       map[*Room]bool
	invitations map[int]*Room
	invite      chan *Client
	accept      chan *Client
	unregister  chan *Room
}

func NewWsServer() *WsServer {
	return &WsServer{
		rooms:       make(map[*Room]bool),
		invitations: make(map[int]*Room),
		invite:      make(chan *Client),
		accept:      make(chan *Client),
		unregister:  make(chan *Room),
	}
}

func (s *WsServer) Run() {
	for {
		select {
		case c := <-s.invite:
			s.inviteClient(c)

		case c := <-s.accept:
			s.acceptClient(c)

		case r := <-s.unregister:
			s.unregisterRoom(r)
		}
	}
}

func (s *WsServer) inviteClient(c *Client) {
	room := NewRoom()
	go room.Run()
	room.register <- c
	s.rooms[room] = true
	s.invitations[c.code] = room
}

func (s *WsServer) acceptClient(c *Client) {
	s.invitations[c.code].register <- c
	delete(s.invitations, c.code)
}

func (s *WsServer) unregisterRoom(room *Room) {
	if exist := s.rooms[room]; exist {
		delete(s.rooms, room)
	}
}
