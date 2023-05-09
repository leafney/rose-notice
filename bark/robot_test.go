/**
 * @Author:      leafney
 * @Date:        2022-10-09 17:55
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package bark

import "testing"

func TestNewRobot(t *testing.T) {

	bot := NewBark("")
	//bot.UseHost("")  // reset host
	//bot.UseToken("") // reset token
	bot.UseSecret("")
	err := bot.
		SetDebug(true).
		SetGroup("test").
		//UseGet().
		//SetBadge(5).
		//SetLevel(1).
		SetSound(SoundChime).
		SetCopy("吃饭啦").
		SetJumpUrl("weixin://").
		SendText("中午啦")
	t.Log(err)
}
