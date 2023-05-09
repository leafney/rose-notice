/**
 * @Author:      leafney
 * @Date:        2022-10-12 12:29
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package chanify

import (
	"github.com/leafney/rose-notify/common/utils"
	"github.com/leafney/rose-notify/common/vars"
	"github.com/leafney/rose-notify/notice"
)

type Chanify struct {
	host  string
	token string
	isGet bool
	debug bool

	isText bool
	isLink bool
	//isImage bool

	copyText  string
	autoCopy  bool
	soundName string
	priority  int
	level     string

	//actions  string
	//timeline string
}

func (r *Chanify) UseHost(url string) notice.Noticer {
	if utils.IsNotEmpty(url) {
		r.host = url
	}
	return r
}

func (r *Chanify) UseToken(token string) notice.Noticer {
	if utils.IsNotEmpty(token) {
		r.token = token
	}
	return r
}

// Deprecated
func (r *Chanify) UseSecret(secret string) notice.Noticer {
	return r
}

func (r *Chanify) SendText(text string) error {
	r.isText = true
	return r.send("", text, "")
}

func (r *Chanify) SendMsg(title, body string) error {
	r.isText = true
	return r.send(title, body, "")
}

func (r *Chanify) SendLink(link string) error {
	r.isLink = true
	return r.send("", "", link)
}

//func (r *Chanify) SendImage(url string) error {
//	r.isImage = true
//	return r.send("", "", url)
//}

// NewChanify host default is `https://api.chanify.net`
func NewChanify(token string) *Chanify {
	return &Chanify{
		host:  vars.HostChanify,
		token: token,
	}
}

func (r *Chanify) SetDebug(debug bool) *Chanify {
	r.debug = debug
	return r
}

func (r *Chanify) SetGet() *Chanify {
	r.isGet = true
	return r
}

// SetCopy Specify what to copy
func (r *Chanify) SetCopy(text string) *Chanify {
	if utils.IsNotEmpty(text) {
		r.copyText = text
	}
	return r
}

// SetAutoCopy auto copy
func (r *Chanify) SetAutoCopy(copy bool) *Chanify {
	r.autoCopy = copy
	return r
}

// SetLevel 0:active，default; 1:time-sensitive; 2:passive
func (r *Chanify) SetLevel(level int) *Chanify {
	switch level {
	case 1:
		// timeSensitive
		r.level = "time-sensitive"
	case 2:
		// passive
		r.level = "passive"
	default:
		// active
		r.level = "active"
	}
	return r
}

func (r *Chanify) SetPriority(priority int) *Chanify {
	r.priority = priority
	return r
}

// SetSound special custom sound
func (r *Chanify) SetSound(soundName string) *Chanify {
	r.soundName = soundName
	return r
}

// SetRingtone enable ringtone
func (r *Chanify) SetRingtone() *Chanify {
	r.soundName = "1"
	return r
}
