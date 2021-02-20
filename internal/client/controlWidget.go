package client

type Control struct {
	BoxWidget
}

func NewControl(row, col, height, width int, context *Context) *Control {
	return &Control{
		BoxWidget{
			WidgetBase{row, col, height, width, context},
		},
	}
}
