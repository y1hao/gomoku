package client

import (
	"fmt"
	"github.com/CoderYihaoWang/gomoku/internal/game"
)

type Board struct {
	WidgetBase
}

func NewBoard(row, col, height, width int, context *Context) *Board {
	return &Board{
		WidgetBase{row, col, height, width, context},
	}
}

func (w *Board) Draw() {
	printStatus(w.context.Game)
}

func (w *Board) Redraw() {
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
