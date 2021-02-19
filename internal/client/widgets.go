package client

type widget interface {
	draw()
	redraw()
}

type widgetBase struct {
	row, col, height, width int
	context *Context
}

type titleWidget struct {
	widgetBase
}

func newTitle(row, col, height, width int, context *Context) *titleWidget {
	return &titleWidget{widgetBase{row, col, height, width, context}}
}

func (w *titleWidget) draw() {}
func (w *titleWidget) redraw() {}

type boardWidget struct {}
func newBoard() *boardWidget {return nil}
func (w *boardWidget) draw() {}
func (w *boardWidget) redraw() {}

type messageWidget struct {}
func newMessage() *messageWidget {return nil}
func (w *messageWidget) draw() {}
func (w *messageWidget) redraw() {}

type boxWidget struct {}
func newBox() *boxWidget {return nil}
func (w *boxWidget) draw() {}
func (w *boxWidget) redraw() {}

type scoreWidget struct {
	boxWidget
}
func newScore() *scoreWidget {return nil}

type historyWidget struct {
	*boxWidget
}
func newHistory() *historyWidget {return nil}

type controlWidget struct {
	*boxWidget
}
func newControl() *controlWidget {return nil}

type chatWidget struct {
	*boxWidget
}
func newChat() *chatWidget {return nil}