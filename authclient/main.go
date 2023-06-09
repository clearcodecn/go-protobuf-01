package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	userProto "grpc-starter/proto"
	"log"
	"sync/atomic"
)

func main() {

	// service a ctx
	cc := NewUserClientPool("localhost:8000", 10)
	//md := metadata.Pairs("key", "value", "token", "123")
	ctx := metadata.AppendToOutgoingContext(context.Background(), "token", "123")

	{
		resp, err := cc.Get().Login(ctx, &userProto.LoginRequest{
			Username: "",
			Password: "admin",
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp.Token, resp.User.Id)
	}

	{
		_, err := cc.Get().Register(ctx, &userProto.RegisterRequest{
			Username: "",
			Password: "",
		})
		if err != nil {
			log.Fatal(err)
		}
	}

}

type userClientPool struct {
	clients []userProto.AuthServiceClient
	index   int64
	//userProto.AuthServiceClient
}

func NewUserClientPool(addr string, size int) *userClientPool {
	var clients []userProto.AuthServiceClient
	for i := 0; i < size; i++ {
		cc, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		clients = append(clients, userProto.NewAuthServiceClient(cc))
	}
	return &userClientPool{clients: clients, index: 0}
}

func (p *userClientPool) Get() userProto.AuthServiceClient {
	// 1. 新增 index.
	index := atomic.AddInt64(&p.index, 1)
	return p.clients[int(index)%len(p.clients)]
}
