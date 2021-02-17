package message

import (
	"github.com/CoderYihaoWang/gomoku/internal/game"
)

const (
	Chat                       = "chat"
	Move                       = "move"
	Status                     = "status"
	Close                      = "close"
	OpponentLeft               = "opponent left"
	InvitationCode             = "invitation code"
	InsufficientInvitationCode = "insufficient invitation code"
	InvalidInvitationCode      = "invalid invitation code"
)

type Message struct {
	Type   string      `json:"type"`
	Info   string      `json:"info,omitempty"`
	Move   *game.Piece `json:"move,omitempty"`
	Status *game.Game  `json:"status,omitempty"`
}

func NewChat(chat string) *Message {
	return &Message{
		Type: Chat,
		Info: chat,
	}
}

func NewMove(move *game.Piece) *Message {
	return &Message{
		Type: Move,
		Move: move,
	}
}

func NewStatus(status *game.Game) *Message {
	return &Message{
		Type:   Status,
		Status: status,
	}
}

func NewClose() *Message {
	return &Message{
		Type: Close,
	}
}

func NewOpponentLeft() *Message {
	return &Message{
		Type: OpponentLeft,
	}
}

func NewInvitationCode(code string) *Message {
	return &Message{
		Type: InvitationCode,
		Info: code,
	}
}

func NewInsufficientInvitationCode() *Message {
	return &Message{
		Type: InsufficientInvitationCode,
	}
}

func NewInvalidInvitationCode(code string) *Message {
	return &Message{
		Type: InvalidInvitationCode,
		Info: code,
	}
}
