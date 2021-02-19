package client

import "github.com/CoderYihaoWang/gomoku/internal/game"

type Context struct {
	AssignedPlayer game.Player
	CurrentPlayer  game.Player
	Status         *game.Game
	History        []*game.Piece
	Chat           string
	Score1, Score2 int
}

func NewContext() *Context {
	return &Context{}
}