package client

import (
	"fmt"
	"strings"
)

type WidgetBase struct {
	row, col, height, width int
}

type BoxWidget struct {
	WidgetBase
	title string
}

func (w *BoxWidget) Draw() {
	pushPosition()
	defer popPosition()

	w.drawTitle()
	w.drawBorder()
}

func (w *BoxWidget) drawTitle() {
	setPosition(w.row, w.col+2)
	print(infoF, secondaryB, w.formatTitle())
}

func (w *BoxWidget) formatTitle() string {
	totalLen := w.width-4
	leftSpaces := (totalLen-len(w.title))/2
	rightSpaces := totalLen-len(w.title)-leftSpaces
	return fmt.Sprintf("%s%s%s",
		strings.Repeat(" ", leftSpaces),
		w.title,
		strings.Repeat(" ", rightSpaces))
}

func (w *BoxWidget) drawBorder() {
	setPosition(w.row, w.col)
	print(mainF, blackB, "┏━")

	setPosition(w.row, w.col+w.width-2)
	print(mainF, blackB, "━┓")

	for i := 0; i < w.height-2; i++ {
		setPosition(w.row+1+i, w.col)
		print(mainF, blackB, "┃")

		setPosition(w.row+1+i, w.col+w.width-1)
		print(mainF, blackB, "┃")
	}

	setPosition(w.row+w.height-1, w.col)
	print(mainF, blackB, "┗")
	for i := 0; i < w.width-2; i++ {
		print(mainF, blackB, "━")
	}
	print(mainF, blackB, "┛")
}