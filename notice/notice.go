package notice

type Noticer interface {
	UseHost(url string) Noticer
	UseToken(token string) Noticer
	UseSecret(secret string) Noticer
	SetDebug(debug bool) Noticer

	SendText(text string) error
	SendMarkdown(title, body string) error
}
