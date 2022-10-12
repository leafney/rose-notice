/**
 * @Author:      leafney
 * @Date:        2022-10-09 17:48
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package bark

type Roboter interface {
	SetHost(host string)
	SetKey(key string)

	SendText(text string) error
	SendMsg(title, body string) error
}

type Robot struct {
	host  string
	key   string
	isGet bool

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

func (r *Robot) SendText(text string) error {
	return r.send("", text)
}

func (r *Robot) SendMsg(title, body string) error {
	return r.send(title, body)
}

func (r *Robot) SetKey(key string) {
	r.key = key
}

func (r *Robot) SetHost(host string) {
	r.host = host
}

// NewRobot host default is `https://api.day.app`
func NewRobot(key, host string) *Robot {
	if len(host) == 0 {
		host = "https://api.day.app"
	}

	return &Robot{
		host: host,
		key:  key,
	}
}

func (r *Robot) SetGet() *Robot {
	r.isGet = true
	return r
}

func (r *Robot) SetGroup(group string) *Robot {
	r.group = group
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

// SetIcon
func (r *Robot) SetIcon(icon string) *Robot {
	r.icon = icon
	return r
}

// SetUrl
func (r *Robot) SetUrl(url string) *Robot {
	r.url = url
	return r
}

func (r *Robot) SetArchive(archive bool) *Robot {
	r.isArchive = archive
	return r
}

// SetLevel 0:active，default; 1:timeSensitive; 2:passive
func (r *Robot) SetLevel(level int) *Robot {
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

func (r *Robot) SetSound(soundName string) *Robot {
	r.soundName = soundName
	return r
}

// SetSilence mute the sounds
func (r *Robot) SetSilence() *Robot {
	r.soundName = SoundSilence
	return r
}

func (r *Robot) SetBadge(badge int64) *Robot {
	r.badge = badge
	return r
}
