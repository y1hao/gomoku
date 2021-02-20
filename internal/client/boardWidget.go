package client

import (
	"github.com/CoderYihaoWang/gomoku/internal/game"
)

type BoardWidget struct {
	WidgetBase
	context *Context
	lastMove *game.Piece
}

func NewBoardWidget(row, col, height, width int, context *Context) *BoardWidget {
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
	w.clearHighLight()
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
	w.lastMove = move
	r, c := w.getPiecePosition(move)

	setPosition(r, c-1)
	print(greenF, yellowB, "[")
	if move.Player == game.Black {
		print(blackF, yellowB, "⬤")
	} else {
		print(whiteF, yellowB, "⬤")
	}
	print(greenF, yellowB, "]")
}

func (w *BoardWidget) clearHighLight() {
	pushPosition()
	defer popPosition()

	if w.lastMove == nil {
		return
	}
	r, c := w.getPiecePosition(w.lastMove)

	setPosition(r, c-1)
	if w.lastMove.Col == 0 {
		printDim(infoF, yellowB, " ")
	} else {
		if w.lastMove.Row == 0 || w.lastMove.Row == 14  {
			printDim(infoF, yellowB, "═")
		} else {
			printDim(infoF, yellowB, "─")
		}
	}

	setPosition(r, c+1)
	if w.lastMove.Col == 14 {
		printDim(infoF, yellowB, " ")
	} else {
		if w.lastMove.Row == 0 || w.lastMove.Row == 14  {
			printDim(infoF, yellowB, "═")
		} else {
			printDim(infoF, yellowB, "─")
		}
	}
}

func (w *BoardWidget) getPiecePosition(p *game.Piece) (row, col int) {
	r, c := 14-p.Row, p.Col
	return w.row+r, w.col+3+2*c
}
