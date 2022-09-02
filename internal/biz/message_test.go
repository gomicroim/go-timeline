package biz

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/durationpb"
	"math"
	"testing"
	"time"
	"timeline-service/internal/conf"
	"timeline-service/internal/data"
)

var (
	user1, user2, user3, user4 = "a", "b", "c", "d"

	group       = "group_a"
	groupMember = []string{user1 /*,user2*/, user3, user4}
)

func mustSetupMessage() MessageUseCase {
	config := &conf.Data{
		Mongo: &conf.Data_Mongo{
			Source:                "mongodb://root:123456@localhost:27017/?authSource=admin&readPreference=primary&directConnection=true&ssl=false",
			ChatDatabase:          "im",
			SyncChatCollection:    "chat_sync",
			HistoryChatCollection: "chat_history",
		},
		Redis: &conf.Data_Redis{
			Addr:         "127.0.0.1:6379",
			Password:     "coffeechat",
			Db:           0,
			ReadTimeout:  durationpb.New(time.Second * 2),
			WriteTimeout: durationpb.New(time.Second * 1),
		},
	}
	client, _, err := data.NewData(config, log.DefaultLogger)
	if err != nil {
		panic(err)
	}
	return NewMessageUseCase(client, config, data.NewSequence(client), log.DefaultLogger)
}

type Msg struct {
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
	IsGroup bool   `json:"is_group,omitempty"`
	Text    string `json:"text,omitempty"`
}

func send(useCase MessageUseCase, from, to, text string) error {
	buffer, err := json.Marshal(&Msg{
		From:    from,
		To:      to,
		IsGroup: false,
		Text:    text,
	})
	if err != nil {
		panic(err.Error())
	}
	_, err = useCase.Send(context.Background(), from, to, string(buffer))
	return err
}

func sendGroup(useCase MessageUseCase, from, group string, groupMember []string, text string) error {
	buffer, err := json.Marshal(&Msg{
		From:    from,
		To:      group,
		IsGroup: true,
		Text:    text,
	})
	if err != nil {
		panic(err.Error())
	}
	_, err = useCase.SendGroup(context.Background(), group, groupMember, string(buffer))
	return err
}

func TestMessageUseCase_Send(t *testing.T) {
	mc := mustSetupMessage()

	// a -> b: 吃了吗？
	// b -> a: 吃了
	//
	// a -> group_a: 大家好
	// c -> group_a: 报三围
	// a -> group_a: 初次见面，多多指教

	assert.NoError(t, send(mc, user1, user2, "吃了吗？"))
	assert.NoError(t, send(mc, user2, user1, "吃了"))

	assert.NoError(t, sendGroup(mc, user1, group, groupMember, "大家好"))
	assert.NoError(t, sendGroup(mc, user3, group, groupMember, "报三围"))
	assert.NoError(t, sendGroup(mc, user1, group, groupMember, "初次见面，多多指教"))
}

func TestMessageUseCase_GetSyncMessage(t *testing.T) {
	mc := mustSetupMessage()
	msgResult, _, err := mc.GetSyncMessage(context.Background(), user4, 0, math.MaxInt64)
	assert.NoError(t, err)
	printResult(t, msgResult)
}

func TestMessageUseCase_GetSingleHistoryMessage(t *testing.T) {
	mc := mustSetupMessage()
	msgResult, err := mc.GetSingleHistoryMessage(context.Background(), user1, user2, 0, 10)
	assert.NoError(t, err)
	printResult(t, msgResult)

	// revert from <- to
	msgResult, err = mc.GetSingleHistoryMessage(context.Background(), user2, user1, 0, 10)
	assert.NoError(t, err)
	printResult(t, msgResult)
}

func TestMessageUseCase_GetGroupHistoryMessage(t *testing.T) {
	mc := mustSetupMessage()
	msgResult, err := mc.GetGroupHistoryMessage(context.Background(), group, 0, 10)
	assert.NoError(t, err)
	printResult(t, msgResult)
}

func printResult(t *testing.T, result []data.Message) {
	for _, v := range result {
		msg := Msg{}
		buffer, err := json.Marshal(v.Message)
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(buffer, &msg))

		t.Logf("[seq=%d] %s -> %s: %s", v.Seq, msg.From, msg.To, msg.Text)
	}
}
