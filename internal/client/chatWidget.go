package client

type Chat struct {
	BoxWidget
}

func NewChat(row, col, height, width int, context *Context) *Chat {
	return &Chat{
		BoxWidget{
			WidgetBase{row, col, height, width, context},
		},
	}
}
