package telegram

import "testing"

func TestNewTelegram(t *testing.T) {
	//bot := NewTelegram("")
	bot := NewTelegramChatId("", "")
	bot.UseToken("")

	err := bot.
		SetDebug(true).
		//SetGet(true).
		SetChatId("").
		SetDisableWebPreview(true).
		//SetSilently(true).
		//SendText("测试消息接收")
		SendMarkdown("*哈哈123* \n[Bing](https://www.bing.com/)")
	t.Log(err)
}
