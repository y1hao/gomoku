package message

import "github.com/CoderYihaoWang/gomoku/internal/game"

type Message struct {
	Type  string      `json:"type"`
	Info  string      `json:"info,omitempty"`
	Move  *game.Piece `json:"move,omitempty"`
	Game  *game.Game  `json:"game,omitempty"`
}

func NewError(error string) *Message {
	return &Message{
		Type: "error",
		Info: error,
	}
}

func NewInfo(info string) *Message {
	return &Message{
		Type: "info",
		Info: info,
	}
}

func NewGame(game *game.Game) *Message {
	return &Message{
		Type: "game",
		Game: game,
	}
}

func NewMove(move *game.Piece) *Message {
	return &Message{
		Type: "move",
		Move: move,
	}
}

func NewChat(chat string) *Message {
	return &Message{
		Type: "chat",
		Info: chat,
	}
}

func NewLeave() *Message {
	return &Message{
		Type: "leave",
	}
}