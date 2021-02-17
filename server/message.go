package main

import "github.com/CoderYihaoWang/gomoku/server/game"

type Message struct {
	Type  string     `json:"type"`
	Info  string     `json:"info,omitempty"`
	Error string     `json:"error,omitempty"`
	Move  *game.Piece `json:"move,omitempty"`
	Game  *game.Game  `json:"game,omitempty"`
}
