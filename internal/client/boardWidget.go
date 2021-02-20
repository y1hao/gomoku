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
	pushPosition()
	defer popPosition()

	w.printEmptyBoard()
	//printStatus(w.context.Game)
}

func (w *BoardWidget) Redraw() {
	pushPosition()
	defer popPosition()

	w.printEmptyBoard()
	//printStatus(w.context.Game)
}

func (w *BoardWidget) printEmptyBoard() {
	setPosition(w.row, w.col)
	printDim(infoF, yellowB, "15 ╔═╤═╤═╤═╤═╤═╤═╤═╤═╤═╤═╤═╤═╤═╗ ")
	for i := 1; i < w.height-1; i++ {
		setPosition(w.row+i, w.col)
		printDimf(infoF, yellowB, "%2d", 15-i)
		printDim(infoF, yellowB, " ╟─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─╢ ")

	}
	setPosition(w.row+w.height-2, w.col)
	printDim(infoF, yellowB, " 1 ╚═╧═╧═╧═╧═╧═╧═╧═╧═╧═╧═╧═╧═╧═╝ ")

	setPosition(w.row+w.height-1, w.col)
	printDim(infoF, yellowB, "   a b c d e f g h i j k l m n o ")
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
