package client

import (
	"fmt"
)

const (
	titleH   = 3
	boardH   = 16
	chatH    = 7
	infoH    = 1
	scoreH   = 4
	historyH = 5
	controlH = 7

	titleR   = 0
	boardR   = titleR + titleH
	chatR    = boardR + boardH
	infoR    = chatR + chatH
	scoreR   = boardR
	historyR = scoreR + scoreH
	controlR = historyR + historyH

	boardC = 0
	sideC  = 35
	endC   = 51
)

type Console struct {
	title   *Title
	board   *Board
	message *Message
	score   *Score
	history *History
	control *Control
	chat    *Chat
}

func NewConsole(context *Context) *Console {
	return &Console{
		title:   NewTitle(titleR, 0, 0, 0, context),
		board:   NewBoard(boardR, 0, 0, 0, context),
		message: NewMessage(0, 0, 0, 0, context),
		score:   NewScore(0, 0, 0, 0, context),
		history: NewHistory(0, 0, 0, 0, context),
		control: NewControl(0, 0, 0, 0, context),
		chat:    NewChat(0, 0, 0, 0, context),
	}
}

func (c *Console) Clear() {
	clear()
}

func (c *Console) DrawAll() {
	c.Clear()
	c.title.Draw()
	c.board.Draw()
	c.message.Draw()
	c.score.Draw()
	c.history.Draw()
	c.control.Draw()
	c.chat.Draw()
}

func (c *Console) UpdateGame() {
	c.board.Redraw()
	c.message.Redraw()
}

func (c *Console) UpdateScore() {
	c.score.Redraw()
}

func (c *Console) UpdateHistory() {
	c.history.Redraw()
}

func (c *Console) UpdateChat() {
	c.chat.Redraw()
}

func (c *Console) DisplayMessage(m string) {
	c.message.Update(m)
	c.message.Redraw()
}

func (c *Console) DisplayError(m string) {
	c.message.Update(fmt.Sprint("err:", m))
	c.message.Redraw()
}
