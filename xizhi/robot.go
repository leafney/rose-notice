package xizhi

import (
	"github.com/leafney/rose-notify/common/notice"
	"github.com/leafney/rose-notify/common/utils"
	"github.com/leafney/rose-notify/common/vars"
)

type XiZhi struct {
	host  string
	token string
	debug bool
}

// Deprecated Manual settings are not supported
func (r *XiZhi) UseHost(url string) notice.Noticer {
	return r
}

func (r *XiZhi) UseToken(token string) notice.Noticer {
	if utils.IsNotEmpty(token) {
		r.token = token
	}
	return r
}

// Deprecated Manual settings are not supported
func (r *XiZhi) UseSecret(secret string) notice.Noticer {
	return r
}

func (r *XiZhi) SendText(text string) error {
	if utils.IsEmpty(text) {
		return vars.ErrParamEmpty
	}
	return r.send(text, "")
}

func (r *XiZhi) SendMarkdown(title, body string) error {
	if utils.IsEmpty(title) {
		return vars.ErrParamEmpty
	}
	return r.send(title, body)
}

func (r *XiZhi) SendMsg(title, body string) error {
	if utils.IsEmpty(title) {
		return vars.ErrParamEmpty
	}
	return r.send(title, body)
}

// NewXiZhi init
func NewXiZhi(token string) *XiZhi {
	return &XiZhi{
		host:  vars.HostXiZhi,
		token: token,
	}
}

func (r *XiZhi) SetDebug(debug bool) *XiZhi {
	r.debug = debug
	return r
}
