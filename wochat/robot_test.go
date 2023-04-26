package wochat

import "testing"

func TestNewRobot(t *testing.T) {
	bot := NewRobot("")

	//err := bot.
	//	DebugMode().
	//	SendText("企业微信消息测试发送")

	err := bot.
		DebugMode().
		SendMarkdown(`## 放假啦

	- <font color="info">粗去玩</font>
	- <font color="comment">次好次哒</font>
	- <font color="warning">拍皂片</font>`)

	t.Log(err)
}
