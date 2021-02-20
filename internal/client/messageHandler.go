package client

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/CoderYihaoWang/gomoku/internal/game"
	"github.com/CoderYihaoWang/gomoku/internal/message"
)

type MessageHandler struct {
	Message chan *message.Message
	Fatal   chan []byte
	Console *Console
	Context *Context
}

func NewMessageHandler(message chan *message.Message, fatal chan []byte, console *Console, context *Context) *MessageHandler {
	return &MessageHandler{
		Message: message,
		Fatal:   fatal,
		Console: console,
		Context: context,
	}
}

func (handler *MessageHandler) Run() {
	for {
		m, ok := <-handler.Message
		if !ok {
			return
		}
		handler.handleMessage(m)
	}
}

func (handler *MessageHandler) handleMessage(m *message.Message) {
	switch m.Type {
	case message.Chat:
		handler.Console.ClearMessage()

		handler.Context.Chat = append(handler.Context.Chat, m.ChatMessage)
		handler.Console.UpdateChat()

	case message.Status:
		handler.Console.ClearMessage()

		handler.Context.Game = m.Status
		handler.Console.UpdateGame()

		if m.Status.LastMove != nil {
			handler.Context.History = append(handler.Context.History, m.Status.LastMove)
		}
		handler.Console.UpdateHistory()

		switch m.Status.Winner {
		case game.None:
		case game.Draw:
			handler.Context.Score1++
			handler.Context.Score2++
			handler.Console.UpdateInfo("DRAW... (Rematch: <Enter>)")

		case handler.Context.Player:
			handler.Context.Score1++
			handler.Console.UpdateWin("WIN !!! (Rematch: <Enter>)")

		default:
			handler.Context.Score2++
			handler.Console.UpdateLose("LOSE... (Rematch: <Enter>)")
		}
		handler.Console.UpdateScore()

	case message.InvitationCode:
		code, _ := strconv.Atoi(strings.TrimSpace(m.Info))
		handler.Console.UpdateInfo(fmt.Sprintf("Your invitation code: %04d", code))

	case message.InsufficientInvitationCode:
		handler.Fatal <- []byte("insufficient invitation code")

	case message.InvalidInvitationCode:
		handler.Fatal <- []byte("invalid invitation code")

	case message.AssignPlayer:
		handler.Console.ClearMessage()

		p, _ := strconv.Atoi(m.Info)
		handler.Context.Player = game.Player(p)
		handler.Console.NewGame()

	case message.OpponentLeft:
		handler.Console.UpdateError("Your opponent has left!")

	case message.InvalidMove:
		handler.Console.UpdateError("Invalid move")

	case message.InvalidOperation:
		handler.Console.UpdateError("Invalid operation")

	case message.InvalidMessageFormat:
		handler.Console.UpdateError("Invalid message format")
	}
}
