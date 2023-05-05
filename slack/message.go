package slack

const (
	msgTypeText = "text"
)

type textMessage struct {
	Text string `json:"text"`
}
