package notify

import (
	"github.com/leafney/rose-notify/wochat"
	"testing"
)

func TestNewNotify(t *testing.T) {

	//n := dingtalk.NewDingTalk("").UseSecret("")
	//n := bark.NewBark("").SetDebug(true).SetAutoCopy(true)
	//n := slack.NewSlack("")
	n := wochat.NewWoChat("")

	ner := NewNotify(n)
	if err := ner.SendNotify("Hello World!"); err != nil {
		t.Log(err)
	}
}
