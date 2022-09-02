package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Message struct {
	// 插入时，务必设置为0，以使用 mongo的objectId，否则会大幅降低写入性能
	//ObjectId string                 `bson:"_id,omitempty"`
	Id      string                 `bson:"id,omitempty"`      // id，非mongo的对象id
	Seq     int64                  `bson:"seq,omitempty"`     // 连续递增序号
	Message map[string]interface{} `bson:"message,omitempty"` // 数据内容
}

func (m Message) String() string {
	return fmt.Sprintf("id=%s,seq=%d,message=%v", m.Id, m.Seq, m.Message)
}

type ScanParameter struct {
	From      int64 // 起始序号
	To        int64 // 结束序号
	MaxCount  int   // 一次性取出的最大数量
	IsForward bool  // 是否向前搜索
}

type MessageRepo interface {
	Store(ctx context.Context, message *Message) (string, error)
	StoreBatch(ctx context.Context, messages []*Message) (failed []*Message, err error)
	Update(ctx context.Context, objId string, message *Message) error
	LoadById(ctx context.Context, objId string) (*Message, error)
	DeleteById(ctx context.Context, objId string) error

	Load(ctx context.Context, id string, seq int64) (*Message, error)
	Scan(ctx context.Context, id string, param ScanParameter) ([]Message, error)
}

type messageRepo struct {
	log            *log.Helper
	chatCollection *mongo.Collection
}

func NewMessageRepo(data *Data, database, collection string, logger log.Logger) MessageRepo {
	return &messageRepo{
		chatCollection: data.Mongo.Database(database).Collection(collection),
		log:            log.NewHelper(logger),
	}
}

func (m *messageRepo) Store(ctx context.Context, message *Message) (string, error) {
	res, err := m.chatCollection.InsertOne(ctx, message)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).String(), nil
}

func (m *messageRepo) StoreBatch(ctx context.Context, messages []*Message) ([]*Message, error) {
	values := make([]interface{}, len(messages))
	for k := range messages {
		values[k] = messages[k]
	}
	_, err := m.chatCollection.InsertMany(ctx, values)
	if err != nil {
		bulkErr, ok := err.(mongo.BulkWriteException)
		if !ok {
			return nil, err
		}
		index := bulkErr.WriteErrors[0].Index
		return messages[index:], nil
	}
	return nil, nil
}

func (m *messageRepo) Update(ctx context.Context, objId string, message *Message) error {
	id, err := primitive.ObjectIDFromHex(objId)
	if err != nil {
		return err
	}
	data := bson.M{"$set": message}
	_, err = m.chatCollection.UpdateByID(ctx, id, data)
	return err
}

func (m *messageRepo) LoadById(ctx context.Context, objId string) (*Message, error) {
	id, err := primitive.ObjectIDFromHex(objId)
	if err != nil {
		return nil, err
	}
	result := m.chatCollection.FindOne(ctx, bson.D{{"_id", id}})
	if err := result.Err(); err != nil {
		return nil, err
	}
	msg := &Message{}
	if err := result.Decode(msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (m *messageRepo) DeleteById(ctx context.Context, objId string) error {
	id, err := primitive.ObjectIDFromHex(objId)
	if err != nil {
		return err
	}
	_, err = m.chatCollection.DeleteOne(ctx, bson.D{{"_id", id}})
	return err
}

func (m *messageRepo) Load(ctx context.Context, id string, seq int64) (*Message, error) {
	filter := bson.M{"id": id, "seq": seq}
	result := m.chatCollection.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return nil, err
	}
	msg := &Message{}
	if err := result.Decode(msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (m *messageRepo) Scan(ctx context.Context, id string, param ScanParameter) ([]Message, error) {
	from := param.From
	to := param.To
	maxCount := int64(param.MaxCount)

	if param.IsForward {
		to = from
		from -= int64(param.MaxCount)
	}

	filter := bson.M{
		"id":  id,
		"seq": bson.M{"$gt": from, "$lt": to},
	}

	opt := &options.FindOptions{
		Sort:  bson.M{"seq": 1},
		Limit: &maxCount,
	}
	cursor, err := m.chatCollection.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	msg := make([]Message, 0)
	if err = cursor.All(ctx, &msg); err != nil {
		return nil, err
	}

	return msg, nil
}
