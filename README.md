# rose-notify

**Webhook-based message notification service**

----

## Support

- [x] Dingtalk
- [x] Feishu
- [x] WoChat (企业微信)
- [x] Bark
- [x] Chanify
- [x] PushDeer
- [ ] 息知
- [ ] Telegram
- [x] Slack


| method   | default host | custom host | set key/token | show title |
|----------| --- | -- | --- | --- |
| Dingtalk | ✅ | ❌ | ✅ | ❌ |
| Feishu   | ✅ | ❌ | ✅ | ❌ | 
| WoChat   | ✅ | ❌ | ✅ | ❌ |
| Bark     | ✅ | ✅ | ✅ | ✅ |
| Chanify  | ✅ | ✅ | ✅ | ✅ |
| PushDeer | ✅ | ✅ | ✅ | ❌ |
| 息知       | ✅ | ❌ | ✅ | ❌ |
| Slack    | ✅ | ❌ | ✅ | ❌ |

----

## How to use

```go
func main() {
	// 初始化
	bot := NewDingTalk("token")
	
	// 链式调用,设置secret
	bot.SetDebug(true).UseSecret("secret")

	// 调用通用方法
	err := bot.SendText("Hello World!")
	
	// 调用独有的方法
	err := bot.SendTextAt("hello", []string{}, true)

	// 使用新的token，调用通用的方法
	err := bot.UseToken("another token").SendText("你好")
}
```

----

### Feature

#### Dingtalk

```go

```

#### Feishu

```go

```

#### Wochat (Work Wechat)

- `SetHost(host string)`
- `SetKey(key string)`
- `SendText(content string) error`
- `SendMarkdown(text string) error`

#### Bark



#### Chanify

- `SetHost(host string)`
- `SetToken(token string)`
- `SendText(text string)`
- `SendMsg(title, body string)`
- `SendLink(link string)`

##

----

#### Slack

- `SetHost(host string)`
- `SendText(text string)`

----
