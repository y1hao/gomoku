package client

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/CoderYihaoWang/gomoku/internal/game"
	"github.com/CoderYihaoWang/gomoku/internal/message"
)

const ExitInput = "exit"

type inputHandler struct {
	Input   chan []byte
	Exit    chan struct{}
	Console *Console
	Context *Context
}

func NewInputHandler(input chan []byte, exit chan struct{}, console *Console, context *Context) *inputHandler {
	return &inputHandler{
		Input:   input,
		Exit:    exit,
		Console: console,
		Context: context,
	}
}

func (handler *inputHandler) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := strings.TrimSpace(string(scanner.Bytes()))
		if !handler.validate(data) {
			handler.Console.UpdateInfo("Invalid input!")
			continue
		}
		handler.Input <- handler.process(data)
	}
}

func (handler *inputHandler) validate(data string) bool {
	// while waiting for rematch, allow empty input for rematch
	if handler.Context.Game.Winner != game.None && len(data) == 0 {
		return true
	}

	// otherwise empty string is false
	if len(data) == 0 {
		return false
	}

	// 'exit'
	if data == ExitInput {
		return true
	}

	// chat
	if strings.HasPrefix(data, "'") && strings.HasSuffix(data, "'") {
		return true
	}

	// move e.g. h8
	if len(data) < 2 || len(data) > 3 {
		return false
	}
	col := data[0]
	if col < 'a' || col > 'o' {
		return false
	}
	row, err := strconv.Atoi(data[1:])
	if err != nil {
		return false
	}
	if row < 1 || row > 15 {
		return false
	}
	return true
}

func (handler *inputHandler) process(m string) (data []byte) {
	// rematch
	if handler.Context.Game.Winner != game.None && len(m) == 0 {
		data, _ = json.Marshal(message.NewNextGame())
		handler.Console.UpdateInfo("Waiting for opponent...")
		handler.Console.WaitForInput()
		return
	}

	switch {
	case m == ExitInput:
		close(handler.Exit)

	case strings.HasPrefix(m, "'") && strings.HasSuffix(m, "'"):
		data, _ = json.Marshal(message.NewChat(&message.ChatMessage{
			Sender:  handler.Context.Player,
			Message: m[1 : len(m)-1],
		}))

	default:
		l, d := m[0], m[1:]
		row, _ := strconv.Atoi(d)
		row--
		col := int(l - 'a')
		data, _ = json.Marshal(message.NewMove(&game.Piece{
			Row:    row,
			Col:    col,
			Player: handler.Context.Player,
		}))
	}
	return data
}
