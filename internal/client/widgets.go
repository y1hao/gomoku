package client

import (
	"fmt"

	"github.com/CoderYihaoWang/gomoku/internal/game"
)

type widget interface {
	draw()
	redraw()
}

type widgetBase struct {
	row, col, height, width int
	context                 *Context
}

type titleWidget struct {
	widgetBase
}

func newTitle(row, col, height, width int, context *Context) *titleWidget {
	return &titleWidget{widgetBase{row, col, height, width, context}}
}

func (w *titleWidget) draw()   {}
func (w *titleWidget) redraw() {}

type boardWidget struct {
	widgetBase
}

func newBoard(row, col, height, width int, context *Context) *boardWidget {
	return &boardWidget{
		widgetBase{row, col, height, width, context},
	}
}

func (w *boardWidget) draw() {
	printStatus(w.context.Game)
}

func (w *boardWidget) redraw() {
	printStatus(w.context.Game)
}

func printStatus(status *game.Game) {
	if status == nil {
		return
	}

	if status.Winner != game.None {
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
}

type messageWidget struct {
	widgetBase
	message string
}

func newMessage(row, col, height, width int, context *Context) *messageWidget {
	return &messageWidget{
		widgetBase: widgetBase{row, col, height, width, context},
	}
}

func (w *messageWidget) update(m string) {
	w.message = m
}

func (w *messageWidget) draw() {
	fmt.Printf("You are: %d\n", w.context.Player)
	if w.context.Game != nil {
		fmt.Printf("%d's turn\n", w.context.Game.Player)
	}
	fmt.Printf("message: %s\n", w.message)
}

func (w *messageWidget) redraw() {
	fmt.Printf("You are: %d\n", w.context.Player)
	if w.context.Game != nil {
		fmt.Printf("%d's turn\n", w.context.Game.Player)
	}
	fmt.Printf("message: %s\n", w.message)
}

type boxWidget struct {
	widgetBase
}

func (w *boxWidget) draw()   {}
func (w *boxWidget) redraw() {}

type scoreWidget struct {
	boxWidget
}

func newScore(row, col, height, width int, context *Context) *scoreWidget {
	return &scoreWidget{
		boxWidget{
			widgetBase{row, col, height, width, context},
		},
	}
}

type historyWidget struct {
	boxWidget
}

func newHistory(row, col, height, width int, context *Context) *historyWidget {
	return &historyWidget{
		boxWidget{
			widgetBase{row, col, height, width, context},
		},
	}
}

type controlWidget struct {
	boxWidget
}

func newControl(row, col, height, width int, context *Context) *controlWidget {
	return &controlWidget{
		boxWidget{
			widgetBase{row, col, height, width, context},
		},
	}
}

type chatWidget struct {
	boxWidget
}

func newChat(row, col, height, width int, context *Context) *chatWidget {
	return &chatWidget{
		boxWidget{
			widgetBase{row, col, height, width, context},
		},
	}
}
