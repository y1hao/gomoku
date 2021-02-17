package message

import (
	"github.com/CoderYihaoWang/gomoku/internal/game"
	"strconv"
)

const (
	Chat                       = "chat"
	Move                       = "move"
	Status                     = "status"
	Leave                      = "leave"
	OpponentLeft               = "opponent left"
	InvitationCode             = "invitation code"
	InsufficientInvitationCode = "insufficient invitation code"
	InvalidInvitationCode      = "invalid invitation code"
	InvalidMove                = "invalid move"
	AssignPlayer               = "assign player"
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

func NewLeave() *Message {
	return &Message{
		Type: Leave,
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

func NewInvalidMove() *Message {
	return &Message{
		Type: InvalidMove,
	}
}

func NewAssignPlayer(p game.Player) *Message {
	return &Message{
		Type: AssignPlayer,
		Info: strconv.Itoa(int(p)),
	}
}
