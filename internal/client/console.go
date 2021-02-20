package client

import "fmt"

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
	title   *titleWidget
	board   *boardWidget
	message *messageWidget
	score   *scoreWidget
	history *historyWidget
	control *controlWidget
	chat    *chatWidget
}

func NewConsole(context *Context) *Console {
	return &Console{
		title:   newTitle(titleR, 0, 0, 0, context),
		board:   newBoard(boardR, 0, 0, 0, context),
		message: newMessage(0, 0, 0, 0, context),
		score:   newScore(0, 0, 0, 0, context),
		history: newHistory(0, 0, 0, 0, context),
		control: newControl(0, 0, 0, 0, context),
		chat:    newChat(0, 0, 0, 0, context),
	}
}

func (c *Console) Clear() {
	clear()
}

func (c *Console) DrawAll() {
	c.Clear()
	c.title.draw()
	c.board.draw()
	c.message.draw()
	c.score.draw()
	c.history.draw()
	c.control.draw()
	c.chat.draw()
}

func (c *Console) UpdateGame() {
	c.board.redraw()
	c.message.redraw()
}

func (c *Console) UpdateScore() {
	c.score.redraw()
}

func (c *Console) UpdateHistory() {
	c.history.redraw()
}

func (c *Console) UpdateChat() {
	c.chat.redraw()
}

func (c *Console) DisplayMessage(m string) {
	c.message.update(m)
	c.message.redraw()
}

func (c *Console) DisplayError(m string) {
	c.message.update(fmt.Sprint("err:", m))
	c.message.redraw()
}
