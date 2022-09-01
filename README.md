# go-timeline

a timeline(消息同步模型) service write by golang, thinks [aliyun tablestore](https://help.aliyun.com/document_detail/89885.html) !!!

depends:

- mongo: >= 5.0
- redis: >= 6.2

## example

### 场景

```text
a -> b: 吃了吗？
b -> a: 吃了

// group_a has member [a, b, c, d]

a -> group_a: 大家好
c -> group_a: 报三围
a -> group_a: 初次见面，多多指教
```

- a给b发送私聊消息
- b回复a
- a在群中发送一条消息
- c在群中发送一条消息
- a在群中发送一条消息

### 发消息

[message_test.go](internal/biz/message_test.go):

```go
var (
	user1, user2, user3, user4 = "a", "b", "c", "d"

	group       = "group_a"
)

assert.NoError(t, send(mc, "a", "b", "吃了吗？"))
assert.NoError(t, send(mc, "b", "a", "吃了"))

assert.NoError(t, sendGroup(mc, "a", "group_a", []string{"a", "c", "d"}, "大家好"))
assert.NoError(t, sendGroup(mc, "c", "group_a", []string{"a", "c", "d"}, "报三围"))
assert.NoError(t, sendGroup(mc, "a", "group_a", []string{"a", "c", "d"}, "初次见面，多多指教"))
```

### 同步消息（扩散写，存N份）

```go
lastRead = 0
msgResult, err := mc.GetSyncMessage(context.Background(), user1, lastRead, math.MaxInt64)
```

- user1(a)

```bash
message_test.go:132: [seq=20] a -> b: 吃了吗？
message_test.go:132: [seq=21] b -> a: 吃了
message_test.go:132: [seq=22] a -> group_a: 大家好
message_test.go:132: [seq=23] c -> group_a: 报三围
message_test.go:132: [seq=24] a -> group_a: 初次见面，多多指教
```

- user2(b)

```bash
message_test.go:132: [seq=16] a -> b: 吃了吗？
message_test.go:132: [seq=17] b -> a: 吃了
```

- user3(c)

```bash
message_test.go:132: [seq=13] a -> group_a: 大家好
message_test.go:132: [seq=14] c -> group_a: 报三围
message_test.go:132: [seq=15] a -> group_a: 初次见面，多多指教
```

- user4(d)

```bash
message_test.go:132: [seq=13] a -> group_a: 大家好
message_test.go:132: [seq=14] c -> group_a: 报三围
message_test.go:132: [seq=15] a -> group_a: 初次见面，多多指教
```

### 查询历史消息（扩散读，只存一份）

- 单聊

```go
lastRead = 0
msgResult, err := mc.GetSingleHistoryMessage(context.Background(), "a", "b", lastRead, 10)
```

```bash
message_test.go:132: [seq=13] a -> b: 吃了吗？
message_test.go:132: [seq=14] b -> a: 吃了
```

- 群聊

```go
lastRead = 0
msgResult, err := mc.GetGroupHistoryMessage(context.Background(), "group_a", lastRead, 10)
```

```bash
message_test.go:132: [seq=22] a -> group_a: 大家好
message_test.go:132: [seq=23] c -> group_a: 报三围
message_test.go:132: [seq=24] a -> group_a: 初次见面，多多指教
```

## screenhost

see [html client](./cmd/client/html/index.html): 

![./doc/img/screenhost.jpg](./doc/img/screenhost.jpg)

## api

### /timeline/send

request:

```bash
curl -X POST http://localhost:8000/timeline/send \
-H 'Content-Type:application/json' \
-d '{"from":"user_a","to":"user_b","message":"{\"from\": \"user_a\",\"to\":\"user_b\",\"text\":\"在吗？\"}"}'
````

response:

```json
{"sequence":"3"}
```

### /timeline/sendGroup

request:

```bash
curl -X POST http://localhost:8000/timeline/sendGroup \
-H 'Content-Type:application/json' \
-d '{"group_name":"test_group","group_members":["user_a","user_b","user_c","user_d"],"message":"{\"from\":\"user_a\",\"to\":\"test_group\",\"is_group\":true,\"text\":\"大家好\"}"}'
```

```json
{
	"group_name": "test_group",
	"group_members": ["user_a", "user_b", "user_c", "user_d"],
	"message": "{\"from\":\"user_a\",\"to\":\"test_group\",\"is_group\":true,\"text\":\"大家好\"}"
}
```

response:

```bash
{"failedMembers":[]}
```

### /timeline/sync

request:

```bash
curl -X GET http://localhost:8000/timeline/sync?member=user_a&last_read=0&count=10
```

response:

```json
{
	"entrySet": [{
		"sequence": "3",
		"message": "{\"from\":\"user_a\",\"text\":\"在吗？\",\"to\":\"user_b\"}"
	}, {
		"sequence": "4",
		"message": "{\"from\":\"user_a\",\"is_group\":true,\"text\":\"大家好\",\"to\":\"test_group\"}"
	}]
}
```

### /timeline/history/single/{from}/{to}

request:

```bash
curl -X GET http://localhost:8000/timeline/history/single/user_a/user_b
```

response:

```json
{
	"entrySet": [{
		"sequence": "7",
		"message": "{\"from\":\"user_a\",\"text\":\"在吗？\",\"to\":\"user_b\"}"
	}]
}
```

### /timeline/history/group/{group}

request:

```bash
curl -X GET http://localhost:8000/timeline/history/group/test_group
```

response:

```json
{
	"entrySet": [{
		"sequence": "1",
		"message": "{\"from\":\"user_a\",\"is_group\":true,\"text\":\"大家好\",\"to\":\"test_group\"}"
	}]
}
```
