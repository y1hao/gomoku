package client

import (
	"fmt"
	"github.com/CoderYihaoWang/gomoku/internal/message"
	"math"
	"strings"
)

type ChatWidget struct {
	BoxWidget
	context *Context
}

func NewChatWidget(row, col, height, width int, context *Context) *ChatWidget {
	return &ChatWidget{
		BoxWidget: BoxWidget{
			WidgetBase: WidgetBase{
				row: row,
				col: col,
				height: height,
				width: width,
			},
			title: "Chat",
		},
		context: context,
	}
}

func (w *ChatWidget) Redraw() {
	pushPosition()
	defer popPosition()

	w.clearMessage()

	h := 0
	var chats []*message.ChatMessage
	for i := len(w.context.Chat)-1; i >=0 && h < w.height-2; i-- {
		h += w.getMessageHeight(w.context.Chat[i])
		chats = append(chats, w.context.Chat[i])
	}
	w.printChats(chats, h)
}

func (w *ChatWidget) printChats(chats []*message.ChatMessage, h int) {
	r := w.row+1+h
	if r > w.row+w.height-1 {
		r = w.row+w.height-1
	}
	for i := 0; i < len(chats); i++ {
		r -= w.getMessageHeight(chats[i])
		w.printMessage(chats[i], r)
	}
}

func (w *ChatWidget) printMessage(m *message.ChatMessage, begR int) {
	rows := w.getMessageRows(m)
	if begR < w.row+1 {
		n := len(rows)-(w.row+1-begR)
		rows = rows[len(rows)-n:]
		if len(rows[0]) < 3 {
			rows[0] = "..."
		} else {
			rows[0] = "..." + rows[0][3:]
		}
		begR = w.row+1
	}

	timeStamp := fmt.Sprintf(" [%02d:%02d:%02d] ", m.Time.Hour(), m.Time.Minute(), m.Time.Second())

	setPosition(begR, w.col+1)
	printBold(highlightF, blackB, timeStamp)

	var color color
	if m.Sender == w.context.Player {
		color = youF
	} else {
		color = oppF
	}

	for i := 0; i < len(rows); i++ {
		setPosition(begR+i, w.col+1+len(" [00:00:00] "))
		print(color, blackB, rows[i])
	}
}

func (w *ChatWidget) getMessageRows(chat *message.ChatMessage) []string {
	m := chat.Message
	var mRows []string
	start, size := 0, w.width-2-len(" [00:00:00] ")
	for start < len(m) {
		end := start + size
		if end > len(m) {
			end = len(m)
		}
		mRows = append(mRows, m[start:end])
		start += size
	}
	return mRows
}

func (w *ChatWidget) getMessageHeight(m *message.ChatMessage) int {
	return int(math.Ceil(float64(len(m.Message))/float64(w.width-2-len(" [00:00:00] "))))
}

func (w *ChatWidget) clearMessage() {
	empty := strings.Repeat(" ", w.width-2-len(" [00:00:00] "))
	for i := w.row+1; i < w.row+w.height-1; i++ {
		setPosition(i, w.col+1)
		print(infoF, blackB, empty)
	}
}