package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/CoderYihaoWang/gomoku/internal/client"
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

var context = client.NewContext()

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws/"}
	if *code >= 0 {
		u.Path = fmt.Sprintf("/ws/%d", *code)
	}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	inboundMessage := make(chan *message.Message)
	messageHandler := client.NewMessageHandler(inboundMessage, context)
	go messageHandler.Run()

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
			inboundMessage <-&m
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
			Sender:  context.AssignedPlayer,
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
			Player: context.AssignedPlayer,
		}))
	case "rematch":
		data, _ = json.Marshal(message.NewNextGame())
	}
	return data
}