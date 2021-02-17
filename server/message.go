package main

import "github.com/CoderYihaoWang/gomoku/server/game"

type Message struct {
	Info          string       `json:"info,omitempty"`
	Error         string       `json:"error,omitempty"`
	Player        game.Player `json:"player,omitempty"`
	Board         game.Board  `json:"board,omitempty"`
	WinningPieces []game.Piece `json:"winningPieces,omitempty"`
}

type Move struct {
	Row    int         `json:"row"`
	Col    int         `json:"col"`
	Player game.Player `json:"player"`
}
