package client

import (
	"fmt"
	"strings"

	"github.com/CoderYihaoWang/gomoku/internal/game"
)

type MessageWidget struct {
	WidgetBase
	context *Context
}

func NewMessage(row, col, height, width int, context *Context) *MessageWidget {
	return &MessageWidget{
		WidgetBase: WidgetBase{
			row:    row,
			col:    col,
			height: height,
			width:  width,
		},
		context: context,
	}
}

func (w *MessageWidget) Draw() {
	pushPosition()
	defer popPosition()

	setPosition(w.row, w.col)
	print(infoF, mainB, strings.Repeat(" ", w.width))
}

func (w *MessageWidget) Redraw() {
	pushPosition()
	defer popPosition()

	setPosition(w.row, w.col+1)
	print(infoF, mainB, "You play ")
	if w.context.Player == game.White {
		print(whiteF, mainB, "⬤")
	} else if w.context.Player == game.Black {
		print(blackF, mainB, "⬤")
	} else {
		print(highlightF, mainB, "?")
	}
	print(infoF, mainB, " ")

	setPosition(w.row, w.col+w.width-len("O's turn "))
	if w.context.Game.Player == game.White {
		print(whiteF, mainB, "⬤")
	} else if w.context.Game.Player == game.Black {
		print(blackF, mainB, "⬤")
	}
	print(infoF, mainB, "'s turn")

	setPosition(w.row, w.col+len(" You play O "))
	var fg, bg color
	switch w.context.Level {
	case none:
		fg, bg = infoF, mainB

	case info:
		fg, bg = infoF, secondaryB

	case error:
		fg, bg = infoF, errorB

	case win:
		fg, bg = infoF, winB

	case lose:
		fg, bg = infoF, loseB
	}

	print(fg, bg,
		fmt.Sprintf(fmt.Sprintf(" %%-%ds", w.width-1-len(" You play O ")-len(" O's turn ")),
			w.context.Message))
}
