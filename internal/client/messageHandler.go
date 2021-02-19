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
		m := m.ChatMessage
		fmt.Printf("[%v] %d: %s\n", m.Time, m.Sender, m.Message)

	case message.Status:
		printStatus(m.Status)
		//handler.Console.DisplayStatus()

	case message.OpponentLeft:
		fmt.Printf("Opponent left\n")

	case message.InvitationCode:
		fmt.Printf("Your invitation code is: %s\n", m.Info)

	case message.InsufficientInvitationCode:
		fmt.Printf("Insufficient invitation code\n")

	case message.InvalidInvitationCode:
		fmt.Printf("Invalid invitation code: %s\n", m.Info)

	case message.InvalidMove:
		fmt.Printf("Invalid move\n")

	case message.AssignPlayer:
		p, _ := strconv.Atoi(m.Info)
		handler.Context.AssignedPlayer = game.Player(p)
		fmt.Printf("You are: %d\n", p)

	case message.InvalidMessageFormat:
		fmt.Printf("Invalid message format\n")

	case message.InvalidOperation:
		fmt.Printf("Invalid operation\n")
	}
}

func printStatus(status *game.Game) {
	if len(status.WinningPieces) != 0 {
		fmt.Printf("%d wins!\n", status.WinningPieces[0].Player)
		return
	}

	for i := range status.Board {
		for j := range status.Board[i] {
			fmt.Printf("%v ", status.Board[i][j])
		}
		fmt.Println()
	}

	if status.LastMove != nil {
		fmt.Printf("Last move: %d: [%d, %d]\n",
			status.LastMove.Player,
			status.LastMove.Row,
			status.LastMove.Col)
	}

	fmt.Printf("%d's turn\n", status.Player)
}