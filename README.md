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
- [x] Slack
- [x] 息知
- [x] Telegram


| method | default host | custom host | need token/key | support secret | support text | support title | support markdown | details                      |
|--------| --- |-------------|--------| --- | --- | --- | --- |------------------------------|
| DingTalk | ✅ | ❌           | ✅      | ✅ | ✅ | ✅ | ✅ | [README](dingtalk/README.md) |
| FeiShu | ✅ | ❌           | ✅      | ✅ | ✅ | ❌ | ❌ | [README](feishu/README.md)   |
| WoChat | ✅ | ❌           | ✅      | ❌ | ✅ | ❌ | ✅ | [README](wochat/README.md)   |
| Bark   | ✅ | ✅           | ✅      | ❌ | ✅ | ✅ | ❌ | [README](bark/README.md)     |
| Chanify | ✅ | ✅           | ✅      | ❌ | ✅ | ✅ | ❌ | [README](chanify/README.md)  |
| PushDeer | ✅ | ✅           | ✅      | ❌ | ✅ | ✅ | ✅ | [README](pushdeer/README.md) |
| 息知     | ✅ | ❌           | ✅      | ❌ | ✅ | ✅ | ✅ | [README](xizhi/README.md)    |
| Slack  | ✅ | ❌           | ✅      | ❌ | ✅ | ❌ | ❌ | [README](slack/README.md)    |
| Telegram | ✅ | ✅           | ✅      | ❌ | ✅ | ❌ | ✅ | [README](telegram/README.md) |

----

## How to use

### single support

```go
func main() {
	// initialization
	bot := NewDingTalk("token")
	
	// Chain call to set secret
	bot.SetDebug(true).UseSecret("secret")

	// Call the generic method
	err := bot.SendText("Hello World!")
	
	// call unique method
	err := bot.SendTextAt("hello", []string{}, true)

	// Use the new token to call the general method
	err := bot.UseToken("another token").SendText("你好")
}
```


### Multiple support

```go
func main() {
    // use:
    n := dingtalk.NewDingTalk("token").UseSecret("secret")
    // or use:
    n := bark.NewBark("token").SetDebug(true).SetAutoCopy(true)
    // or use: 
    n := slack.NewSlack("token")
    
    ner := NewNotify(n)
    if err := ner.SendNotify("Hello World!"); err != nil {
    t.Log(err)
    }
}

```

----
