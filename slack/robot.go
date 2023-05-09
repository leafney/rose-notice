package slack

import (
	"github.com/leafney/rose-notify/common/notice"
	"github.com/leafney/rose-notify/common/utils"
	"github.com/leafney/rose-notify/common/vars"
)

//// Roboter is the interface implemented by Slack that can send multiple types of messages.
//type Roboter interface {
//	SetHost(host string)
//	//SetSecret(secret string)
//
//	SendText(text string) error
//}

// Slack represents a feishu custom robot that can send messages.
type Slack struct {
	host  string
	token string
	debug bool
}

// Deprecated
func (r *Slack) UseHost(url string) notice.Noticer {
	//if utils.IsNotEmpty(url) {
	//	r.host = url
	//}
	return r
}

func (r *Slack) UseToken(token string) notice.Noticer {
	if utils.IsNotEmpty(token) {
		r.token = token
	}
	return r
}

// Deprecated
func (r *Slack) UseSecret(secret string) notice.Noticer {
	return r
}

func (r *Slack) SendText(text string) error {
	if utils.IsEmpty(text) {
		return vars.ErrParamEmpty
	}

	params := &textMessage{
		Text: text,
	}

	return r.send(params)
}

// NewSlack init
func NewSlack(token string) *Slack {
	return &Slack{
		host:  vars.HostSlack,
		token: token,
	}
}

func (r *Slack) SetDebug(debug bool) *Slack {
	r.debug = debug
	return r
}
