package client

type Title struct {
	WidgetBase
}

func NewTitle(row, col, height, width int, context *Context) *Title {
	return &Title{WidgetBase{row, col, height, width, context}}
}

func (w *Title) Draw()   {}
func (w *Title) Redraw() {}
