/**
 * @Author:      leafney
 * @Date:        2022-10-09 09:44
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package dingtalk

import (
	"github.com/leafney/rose-notify/common/utils"
	"github.com/leafney/rose-notify/common/vars"
	"github.com/leafney/rose-notify/notice"
)

//
//// Roboter is the interface implemented by DingTalk that can send multiple types of messages.
//type Roboter interface {
//	UseHost(url string)
//	UseToken(token string)
//	UseSecret(secret string)
//
//	SendText(content string) error
//	SendTextAt(content string, atMobiles []string, isAtAll bool) error
//	SendLink(title, text, messageURL, picURL string) error
//	SendMarkdown(title, text string) error
//	SendMarkdownAt(title, text string, atMobiles []string, isAtAll bool) error
//	//SendActionCard(title, text, singleTitle, singleURL, btnOrientation, hideAvatar string) error
//	//SendActionCards(title, text string)error
//	//SendFeedCard(title string)
//}

// DingTalk represents a dingtalk custom robot that can send messages to groups.
type DingTalk struct {
	host   string
	token  string
	secret string
	debug  bool
}

func (r *DingTalk) UseHost(url string) notice.Noticer {
	if utils.IsNotEmpty(url) {
		r.host = url
	}
	return r
}

func (r *DingTalk) UseToken(token string) notice.Noticer {
	if utils.IsNotEmpty(token) {
		r.token = token
	}
	return r
}

func (r *DingTalk) UseSecret(secret string) notice.Noticer {
	if utils.IsNotEmpty(secret) {
		r.secret = secret
	}
	return r
}

func (r *DingTalk) SetDebug(debug bool) notice.Noticer {
	r.debug = debug
	return r
}

func (r *DingTalk) SendText(text string) error {
	if utils.IsEmpty(text) {
		return vars.ErrParamEmpty
	}

	return r.send(&textMessage{
		MsgType: msgTypeText,
		Text: textParams{
			Content: text,
		},
	})
}

func (r *DingTalk) SendTextAt(text string, atMobiles []string, isAtAll bool) error {
	if utils.IsEmpty(text) {
		return vars.ErrParamEmpty
	}

	return r.send(&textMessage{
		MsgType: msgTypeText,
		Text: textParams{
			Content: text,
		},
		At: atParams{
			AtMobiles: atMobiles,
			IsAtAll:   isAtAll,
		},
	})
}

func (r *DingTalk) SendLink(title, text, messageURL, picURL string) error {
	if utils.IsEmpty(title) || utils.IsEmpty(text) || utils.IsEmpty(messageURL) {
		return vars.ErrParamEmpty
	}

	return r.send(&linkMessage{
		MsgType: msgTypeLink,
		Link: linkParams{
			Title:      title,
			Text:       text,
			MessageURL: messageURL,
			PicURL:     picURL,
		},
	})
}

func (r *DingTalk) SendMarkdown(title, body string) error {
	if utils.IsEmpty(title) || utils.IsEmpty(body) {
		return vars.ErrParamEmpty
	}

	return r.send(&markdownMessage{
		MsgType: msgTypeMarkdown,
		Markdown: markdownParams{
			Title: title,
			Text:  body,
		},
	})
}

func (r *DingTalk) SendMarkdownAt(title, body string, atMobiles []string, isAtAll bool) error {
	if utils.IsEmpty(title) || utils.IsEmpty(body) {
		return vars.ErrParamEmpty
	}

	return r.send(&markdownMessage{
		MsgType: msgTypeMarkdown,
		Markdown: markdownParams{
			Title: title,
			Text:  body,
		},
		At: atParams{
			AtMobiles: atMobiles,
			IsAtAll:   isAtAll,
		},
	})
}

// NewDingTalk returns a roboter that can send messages.
func NewDingTalk(token string) *DingTalk {
	return &DingTalk{
		host:  vars.HostDingTalk,
		token: token,
	}
}
