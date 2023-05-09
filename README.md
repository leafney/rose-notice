# rose-notify

**Webhook-based message notification service**

----

## Support

- [x] DingTalk
- [x] FeiShu
- [x] WoChat (企业微信)
- [x] Bark
- [x] Chanify
- [x] PushDeer
- [ ] 息知
- [x] Slack
- [ ] Telegram



| method   | default host | custom host | need token/key | support secret | support text | support title | support markdown | details |
|----------| --- |-------------| --- | --- | --- | --- | --- | --- |
| DingTalk | ✅ | ❌           | ✅ | ✅ | ✅ | ✅ | ✅ | [README](dingtalk/README.md) |
| FeiShu   | ✅ | ❌           | ✅ | ✅ | ✅ | ❌ | ❌ | [README](feishu/README.md) |
| WoChat   | ✅ | ❌           | ✅ | ❌ | ✅ | ❌ | ✅ | [README](wochat/README.md) |
| Bark     | ✅ | ✅           | ✅ | ❌ | ✅ | ✅ | ❌ | [README](bark/README.md) |
| Chanify  | ✅ | ✅           | ✅ | ❌ | ✅ | ✅ | ❌ | [README](chanify/README.md) |
| PushDeer | ✅ | ✅           | ✅ | ❌ | ✅ | ✅ | ✅ | [README](pushdeer/README.md) |
| 息知      | ✅ | ❌           | ✅ | ❌ | ✅ | ✅ | ✅ | [README](xizhi/README.md) |
| Slack    | ✅ | ❌           | ✅ | ❌ | ✅ | ❌ | ❌ | [README](slack/README.md) |


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
