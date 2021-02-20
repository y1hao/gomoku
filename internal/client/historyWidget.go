package client

type HistoryWidget struct {
	BoxWidget
	context *Context
}

func NewHistory(row, col, height, width int, context *Context) *HistoryWidget {
	return &HistoryWidget{
		BoxWidget: BoxWidget{
			WidgetBase: WidgetBase{
				row: row,
				col: col,
				height: height,
				width: width,
			},
			title: "HistoryWidget",
		},
		context: context,
	}
}

func (w *HistoryWidget) Redraw() {

}
