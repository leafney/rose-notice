package wochat

import (
	"github.com/leafney/rose-notify/common/notice"
	"github.com/leafney/rose-notify/common/utils"
	"github.com/leafney/rose-notify/common/vars"
)

//// Roboter is the interface implemented by WoChat that can send multiple types of messages.
//type Roboter interface {
//	SetHost(host string)
//	SetKey(token string)
//
//	SendText(content string) error
//	SendMarkdown(text string) error
//}

// WoChat
type WoChat struct {
	host  string
	token string
	debug bool
}

// Deprecated
func (r *WoChat) UseHost(url string) notice.Noticer {
	//if utils.IsNotEmpty(url) {
	//	r.host = url
	//}
	return r
}

func (r *WoChat) UseToken(token string) notice.Noticer {
	if utils.IsNotEmpty(token) {
		r.token = token
	}
	return r
}

// Deprecated
func (r *WoChat) UseSecret(secret string) notice.Noticer {
	return r
}

func (r *WoChat) SendText(text string) error {
	if utils.IsEmpty(text) {
		return vars.ErrParamEmpty
	}

	params := &textMessage{
		MsgType: msgTypeText,
		Text: textParams{
			Content: text,
		},
	}

	return r.send(params)
}

func (r *WoChat) SendMarkdown(body string) error {
	params := &markdownMessage{
		MsgType: msgTypeMarkdown,
		Markdown: markdownParams{
			Content: body,
		},
	}
	return r.send(params)
}

// NewWoChat returns a roboter that can send messages.
func NewWoChat(token string) *WoChat {
	return &WoChat{
		host:  vars.HostWochat,
		token: token,
	}
}

func (r *WoChat) SetDebug(debug bool) *WoChat {
	r.debug = debug
	return r
}
