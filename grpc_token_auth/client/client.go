package main

import (
	"RPCLearning/grpc_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {

	// 设置拦截器
	interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("耗时：%s\n", time.Since(start))
		return err
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))

	// 建立连接并通过连接实例化一个客户端对象
	conn, err := grpc.Dial("localhost:50052", opts...)
	if err != nil {
		panic(err)
	}
	client := proto.NewGreeterClient(conn)

	// 调用服务
	helloReply, err := client.SayHello(context.Background(), &proto.HelloRequest{
		Name: "koweiyi",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(helloReply.Message)
}
