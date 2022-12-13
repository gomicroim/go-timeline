package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "github.com/gomicroim/go-timeline/api/v1"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
)

var (
	endpoints = flag.String("etcd", "127.0.0.1:2379", "-etcd=127.0.0.1:2379")
)

func main() {
	flag.Parse()

	// init etcd conn
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{*endpoints},
	})
	if err != nil {
		panic(err)
	}
	dis := etcd.New(etcdClient)
	log.Println("connect etcd", "etcd", *endpoints)

	// new timeline client
	client, err := NewTimelineClient("go-timeline", dis)
	if err != nil {
		panic(err)
	}

	log.Println("send msg")

	msgData := map[string]interface{}{"text:": "hello"}
	buffer, _ := json.Marshal(msgData)

	// call grpc `send`
	res, err := client.Send(context.Background(), &v1.SendMessageRequest{
		From:    "user_a",
		To:      "user_z",
		Message: string(buffer),
	})
	if err != nil {
		panic(err)
	}
	log.Println("send success,result:", res)
}

func NewTimelineClient(serviceName string, discovery registry.Discovery) (v1.TimelineClient, error) {
	rpcUserEndpoint := "discovery:///" + serviceName
	conn, err := grpc.DialInsecure(context.Background(),
		grpc.WithEndpoint(rpcUserEndpoint),
		grpc.WithDiscovery(discovery),
	)
	if err != nil {
		return nil, err
	}
	return v1.NewTimelineClient(conn), nil
}
