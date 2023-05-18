package telegram

import (
	"github.com/leafney/rose-notify/common/notice"
	"github.com/leafney/rose-notify/common/utils"
	"github.com/leafney/rose-notify/common/vars"
)

type Telegram struct {
	host  string
	token string
	isGet bool
	debug bool

	chatId             string
	proxy              string
	isMarkdown         bool
	isHtml             bool
	isSilent           bool
	disablePagePreview bool
}

func (r *Telegram) UseHost(url string) notice.Noticer {
	if utils.IsNotEmpty(url) {
		r.host = url
	}
	return r
}

func (r *Telegram) UseToken(token string) notice.Noticer {
	if utils.IsNotEmpty(token) {
		r.token = token
	}
	return r
}

// Deprecated
func (r *Telegram) UseSecret(secret string) notice.Noticer {
	return r
}

func (r *Telegram) SendText(text string) error {
	if utils.IsEmpty(text) {
		return vars.ErrParamEmpty
	}

	return r.send(text)
}

// SendMarkdown The supported Markdown syntax is detailed in https://core.telegram.org/bots/api#markdownv2-style
func (r *Telegram) SendMarkdown(body string) error {
	if utils.IsEmpty(body) {
		return vars.ErrParamEmpty
	}
	r.isMarkdown = true

	return r.send(body)
}

func (r *Telegram) SendHtml(body string) error {
	if utils.IsEmpty(body) {
		return vars.ErrParamEmpty
	}
	r.isHtml = true

	return r.send(body)
}

func NewTelegram(token string) *Telegram {
	return &Telegram{
		host:       vars.HostTelegram,
		token:      token,
		isMarkdown: false,
		isHtml:     false,
	}
}

func NewTelegramChatId(token, chatId string) *Telegram {
	return &Telegram{
		host:       vars.HostTelegram,
		token:      token,
		isMarkdown: false,
		isHtml:     false,
		chatId:     chatId,
	}
}

func (r *Telegram) SetDebug(debug bool) *Telegram {
	r.debug = debug
	return r
}

func (r *Telegram) SetGet(get bool) *Telegram {
	r.isGet = get
	return r
}

// SetSilently Sends the message silently. Users will receive a notification with no sound.
func (r *Telegram) SetSilently(silent bool) *Telegram {
	r.isSilent = silent
	return r
}

// SetDisableWebPreview Disables link previews for links in this message
func (r *Telegram) SetDisableWebPreview(disable bool) *Telegram {
	r.disablePagePreview = disable
	return r
}

// SetChatId Integer or String, Unique identifier for the target chat or username of the target channel (in the format @channelusername)
func (r *Telegram) SetChatId(chatId string) *Telegram {
	if utils.IsNotEmpty(chatId) {
		r.chatId = chatId
	}
	return r
}

func (r *Telegram) SetProxy(proxy string) *Telegram {
	if utils.IsNotEmpty(proxy) {
		r.proxy = proxy
	}
	return r
}
