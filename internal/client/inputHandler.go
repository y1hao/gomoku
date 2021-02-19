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

type inputHandler struct {
	Input chan []byte
	Exit chan struct{}
	Console *Console
	Context *Context
}

func NewInputHandler(input chan []byte, exit chan struct{}, console *Console, context *Context) *inputHandler {
	return &inputHandler{
		Input: input,
		Exit: exit,
		Console: console,
		Context: context,
	}
}

func (handler *inputHandler) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := strings.TrimSpace(string(scanner.Bytes()))
		if len(data) == 0 {
			continue
		}
		handler.Input <- handler.formatMessage(data)
	}
}

func (handler *inputHandler) formatMessage(m string) []byte {
	fields := strings.Split(m, " ")
	if len(fields) < 1 {
		return nil
	}
	var data []byte
	switch fields[0] {
	case "exit":
		close(handler.Exit)
	case "chat":
		data, _ = json.Marshal(message.NewChat(&message.ChatMessage{
			Sender:  handler.Context.AssignedPlayer,
			Message: strings.Join(fields[1:], " "),
		}))
	case "move":
		if len(fields) < 3 {
			break
		}
		row, err := strconv.Atoi(fields[1])
		if err != nil {
			break
		}
		col, err := strconv.Atoi(fields[2])
		if err != nil {
			break
		}
		data, _ = json.Marshal(message.NewMove(&game.Piece{
			Row:    row,
			Col:    col,
			Player: handler.Context.AssignedPlayer,
		}))
	case "rematch":
		data, _ = json.Marshal(message.NewNextGame())
	}
	return data
}