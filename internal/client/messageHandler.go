package client

import (
	"fmt"
	"strconv"

	"github.com/CoderYihaoWang/gomoku/internal/game"
	"github.com/CoderYihaoWang/gomoku/internal/message"
)

type MessageHandler struct {
	Message chan *message.Message
	Console *Console
	Context *Context
}

func NewMessageHandler(message chan *message.Message, console *Console, context *Context) *MessageHandler {
	return &MessageHandler{
		Message: message,
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
		chat := m.ChatMessage
		handler.Console.UpdateChat()
		fmt.Printf("[%v] %d: %s\n", chat.Time, chat.Sender, chat.Message)

	case message.Status:
		handler.Context.Game = m.Status
		handler.Console.UpdateBoard()
		handler.Console.UpdateInfo()

		if m.Status.LastMove != nil {
			handler.Context.History = append(handler.Context.History, m.Status.LastMove)
		}
		handler.Console.UpdateHistory()

		switch m.Status.Winner {
		case game.None:
		case game.Draw:

		case handler.Context.Player:
			handler.Context.Score1++
			handler.Console.DisplayMessage("WIN !!! (Rematch: <Enter>)")

		default:
			handler.Context.Score2++
			handler.Console.DisplayError("LOSE... (Rematch: <Enter>)")
		}
		handler.Console.UpdateScore()

	case message.InvitationCode:
		handler.Console.DisplayMessage(fmt.Sprintf("Your invitation code: %s", m.Info))

	case message.InsufficientInvitationCode:
		handler.Console.DisplayFatal("Sorry, too many players online at the moment.\nPlease try again later!")

	case message.AssignPlayer:
		p, _ := strconv.Atoi(m.Info)
		handler.Context.Player = game.Player(p)
		handler.Console.UpdateInfo()

	case message.OpponentLeft:
		handler.Console.DisplayError("Your opponent has left!")

	case message.InvalidInvitationCode:
		handler.Console.DisplayError("Invalid invitation code: %s")

	case message.InvalidMove:
		handler.Console.DisplayError("Invalid move")

	case message.InvalidMessageFormat:
		handler.Console.DisplayError("Invalid message format")

	case message.InvalidOperation:
		handler.Console.DisplayError("Invalid operation")
	}
}