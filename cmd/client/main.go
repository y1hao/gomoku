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

	context := client.NewContext()

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws/"}
	if *code >= 0 {
		u.Path = fmt.Sprintf("/ws/%d", *code)
	}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	console := client.NewConsole(context)
	console.DrawAll()

	messages := make(chan *message.Message)
	fatal := make(chan []byte)
	messageHandler := client.NewMessageHandler(messages, fatal, console, context)
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
			messages <- &m
		}
	}()

	handleExit := func() {
		err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("Connection was lost")
			return
		}
		select {
		case <-done:
		case <-time.After(time.Second):
		}
		console.Clear()
		return
	}

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			handleExit()
			return
		case <-exit:
			handleExit()
			return
		case m := <-fatal:
			handleExit()
			log.Printf("GOMOKU has exited due to %s\nPlease try again!", m)
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
