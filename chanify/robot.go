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
)

type Roboter interface {
	SetHost(host string)
	SetToken(token string)

	SendText(text string) error
	SendMsg(title, body string) error
	SendLink(link string) error

	//SendImage(url string) error
	//SendAudio(url string) error
	//SendFile(url, desc string) error
	//SendAction(title, text string) error
	//SendTimeline(code string) error
}

type Robot struct {
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

func (r *Robot) SetHost(host string) {
	r.host = host
}

func (r *Robot) SetToken(token string) {
	r.token = token
}

func (r *Robot) SendText(text string) error {
	r.isText = true
	return r.send("", text, "")
}

func (r *Robot) SendMsg(title, body string) error {
	r.isText = true
	return r.send(title, body, "")
}

func (r *Robot) SendLink(link string) error {
	r.isLink = true
	return r.send("", "", link)
}

//func (r *Robot) SendImage(url string) error {
//	r.isImage = true
//	return r.send("", "", url)
//}

// NewRobot host default is `https://api.chanify.net`
func NewRobot(token, host string) *Robot {
	if !utils.IsNotEmpty(host) {
		host = "https://api.chanify.net"
	}

	return &Robot{
		host:  host,
		token: token,
	}
}

func (r *Robot) SetGet() *Robot {
	r.isGet = true
	return r
}

// SetCopy Specify what to copy
func (r *Robot) SetCopy(text string) *Robot {
	r.copyText = text
	return r
}

// SetAutoCopy auto copy
func (r *Robot) SetAutoCopy(copy bool) *Robot {
	r.autoCopy = copy
	return r
}

// SetLevel 0:active，default; 1:time-sensitive; 2:passive
func (r *Robot) SetLevel(level int) *Robot {
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

func (r *Robot) SetPriority(priority int) *Robot {
	r.priority = priority
	return r
}

// SetSound special custom sound
func (r *Robot) SetSound(soundName string) *Robot {
	r.soundName = soundName
	return r
}

// SetDefSound enable default sound
func (r *Robot) SetDefSound() *Robot {
	r.soundName = "1"
	return r
}

func (r *Robot) DebugMode() *Robot {
	r.debug = true
	return r
}
