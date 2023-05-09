package notify

import "github.com/leafney/rose-notify/common/notice"

type (
	Notify interface {
		SendNotify(msg string) error
	}

	defaultNotify struct {
		name notice.Noticer
	}
)

func (c *defaultNotify) SendNotify(msg string) error {
	return c.name.SendText(msg)
}

func NewNotify(n notice.Noticer) Notify {
	return &defaultNotify{
		name: n,
	}
}
