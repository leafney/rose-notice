## telegram

- [Telegram Bot API](https://core.telegram.org/bots/api#sendmessage)


- Api Format: `https://api.telegram.org/bot<token>/METHOD_NAME`
- Token Format: `123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11`

## Feature

- `NewTelegram(token string)`
- `NewTelegramChatId(token, chatId string)`
- `UseHost(url string)`
- `UseToken(token string)`
- `SendText(text string)`
- `SendMarkdown(body string)` The supported Markdown syntax is detailed in [Telegram Bot API](https://core.telegram.org/bots/api#markdownv2-style)
- `SendHtml(body string)`
- `SetDebug(debug bool)`
- `SetGet(get bool)`
- `SetSilently(silent bool)`
- `SetDisableWebPreview(disable bool)`
- `SetChatId(chatId string)`


----

How to get chat_id:

```
https://api.telegram.org/bot<token>/getUpdates

https://api.telegram.org/bot<token>/deleteWebhook?drop_pending_updates=true
```

----
