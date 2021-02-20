package client

import (
	"fmt"
	"github.com/CoderYihaoWang/gomoku/internal/game"
)

type BoardWidget struct {
	WidgetBase
	context *Context
}

func NewBoard(row, col, height, width int, context *Context) *BoardWidget {
	return &BoardWidget{
		WidgetBase: WidgetBase{
			row: row,
			col: col,
			height: height,
			width: width,
		},
		context: context,
	}
}

func (w *BoardWidget) Draw() {
	printStatus(w.context.Game)
}

func (w *BoardWidget) Redraw() {
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
