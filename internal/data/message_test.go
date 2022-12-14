package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomicroim/go-timeline/internal/conf"
	"github.com/gomicroim/go-timeline/internal/data/model"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/durationpb"
	"math"
	"testing"
	"time"
)

var (
	testId = "room_2022"
)

func mustSetupMessage() MessageRepo {
	config := &conf.Data{
		Mongo: &conf.Data_Mongo{
			Source:                "mongodb://root:123456@localhost:27017/?authSource=admin&readPreference=primary&directConnection=true&ssl=false",
			ChatDatabase:          "im",
			SyncChatCollection:    "unit_chat_sync",
			HistoryChatCollection: "unit_chat_history",
		},
		Redis: &conf.Data_Redis{
			Addr:         "127.0.0.1:6379",
			Password:     "coffeechat",
			Db:           0,
			ReadTimeout:  durationpb.New(time.Second * 2),
			WriteTimeout: durationpb.New(time.Second * 1),
		},
	}
	data, _, err := NewData(config, log.DefaultLogger)
	if err != nil {
		panic(err)
	}
	return NewMessageRepo(data, config.Mongo.ChatDatabase, config.Mongo.SyncChatCollection, log.DefaultLogger)
}

func TestMessageRepo_Store(t *testing.T) {
	repo := mustSetupMessage()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	msgData := model.NewColumnMap()
	msgData.AddInt64Column("from", 10086)
	msgData.AddInt64Column("to", 17300000000)
	msgData.AddStringColumn("msg", "hello world")

	_, err := repo.Store(ctx, &Message{
		Id:      testId,
		Seq:     11,
		Message: msgData.ToMap(),
	})
	assert.NoError(t, err)
}

func TestMessageRepo_StoreBatch(t *testing.T) {
	repo := mustSetupMessage()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	r, err := repo.StoreBatch(ctx, []*Message{
		{
			//ObjectId: 1,
			Id:      "user_a",
			Seq:     1,
			Message: map[string]interface{}{"from": "1"},
		},
		{
			//ObjectId: 1,
			Id:      "user_a",
			Seq:     2,
			Message: map[string]interface{}{"from": "1"},
		},
		{
			//ObjectId: 2,
			Id:      "user_a",
			Seq:     3,
			Message: map[string]interface{}{"from": "1"},
		},
	})
	if err != nil {
		t.Logf("%v", r)
	}
}

func TestMessageRepo_Load(t *testing.T) {
	repo := mustSetupMessage()
	msg, err := repo.LoadById(context.Background(), "630c7067aec14041bbcba222")
	assert.NoError(t, err)
	t.Log(msg)
}

func TestMessageRepo_Update(t *testing.T) {
	repo := mustSetupMessage()

	msgData := model.NewColumnMap()
	msgData.AddInt64Column("from", 1008600)
	msgData.AddStringColumn("msg333", "hello world")

	err := repo.Update(context.Background(), "630c7067aec14041bbcba222", &Message{
		Id:      testId,
		Seq:     3,
		Message: msgData.ToMap(),
	})
	assert.NoError(t, err)
}

func TestMessageRepo_Delete(t *testing.T) {
	// delete
	repo := mustSetupMessage()
	err := repo.DeleteById(context.Background(), "630c7067aec14041bbcba222")
	assert.NoError(t, err)
}

func TestMessageRepo_Scan(t *testing.T) {
	repo := mustSetupMessage()

	result, err := repo.Scan(context.Background(), testId, ScanParameter{
		From:      2,
		To:        math.MaxInt64,
		MaxCount:  4,
		IsForward: false,
	})
	assert.NoError(t, err)
	t.Logf("%v", result)
}
