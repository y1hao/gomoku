package client

type ControlWidget struct {
	BoxWidget
}

func NewControlWidget(row, col, height, width int) *ControlWidget {
	return &ControlWidget{
		BoxWidget: BoxWidget{
			WidgetBase: WidgetBase{
				row: row,
				col: col,
				height: height,
				width: width,
			},
			title: "Control",
		},
	}
}
