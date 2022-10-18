/**
 * @Author:      leafney
 * @Date:        2022-10-12 10:13
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package pushdeer

import "github.com/leafney/rose-notify/utils"

type Roboter interface {
	SetHost(host string)
	SetKey(key string)

	SendText(text string) error
	SendMsg(title, body string) error
	SendImage(url string) error
	SendMarkdown(title, body string) error
}

type Robot struct {
	host  string
	key   string
	isGet bool
}

func (r *Robot) SetHost(host string) {
	r.host = host
}

func (r *Robot) SetKey(key string) {
	r.key = key
}

func (r *Robot) SendText(text string) error {
	return r.send("text", text, "")
}

func (r *Robot) SendMsg(title, body string) error {
	return r.send("text", title, body)
}

func (r *Robot) SendImage(url string) error {
	return r.send("image", url, "")
}

func (r *Robot) SendMarkdown(title, body string) error {
	return r.send("markdown", title, body)
}

// NewRobot host default is `https://api2.pushdeer.com`
func NewRobot(key, host string) *Robot {
	if !utils.IsNotEmpty(host) {
		host = "https://api2.pushdeer.com"
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
