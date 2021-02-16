package main

type Invitation struct {
	Code   int
	Client *Client
}

type WsServer struct {
	rooms       map[*Room]bool
	invitations map[int]*Room
	invite      chan Invitation
	accept      chan Invitation
	unregister  chan *Room
}

func NewWsServer() *WsServer {
	return &WsServer{
		rooms:       make(map[*Room]bool),
		invitations: make(map[int]*Room),
		invite:      make(chan Invitation),
		accept:      make(chan Invitation),
		unregister:  make(chan *Room),
	}
}

func (s *WsServer) Run() {
	for {
		select {
		case inv := <-s.invite:
			s.inviteClient(inv.Client, inv.Code)

		case inv := <-s.accept:
			s.acceptClient(inv.Client, inv.Code)

		case room := <-s.unregister:
			s.unregisterRoom(room)
		}
	}
}

func (s *WsServer) inviteClient(client *Client, code int) {
	room := NewRoom()
	go room.Run()
	room.register <- client
	s.rooms[room] = true
	s.invitations[code] = room
}

func (s *WsServer) acceptClient(client *Client, code int) {
	if room, ok := s.invitations[code]; ok {
		room.register <- client
		delete(s.invitations, code)
	}
}

func (s *WsServer) unregisterRoom(room *Room) {
	if exist := s.rooms[room]; exist {
		delete(s.rooms, room)
	}
}
