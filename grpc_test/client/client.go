package main

import (
	"RPCLearning/grpc_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	client := proto.NewGreeterClient(conn)

	// 设置超时机制
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	helloReply, err := client.SayHello(ctx, &proto.HelloRequest{
		Name: "koweiyi",
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			fmt.Println("解析错误失败")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
		return
	}

	fmt.Println(helloReply.Message)
}
