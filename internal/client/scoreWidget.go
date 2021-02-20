package client

type ScoreWidget struct {
	BoxWidget
	context *Context
}

func NewScoreWidget(row, col, height, width int, context *Context) *ScoreWidget {
	return &ScoreWidget{
		BoxWidget: BoxWidget{
			WidgetBase: WidgetBase{
				row: row,
				col: col,
				height: height,
				width: width,
			},
			title: "ScoreWidget",
		},
		context: context,
	}
}

func (w *ScoreWidget) Redraw() {

}