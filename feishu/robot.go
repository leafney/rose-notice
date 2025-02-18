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

	"github.com/leafney/rose-notify/common/notice"
	"github.com/leafney/rose-notify/common/utils"
	"github.com/leafney/rose-notify/common/vars"
)

// FeiShu represents a feishu custom robot that can send messages.
type FeiShu struct {
	host   string
	token  string
	secret string
	debug  bool
}

// Deprecated
func (r *FeiShu) UseHost(url string) notice.Noticer {
	//if utils.IsNotEmpty(url) {
	//	r.host = url
	//}
	return r
}

func (r *FeiShu) UseToken(token string) notice.Noticer {
	if utils.IsNotEmpty(token) {
		r.token = token
	}
	return r
}

func (r *FeiShu) UseSecret(secret string) notice.Noticer {
	if utils.IsNotEmpty(secret) {
		r.secret = secret
	}
	return r
}

func (r *FeiShu) SendText(text string) error {
	if utils.IsEmpty(text) {
		return vars.ErrParamEmpty
	}

	params := &textMessage{
		MsgType: msgTypeText,
		Content: textParams{
			Text: text,
		},
	}

	if utils.IsNotEmpty(r.secret) {
		timestamp := fmt.Sprintf("%d", time.Now().Unix())
		signMsg := fmt.Sprintf("%s\n%s", timestamp, r.secret)
		params.Timestamp = timestamp
		params.Sign = sign(signMsg)
	}

	return r.send(params)
}

func (r *FeiShu) SendFullText(title string, elements [][]postElement) error {
	if utils.IsEmpty(title) || len(elements) == 0 {
		return vars.ErrParamEmpty
	}

	params := &postMessage{
		MsgType: msgTypePost,
		Content: postParams{
			Post: postContent{
				ZhCn: postLanguageContent{
					Title:   title,
					Content: elements,
				},
			},
		},
	}

	if utils.IsNotEmpty(r.secret) {
		timestamp := fmt.Sprintf("%d", time.Now().Unix())
		signMsg := fmt.Sprintf("%s\n%s", timestamp, r.secret)
		params.Timestamp = timestamp
		params.Sign = sign(signMsg)
	}

	return r.send(params)
}

// NewFeiShu returns a roboter that can send messages.
func NewFeiShu(token string) *FeiShu {
	return &FeiShu{
		host:  vars.HostFeiShu,
		token: token,
	}
}

func (r *FeiShu) SetDebug(debug bool) *FeiShu {
	r.debug = debug
	return r
}

func sign(message string) string {
	var data []byte
	hmac256 := hmac.New(sha256.New, []byte(message))
	hmac256.Write(data)
	return base64.StdEncoding.EncodeToString(hmac256.Sum(nil))
}
