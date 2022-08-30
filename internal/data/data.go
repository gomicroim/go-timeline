package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"timeline-service/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewSequence)

// Data .
type Data struct {
	Redis *redis.Client
	Mongo *mongo.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return nil, nil, err
	}
	logger.Log(log.LevelInfo, "msg", "connection redis success", "addr", c.Redis.Addr)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(c.Mongo.Source))
	if err != nil {
		return nil, nil, err
	}
	logger.Log(log.LevelInfo, "msg", "connection mongo success", "addr", c.Mongo.Source)

	return &Data{
		Redis: redisClient,
		Mongo: mongoClient,
	}, cleanup, nil
}
