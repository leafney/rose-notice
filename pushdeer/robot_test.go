/**
 * @Author:      leafney
 * @Date:        2022-10-12 10:18
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package pushdeer

import "testing"

func TestNewRobot(t *testing.T) {
	bot := NewRobot("ef2cce2e-2d00-4493-9a36-2c0c0592eb21", "")
	//bot.SetHost("") // reset host
	//bot.SetKey("")  // reset key

	err := bot.
		DebugMode().
		//SetGet().
		//SendText("我饿啦，我饿啦")
		//SendMsg("中午啦", "我饿啦")
		//SendImage("https://www.baidu.com/img/flexible/logo/pc/index@2.png")
		SendMarkdown("中午啦", "**今天吃什么**\n\n- 鱼香肉丝\n- 宫保鸡丁\n- 麻辣小龙虾")
	//SendMarkdown("### 天气预报", "#### 杭州天气  \n > 9度，@1825718XXXX 西北风1级，空气良89，相对温度73%\n\n > ![screenshot](http://i01.lw.aliimg.com/media/lALPBbCc1ZhJGIvNAkzNBLA_1200_588.png)\n  > ###### 10点20分发布 [天气](http://www.thinkpage.cn/) ")
	t.Log(err)
}
