package client

import (
	"strings"
)

type TitleWidget struct {
	WidgetBase
}

func NewTitleWidget(row, col, height, width int) *TitleWidget {
	return &TitleWidget{
		WidgetBase{
			row:    row,
			col:    col,
			height: height,
			width:  width,
		},
	}
}

func (w *TitleWidget) Draw() {
	pushPosition()
	defer popPosition()

	defer resetColor()

	setPosition(w.row, w.col)
	printf(infoF, mainB, "%s%s%s", "╔", strings.Repeat("═", w.width-2), "╗")

	spaces := w.width - 2 - len("G O M O K U")
	leftSpaces, rightSpaces := spaces/2, spaces-spaces/2

	setPosition(w.row+1, w.col)
	print(infoF, mainB, "║")
	print(infoF, mainB, strings.Repeat(" ", leftSpaces))
	print(infoF, mainB, "G ")
	print(whiteF, mainB, "⬤")
	print(infoF, mainB, " M ")
	print(blackF, mainB, "⬤")
	print(infoF, mainB, " K U")
	print(infoF, mainB, strings.Repeat(" ", rightSpaces))
	print(infoF, mainB, "║")

	setPosition(w.row+2, w.col)
	printf(infoF, mainB, "%s%s%s", "╚", strings.Repeat("═", w.width-2), "╝")
}
