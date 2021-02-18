package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/CoderYihaoWang/gomoku/internal/game"
	"github.com/CoderYihaoWang/gomoku/internal/message"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var code = flag.Int("code", -1, "invitation code")

var player game.Player

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws/"}
	if *code >= 0 {
		u.Path = fmt.Sprintf("/ws/%d", *code)
	}

	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, data, err := c.ReadMessage()
			if err != nil {
				// closed
				return
			}
			var m message.Message
			err = json.Unmarshal(data, &m)
			if err != nil {
				continue
			}
			handleMessage(&m)
		}
	}()

	input := make(chan []byte)
	go func() {
		defer close(input)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data := strings.TrimSpace(string(scanner.Bytes()))
			if len(data) == 0 {
				continue
			}
			input <- formatMessage(data)
		}
	}()

	for {
		select {
		case <-done:
			return
		case m := <-input:
			err := c.WriteMessage(websocket.TextMessage, m)
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func formatMessage(m string) []byte {
	fields := strings.Split(m, " ")
	if len(fields) < 1 {
		return nil
	}
	var data []byte
	switch fields[0] {
	case "chat":
		data, _ = json.Marshal(message.NewChat(&message.ChatMessage{
			Sender:  player,
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
			Player: player,
		}))
	case "rematch":
		data, _ = json.Marshal(message.NewNextGame())
	}
	return data
}

func handleMessage(m *message.Message) {
	switch m.Type {
	case message.Chat:
		m := m.ChatMessage
		fmt.Printf("[%v] %d: %s\n", m.Time, m.Sender, m.Message)

	case message.Status:
		printStatus(m.Status)

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
		player = game.Player(p)
		fmt.Printf("You are: %d\n", player)

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
