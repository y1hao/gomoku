package client

import (
	"fmt"
)

const (
	info = iota
	error
	win
	lose
)

type level int

type MessageWidget struct {
	WidgetBase
	context *Context
	level level
}

func NewMessage(row, col, height, width int, context *Context) *MessageWidget {
	return &MessageWidget{
		WidgetBase: WidgetBase{
			row: row,
			col: col,
			height: height,
			width: width,
		},
		context: context,
	}
}

func (w *MessageWidget) Draw() {
	fmt.Printf("You are: %d\n", w.context.Player)
	if w.context.Game != nil {
		fmt.Printf("%d's turn\n", w.context.Game.Player)
	}
	fmt.Printf("message: %s\n", w.context.Message)
}

func (w *MessageWidget) Redraw() {
	fmt.Printf("You are: %d\n", w.context.Player)
	if w.context.Game != nil {
		fmt.Printf("%d's turn\n", w.context.Game.Player)
	}
	fmt.Printf("message: %s\n", w.context.Message)
}
