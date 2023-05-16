package xizhi

import "testing"

func TestNewXiZhi(t *testing.T) {
	bot := NewXiZhi("")

	bot.UseToken("")

	err := bot.SetDebug(true).
		SendMarkdown("测试markdown", "- 第一 \n - 第二 \n - 第三")
	//SendMsg("让我看看", "看个得儿")
	//SendText("你好")
	t.Log(err)
}
