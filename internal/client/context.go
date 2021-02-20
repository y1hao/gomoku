package client

import (
	"github.com/CoderYihaoWang/gomoku/internal/game"
)

type level int

const (
	none = iota
	info
	error
	win
	lose
)

type Context struct {
	Player         game.Player
	Score1, Score2 int
	Game           *game.Game
	History        []*game.Piece
	Chat           string
	Message string
	Level level
}

func NewContext() *Context {
	return &Context{
		History: make([]*game.Piece, 0),
		Game:    game.New(),
	}
}
