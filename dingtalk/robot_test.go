/**
 * @Author:      leafney
 * @Date:        2022-10-09 09:45
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package dingtalk

import "testing"

func TestNewRobot(t *testing.T) {

	bot := NewDingTalk("")
	// 初始化的设置
	bot.SetDebug(true).UseSecret("")

	// 调用通用方法
	//err := bot.SendText("test 你好，这是测试内容")

	// 使用新的token，调用通用的方法
	//err := bot.UseToken("").SendText("你好呀")

	// 调用独有的方法
	//err := bot.SendTextAt("hello", []string{}, true)

	err := bot.
		SetDebug(true).
		SendMarkdown("### 这是标题", "#### 杭州天气  \n > 9度，@1825718XXXX 西北风1级，空气良89，相对温度73%\n\n > ![screenshot](http://i01.lw.aliimg.com/media/lALPBbCc1ZhJGIvNAkzNBLA_1200_588.png)\n  > ###### 10点20分发布 [天气](http://www.thinkpage.cn/) ")

	t.Log(err)
}
