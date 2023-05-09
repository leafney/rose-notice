package notice

type Noticer interface {
	UseHost(url string) Noticer
	UseToken(token string) Noticer
	UseSecret(secret string) Noticer
	SendText(text string) error
}
