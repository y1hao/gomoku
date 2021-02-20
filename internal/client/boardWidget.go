package client

import (
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
	w.printEmptyBoard()
}

func (w *BoardWidget) Redraw() {
	w.printLastMove()
}

func (w *BoardWidget) printEmptyBoard() {
	pushPosition()
	defer popPosition()

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

func (w *BoardWidget) printLastMove() {
	pushPosition()
	defer popPosition()

	if w.context.Game.LastMove == nil {
		return
	}
	move := w.context.Game.LastMove
	r, c := 14-move.Row, move.Col

	setPosition(w.row+r, w.col+3+c*2)
	if move.Player == game.Black {
		print(blackF, yellowB, "⬤")
	} else {
		print(whiteF, yellowB, "⬤")
	}
}
