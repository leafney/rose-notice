/**
 * @Author:      leafney
 * @Date:        2022-10-09 11:18
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package feishu

import "testing"

func TestNewRobot(t *testing.T) {
	bot := NewFeiShu("")
	bot.UseSecret("")

	err := bot.
		SetDebug(true).
		SendText("消息发送测试\n测试")
	t.Log(err)
}

func TestPostElementBuilder(t *testing.T) {
	bot := NewFeiShu("")
	bot.UseSecret("")

	// 创建富文本消息构建器
	builder := NewPostElementBuilder()

	// 构建富文本消息
	elements := builder.
		AddText("这是第一行文本").AddLink("点击这里", "https://example.com").
		AddLine().
		AddText("这是第二行文本").
		Build()

	// 发送富文本消息
	err := bot.
		SetDebug(true).
		SendFullText("测试富文本消息", elements)
	t.Log(err)
}
