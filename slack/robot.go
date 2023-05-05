package slack

// Roboter is the interface implemented by Robot that can send multiple types of messages.
type Roboter interface {
	SetHost(host string)
	//SetSecret(secret string)

	SendText(text string) error
}

// Robot represents a feishu custom robot that can send messages.
type Robot struct {
	host   string
	token  string
	secret string
	debug  bool
}

//
//func (r *Robot) SetSecret(secret string) {
//	r.secret = secret
//}

func (r *Robot) SetHost(host string) {
	r.host = host
}

func (r *Robot) SendText(text string) error {
	params := &textMessage{
		Text: text,
	}

	return r.send(params)
}

// NewRobot returns a roboter that can send messages.
func NewRobot(token string) *Robot {
	return &Robot{
		host:  "https://hooks.slack.com/services",
		token: token,
	}
}

func (r *Robot) SetDebug(debug bool) *Robot {
	r.debug = debug
	return r
}
