package client

type History struct {
	BoxWidget
}

func NewHistory(row, col, height, width int, context *Context) *History {
	return &History{
		BoxWidget{
			WidgetBase{row, col, height, width, context},
		},
	}
}
