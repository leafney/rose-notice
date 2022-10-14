/**
 * @Author:      leafney
 * @Date:        2022-10-09 10:53
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package feishu

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

// Roboter is the interface implemented by Robot that can send multiple types of messages.
type Roboter interface {
	SetSecret(secret string)
	SetHost(host string)

	SendText(text string) error
}

// Robot represents a feishu custom robot that can send messages.
type Robot struct {
	host   string
	token  string
	secret string
}

func (r *Robot) SetSecret(secret string) {
	r.secret = secret
}

func (r *Robot) SetHost(host string) {
	r.host = host
}

func (r *Robot) SendText(text string) error {
	params := &textMessage{
		MsgType: msgTypeText,
		Content: textParams{
			Text: text,
		},
	}

	if len(r.secret) > 0 {
		timestamp := fmt.Sprintf("%d", time.Now().Unix())
		signMsg := fmt.Sprintf("%s\n%s", timestamp, r.secret)
		params.Timestamp = timestamp
		params.Sign = sign(signMsg)
	}

	return r.send(params)
}

// NewRobot returns a roboter that can send messages.
func NewRobot(token string) Roboter {
	return &Robot{
		host:  "https://open.feishu.cn/open-apis/bot/v2/hook",
		token: token,
	}
}

func sign(message string) string {
	var data []byte
	hmac256 := hmac.New(sha256.New, []byte(message))
	hmac256.Write(data)
	return base64.StdEncoding.EncodeToString(hmac256.Sum(nil))
}
