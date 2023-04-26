package wochat

import (
	"errors"
	"github.com/leafney/rose-notify/utils"
)

// Roboter is the interface implemented by Robot that can send multiple types of messages.
type Roboter interface {
	SetHost(host string)
	SetKey(key string)

	SendText(content string) error
	SendMarkdown(text string) error
}

// Robot represents a feishu custom robot that can send messages.
type Robot struct {
	host  string
	key   string
	debug bool
}

func (r *Robot) SetHost(host string) *Robot {
	r.host = host
	return r
}

func (r *Robot) SetKey(key string) *Robot {
	r.key = key
	return r
}

func (r *Robot) SendText(content string) error {
	if !utils.IsNotEmpty(content) {
		return errors.New("content empty")
	}

	params := &textMessage{
		MsgType: msgTypeText,
		Text: textParams{
			Content: content,
		},
	}

	return r.send(params)
}

func (r *Robot) SendMarkdown(text string) error {
	params := &markdownMessage{
		MsgType: msgTypeMarkdown,
		Markdown: markdownParams{
			Content: text,
		},
	}
	return r.send(params)
}

// NewRobot returns a roboter that can send messages.
func NewRobot(key string) *Robot {
	return &Robot{
		host: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send",
		key:  key,
	}
}

func (r *Robot) DebugMode() *Robot {
	r.debug = true
	return r
}
