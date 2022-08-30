package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"math"
	"timeline-service/internal/conf"
	"timeline-service/internal/data"
)

type messageUseCase struct {
	syncRepo    data.MessageRepo
	historyRepo data.MessageRepo
	seqCache    data.SequenceCache
	log         *log.Helper
}

var (
	kSeqCachePrefix = "seq"
)

type GroupResult struct {
	GroupMember string
	Err         error
	Seq         int64
}

type MessageUseCase interface {
	Send(ctx context.Context, from, to, jsonMsg string) (seq int64, err error)
	SendGroup(ctx context.Context, group string, groupMember []string, jsonMsg string) (result []*GroupResult, err error)
	GetSyncMessage(ctx context.Context, member string, lastRead int64, limit int) ([]data.Message, error)

	GetSingleHistoryMessage(ctx context.Context, from, to string, lastRead int64, numOfHistory int32) ([]data.Message, error)
	GetGroupHistoryMessage(ctx context.Context, group string, lastRead int64, numOfHistory int32) ([]data.Message, error)
}

// NewMessageUseCase new a MessageUseCase
func NewMessageUseCase(client *data.Data, config *conf.Data, cache data.SequenceCache, logger log.Logger) MessageUseCase {
	msg := &messageUseCase{
		syncRepo:    data.NewMessageRepo(client, config.Mongo.ChatDatabase, config.Mongo.SyncChatCollection, logger),
		historyRepo: data.NewMessageRepo(client, config.Mongo.ChatDatabase, config.Mongo.HistoryChatCollection, logger),
		log:         log.NewHelper(logger),
		seqCache:    cache,
	}
	msg.log.Info("NewMessageUseCase", ",database:", config.Mongo.ChatDatabase,
		",syncStore:", config.Mongo.SyncChatCollection, ",historyStore:", config.Mongo.HistoryChatCollection)
	return msg
}

func (m *messageUseCase) Send(ctx context.Context, from, to, jsonMsg string) (seq int64, err error) {
	// history
	history := buildSingSeqKey(from, to)
	if seq, err = m.store(ctx, history, jsonMsg, m.historyRepo); err != nil {
		return
	}

	sender := from
	if seq, err = m.store(ctx, sender, jsonMsg, m.syncRepo); err != nil {
		return
	}

	// only return receiver sequence
	receiver := to
	return m.store(ctx, receiver, jsonMsg, m.syncRepo)
}

func (m *messageUseCase) SendGroup(ctx context.Context, group string, groupMember []string, jsonMsg string) (result []*GroupResult, err error) {
	// history
	sender := buildGroupSeqKey(group)
	if _, err = m.store(ctx, sender, jsonMsg, m.historyRepo); err != nil {
		return nil, err
	}

	msgObj := make(map[string]interface{}, 0)
	if err = json.Unmarshal([]byte(jsonMsg), &msgObj); err != nil {
		return nil, err
	}

	result = make([]*GroupResult, len(groupMember))

	// write to all group member receiver
	for i, member := range groupMember {
		result[i] = &GroupResult{}
		receiver := member
		result[i].Seq, result[i].Err = m.store(ctx, receiver, jsonMsg, m.syncRepo)
	}
	return result, nil
}

func (m *messageUseCase) GetSyncMessage(ctx context.Context, member string, lastRead int64, limit int) ([]data.Message, error) {
	return m.syncRepo.Scan(ctx, member, data.ScanParameter{
		From:      lastRead,
		To:        math.MaxInt64,
		MaxCount:  limit,
		IsForward: false,
	})
}

func (m *messageUseCase) GetSingleHistoryMessage(ctx context.Context, from, to string, lastRead int64, numOfHistory int32) ([]data.Message, error) {
	sender := buildSingSeqKey(from, to)
	return m.historyRepo.Scan(ctx, sender, data.ScanParameter{
		From:      lastRead,
		To:        math.MaxInt64,
		MaxCount:  int(numOfHistory),
		IsForward: false,
	})
}

func (m *messageUseCase) GetGroupHistoryMessage(ctx context.Context, group string, lastRead int64, numOfHistory int32) ([]data.Message, error) {
	sender := buildGroupSeqKey(group)
	return m.historyRepo.Scan(ctx, sender, data.ScanParameter{
		From:      lastRead,
		To:        math.MaxInt64,
		MaxCount:  int(numOfHistory),
		IsForward: false,
	})
}

func (m *messageUseCase) store(ctx context.Context, id, msg string, repo data.MessageRepo) (int64, error) {
	seq, err := m.seqCache.Incr(ctx, id)
	if err != nil {
		return 0, err
	}

	msgObj := make(map[string]interface{})
	if err = json.Unmarshal([]byte(msg), &msgObj); err != nil {
		return 0, err
	}

	_, err = repo.Store(ctx, &data.Message{
		Id:      id,
		Seq:     seq,
		Message: msgObj,
	})
	return seq, err
}

func buildSingSeqKey(userA, userB string) string {
	if userA > userB {
		return fmt.Sprintf("%s:%s:%s", kSeqCachePrefix, userB, userA)
	}
	return fmt.Sprintf("%s:%s:%s", kSeqCachePrefix, userA, userB)
}

func buildGroupSeqKey(group string) string {
	return fmt.Sprintf("%s:group:%s", kSeqCachePrefix, group)
}
