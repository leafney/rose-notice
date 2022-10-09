/**
 * @Author:      leafney
 * @Date:        2022-10-09 11:19
 * @Project:     rose-notice
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package feishu

const (
	msgTypeText = "text"
)

type textMessage struct {
	Timestamp string     `json:"timestamp,omitempty"`
	Sign      string     `json:"sign,omitempty"`
	MsgType   string     `json:"msg_type"`
	Content   textParams `json:"content"`
}

type textParams struct {
	Text string `json:"text"`
}
