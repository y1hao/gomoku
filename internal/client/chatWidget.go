package client

type ChatWidget struct {
	BoxWidget
	context *Context
}

func NewChatWidget(row, col, height, width int, context *Context) *ChatWidget {
	return &ChatWidget{
		BoxWidget: BoxWidget{
			WidgetBase: WidgetBase{
				row: row,
				col: col,
				height: height,
				width: width,
			},
			title: "ChatWidget",
		},
		context: context,
	}
}

func (w *ChatWidget) Redraw() {

}
