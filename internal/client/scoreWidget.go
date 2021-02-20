package client

type Score struct {
	BoxWidget
}

func NewScore(row, col, height, width int, context *Context) *Score {
	return &Score{
		BoxWidget{
			WidgetBase{row, col, height, width, context},
		},
	}
}
