/**
 * @Author:      leafney
 * @Date:        2022-10-09 11:18
 * @Project:     rose-notice
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package feishu

import "testing"

func TestNewRobot(t *testing.T) {
	bot := NewRobot("")
	bot.SetSecret("")
	err := bot.SendText("消息发送测试")
	t.Log(err)
}
