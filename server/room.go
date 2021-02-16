package main

type Room struct {
	clients map[*Client]bool
	register chan *Client
	unregister chan *Client
	broadcast chan []byte
}

