/**
 * @Author:      leafney
 * @Date:        2022-10-09 11:19
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package feishu

const (
	msgTypeText = "text"
	msgTypePost = "post"
)

type textMessage struct {
	Timestamp string     `json:"timestamp,omitempty"`
	Sign      string     `json:"sign,omitempty"`
	MsgType   string     `json:"msg_type"`
	Content   textParams `json:"content"`
}

type textParams struct {
	Text string `json:"text"`
}

type postMessage struct {
	Timestamp string     `json:"timestamp,omitempty"`
	Sign      string     `json:"sign,omitempty"`
	MsgType   string     `json:"msg_type"`
	Content   postParams `json:"content"`
}

type postParams struct {
	Post postContent `json:"post"`
}

type postContent struct {
	ZhCn postLanguageContent `json:"zh_cn"`
}

type postLanguageContent struct {
	Title   string          `json:"title"`
	Content [][]postElement `json:"content"`
}

type postElement struct {
	Tag    string `json:"tag"`
	Text   string `json:"text,omitempty"`
	Href   string `json:"href,omitempty"`
	UserId string `json:"user_id,omitempty"`
}

// PostElementBuilder 用于构建富文本消息元素
type PostElementBuilder struct {
	elements [][]postElement
}

// NewPostElementBuilder 创建一个新的富文本消息构建器
func NewPostElementBuilder() *PostElementBuilder {
	return &PostElementBuilder{}
}

// AddLine 添加一个新的行
func (b *PostElementBuilder) AddLine() *PostElementBuilder {
	b.elements = append(b.elements, []postElement{})
	return b
}

// AddText 在当前行添加文本元素
func (b *PostElementBuilder) AddText(text string) *PostElementBuilder {
	if len(b.elements) == 0 {
		b.AddLine()
	}
	lastLine := len(b.elements) - 1
	b.elements[lastLine] = append(b.elements[lastLine], postElement{
		Tag:  "text",
		Text: text,
	})
	return b
}

// AddLink 在当前行添加链接元素
func (b *PostElementBuilder) AddLink(text, href string) *PostElementBuilder {
	if len(b.elements) == 0 {
		b.AddLine()
	}
	lastLine := len(b.elements) - 1
	b.elements[lastLine] = append(b.elements[lastLine], postElement{
		Tag:  "a",
		Text: text,
		Href: href,
	})
	return b
}

// AddAtUser 在当前行添加@用户元素
func (b *PostElementBuilder) AddAtUser(userId string) *PostElementBuilder {
	if len(b.elements) == 0 {
		b.AddLine()
	}
	lastLine := len(b.elements) - 1
	b.elements[lastLine] = append(b.elements[lastLine], postElement{
		Tag:    "at",
		UserId: userId,
	})
	return b
}

// Build 构建富文本消息元素
func (b *PostElementBuilder) Build() [][]postElement {
	return b.elements
}
