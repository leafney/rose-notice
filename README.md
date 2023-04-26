# rose-notify

**Webhook-based message notification service**

----

### Support

- [x] Dingtalk
- [x] Feishu
- [x] WoChat (企业微信)
- [x] Bark
- [x] Chanify
- [x] PushDeer
- [ ] 息知
- [ ] Telegram


| method | default host | custom host | set key/token | show title |
| --- | --- | --- | --- | --- |
| Dingtalk | ✅ | ❌ | ✅ | ❌ |
| Feishu | ✅ | ❌ | ✅ | ❌ | 
| WoChat | ✅ | ❌ | ✅ | ❌ |
| Bark | ✅ | ✅ | ✅ | ✅ |
| Chanify | ✅ | ✅ | ✅ | ✅ |
| PushDeer | ✅ | ✅ | ✅ | ❌ |

----

### Feature

#### Dingtalk

- `SetHost(url string)`
- `SetSecret(secret string)`
- `SendText(content string)`
- `SendTextAt(content string, atMobiles []string, isAtAll bool)`
- `SendLink(title, text, messageURL, picURL string)`
- `SendMarkdown(title, text string)`
- `SendMarkdownAt(title, text string, atMobiles []string, isAtAll bool)`

#### Feishu

- `SetHost(host string)`
- `SetSecret(secret string)`
- `SendText(text string)`

#### Wochat (Work Wechat)

- `SetHost(host string)`
- `SetKey(key string)`
- `SendText(content string) error`
- `SendMarkdown(text string) error`

#### Bark

- `SetHost(host string)`
- `SetKey(key string)`
- `SendText(text string)`
- `SendMsg(title, body string)`

#### Chanify

- `SetHost(host string)`
- `SetToken(token string)`
- `SendText(text string)`
- `SendMsg(title, body string)`
- `SendLink(link string)`

#### PushDeer

- `SetHost(host string)`
- `SetKey(key string)`
- `SendText(text string)`
- `SendMsg(title, body string)`
- `SendImage(url string)`
- `SendMarkdown(title, body string)`

----
