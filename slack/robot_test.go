package slack

import "testing"

func TestNewRobot(t *testing.T) {
	bot := NewRobot("")

	err := bot.
		SetDebug(false).
		SendText("Hello World!")
	t.Log(err)
}
