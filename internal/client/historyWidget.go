package client

import (
	"fmt"
	"github.com/CoderYihaoWang/gomoku/internal/game"
)

type HistoryWidget struct {
	BoxWidget
	context *Context
}

func NewHistoryWidget(row, col, height, width int, context *Context) *HistoryWidget {
	return &HistoryWidget{
		BoxWidget: BoxWidget{
			WidgetBase: WidgetBase{
				row: row,
				col: col,
				height: height,
				width: width,
			},
			title: "History",
		},
		context: context,
	}
}

func (w *HistoryWidget) Redraw() {
	pushPosition()
	defer popPosition()

	var history []*game.Piece
	count := 0
	for i := len(w.context.History)-1; i >= 0 && count < 3; i-- {
		history = append(history, w.context.History[i])
		count++
	}

	index := len(w.context.History)-1
	i := 1
	for _, move := range history {
		setPosition(w.row+i, w.col+1)
		print(infoF, blackB, fmt.Sprintf("%4d: ", index))
		index--
		i++
		w.drawMove(move)
	}
}

func (w *HistoryWidget) drawMove(p *game.Piece) {
	if p.Player == game.Black {
		print(blackF, yellowB, " ⬤ ")
	} else {
		print(whiteF, yellowB, " ⬤ ")
	}
	print(infoF, blackB, " ")
	if p.Player == w.context.Player {
		print(youF, blackB, w.getPieceCode(p))
	} else {
		print(oppF, blackB, w.getPieceCode(p))
	}
}

func (w HistoryWidget) getPieceCode(p *game.Piece) string {
	str := fmt.Sprintf("%c%d", p.Col+'a', p.Row+1)
	return fmt.Sprintf("%-3s", str)
}
