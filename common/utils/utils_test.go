/**
 * @Author:      leafney
 * @Date:        2022-10-10 11:23
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package utils

import "testing"

func TestJoinPath(t *testing.T) {
	t.Log(JoinPath("https://www.baidu.com", "aaa", "", "ccc"))

	t.Log(IsNumber("-2344566"))
}
