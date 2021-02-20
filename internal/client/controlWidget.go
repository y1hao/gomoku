package client

type ControlWidget struct {
	BoxWidget
}

func NewControlWidget(row, col, height, width int) *ControlWidget {
	return &ControlWidget{
		BoxWidget: BoxWidget{
			WidgetBase: WidgetBase{
				row: row,
				col: col,
				height: height,
				width: width,
			},
			title: "Control",
		},
	}
}

func (w *ControlWidget) Draw() {
	w.BoxWidget.Draw()

	pushPosition()
	defer popPosition()

	setPosition(w.row+1, w.col+1)
	print(highlightF, blackB, "Exit")
	print(infoF, blackB, ": ")
	printBold(greenF, blackB, "exit")

	setPosition(w.row+2, w.col+1)
	print(highlightF, blackB, "Move")
	print(infoF, blackB, ": <")
	print(highlightF, blackB, "c")
	print(infoF, blackB, "><")
	print(highlightF, blackB, "r")
	print(infoF, blackB, ">")

	setPosition(w.row+3, w.col+3)
	print(infoF, blackB, "eg: ")
	printBold(greenF, blackB, "h8")

	setPosition(w.row+4, w.col+1)
	print(highlightF, blackB, "Chat")
	print(infoF, blackB, ": ")
	printBold(greenF, blackB, "'")
	print(highlightF, blackB, "...")
	printBold(greenF, blackB, "'")

	setPosition(w.row+5, w.col+3)
	print(infoF, blackB, "eg: ")
	printBold(greenF, blackB, "'Yo!'")
}
