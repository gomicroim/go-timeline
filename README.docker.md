# go-timeline

a timeline(消息同步模型) service write by golang, thinks [aliyun tablestore](https://help.aliyun.com/document_detail/89885.html) !!!

depends:

- mongo: >= 5.0
- redis: >= 6.2

framework:

- [kratos](https://github.com/go-kratos/kratos)

## architecture

### data

mongodb主要有2个文档：

- `chat_sync`: 离线消息库（同步库），扩散写。一个消息写N份，当客户端成功拉取消息后，应删除离线消息，释放存储（目前为了调试方便，不会删除）
- `chat_history`: 持久化存储库（一般使用mysql，这里为了简单直接使用mongo实现），扩散读，一个消息只写一份。有2个作用：
    - 消息漫游。参考QQ消息漫游功能，可以任意查看几个月前，甚至一年前或者任意时间的历史消息。
    - web支持。因为web没有存储能力（无法缓存timeline同步位点），所以可以不用从 chat_sync 走同步流程，直接通过 chat_history 拉历史消息显示即可。

文档结构：

```go
type Message struct {
	Id      string                 `bson:"id,omitempty"`      // id，非mongo的对象id
	Seq     int64                  `bson:"seq,omitempty"`     // 连续递增序号
	Message map[string]interface{} `bson:"message,omitempty"` // 数据内容
}
```

- 离线消息库中：id代表收件人的id，具体的发送者信息、消息内容等需要自己解析Message，timeline服务并不限制存储的结构。
- 持久存储库：id代表会话（会话关系timeline不存储，需要上游服务自行实现），私聊的会话的id为：`samllUserId:bigUserId`，群聊会话的ID就是 `groupId`。故查询的时候，直接按照该规则查询即可。

### 消息序号生成

timeline 序号生成直接采用 redis 实现，保证同一个id下，seq严格递增即可（且连续）。

## example

### 场景

假设有如下聊天场景：

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

同步 a 的消息列表得到如下结果：

```bash
message_test.go:132: [seq=20] a -> b: 吃了吗？
message_test.go:132: [seq=21] b -> a: 吃了
message_test.go:132: [seq=22] a -> group_a: 大家好
message_test.go:132: [seq=23] c -> group_a: 报三围
message_test.go:132: [seq=24] a -> group_a: 初次见面，多多指教
```

## api

see [https://github.com/gomicroim/go-timeline](https://github.com/gomicroim/go-timeline) 

- /timeline/send
- /timeline/sendGroup
- /timeline/sync
- /timeline/history/single/{from}/{to}
- /timeline/history/group/{group}

## Usage

### Mongo & Redis & Go-Timeline

see detail: [docker-compose.yml](https://github.com/gomicroim/go-timeline/docker-compose.yml)

```yaml
version: '3.4'

services:
  timeline:
    image: 'xmcy0011/go-timeline:v0.1'
    container_name: go-timeline
    volumes:
      - ./configs:/data/conf
    environment:
      - TZ=Asia/Shanghai
    deploy:
      restart_policy:
        condition: on-failure
        max_attempts: 3
    networks:
      - timeline
    ports:
      - "9321:8000"
      - "9322:9000"
  mongo:
    image: 'mongo:4.4.16'
    container_name: tl_mongo
    restart: on-failure
    volumes:
      - im_mongo_data:/data/db
      - im_mongo_data:/data/configdb
    ports:
      - '27017:27017'
    networks:
      - timeline
    environment:
      - MONGO_INITDB_ROOT_USERNAME=root
      - MONGO_INITDB_ROOT_PASSWORD=123456
  redis:
    build: deploy/redis
    container_name: tl_redis
    restart: on-failure
    volumes:
      - im_redis_data:/data  # 持久化数据，其他如密码等，在redis.conf里已配置
    networks:
      - timeline
    ports:
      - "6379:6379"     # 端口改掉，预防攻击
volumes: # 声明使用的数据卷，不使用宿主机目录，可以自行更改
  im_mysql_data:
    driver: local
  im_redis_data:
    driver: local
  im_mongo_data:
    driver: local
# 为了简单，使用默认侨接网络 + DNS连接内部服务器方式
# 根据https://docs.docker.com/compose/networking/，docker compose默认创建了网络，但是创建自己的侨接网络更安全
# 另外一个Host主机网络只支持Linux，虽然性能更高
networks:
  timeline:
    driver: bridge # 侨接网络，此时需要使用host访问（service下面的服务名）
    ipam: # see: https://docs.docker.com/compose/compose-file/compose-file-v3/#network
      config:
        - subnet: 173.16.0.0/24
          gateway: 173.16.0.1
```

## Redis config

see: [redis.conf](https://github.com/gomicroim/go-timeline/tree/main/deploy/redis/redis.conf)