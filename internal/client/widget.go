package client

type Widget interface {
	Draw()
	Redraw()
}

type WidgetBase struct {
	row, col, height, width int
	context                 *Context
}

type BoxWidget struct {
	WidgetBase
}

func (w *BoxWidget) Draw()   {}
func (w *BoxWidget) Redraw() {}
