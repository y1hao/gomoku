package client

import (
	"github.com/CoderYihaoWang/gomoku/internal/game"
)

type Context struct {
	Player         game.Player
	Score1, Score2 int
	Game           *game.Game
	History        []*game.Piece
	Chat           string
}

func NewContext() *Context {
	return &Context{}
}