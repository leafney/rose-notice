package notify

import (
	"github.com/leafney/rose-notify/dingtalk"
	"testing"
)

func TestNewNotify(t *testing.T) {

	ner := dingtalk.NewDingTalk("").UseSecret("")
	//ner := bark.NewBark("").SetDebug(true).SetAutoCopy(true)
	//ner := slack.NewSlack("")

	n := NewNotify(ner)
	if err := n.SendNotify("Hello World!"); err != nil {
		t.Log(err)
	}
}
