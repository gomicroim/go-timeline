package data

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type SequenceCache interface {
	Incr(ctx context.Context, key string) (int64, error)
	GetLatest(ctx context.Context, key string) (int64, error)
}

type sequence struct {
	client *redis.Client
}

var _ SequenceCache = (*sequence)(nil)

func NewSequence(data *Data) SequenceCache {
	return &sequence{client: data.Redis}
}

func (s *sequence) Incr(ctx context.Context, key string) (int64, error) {
	return s.client.Incr(ctx, key).Result()
}

func (s *sequence) GetLatest(ctx context.Context, key string) (int64, error) {
	return s.client.Get(ctx, key).Int64()
}
