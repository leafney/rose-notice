/**
 * @Author:      leafney
 * @Date:        2022-10-12 12:29
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package chanify

import "testing"

func TestNewRobot(t *testing.T) {
	bot := NewRobot("", "")
	//bot.SetHost("")  // reset host
	//bot.SetToken("") // reset token

	err := bot.
		//SetGet().
		//SetCopy("三拼霸霸奶茶").
		//SetAutoCopy(true).
		//SetLevel(1).
		//SetSound("shake").
		//SetDefSound().
		//SendText("下午茶时间")
		SendMsg("下午茶喝什么", "蜜雪冰城")
		//SendLink("https://www.baidu.com")
		//SendLink("https://www.baidu.com/img/flexible/logo/pc/index@2.png")
		//SendLink("weixin://")
	if err != nil {
		t.Log(err)
	}
}
