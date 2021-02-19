package client

import (
	"fmt"
)

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
		board: newBoard(0,0,0,0,context),
		message: newMessage(0,0,0,0,context),
		score: newScore(),
		history: newHistory(),
		control: newControl(),
		chat: newChat(),
	}
}

func (c *Console) DrawAll() {
	c.title.draw()
	c.board.draw()
	c.message.draw()
	c.score.draw()
	c.history.draw()
	c.control.draw()
	c.chat.draw()
}

func (c *Console) UpdateBoard() {
	c.board.redraw()
}

func (c *Console) UpdateScore() {}

func (c *Console) UpdateHistory() {}

func (c *Console) UpdateChat() {}

func (c *Console) UpdateInfo() {
	c.message.redraw()
}

func (c *Console) DisplayMessage(m string) {
	fmt.Println(m)
}

func (c *Console) DisplayError(m string) {
	fmt.Println(m)
}

func (c *Console) DisplayFatal(m string) {
	clear()
	fmt.Println(m)
}
