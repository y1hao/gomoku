package client

import "fmt"

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
			title: "Score",
		},
		context: context,
	}
}

func (w *ScoreWidget) Draw() {
	w.BoxWidget.Draw()

	pushPosition()
	defer popPosition()

	setPosition(w.row+1, w.col+2)
	print(infoF, blackB, "You:")
	print(youF, blackB, fmt.Sprintf("% 6d", w.context.Score1))

	setPosition(w.row+2, w.col+2)
	print(infoF, blackB, "Opp:")
	print(oppF, blackB, fmt.Sprintf("% 6d", w.context.Score2))
}

func (w *ScoreWidget) Redraw() {
	setPosition(w.row+1, w.col+2+len("You:"))
	print(youF, blackB, fmt.Sprintf("% 6d", w.context.Score1))

	setPosition(w.row+2, w.col+2+len("Opp:"))
	print(oppF, blackB, fmt.Sprintf("% 6d", w.context.Score2))
}