package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomicroim/go-timeline/internal/conf"
	"github.com/gomicroim/go-timeline/internal/data"
	"math"
	"time"
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

const (
	kWarnStoreBatchCostTime = 500 // SendGroup 批量写mongo耗时警告，ms
)

type GroupResult struct {
	GroupMember string
	Err         error
	Seq         int64
}

type MessageUseCase interface {
	Send(ctx context.Context, from, to, jsonMsg string) (seq int64, err error)
	SendGroup(ctx context.Context, group string, groupMember []string, jsonMsg string) (result []*GroupResult, err error)
	GetSyncMessage(ctx context.Context, member string, lastRead int64, limit int) ([]data.Message, int64, error)

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

	result = make([]*GroupResult, 0)
	messageArr := make([]*data.Message, len(groupMember))

	// write to all group member receiver
	for k, member := range groupMember {
		seq, err := m.seqCache.Incr(ctx, member)
		if err != nil {
			return nil, err
		}
		messageArr[k] = &data.Message{
			Id:      member,
			Seq:     seq,
			Message: msgObj,
		}
	}
	// write batch
	t1 := time.Now()
	failedMsg, err := m.syncRepo.StoreBatch(ctx, messageArr)
	if err != nil {
		m.log.Warn(err.Error())
		return nil, err
	}
	cost := time.Now().Sub(t1).Milliseconds()
	if cost > kWarnStoreBatchCostTime {
		m.log.Warn("SendGroup cost too time", "time:", cost)
	} else {
		m.log.Debug("SendGroup cost ", "time:", cost)
	}

	if failedMsg != nil {
		for _, v := range failedMsg {
			result = append(result, &GroupResult{
				GroupMember: v.Id,
				Err:         err,
				Seq:         v.Seq,
			})
		}
	}
	return result, nil
}

func (m *messageUseCase) GetSyncMessage(ctx context.Context, member string, lastRead int64, limit int) ([]data.Message, int64, error) {
	seq, err := m.seqCache.GetLatest(ctx, member)
	if err != nil {
		m.log.Warn(err.Error())
		return nil, 0, err
	}

	msgArr, err := m.syncRepo.Scan(ctx, member, data.ScanParameter{
		From:      lastRead,
		To:        math.MaxInt64,
		MaxCount:  limit,
		IsForward: false,
	})
	return msgArr, seq, err
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
