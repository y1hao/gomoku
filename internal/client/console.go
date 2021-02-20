package client

import "fmt"

const (
	titleH   = 3
	boardH   = 16
	chatH    = 5
	messageH = 1
	scoreH   = 4
	historyH = 5
	controlH = 7

	titleR   = 1
	boardR   = titleR + titleH + 1
	chatR    = boardR + boardH + 1
	messageR = chatR + chatH + 1
	scoreR   = boardR
	historyR = scoreR + scoreH
	controlR = historyR + historyH
	inputR   = messageR + messageH + 1

	titleW   = 50
	boardW   = 35
	scoreW   = titleW - boardW
	historyW = scoreW
	controlW = scoreW
	chatW    = titleW
	messageW = titleW

	titleC   = 5
	boardC   = titleC
	messageC = titleC
	chatC    = titleC
	scoreC   = titleC + boardW
	historyC = scoreC
	controlC = scoreC
	inputC   = titleC
)

type Console struct {
	title   Widget
	board   Widget
	message Widget
	score   Widget
	history Widget
	control Widget
	chat    Widget
	context *Context
}

func NewConsole(context *Context) *Console {
	return &Console{
		title:   NewTitleWidget(titleR, titleC, titleH, titleW),
		board:   NewBoardWidget(boardR, boardC, boardH, boardW, context),
		message: NewMessage(messageR, messageC, messageH, messageW, context),
		score:   NewScoreWidget(scoreR, scoreC, scoreH, scoreW, context),
		history: NewHistoryWidget(historyR, historyC, historyH, historyW, context),
		control: NewControlWidget(controlR, controlC, controlH, controlW),
		chat:    NewChatWidget(chatR, chatC, chatH, chatW, context),
		context: context,
	}
}

func (c *Console) Clear() {
	clear()
}

func (c *Console) DrawAll() {
	clear()
	c.title.Draw()
	c.board.Draw()
	c.message.Draw()
	c.score.Draw()
	c.history.Draw()
	c.control.Draw()
	c.chat.Draw()
	c.WaitForInput()
}

func (c *Console) WaitForInput() {
	setPosition(inputR, inputC)
	fmt.Print("> ")
	clearToEnd()
}

func (c *Console) NewGame() {
	c.board.Draw()
	c.message.Redraw()
	c.WaitForInput()
}

func (c *Console) UpdateGame() {
	c.board.Redraw()
	c.message.Redraw()
	c.WaitForInput()
}

func (c *Console) UpdateScore() {
	c.score.Redraw()
	c.WaitForInput()
}

func (c *Console) UpdateHistory() {
	c.history.Redraw()
	c.WaitForInput()
}

func (c *Console) UpdateChat() {
	c.chat.Redraw()
	c.WaitForInput()
}

func (c *Console) ClearMessage() {
	c.context.Level = none
	c.context.Message = ""
	c.message.Redraw()
	c.WaitForInput()
}

func (c *Console) UpdateInfo(m string) {
	c.context.Level = info
	c.context.Message = m
	c.message.Redraw()
	c.WaitForInput()
}

func (c *Console) UpdateError(m string) {
	c.context.Level = error
	c.context.Message = m
	c.message.Redraw()
	c.WaitForInput()
}

func (c *Console) UpdateWin(m string) {
	c.context.Level = win
	c.context.Message = m
	c.message.Redraw()
	c.WaitForInput()
}

func (c *Console) UpdateLose(m string) {
	c.context.Level = lose
	c.context.Message = m
	c.message.Redraw()
	c.WaitForInput()
}
