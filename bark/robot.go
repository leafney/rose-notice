/**
 * @Author:      leafney
 * @Date:        2022-10-09 17:48
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package bark

import (
	"github.com/leafney/rose-notify/common/utils"
	"github.com/leafney/rose-notify/common/vars"
	"github.com/leafney/rose-notify/notice"
)

//type Roboter interface {
//	SetHost(host string)
//	SetKey(token string)
//
//	SendText(text string) error
//	SendMsg(title, body string) error
//}

type Bark struct {
	host  string
	token string
	isGet bool
	debug bool

	group     string
	copyText  string
	autoCopy  bool
	icon      string
	url       string
	isArchive bool
	level     string
	soundName string
	badge     int64
}

func (r *Bark) UseHost(url string) notice.Noticer {
	if utils.IsNotEmpty(url) {
		r.host = url
	}
	return r
}

func (r *Bark) UseToken(token string) notice.Noticer {
	if utils.IsNotEmpty(token) {
		r.token = token
	}
	return r
}

func (r *Bark) UseSecret(secret string) notice.Noticer {
	return r
}

func (r *Bark) SendText(text string) error {
	return r.send("", text)
}

func (r *Bark) SendMsg(title, body string) error {
	return r.send(title, body)
}

// NewBark host default is `https://api.day.app`
func NewBark(token string) *Bark {
	return &Bark{
		host:  vars.HostBark,
		token: token,
	}
}

func (r *Bark) SetDebug(debug bool) *Bark {
	r.debug = debug
	return r
}

func (r *Bark) SetGet() *Bark {
	r.isGet = true
	return r
}

func (r *Bark) SetGroup(group string) *Bark {
	if utils.IsNotEmpty(group) {
		r.group = group
	}
	return r
}

// SetCopy Specify what to copy
func (r *Bark) SetCopy(text string) *Bark {
	if utils.IsNotEmpty(text) {
		r.copyText = text
	}
	return r
}

// SetAutoCopy auto copy
func (r *Bark) SetAutoCopy(copy bool) *Bark {
	r.autoCopy = copy
	return r
}

// SetIcon icon url
func (r *Bark) SetIcon(icon string) *Bark {
	if utils.IsNotEmpty(icon) {
		r.icon = icon
	}
	return r
}

// SetJumpUrl Jump to the address of URL
func (r *Bark) SetJumpUrl(url string) *Bark {
	if utils.IsNotEmpty(url) {
		r.url = url
	}
	return r
}

func (r *Bark) SetArchive(archive bool) *Bark {
	r.isArchive = archive
	return r
}

// SetLevel 0:active，default; 1:timeSensitive; 2:passive
func (r *Bark) SetLevel(level int) *Bark {
	switch level {
	case 1:
		// timeSensitive
		r.level = "timeSensitive"
	case 2:
		// passive
		r.level = "passive"
	default:
		// active
		r.level = "active"
	}
	return r
}

func (r *Bark) SetSound(name Sound) *Bark {
	r.soundName = name.String()
	return r
}

// SetSilence mute the sounds
func (r *Bark) SetSilence() *Bark {
	r.soundName = SoundSilence
	return r
}

func (r *Bark) SetBadge(badge int64) *Bark {
	r.badge = badge
	return r
}
