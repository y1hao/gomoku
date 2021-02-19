package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"

	"github.com/CoderYihaoWang/gomoku/internal/client"
	"github.com/CoderYihaoWang/gomoku/internal/message"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var code = flag.Int("code", -1, "invitation code")

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

	context := client.NewContext()
	console := client.NewConsole(context)

	messages := make(chan *message.Message)
	messageHandler := client.NewMessageHandler(messages, console, context)
	go messageHandler.Run()

	input := make(chan []byte)
	exit := make(chan struct{})
	inputHandler := client.NewInputHandler(input, exit, console, context)
	go inputHandler.Run()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, data, err := c.ReadMessage()
			if err != nil {
				return
			}
			var m message.Message
			err = json.Unmarshal(data, &m)
			if err != nil {
				continue
			}
			messages <-&m
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			handleExit(c, done)
			return
		case <-exit:
			handleExit(c, done)
			return
		case m := <-input:
			err := c.WriteMessage(websocket.TextMessage, m)
			if err != nil {
				log.Println("Connection was lost")
				return
			}
		}
	}
}

func handleExit(c *websocket.Conn, done chan struct{}) {
	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("Connection was lost")
		return
	}
	select {
	case <-done:
	case <-time.After(time.Second):
	}
	return
}