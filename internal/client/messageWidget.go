package client

import (
	"fmt"
)

type Message struct {
	WidgetBase
	message string
}

func NewMessage(row, col, height, width int, context *Context) *Message {
	return &Message{
		WidgetBase: WidgetBase{row, col, height, width, context},
	}
}

func (w *Message) Update(m string) {
	w.message = m
}

func (w *Message) Draw() {
	fmt.Printf("You are: %d\n", w.context.Player)
	if w.context.Game != nil {
		fmt.Printf("%d's turn\n", w.context.Game.Player)
	}
	fmt.Printf("message: %s\n", w.message)
}

func (w *Message) Redraw() {
	fmt.Printf("You are: %d\n", w.context.Player)
	if w.context.Game != nil {
		fmt.Printf("%d's turn\n", w.context.Game.Player)
	}
	fmt.Printf("message: %s\n", w.message)
}
