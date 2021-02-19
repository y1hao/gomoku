package client

const (
	titleH = 3
	boardH = 16
	chatH = 7
	infoH = 1
	scoreH = 4
	historyH = 5
	controlH = 7

	titleR = 0
	boardR = titleR + titleH
	chatR = boardR + boardH
	infoR = chatR + chatH
	scoreR = boardR
	historyR = scoreR + scoreH
	controlR = historyR + historyH

	boardC = 0
	sideC = 35
	endC = 51
)

type Console struct {
	title widget
	board widget
	message widget
	score widget
	history widget
	control widget
	chat widget
}

func NewConsole(context *Context) *Console {
	return &Console{
		title: newTitle(0,0,0,0, context),
		board: newBoard(),
		message: newMessage(),
		score: newScore(),
		history: newHistory(),
		control: newControl(),
		chat: newChat(),
	}
}

func (c *Console) DisplayStatus() {}

func (c *Console) DisplayMessage(m string) {}


