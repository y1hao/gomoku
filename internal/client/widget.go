package client

type WidgetBase struct {
	row, col, height, width int
}

type BoxWidget struct {
	WidgetBase
	title string
}

func (w *BoxWidget) Draw() {}
