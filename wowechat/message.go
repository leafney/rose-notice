package wowechat

const (
	msgTypeText     = "text"
	msgTypeMarkdown = "markdown"
)

type textMessage struct {
	MsgType string     `json:"msgtype"`
	Text    textParams `json:"text"`
}

type textParams struct {
	Content             string   `json:"content"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

type markdownMessage struct {
	MsgType  string         `json:"msgtype"`
	Markdown markdownParams `json:"markdown"`
}

type markdownParams struct {
	Content string `json:"content"`
}
