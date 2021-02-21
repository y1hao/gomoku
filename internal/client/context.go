package client

import (
	"github.com/CoderYihaoWang/gomoku/internal/game"
	"github.com/CoderYihaoWang/gomoku/internal/message"
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
	Chat           []*message.ChatMessage
	Message        string
	Level          level
}

func NewContext() *Context {
	return &Context{
		History: make([]*game.Piece, 0),
		Chat:    make([]*message.ChatMessage, 0),
		Game:    game.New(),
	}
}
