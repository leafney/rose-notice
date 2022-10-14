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

	bot := NewRobot("")
	bot.SetSecret("")
	//bot.SetHost("")
	//err := robot.SendText("test 你好，这是测试内容")
	err := bot.SendMarkdown("### 这是标题", "#### 杭州天气  \n > 9度，@1825718XXXX 西北风1级，空气良89，相对温度73%\n\n > ![screenshot](http://i01.lw.aliimg.com/media/lALPBbCc1ZhJGIvNAkzNBLA_1200_588.png)\n  > ###### 10点20分发布 [天气](http://www.thinkpage.cn/) ")
	t.Log(err)
}
