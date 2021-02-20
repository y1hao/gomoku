package client

import (
	"fmt"
)

const (
	titleH   = 3
	boardH   = 16
	chatH    = 7
	messageH = 1
	scoreH   = 4
	historyH = 5
	controlH = 7

	titleR   = 1
	boardR   = titleR + titleH
	chatR    = boardR + boardH
	messageR = chatR + chatH
	scoreR   = boardR
	historyR = scoreR + scoreH
	controlR = historyR + historyH

	titleW   = 50
	boardW   = 35
	scoreW   = titleW - boardW
	historyW = scoreW
	controlW = scoreW
	chatW    = titleW
	messageW = titleW

	titleC   = 1
	boardC   = titleC
	messageC = titleC
	chatC    = titleC
	scoreC   = titleC + boardW
	historyC = scoreC
	controlC = scoreC
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
		title:   NewTitle(titleR, titleC, titleH, titleW, context),
		board:   NewBoard(boardR, boardC, boardH, boardW, context),
		message: NewMessage(messageR, messageC, messageH, messageW, context),
		score:   NewScore(scoreR, scoreC, scoreH, scoreW, context),
		history: NewHistory(historyR, historyC, historyH, historyW, context),
		control: NewControl(controlR, controlC, controlH, controlW, context),
		chat:    NewChat(chatR, chatC, chatH, chatW, context),
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
