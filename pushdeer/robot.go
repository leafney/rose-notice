/**
 * @Author:      leafney
 * @Date:        2022-10-12 10:13
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package pushdeer

import (
	"github.com/leafney/rose-notify/common/utils"
	"github.com/leafney/rose-notify/common/vars"
	"github.com/leafney/rose-notify/notice"
)

//
//type Roboter interface {
//	SetHost(host string)
//	SetKey(token string)
//
//	SendText(text string) error
//	SendMsg(title, body string) error
//	SendImage(url string) error
//	SendMarkdown(title, body string) error
//}

type PushDeer struct {
	host  string
	token string
	isGet bool
	debug bool
}

func (r *PushDeer) UseHost(url string) notice.Noticer {
	if utils.IsNotEmpty(url) {
		r.host = url
	}
	return r
}

func (r *PushDeer) UseToken(token string) notice.Noticer {
	if utils.IsNotEmpty(token) {
		r.token = token
	}
	return r
}

func (r *PushDeer) UseSecret(secret string) notice.Noticer {
	return r
}

// Deprecated
func (r *PushDeer) SendText(text string) error {
	if utils.IsEmpty(text) {
		return vars.ErrParamEmpty
	}

	return r.send("text", text, "")
}

func (r *PushDeer) SetDebug(debug bool) *PushDeer {
	r.debug = debug
	return r
}

func (r *PushDeer) SetGet() *PushDeer {
	r.isGet = true
	return r
}

func (r *PushDeer) SendMarkdown(title, body string) error {
	return r.send("markdown", title, body)
}

func (r *PushDeer) SendMsg(title, body string) error {
	return r.send("text", title, body)
}

func (r *PushDeer) SendImage(url string) error {
	return r.send("image", url, "")
}

// NewPushDeer
func NewPushDeer(token string) *PushDeer {
	return &PushDeer{
		host:  vars.HostPushDeer,
		token: token,
	}
}
